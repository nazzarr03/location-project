package main

import (
	"log"

	"github.com/joho/godotenv"
	config "github.com/nazzarr03/location-project/configs"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print(".env file not found")
	}

	config.ConnectDB()
}

