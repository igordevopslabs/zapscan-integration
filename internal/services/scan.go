package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/igordevopslabs/zapscan-integration/pkg/logger"

	"github.com/igordevopslabs/zapscan-integration/internal/models"
	"github.com/igordevopslabs/zapscan-integration/internal/repository"
	"go.uber.org/zap"
)

type ScanResponse struct {
	Scan string `json:"scan" `
}

type ScanResult struct {
	Progress string `json:"progress"`
	State    string `json:"state"`
	Alerts   []struct {
		Risk string `json:"risk"`
		Name string `json:"name"`
	} `json:"alerts"`
}

var zapApiKey string
var zapEndpoint string

func init() {
	zapApiKey = os.Getenv("ZAP_KEY")
	zapEndpoint = os.Getenv("ZAP_EDP")
}

func CreateSite(urls []string) ([]string, error) {
	logger.LogInfo("service", zap.String("operation", "service.create_site"))

	scanIDs := []string{}

	//valida a existencia das variaveis de ambiente
	if zapApiKey == "" {
		return nil, errors.New("ZAP_KEY environment variable is not set")
	}
	if zapEndpoint == "" {
		return nil, errors.New("ZAP_EDP environment variable is not set")
	}

	//itera em cima da req do user sobre os site urls passados
	for _, url := range urls {
		zapUrl := fmt.Sprintf("%s/JSON/spider/action/scan/?apikey=%s&url=%s", zapEndpoint, zapApiKey, url)
		resp, err := http.Get(zapUrl)
		if err != nil || resp.StatusCode != http.StatusOK {
			return nil, errors.New("failed to add site to scan tree")
		}

		defer resp.Body.Close()

		//recebe o corpo da requisição a parir da resposta
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response body: %v", err)
			return nil, errors.New("failed to read response body")
		}

		//pega os dados a partir do body e joga na variavel scanResponse
		var scanResponse ScanResponse
		if err := json.Unmarshal(body, &scanResponse); err != nil {
			log.Printf("Failed to parse response JSON: %v", err)
			return nil, errors.New("failed to parse response JSON")
		}

		scanIDs = append(scanIDs, scanResponse.Scan)
	}
	return scanIDs, nil
}

func StartScan(urls []string) ([]string, error) {
	logger.LogInfo("service", zap.String("operation", "service.start_scan"))
	scanIDs := []string{}

	if zapApiKey == "" {
		return nil, errors.New("ZAP_KEY environment variable is not set")
	}
	if zapEndpoint == "" {
		return nil, errors.New("ZAP_EDP environment variable is not set")
	}

	for _, url := range urls {
		zapUrl := fmt.Sprintf("%s/JSON/ascan/action/scan/?apikey=%s&url=%s", zapEndpoint, zapApiKey, url)

		resp, err := http.Get(zapUrl)
		if err != nil || resp.StatusCode != http.StatusOK {
			return nil, errors.New("failed to start scan")
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.LogError("Failed to read body", err)
			return nil, errors.New("failed to start scan")
		}

		var scanResponse ScanResponse
		if err := json.Unmarshal(body, &scanResponse); err != nil {
			logger.LogError("Failed to parse JSON", err)
			return nil, errors.New("failed to start scan")
		}
		if resp.StatusCode != http.StatusOK {
			logger.LogInfo("Problem", zap.String("http_status", resp.Status))
			return nil, errors.New("failed to start scan")
		}

		scanIDs = append(scanIDs, scanResponse.Scan)
	}
	return scanIDs, nil
}

func ListScans() ([]models.Scan, error) {
	logger.LogInfo("service", zap.String("operation", "service.list_scans"))
	return repository.GetAllScans()
}

func GetScanResult(scanId string) (models.Scan, error) {
	logger.LogInfo("service", zap.String("operation", "service.get_scan_result"))
	//atribui scanID no campo ScanID no banco
	scan := models.Scan{ScanID: scanId}

	// Verificar o status do scan
	zapUrl := fmt.Sprintf("%s/JSON/ascan/view/status/?apikey=%s&scanId=%s", zapEndpoint, zapApiKey, scanId)
	resp, err := http.Get(zapUrl)
	if err != nil {
		return scan, errors.New("failed to get scan status from ZAP API")
	}
	defer resp.Body.Close()

	//nova variavel para receber o body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return scan, errors.New("failed to read response body")
	}

	//passa os dados do body para a variavel statusResponse
	var statusResponse map[string]interface{}
	if err := json.Unmarshal(body, &statusResponse); err != nil {
		return scan, errors.New("failed to parse response JSON")
	}

	//converte status para string para facilitar a manipulção
	status, ok := statusResponse["status"].(string)
	if !ok {
		return scan, errors.New("status not found or is not a string in response")
	}

	//atribui status no campo status no banco
	scan.Status = status

	// Verificação de status simplificada
	if status != "100" {
		return scan, errors.New("scan not completed")
	}

	// Obter os resultados do scan
	zapUrl = fmt.Sprintf("%s/JSON/ascan/view/scanProgress/?apikey=%s&scanId=%s", zapEndpoint, zapApiKey, scanId)
	resp, err = http.Get(zapUrl)
	if err != nil {
		return scan, errors.New("failed to get scan result from ZAP API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return scan, fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return scan, errors.New("failed to read response body")
	}

	var scanResult ScanResult
	if err := json.Unmarshal(body, &scanResult); err != nil {
		return scan, errors.New("failed to parse response JSON")
	}

	// Salvar a URL, Scan ID e resultados no banco de dados
	scan.Results = string(body)

	// Verificar se o scan já existe no banco de dados
	existingScan, _ := repository.GetScanByScanID(scanId)

	// Se o scan já existe, atualizar o registro existente
	if existingScan != nil {
		existingScan.Status = "100"
		existingScan.Results = string(body)
		if err := repository.UpdateScan(existingScan); err != nil {
			return scan, errors.New("failed to update scan in database")
		}
		return *existingScan, nil
	}

	// Se o scan não existe, criar um novo registro
	if err := repository.SaveScan(&scan); err != nil {
		return scan, errors.New("failed to save scan to database")
	}

	return scan, nil
}
