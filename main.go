package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api-demo/routes"
)

func main() {
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
