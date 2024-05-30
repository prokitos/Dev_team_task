package app

import (
	"fmt"
	"module/internal/config"
	"module/internal/server"

	log "github.com/sirupsen/logrus"
)

func RunApp() {

	// получение конфигов
	cfg := config.ConfigMustLoad()
	fmt.Println(cfg)

	// установка логов. установка чтобы показывать логи debug уровня
	log.SetLevel(log.DebugLevel)
	log.Info("the server is starting")

	// миграция и подключение к бд. в горутине.

	// запуск сервера.
	server.ServerStart(cfg.Server)
}
