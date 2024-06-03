package repository

import (
	"github.com/igordevopslabs/zapscan-integration/config"
	"github.com/igordevopslabs/zapscan-integration/models"
)

func GetAllScans() ([]models.Scan, error) {
	var scans []models.Scan
	result := config.DB.Find(&scans)
	return scans, result.Error
}

func SaveScan(scan *models.Scan) error {
	result := config.DB.Create(scan)
	return result.Error
}

func UpdateScan(scan *models.Scan) error {
	result := config.DB.Save(scan)
	return result.Error
}

func GetScanByScanID(scanID string) (*models.Scan, error) {
	var scan models.Scan
	result := config.DB.Where("scan_id = ?", scanID).First(&scan)
	if result.Error != nil {
		return nil, result.Error
	}
	return &scan, nil
}
