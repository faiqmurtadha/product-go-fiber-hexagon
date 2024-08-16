package http

import (
	"product-go-fiber-hexagon/internal/core/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, productService *service.ProductService) {
	productHandler := NewProductHandler(productService)

	app.Post("/products", productHandler.Create)
	app.Get("/products", productHandler.GetAll)
	app.Get("/products/:id", productHandler.FindById)
	app.Put("/products/:id", productHandler.Update)
	app.Delete("/products/:id", productHandler.Delete)
}
