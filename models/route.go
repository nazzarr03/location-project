package models

import "gorm.io/gorm"

type Route struct {
	gorm.Model
	StartLocationID uint    `json:"start_location_id" gorm:"not null"`
	EndLocationID   uint    `json:"end_location_id" gorm:"not null"`
	Distance        float64 `json:"distance" gorm:"not null"`
}
