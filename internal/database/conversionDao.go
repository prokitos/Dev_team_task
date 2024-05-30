package database

import (
	"module/internal/models"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func CreateData(c *fiber.Ctx, curConversion *models.Test_Conversion) error {

	// создать запись, и вывести ошибку если всё плохо
	if result := GlobalHandler.Create(&curConversion); result.Error != nil {
		log.Debug("create record error!")
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	// отправить что всё хорошо
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": curConversion}})
}
