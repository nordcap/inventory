package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Number int    `gorm:"index" json:"number"`
	Name   string `gorm:"index" json:"name"`
}
