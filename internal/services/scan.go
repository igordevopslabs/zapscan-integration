package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/igordevopslabs/zapscan-integration/config"
	"github.com/igordevopslabs/zapscan-integration/internal/models"
	"github.com/igordevopslabs/zapscan-integration/internal/repository"
)

type ActiveScan struct {
	ID string `json:"id"`
}

type ActiveScansResponse struct {
	Scans []ActiveScan `json:"scans"`
}

type ScanResponse struct {
	Scan string `json:"scan"`
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
	config.LoadEnvs()
	zapApiKey = os.Getenv("ZAP_KEY")
	zapEndpoint = os.Getenv("ZAP_EDP")
}

func CreateSite(urls []string) ([]string, error) {

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

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response body: %v", err)
			return nil, errors.New("failed to read response body")
		}

		//variavel para receber o scanID da APi do zaproxy
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

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response body: %v", err)
			return nil, errors.New("failed to start scan")
		}

		var scanResponse ScanResponse
		if err := json.Unmarshal(body, &scanResponse); err != nil {
			log.Printf("Failed to parse response JSON: %v", err)
			return nil, errors.New("failed to start scan")
		}

		fmt.Println("ZAP API response: Scan ID =", scanResponse.Scan)

		if resp.StatusCode != http.StatusOK {
			log.Printf("Non-OK HTTP status: %s", resp.Status)
			return nil, errors.New("failed to start scan")
		}

		scanIDs = append(scanIDs, scanResponse.Scan)
	}
	return scanIDs, nil
}

func ListActiveScans() ([]ActiveScan, error) {
	zapUrl := fmt.Sprintf("%s/JSON/ascan/view/scans/?apikey=%s", zapEndpoint, zapApiKey)
	resp, err := http.Get(zapUrl)
	if err != nil {
		return nil, errors.New("failed to get active scans from ZAP API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("failed to read response body")
	}

	var activeScansResponse ActiveScansResponse
	if err := json.Unmarshal(body, &activeScansResponse); err != nil {
		return nil, errors.New("failed to parse response JSON")
	}

	log.Printf("Active scans: %v", activeScansResponse.Scans)

	return activeScansResponse.Scans, nil
}

func CheckScanStatus(scanId string) (string, error) {
	zapUrl := fmt.Sprintf("%s/JSON/ascan/view/status/?apikey=%s&scanId=%s", zapEndpoint, zapApiKey, scanId)
	resp, err := http.Get(zapUrl)
	if err != nil {
		return "", errors.New("failed to get scan status from ZAP API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("failed to read response body")
	}

	var statusResponse map[string]interface{}
	if err := json.Unmarshal(body, &statusResponse); err != nil {
		return "", errors.New("failed to parse response JSON")
	}

	if status, ok := statusResponse["status"]; ok {
		return status.(string), nil
	}
	return "", errors.New("status not found in response")
}

func CheckScanCompletion(scanId string) (bool, error) {
	zapUrl := fmt.Sprintf("%s/JSON/ascan/view/status/?apikey=%s&scanId=%s", zapEndpoint, zapApiKey, scanId)
	resp, err := http.Get(zapUrl)
	if err != nil {
		return false, errors.New("failed to get scan status from ZAP API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, errors.New("failed to read response body")
	}

	var statusResponse map[string]interface{}
	if err := json.Unmarshal(body, &statusResponse); err != nil {
		return false, errors.New("failed to parse response JSON")
	}

	if status, ok := statusResponse["status"]; ok {
		return status.(string) == "100", nil
	}
	return false, errors.New("status not found in response")
}

func ListScans() ([]models.Scan, error) {
	return repository.GetAllScans()
}

func GetScanResult(scanId string) (models.Scan, error) {
	// Verificar o status do scan

	zapUrl := fmt.Sprintf("%s/JSON/ascan/view/status/?apikey=%s&scanId=%s", zapEndpoint, zapApiKey, scanId)
	resp, err := http.Get(zapUrl)
	if err != nil {
		return models.Scan{}, errors.New("failed to get scan status from ZAP API")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Scan{}, errors.New("failed to read response body")
	}

	var statusResponse map[string]interface{}
	if err := json.Unmarshal(body, &statusResponse); err != nil {
		return models.Scan{}, errors.New("failed to parse response JSON")
	}

	if status, ok := statusResponse["status"]; ok && status.(string) != "100" {
		return models.Scan{}, errors.New("scan not completed")
	}

	// Verificar se o scan já existe no banco de dados
	existingScan, err := repository.GetScanByScanID(scanId)
	if err == nil && existingScan != nil {
		return *existingScan, nil
	}

	// Obter os resultados do scan
	zapUrl = fmt.Sprintf("%s/JSON/ascan/view/scanProgress/?apikey=%s&scanId=%s", zapEndpoint, zapApiKey, scanId)
	resp, err = http.Get(zapUrl)
	if err != nil {
		return models.Scan{}, errors.New("failed to get scan result from ZAP API")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Scan{}, fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Scan{}, errors.New("failed to read response body")
	}

	var scanResult ScanResult
	if err := json.Unmarshal(body, &scanResult); err != nil {
		return models.Scan{}, errors.New("failed to parse response JSON")
	}

	// Salvar a URL, Scan ID e resultados no banco de dados

	scan := models.Scan{
		ScanID:  scanId,
		Status:  "completed",
		Results: string(body),
	}

	// Se o scan já existe, atualizar o registro existente
	if existingScan != nil {
		existingScan.Status = "completed"
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
