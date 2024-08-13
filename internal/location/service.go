package location

import (
	"github.com/nazzarr03/location-project/db/entity"
	"github.com/nazzarr03/location-project/pkg/utils"
	"github.com/pkg/errors"
)

type LocationService struct {
	LocationRepository LocationRepository
}

func NewLocationService(locationRepository *LocationRepository) *LocationService {
	return &LocationService{LocationRepository: *locationRepository}
}

func (s *LocationService) CreateLocation(locationDTO *CreateLocationRequest) (*entity.Location, error) {
	location := new(entity.Location)
	err := utils.DTOtoJSON(locationDTO, location)
	if err != nil {
		return nil, errors.New("Failed to convert locationDTO to location")
	}

	createdLocation, err := s.LocationRepository.CreateLocation(location)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create location")
	}

	return createdLocation, nil
}

func (s *LocationService) GetLocations(req *BaseRequest) (*LocationResponseDTO, error) {
	locations, err := s.LocationRepository.GetLocations(req)
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
