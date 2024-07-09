package routes

import (
	"github.com/gofiber/fiber/v2"
	"sample-golang-fiber-started/controllers"
)

func TodoRoute(route fiber.Router) {
	route.Get("", controllers.GetTodos)
}
