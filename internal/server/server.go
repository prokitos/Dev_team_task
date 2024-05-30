package server

import (
	"log"
	"module/internal/models"

	"github.com/gofiber/fiber/v2"
)

func ServerStart(port models.ServerConfig) {

	app := fiber.New()

	handlers(app)

	log.Fatal(app.Listen(port.Port))
}

func handlers(instance *fiber.App) {

	instance.Post("/convert", routeConvert)

}
