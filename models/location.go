package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	PlaceName string    `gorm:"index:idx_placename,unique" json:"place_name"`
	Storage   []Storage //связь с табл storage

}
