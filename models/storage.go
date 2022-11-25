package models

import "gorm.io/gorm"

type Storage struct {
	gorm.Model
	FileName   string `gorm:"index:idx_filename,unique" json:"file_name"`
	ProductID  uint   `json:"product_id"`
	LocationID uint   `json:"location_id"`
}
