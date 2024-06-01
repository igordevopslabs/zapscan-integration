package main

import (
	"github.com/gin-gonic/gin"
	"github.com/igordevopslabs/zapscan-integration/config"
	docs "github.com/igordevopslabs/zapscan-integration/docs"
	"github.com/igordevopslabs/zapscan-integration/migrations"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	config.LoadEnvs()
	config.ConnectToDB()
	migrations.SyncDatabase()
}

// @title API ZapScan Integration
// @version 1.0
// @description A simple REST API to integration a ZAProxy vulnerability scans
// @host localhost:3000
// @BasePath /
// @SecurityDefinitions BasicAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	//Documentation
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()
}
