package location_test

import (
	"testing"

	"github.com/nazzarr03/location-project/db/entity"
	"github.com/nazzarr03/location-project/internal/location"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockLocationRepository struct {
	mock.Mock
}

func (m *MockLocationRepository) CreateLocation(location *entity.Location) (*entity.Location, error) {
	args := m.Called(location)
	result := args.Get(0)
	return result.(*entity.Location), args.Error(1)
}

func (m *MockLocationRepository) GetLocations(req *location.BaseRequest) ([]entity.Location, error) {
	args := m.Called(req)
	result := args.Get(0)
	return result.([]entity.Location), args.Error(1)
}

func (m *MockLocationRepository) GetLocationByID(id uint) (*entity.Location, error) {
	args := m.Called(id)
	result := args.Get(0)
	return result.(*entity.Location), args.Error(1)
}

func (m *MockLocationRepository) UpdateLocation(id uint, location *entity.Location) (*entity.Location, error) {
	args := m.Called(id, location)
	result := args.Get(0)
	return result.(*entity.Location), args.Error(1)
}

func TestServiceCreateLocation(t *testing.T) {
	mockRepo := new(MockLocationRepository)

	locationEntity := &entity.Location{
		Model:     gorm.Model{ID: 1},
		Name:      "test",
		Latitude:  40.75351,
		Longitude: 74.8531,
		Color:     "#FF0000",
	}

	locationDTO := &location.CreateLocationRequest{
		Name:      "test",
		Latitude:  40.75351,
		Longitude: 74.8531,
		Color:     "#FF0000",
	}

	mockRepo.On("CreateLocation", mock.AnythingOfType("*entity.Location")).Return(locationEntity, nil)

	service := location.NewLocationService(mockRepo)

	createdLocation, err := service.CreateLocation(locationDTO)

	assert.NoError(t, err)
	assert.Equal(t, locationEntity, createdLocation)
	mockRepo.AssertExpectations(t)
}

func TestServiceGetLocations(t *testing.T) {
	mockRepo := new(MockLocationRepository)

	req := &location.BaseRequest{}

	locations := []entity.Location{
		{
			Model:     gorm.Model{ID: 1},
			Name:      "test1",
			Latitude:  40.75351,
			Longitude: 74.8531,
			Color:     "#FF0000",
		},
		{
			Model:     gorm.Model{ID: 2},
			Name:      "test2",
			Latitude:  40.75351,
			Longitude: 74.8531,
			Color:     "#eeeeee",
		},
	}

	mockRepo.On("GetLocations", req).Return(locations, nil)

	service := location.NewLocationService(mockRepo)

	locationResponseDTO, err := service.GetLocations(req)

	assert.NoError(t, err)
	assert.Equal(t, len(locations), locationResponseDTO.Count)
	mockRepo.AssertExpectations(t)
}

func TestServiceGetLocationByID(t *testing.T) {
	mockRepo := new(MockLocationRepository)

	service := location.NewLocationService(mockRepo)

	location := &entity.Location{
		Model:     gorm.Model{ID: 1},
		Name:      "test",
		Latitude:  40.75351,
		Longitude: 74.8531,
		Color:     "#FF0000",
	}

	mockRepo.On("GetLocationByID", location.ID).Return(location, nil)

	locationDTO, err := service.GetLocationByID(location.ID)

	assert.NoError(t, err)
	assert.Equal(t, location.Name, locationDTO.Name)
	mockRepo.AssertExpectations(t)
}

func TestServiceCreateRouteByID(t *testing.T) {
	mockRepo := new(MockLocationRepository)

	service := location.NewLocationService(mockRepo)

	startLocation := &entity.Location{
		Model:     gorm.Model{ID: 1},
		Name:      "test",
		Latitude:  40.7128,
		Longitude: -74.0060,
		Color:     "#FF0000",
	}

	locations := []entity.Location{
		{
			Model:     gorm.Model{ID: 2},
			Name:      "test1",
			Latitude:  40.730610,
			Longitude: -73.935242,
			Color:     "#FF0000",
		},
		{
			Model:     gorm.Model{ID: 3},
			Name:      "test2",
			Latitude:  40.6643,
			Longitude: -73.9385,
			Color:     "#FF0000",
		},

		{
			Model:     gorm.Model{ID: 4},
			Name:      "test3",
			Latitude:  40.730610,
			Longitude: -73.935242,
			Color:     "#FF0000",
		},
	}

	mockRepo.On("GetLocationByID", startLocation.ID).Return(startLocation, nil)
	mockRepo.On("GetLocations", mock.AnythingOfType("*location.BaseRequest")).Return(locations, nil)

	locationResponseDTO, err := service.CreateRouteByID(startLocation.ID)

	assert.NoError(t, err)
	assert.Equal(t, 3, locationResponseDTO.Count)

	expectedOrder := []string{"test1", "test3", "test2"}
	for i, loc := range locationResponseDTO.Data {
		assert.Equal(t, expectedOrder[i], loc.Name)
	}

	mockRepo.AssertExpectations(t)
}
