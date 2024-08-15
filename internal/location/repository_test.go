package location_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nazzarr03/location-project/db/entity"
	"github.com/nazzarr03/location-project/internal/location"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateLocation(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	mock.ExpectExec("INSERT INTO locations").
		WithArgs("test", 40.75351, 74.8531, "red").
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := location.NewLocationRepository(gormDB)

	locationEntity := &entity.Location{
		Name:      "test",
		Latitude:  40.75351,
		Longitude: 74.8531,
		Color:     "#eeeeee",
	}

	_, err := repo.CreateLocation(locationEntity)

	assert.NoError(t, err)
	mock.ExpectationsWereMet()
}

func TestGetLocations(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	rows := sqlmock.NewRows([]string{"id", "name", "latitude", "longitude", "color"}).
		AddRow(1, "test1", 40.75351, 74.8531, "red").
		AddRow(2, "test2", 40.75351, 74.8531, "blue")

	mock.ExpectQuery("SELECT \\* FROM locations").
		WillReturnRows(rows)

	repo := location.NewLocationRepository(gormDB)

	locations, err := repo.GetLocations(&location.BaseRequest{})

	assert.NoError(t, err)
	assert.Len(t, locations, 2)
	assert.Equal(t, "test1", locations[0].Name)
	assert.Equal(t, "test2", locations[1].Name)
	mock.ExpectationsWereMet()
}

func TestGetLocationByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	rows := sqlmock.NewRows([]string{"id", "name", "latitude", "longitude", "color"}).
		AddRow(1, "test", 40.75351, 74.8531, "red")

	mock.ExpectQuery("SELECT \\* FROM locations WHERE id = \\$1").
		WithArgs(1).
		WillReturnRows(rows)

	repo := location.NewLocationRepository(gormDB)

	location, err := repo.GetLocationByID(1)

	assert.NoError(t, err)
	assert.Equal(t, "test", location.Name)
	mock.ExpectationsWereMet()
}

func TestUpdateLocation(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	repo := location.NewLocationRepository(gormDB)

	rows := sqlmock.NewRows([]string{"id", "name", "latitude", "longitude", "color"}).
		AddRow(1, "test", 40.75351, 74.8531, "red")

	mock.ExpectQuery("SELECT \\* FROM locations WHERE id = \\$1").
		WithArgs(1).
		WillReturnRows(rows)

	locationEntity := &entity.Location{
		Name:      "test",
		Latitude:  40.75351,
		Longitude: 74.8531,
		Color:     "red",
	}

	mock.ExpectExec("UPDATE locations").
		WithArgs("test", 40.75351, 74.8531, "red", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	updatedLocation, err := repo.UpdateLocation(1, locationEntity)

	assert.NoError(t, err)
	assert.Equal(t, "test", updatedLocation.Name)
	mock.ExpectationsWereMet()
}
