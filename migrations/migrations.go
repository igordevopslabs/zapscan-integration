package migrations

import (
	"github.com/igordevopslabs/zapscan-integration/config"
	"github.com/igordevopslabs/zapscan-integration/models"
)

func SyncDatabase() {

	config.DB.AutoMigrate(&models.Scan{})

}
