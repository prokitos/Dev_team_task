package server

import (
	"log"
	"module/internal/models"
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
)

func ServerStart(port models.ServerConfig) {

	app := fiber.New()

	handlers(app)

	log.Fatal(app.Listen(port.Port))
}

func handlers(instance *fiber.App) {

	instance.Post("/convert", services.Something)

}
