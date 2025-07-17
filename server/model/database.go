package model

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	Time string
	Message string
	Type string
}