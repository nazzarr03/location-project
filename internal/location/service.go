package location

import (
	"sort"

	"github.com/nazzarr03/location-project/db/entity"
	"github.com/nazzarr03/location-project/pkg/utils"
	"github.com/pkg/errors"
)

type LocationService struct {
	LocationRepositoryInterface LocationRepositoryInterface
}

type LocationServiceInterface interface {
	CreateLocation(locationDTO *CreateLocationRequest) (*entity.Location, error)
	GetLocations(req *BaseRequest) (*LocationResponseDTO, error)
	GetLocationByID(id uint) (*LocationDTO, error)
	UpdateLocation(id uint, locationDTO *UpdateLocationRequest) (*entity.Location, error)
	CreateRouteByID(id uint) (*LocationResponseDTO, error)
}

func NewLocationService(locationRepository LocationRepositoryInterface) *LocationService {
	return &LocationService{LocationRepositoryInterface: locationRepository}
}

func (s *LocationService) CreateLocation(locationDTO *CreateLocationRequest) (*entity.Location, error) {
	location := new(entity.Location)
	err := utils.DTOtoJSON(locationDTO, location)
	if err != nil {
		return nil, errors.New("Failed to convert locationDTO to location")
	}

	createdLocation, err := s.LocationRepositoryInterface.CreateLocation(location)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create location")
	}

	return createdLocation, nil
}

func (s *LocationService) GetLocations(req *BaseRequest) (*LocationResponseDTO, error) {
	locations, err := s.LocationRepositoryInterface.GetLocations(req)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get locations")
	}
	locationDTOs := []LocationDTO{}
	for i := range locations {
		locationDTO := new(LocationDTO)
		err := utils.JSONtoDTO(&locations[i], locationDTO)
		if err != nil {
			return nil, errors.New("Failed to convert location to locationDTO")
		}
		locationDTOs = append(locationDTOs, *locationDTO)
	}

	var resultDTO LocationResponseDTO
	resultDTO.Count = len(locationDTOs)
	resultDTO.Data = locationDTOs

	return &resultDTO, nil
}

func (s *LocationService) GetLocationByID(id uint) (*LocationDTO, error) {
	location, err := s.LocationRepositoryInterface.GetLocationByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get location by id")
	}

	locationDTO := new(LocationDTO)
	err = utils.JSONtoDTO(location, locationDTO)
	if err != nil {
		return nil, errors.New("Failed to convert location to locationDTO")
	}

	return locationDTO, nil
}

func (s *LocationService) UpdateLocation(id uint, locationDTO *UpdateLocationRequest) (*entity.Location, error) {
	location := new(entity.Location)
	err := utils.DTOtoJSON(locationDTO, location)
	if err != nil {
		return nil, errors.New("Failed to convert locationDTO to location")
	}

	updatedLocation, err := s.LocationRepositoryInterface.UpdateLocation(id, location)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to update location")
	}

	return updatedLocation, nil
}

func (s *LocationService) CreateRouteByID(id uint) (*LocationResponseDTO, error) {
	startLocation, err := s.LocationRepositoryInterface.GetLocationByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get start location")
	}

	locations, err := s.LocationRepositoryInterface.GetLocations(&BaseRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get locations")
	}

	type LocationDistance struct {
		Location entity.Location
		Distance float64
	}

	var sortedLocations []LocationDistance
	for i := range locations {
		if locations[i].ID == id {
			continue
		}

		distance := utils.HaversineDistance(startLocation.Latitude, startLocation.Longitude, locations[i].Latitude, locations[i].Longitude)
		sortedLocations = append(sortedLocations, LocationDistance{Location: locations[i], Distance: distance})
	}

	sort.Slice(sortedLocations, func(i, j int) bool {
		return sortedLocations[i].Distance < sortedLocations[j].Distance
	})

	var locationDTOs []LocationDTO
	for i := range sortedLocations {
		locationDTO := new(LocationDTO)
		err := utils.JSONtoDTO(&sortedLocations[i].Location, locationDTO)
		if err != nil {
			return nil, errors.New("Failed to convert location to locationDTO")
		}
		locationDTOs = append(locationDTOs, *locationDTO)
	}

	var resultDTO LocationResponseDTO
	resultDTO.Count = len(locationDTOs)
	resultDTO.Data = locationDTOs

	return &resultDTO, nil
}
