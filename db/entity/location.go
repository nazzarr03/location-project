package entity

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Name      string  `json:"name" gorm:"unique" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"latitude, required"`
	Longitude float64 `json:"longitude" validate:"longitude, required"`
	Color     string  `json:"color" validate:"hexcolor, required"`
}
