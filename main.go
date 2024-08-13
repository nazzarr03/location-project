package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	db "github.com/nazzarr03/location-project/db"
	"github.com/nazzarr03/location-project/internal/location"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print(".env file not found")
	}

	db.ConnectDB()

	database := db.Db

	locationRepository := location.NewLocationRepository(database)
	locationService := location.NewLocationService(locationRepository)
	locationHandler := location.NewLocationHandler(locationService)

	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})

	api := app.Group("/api/v1")

	locationApi := api.Group("/locations")
	locationApi.Post("/", locationHandler.CreateLocation)
	locationApi.Get("/", locationHandler.GetLocations)
	locationApi.Get("/:id", locationHandler.GetLocationByID)
	locationApi.Put("/", locationHandler.UpdateLocation)

	app.Listen(":8081")
}
