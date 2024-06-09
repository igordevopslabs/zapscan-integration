package repository

import (
	"github.com/igordevopslabs/zapscan-integration/pkg/database"
	"github.com/igordevopslabs/zapscan-integration/pkg/logger"

	"github.com/igordevopslabs/zapscan-integration/internal/models"
	"go.uber.org/zap"
)

func GetAllScans() ([]models.Scan, error) {
	logger.LogInfo("repository", zap.String("operation", "repository.get_all_scans"))
	var scans []models.Scan
	result := database.DB.Find(&scans)
	return scans, result.Error
}

func SaveScan(scan *models.Scan) error {
	logger.LogInfo("repository", zap.String("operation", "repository.save_scan"))
	result := database.DB.Create(scan)
	return result.Error
}

func UpdateScan(scan *models.Scan) error {
	logger.LogInfo("repository", zap.String("operation", "repository.update_scan"))
	result := database.DB.Save(scan)
	return result.Error
}

func GetScanByScanID(scanID string) (*models.Scan, error) {
	logger.LogInfo("repository", zap.String("operation", "repository.get_scan_by_id"))
	var scan models.Scan
	result := database.DB.Where("scan_id = ?", scanID).First(&scan)
	if result.Error != nil {
		return nil, result.Error
	}
	return &scan, nil
}
