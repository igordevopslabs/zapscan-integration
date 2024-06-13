package main

import (
	"github.com/gin-gonic/gin"
	docs "github.com/igordevopslabs/zapscan-integration/docs"
	"github.com/igordevopslabs/zapscan-integration/internal/controllers"
	"github.com/igordevopslabs/zapscan-integration/pkg/database"
	"github.com/igordevopslabs/zapscan-integration/pkg/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	database.ConnectToDB()
}

// @title API ZapScan Integration
// @version 1.1
// @description A simple REST API to integration a ZAProxy vulnerability scans
// @host localhost:9000
// @BasePath /
// @SecurityDefinitions BasicAuth
// @in header
// @name Authorization
func main() {

	r := gin.Default()

	//Documentation
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.POST("/create", middleware.BasicAuth(), controllers.CreateSite)
	r.POST("/start", middleware.BasicAuth(), controllers.StartScan)
	r.GET("/list", middleware.BasicAuth(), controllers.ListScans)
	r.GET("/results/:scanId", middleware.BasicAuth(), controllers.GetScanResult)

	r.Run()
}
