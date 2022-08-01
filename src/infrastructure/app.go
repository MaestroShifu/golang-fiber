package infrastructure

import (
	"github.com/MaestroShifu/golang-fiber/src/infrastructure/interface/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
)

func basicConfigApp(app *fiber.App) {
	// Sirve para comprimir la respuesta usando [gzip]
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	// Sirve para activar los cors en nuestra aplicacion
	app.Use(cors.New(cors.ConfigDefault))

	// Conjunto de middlewares (Cositas para seguridad)
	app.Use(helmet.New())

	// Configuracion para tener logs en nuestro servidor
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
}

func StartApp() *fiber.App {
	config := fiber.Config{
		Prefork:       false,               // Genera multiples procesos en el mismo puerto, Con docker correr de forma CMD ["sh", "-c", "/app"]
		CaseSensitive: true,                // Valida que las rutas sean escritas de forma estricta... sensible a mayusculas o minusculas
		StrictRouting: true,                // Valida de forma estricta el path
		AppName:       "Golang api v1.0.1", // Configura un nombre de la aplicacion
	}

	app := fiber.New(config)

	// Add configuracion
	basicConfigApp(app)

	// Add controllers
	controllers.AddMonitorController(app)
	controllers.AddProductController(app)

	/* 	app.Get("/", func(c *fiber.Ctx) error {
		configApp := GetConfigSystem()

		return c.JSON(configApp)
	}) */

	return app
}
