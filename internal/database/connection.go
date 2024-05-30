package database

import (
	"fmt"
	"log"
	"module/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GlobalHandler *gorm.DB

func OpenConnection(config models.ConnectConfig) {
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Pass, config.Host, config.Port, config.Name)

	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	GlobalHandler = db
}

func StartMigration() {
	GlobalHandler.AutoMigrate(models.Test_Conversion{})
}
