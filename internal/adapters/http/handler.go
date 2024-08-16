package http

import (
	"product-go-fiber-hexagon/internal/core/model"
	"product-go-fiber-hexagon/internal/core/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) Create(c *fiber.Ctx) error {
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	createdProduct, err := h.service.Create(&product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Return a response with message
	response := fiber.Map{
		"message": "Product created successfully",
		"product": createdProduct,
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	page, err := strconv.ParseInt(c.Query("page", "1"), 10, 64)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.ParseInt(c.Query("limit", "10"), 10, 64)
	if err != nil || limit < 1 {
		limit = 10
	}

	name := c.Query("name", "")

	paginatedProducts, err := h.service.GetAll(page, limit, name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(paginatedProducts)
}

func (h *ProductHandler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := h.service.FindById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return c.JSON(product)
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	updatedProduct, err := h.service.Update(id, &product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Return a response with message
	response := fiber.Map{
		"message": "Product updated successfully",
		"product": updatedProduct,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.service.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("Product deleted successfully")
}
