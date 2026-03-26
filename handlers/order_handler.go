package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go-fiber-api-demo/models"
)

var orders []models.Order
var idCounter = 1

func CreateOrder(c *fiber.Ctx) error {
	order := new(models.Order)

	if err := c.BodyParser(order); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	order.ID = idCounter
	order.Status = "pending"
	idCounter++

	orders = append(orders, *order)

	return c.JSON(order)
}

func GetOrders(c *fiber.Ctx) error {
	return c.JSON(orders)
}

func UpdateOrder(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	for i, o := range orders {
		if o.ID == id {
			orders[i].Status = "completed"
			return c.JSON(orders[i])
		}
	}

	return c.Status(404).JSON(fiber.Map{"error": "Order not found"})
}

func DeleteOrder(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	for i, o := range orders {
		if o.ID == id {
			orders = append(orders[:i], orders[i+1:]...)
			return c.JSON(fiber.Map{"message": "Deleted"})
		}
	}

	return c.Status(404).JSON(fiber.Map{"error": "Order not found"})
}
