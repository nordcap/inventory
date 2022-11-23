package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nordcap/inventory/database"
	"github.com/nordcap/inventory/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("стартовая страница")
	})

	app.Post("/api/products", routes.CreateProduct)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)
	log.Fatal(app.Listen(":5000"))
}
