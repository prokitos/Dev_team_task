package services

import (
	"module/internal/database"
	"module/internal/models"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func ConversionStart(c *fiber.Ctx) error {

	var curOperation models.Test_Conversion

	// получение из BODY
	if err := c.BodyParser(&curOperation); err != nil {
		log.Debug("body parse error")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// получение из URL
	curOperation.Base_currency = c.Query("base_currency", "")
	curOperation.Target_currency = c.Query("target_currency", "")

	var rate float64 = 0.1
	// обращение к внешнему API и получение коэфициента

	curOperation.Result_amount = curOperation.Amount * rate

	// запись в бд
	database.CreateData(c, &curOperation)

	return c.SendStatus(fiber.StatusAccepted)
}
