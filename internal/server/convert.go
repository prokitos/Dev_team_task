package server

import (
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
)

func routeConvert(c *fiber.Ctx) error {

	// тут провека токена если надо

	// а потом переход к самому сервису
	return services.ConversionStart(c)

}
