package location

import (
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
