package controllers

import (
	"github.com/MaestroShifu/golang-fiber/src/infrastructure/database"
	"github.com/gofiber/fiber/v2"
)

func AddProductController(app *fiber.App) {
	app.Get("/products", func(c *fiber.Ctx) error {
		conn, _ := database.Connect()
		rows, _ := conn.Raw("select version();").Rows()
		defer rows.Close()
		return c.SendString("root")
	})
}
