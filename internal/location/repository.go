package location

import (
	"github.com/nazzarr03/location-project/db/entity"
	"gorm.io/gorm"
)

type LocationRepository struct {
	Db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{Db: db}
}

func (r *LocationRepository) CreateLocation(location *entity.Location) (*entity.Location, error) {
	if err := r.Db.Create(location).Error; err != nil {
		return nil, err
	}

	return location, nil
}

func (r *LocationRepository) GetLocations(req *BaseRequest) ([]entity.Location, error) {
	var locations []entity.Location
	query := r.Db
	if req.Limit != 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset != 0 {
		query = query.Offset(req.Offset)
	}

	if err := query.Find(&locations).Error; err != nil {
		return nil, err
	}

	return locations, nil
}
