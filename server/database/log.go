package database

import (
	"filesword/model"
	"time"
)

func Log(message, types string) error {
	log := model.Log{
		Time: time.Now().Format("2006-01-02 15:04:05"),
		Message: message,
		Type: types,
	}
	return db.Create(&log).Error
}