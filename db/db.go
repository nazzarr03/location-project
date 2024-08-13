package db

import (
	"fmt"
	"log"
	"os"

	"github.com/nazzarr03/location-project/db/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := Db.AutoMigrate(&entity.Location{}, &entity.Route{}); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	log.Println("Successfully connected to database")
}
