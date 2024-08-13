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
