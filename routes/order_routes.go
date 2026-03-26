package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api-demo/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/orders", handlers.CreateOrder)
	api.Get("/orders", handlers.GetOrders)
	api.Put("/orders/:id", handlers.UpdateOrder)
	api.Delete("/orders/:id", handlers.DeleteOrder)
}
