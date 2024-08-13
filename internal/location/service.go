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
