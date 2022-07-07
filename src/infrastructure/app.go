package infrastructure

import (
	"github.com/gofiber/fiber/v2"
)

func StartApp() *fiber.App {
	config := fiber.Config{
		Prefork:       false,               // Genera multiples procesos en el mismo puerto, Con docker correr de forma CMD ["sh", "-c", "/app"]
		CaseSensitive: true,                // Valida que las rutas sean escritas de forma estricta... sensible a mayusculas o minusculas
		StrictRouting: true,                // Valida de forma estricta el path
		AppName:       "Golang api v1.0.1", // Configura un nombre de la aplicacion
	}

	app := fiber.New(config)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!!")
	})

	return app
}
