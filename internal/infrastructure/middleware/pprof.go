package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func SetupPprof(app *fiber.App) {
	app.Use(pprof.New())
}
