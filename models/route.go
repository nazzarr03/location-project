package models

import "gorm.io/gorm"

type Route struct {
	gorm.Model
	StartLocationID uint    `json:"start_location_id" validate:"required"`
	EndLocationID   uint    `json:"end_location_id" validate:"required"`
	Distance        float64 `json:"distance" validate:"required, positive"`
}
