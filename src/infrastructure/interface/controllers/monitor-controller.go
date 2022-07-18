package controllers

import (
	"fmt"

	"github.com/MaestroShifu/golang-fiber/src/infrastructure/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func AddMonitorController(app *fiber.App) {
	config := config.GetConfigSystem()
	title := fmt.Sprintf("My metrics of %s", config.NAME_APP)

	app.Get("/metrics", monitor.New(monitor.Config{
		Title: title,
	}))
}
