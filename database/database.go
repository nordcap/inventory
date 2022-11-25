package database

import (
	"log"
	"os"

	"github.com/nordcap/inventory/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка доступа к базе данных \n", err.Error())
		os.Exit(1)
	}

	log.Println("Соединение к БД прошло успешно")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Запуск миграций")
	db.AutoMigrate(&models.Product{}, &models.Location{}, &models.Storage{})

	Database = DbInstance{Db: db}
}
