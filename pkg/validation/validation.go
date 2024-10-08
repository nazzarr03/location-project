package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/nazzarr03/location-project/db/entity"
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
	validate.RegisterValidation("latitude", LatitudeValidation)
	validate.RegisterValidation("longitude", LongitudeValidation)
	validate.RegisterValidation("hexcolor", HexColorValidation)
	validate.RegisterValidation("positive", PositiveValidation)
}

func ValidateStruct(s interface{}) error {
	if validate == nil {
		validate = validator.New()
	}

	return validate.Struct(s)
}

func LatitudeValidation(fl validator.FieldLevel) bool {
	value := fl.Field().Float()
	return value >= -90 && value <= 90
}

func LongitudeValidation(fl validator.FieldLevel) bool {
	value := fl.Field().Float()
	return value >= -180 && value <= 180
}

func HexColorValidation(fl validator.FieldLevel) bool {
	color := fl.Field().String()
	if len(color) != 7 || color[0] != '#' {
		return false
	}

	for _, c := range color[1:] {
		if !(('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')) {
			return false
		}
	}

	return true
}

func PositiveValidation(fl validator.FieldLevel) bool {
	value := fl.Field().Float()
	return value > 0
}

func ValidateLocation(location *entity.Location) error {
	return ValidateStruct(location)
}
