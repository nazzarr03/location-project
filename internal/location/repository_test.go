package location_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nazzarr03/location-project/db/entity"
	"github.com/nazzarr03/location-project/internal/location"
	"github.com/pkg/errors"
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

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"locations\"").
		WithArgs("test", 40.75351, 74.8531, "#FF0000").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectRollback()

	repo := location.NewLocationRepository(gormDB)

	locationEntity := &entity.Location{
		Name:      "test",
		Latitude:  40.75351,
		Longitude: 74.8531,
		Color:     "#FF0000",
	}

	_, err := repo.CreateLocation(locationEntity)

	assert.Error(t, err)
	mock.ExpectationsWereMet()
}

func TestGetLocations(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	rows := sqlmock.NewRows([]string{"id", "name", "latitude", "longitude", "color"}).
		AddRow(1, "test1", 40.75351, 74.8531, "#FF0000").
		AddRow(2, "test2", 40.75351, 74.8531, "#eeeeee")

	mock.ExpectQuery("SELECT \\* FROM \"locations\"").
		WillReturnRows(rows)

	repo := location.NewLocationRepository(gormDB)

	locations, err := repo.GetLocations(&location.BaseRequest{})

	assert.NoError(t, err)
	assert.NotEmpty(t, locations)
	assert.Len(t, locations, 2)
	if len(locations) > 0 {
		assert.Equal(t, "test1", locations[0].Name)
	}
	if len(locations) > 1 {
		assert.Equal(t, "test2", locations[1].Name)
	}
	mock.ExpectationsWereMet()
}

func TestGetLocationByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	rows := sqlmock.NewRows([]string{"id", "name", "latitude", "longitude", "color"}).
		AddRow(1, "test", 40.75351, 74.8531, "#FF0000")

	mock.ExpectQuery("SELECT \\* FROM \"locations\" WHERE \"locations\".\"id\" = \\$1 AND \"locations\".\"deleted_at\" IS NULL ORDER BY \"locations\".\"id\" LIMIT \\$2").
		WithArgs(1, 1).
		WillReturnRows(rows)

	repo := location.NewLocationRepository(gormDB)

	location, err := repo.GetLocationByID(1)

	assert.NoError(t, err)
	if location != nil {
		assert.Equal(t, "test", location.Name)
	}
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
		AddRow(1, "test", 40.75351, 74.8531, "#FF0000")

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT \\* FROM locations WHERE id = \\$1").
		WithArgs(1).
		WillReturnRows(rows)

	locationEntity := &entity.Location{
		Name:      "test",
		Latitude:  40.75351,
		Longitude: 74.8531,
		Color:     "#FF0000",
	}

	mock.ExpectExec("UPDATE \"locations\" SET \"updated_at\"=\\$1,\"name\"=\\$2,\"latitude\"=\\$3,\"longitude\"=\\$4,\"color\"=\\$5 WHERE id = \\$6 AND \"locations\".\"deleted_at\" IS NULL").
		WithArgs(sqlmock.AnyArg(), "test", 40.75351, 74.8531, "#FF0000", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit().WillReturnError(errors.New("forced error"))
	mock.ExpectRollback()

	updatedLocation, err := repo.UpdateLocation(1, locationEntity)

	assert.Error(t, err)
	if updatedLocation != nil {
		assert.Equal(t, "test", updatedLocation.Name)
	}
	mock.ExpectationsWereMet()
}
