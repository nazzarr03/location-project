package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Name      string  `json:"name" gorm:"unique;type:varchar(100);not null"`
	Latitude  float64 `json:"latitude" gorm:"type:decimal(9,6);not null"`
	Longitude float64 `json:"longitude" gorm:"type:decimal(9,6);not null"`
	Color     string  `json:"color" gorm:"type:varchar(7);not null"`
}
