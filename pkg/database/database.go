package database

import (
	"log"
	"os"

	"github.com/igordevopslabs/zapscan-integration/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error to connect database")
	}

	DB.AutoMigrate(&models.Scan{})
}
