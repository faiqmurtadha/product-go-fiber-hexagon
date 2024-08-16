package main

import (
	"product-go-fiber-hexagon/internal/adapters/http"
	"product-go-fiber-hexagon/internal/core/service"
	"product-go-fiber-hexagon/internal/infrastructure/database"
	"product-go-fiber-hexagon/internal/infrastructure/middleware"
	"product-go-fiber-hexagon/internal/infrastructure/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middleware.SetupPprof(app)

	productCollection := database.InitMongoDB()

	productRepository := repository.NewProductRepository(productCollection)
	productService := service.NewProductService(productRepository)

	// Register adapters
	http.SetupRoutes(app, productService)

	app.Listen(":3000")
}
