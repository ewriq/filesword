package database

import (
	"filesword/model"
	"filesword/utils"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() error {
	config := utils.LoadConfig("./config.ini")
	database, err := gorm.Open(sqlite.Open(config.Database), &gorm.Config{})
	if err != nil {
		panic("Veritabanına bağlanılamadı!")
	}

	err = database.AutoMigrate(&model.Log{})
	if err != nil {
        panic("Migration başarısız!")
    }

	db = database
	fmt.Println("Veritabanı bağlantısı başarılı.")
	return nil
}