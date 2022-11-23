package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nordcap/inventory/database"
	"github.com/nordcap/inventory/models"
)

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	database.Database.Db.Create(&product)

	return c.Status(fiber.StatusOK).SendString("Создание МЦ прошло успешно!")

}
