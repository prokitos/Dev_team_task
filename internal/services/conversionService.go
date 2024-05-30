package services

import "github.com/gofiber/fiber/v2"

func Something(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusAccepted)
}
