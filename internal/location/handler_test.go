package location_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/nazzarr03/location-project/db/entity"
	"github.com/nazzarr03/location-project/internal/location"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLocationService struct {
	mock.Mock
}

func (m *MockLocationService) CreateLocation(location *location.CreateLocationRequest) (*entity.Location, error) {
	args := m.Called(location)
	result := args.Get(0)
	return result.(*entity.Location), args.Error(1)
}

func (m *MockLocationService) GetLocations(req *location.BaseRequest) (*location.LocationResponseDTO, error) {
	args := m.Called(req)
	result := args.Get(0)
	return result.(*location.LocationResponseDTO), args.Error(1)
}

func (m *MockLocationService) GetLocationByID(id uint) (*location.LocationDTO, error) {
	args := m.Called(id)
	result := args.Get(0)
	return result.(*location.LocationDTO), args.Error(1)
}

func (m *MockLocationService) UpdateLocation(id uint, location *location.UpdateLocationRequest) (*entity.Location, error) {
	args := m.Called(id, location)
	result := args.Get(0)
	return result.(*entity.Location), args.Error(1)
}

func (m *MockLocationService) CreateRouteByID(id uint) (*location.LocationResponseDTO, error) {
	args := m.Called(id)
	result := args.Get(0)
	return result.(*location.LocationResponseDTO), args.Error(1)
}

func TestHandlerCreateLocation(t *testing.T) {
	mockService := new(MockLocationService)

	handler := location.NewLocationHandler(mockService)

	app := fiber.New()
	app.Post("/api/v1/locations", handler.CreateLocation)

	req := location.CreateLocationRequest{
		Name:      "test",
		Latitude:  40.75351,
		Longitude: 74.8531,
		Color:     "#FF0000",
	}

	mockService.On("CreateLocation", &req).Return(&entity.Location{}, nil)

	resp, err := app.Test(httptest.NewRequest("POST", "/api/v1/locations", bytes.NewReader([]byte(`{"name":"test","latitude":40.75351,"longitude":74.8531,"color":"#FF0000"}`))))

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestHandlerGetLocations(t *testing.T) {
	mockService := new(MockLocationService)

	handler := location.NewLocationHandler(mockService)

	app := fiber.New()
	app.Get("/api/v1/locations", handler.GetLocations)

	req := location.BaseRequest{}

	mockService.On("GetLocations", &req).Return(&location.LocationResponseDTO{}, nil)

	resp, err := app.Test(httptest.NewRequest("GET", "/api/v1/locations", nil))

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestHandlerGetLocationByID(t *testing.T) {
	mockService := new(MockLocationService)

	handler := location.NewLocationHandler(mockService)

	app := fiber.New()
	app.Get("/api/v1/locations/:id", handler.GetLocationByID)

	mockService.On("GetLocationByID", uint(1)).Return(&location.LocationDTO{}, nil)

	resp, err := app.Test(httptest.NewRequest("GET", "/api/v1/locations/1", nil))

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	mockService.AssertExpectations(t)
}

func TestHandlerCreateRouteByID(t *testing.T) {
	mockService := new(MockLocationService)

	handler := location.NewLocationHandler(mockService)

	app := fiber.New()
	app.Get("/api/v1/route/:id", handler.CreateRouteByID)

	startLocationID := uint(1)

	routeResponseDTO := &location.LocationResponseDTO{}
	routeResponseDTO.Count = 3
	routeResponseDTO.Data = []location.LocationDTO{
		{Name: "test1"},
		{Name: "test2"},
		{Name: "test3"},
	}

	mockService.On("CreateRouteByID", startLocationID).Return(routeResponseDTO, nil)

	req := httptest.NewRequest("GET", "/api/v1/route/1", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var responseBody location.LocationResponseDTO
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 3, responseBody.Count)

	for i, loc := range responseBody.Data {
		assert.Equal(t, routeResponseDTO.Data[i].Name, loc.Name)
	}

	mockService.AssertExpectations(t)
}
