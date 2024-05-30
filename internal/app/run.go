package app

import (
	log "github.com/sirupsen/logrus"
)

func RunApp() {

	// установка логов. установка чтобы показывать логи debug уровня
	log.SetLevel(log.DebugLevel)
	log.Info("the server is starting")

	// миграция и подключение к бд. в горутине.

	// запуск сервера.

}
