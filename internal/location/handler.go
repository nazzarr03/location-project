package location

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nazzarr03/location-project/db/entity"
	"github.com/nazzarr03/location-project/pkg/validation"
)

type LocationHandler struct {
	LocationServiceInterface LocationServiceInterface
}

func NewLocationHandler(locationService LocationServiceInterface) *LocationHandler {
	return &LocationHandler{LocationServiceInterface: locationService}
}

func (h *LocationHandler) CreateLocation(c *fiber.Ctx) error {
	req := new(CreateLocationRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	location := &entity.Location{
		Name:      req.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Color:     req.Color,
	}

	if err := validation.ValidateLocation(location); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	location, err := h.LocationServiceInterface.CreateLocation(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(location)
}

func (h *LocationHandler) GetLocations(c *fiber.Ctx) error {
	req := new(BaseRequest)
	if err := c.QueryParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	locations, err := h.LocationServiceInterface.GetLocations(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(locations)
}

func (h *LocationHandler) GetLocationByID(c *fiber.Ctx) error {
	id := c.Params("id")

	locationID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	location, err := h.LocationServiceInterface.GetLocationByID(uint(locationID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(location)
}

func (h *LocationHandler) UpdateLocation(c *fiber.Ctx) error {
	id := c.Params("id")

	locationID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	req := new(UpdateLocationRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	location := &entity.Location{
		Name:      req.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Color:     req.Color,
	}

	if err := validation.ValidateLocation(location); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	location, err = h.LocationServiceInterface.UpdateLocation(uint(locationID), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(location)
}
