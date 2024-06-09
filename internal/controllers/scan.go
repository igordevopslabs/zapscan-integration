package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igordevopslabs/zapscan-integration/internal/services"
)

//Definição das structs para receber os parametros

// struct para receber os inputs de create scan (spidering tree)
type CreateSiteRequest struct {
	URLs []string `json:"urls"`
}

// struct para receber os intputs de start scan (active scan)
type StartScanRequest struct {
	URLs []string `json:"urls"`
}

// @Summary     Create Site
// @Description Create Site for new scan
// @ID          create-scan
// @Tags  	    create-scans
// @Accept      json
// @Produce     json
// @Param       Authorization header string true "Authorization header"
// @Security    BasicAuth
// @Success     200
// @Failure     500
// @Router      /create [post]
func CreateSite(c *gin.Context) {
	var req CreateSiteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scanIDs, err := services.CreateSite(req.URLs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Sites added to scan tree",
		"site":    req.URLs,
		"site_id": scanIDs,
	})
}

// @Summary     Start Scan
// @Description Start a new active scan
// @ID          post-scans
// @Tags  	    start-scans
// @Accept      json
// @Produce     json
// @Param       Authorization header string true "Authorization header"
// @Security    BasicAuth
// @Success     200
// @Failure     500
// @Router      /start [post]
func StartScan(c *gin.Context) {
	var req StartScanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scanIDs, err := services.StartScan(req.URLs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "New scan started for site.",
		"site":     req.URLs,
		"scan_ids": scanIDs,
	})
}

// @Summary     List all scans
// @Description List All existing scansIds
// @ID          list-all
// @Tags  	    get-scans
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     500
// @Router      /list [get]
func ListScans(c *gin.Context) {
	scans, err := services.ListScans()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"scans": scans})
}

// @Summary     Show results by scanId
// @Description Show scan results by scanId
// @ID          show-results
// @Tags  	    get-scans
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     500
// @Router      /results/:scanId [get]
func GetScanResult(c *gin.Context) {
	scanId := c.Param("scanId")
	result, err := services.GetScanResult(scanId)
	if err != nil {
		if err.Error() == "scan not completed" {
			c.JSON(http.StatusOK, gin.H{
				"message": "Scan still in progress",
				"ScanID":  scanId,
				"Status":  result.Status,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}
