package config

import (
	"github.com/igordevopslabs/zapscan-integration/internal/models"
)

func SyncDatabase() {

	DB.AutoMigrate(&models.Scan{})

}
