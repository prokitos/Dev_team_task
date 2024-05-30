package app

import (
	"module/internal/config"
	"module/internal/database"
	"module/internal/server"

	log "github.com/sirupsen/logrus"
)

func RunApp() {

	// получение конфигов
	cfg := config.ConfigMustLoad()

	// установка логов. установка чтобы показывать логи debug уровня
	log.SetLevel(log.DebugLevel)
	log.Info("the server is starting")

	// миграция и подключение к бд. в горутине.
	database.OpenConnection(cfg.Connect)
	go database.StartMigration()

	// запуск сервера.
	server.ServerStart(cfg.Server)
}
