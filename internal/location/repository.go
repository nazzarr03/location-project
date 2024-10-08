package location

import (
	"github.com/nazzarr03/location-project/db/entity"
	"gorm.io/gorm"
)

type LocationRepository struct {
	Db *gorm.DB
}

type LocationRepositoryInterface interface {
	CreateLocation(location *entity.Location) (*entity.Location, error)
	GetLocations(req *BaseRequest) ([]entity.Location, error)
	GetLocationByID(id uint) (*entity.Location, error)
	UpdateLocation(id uint, location *entity.Location) (*entity.Location, error)
}

func NewLocationRepository(db *gorm.DB) LocationRepositoryInterface {
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

func (r *LocationRepository) GetLocationByID(id uint) (*entity.Location, error) {
	location := new(entity.Location)
	if err := r.Db.First(location, id).Error; err != nil {
		return nil, err
	}
	return location, nil
}

func (r *LocationRepository) UpdateLocation(id uint, location *entity.Location) (*entity.Location, error) {
	if err := r.Db.Model(&entity.Location{}).Where("id = ?", id).Updates(location).Error; err != nil {
		return nil, err
	}
	return location, nil
}
