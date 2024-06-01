package models

import "gorm.io/gorm"

type Scan struct {
	gorm.Model
	URL     string
	Status  string
	Results string `gorm:"type:text"`
}
