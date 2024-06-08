package models

import "gorm.io/gorm"

type Scan struct {
	gorm.Model
	Status  string
	ScanID  string
	Results string `gorm:"type:text"`
}
