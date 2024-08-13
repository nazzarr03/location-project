package location

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nazzarr03/location-project/db/entity"
	"github.com/nazzarr03/location-project/pkg/validation"
)

type LocationHandler struct {
	LocationService LocationService
}

func NewLocationHandler(locationService *LocationService) *LocationHandler {
	return &LocationHandler{LocationService: *locationService}
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

	location, err := h.LocationService.CreateLocation(req)
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

	locations, err := h.LocationService.GetLocations(req)
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

	location, err := h.LocationService.GetLocationByID(uint(locationID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(location)
}

func (h *LocationHandler) UpdateLocation(c *fiber.Ctx) error {
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

	locations, err := h.LocationService.UpdateLocation(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(locations)
}
