package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"sample-golang-fiber-started/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	SetupRoutes(app)

	err := app.Listen(":3000")

	if err != nil {
		panic(err)
	}
}

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the endpoint",
		})
	})

	api := app.Group("/api")

	api.Get("", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "You are at the api endpoint",
		})
	})

	routes.TodoRoute(api.Group("/todos"))
}
