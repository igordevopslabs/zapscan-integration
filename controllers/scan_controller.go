package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igordevopslabs/zapscan-integration/services"
)

//Definição das structs para receber os parametros

type CreateSiteRequest struct {
	URLs []string `json:"urls"`
}

type StartScanRequest struct {
	URLs []string `json:"urls"`
}

func CreateSite(c *gin.Context) {
	var req CreateSiteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateSite(req.URLs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sites added to scan tree"})
}

func StartScan(c *gin.Context) {
	var req StartScanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	go func() {
		if err := services.StartScan(req.URLs); err != nil {
			log.Printf("Error starting scan: %v", err)
		}
	}()
	c.JSON(http.StatusOK, gin.H{"message": "Scan started"})
}

func ListAllActiveScans(c *gin.Context) {
	scans, err := services.ListActiveScans()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"scans": scans})
}

func GetScanStatus(c *gin.Context) {
	scanId := c.Param("scanId")
	status, err := services.CheckScanStatus(scanId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": status})
}

func ListScans(c *gin.Context) {
	scans, err := services.ListScans()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"scans": scans})
}

func GetScanResult(c *gin.Context) {
	scanId := c.Param("scanId")
	result, err := services.GetScanResult(scanId)
	if err != nil {
		if err.Error() == "scan not completed" {
			c.JSON(http.StatusOK, gin.H{"message": "Scan not completed yet"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}
