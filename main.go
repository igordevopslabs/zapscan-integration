package main

import (
	"github.com/gin-gonic/gin"
	"github.com/igordevopslabs/zapscan-integration/internal/controllers"
	"github.com/igordevopslabs/zapscan-integration/pkg/database"
)

func init() {
	database.ConnectToDB()
}

// @title API ZapScan Integration
// @version 1.1
// @description A simple REST API to integration a ZAProxy vulnerability scan
// @host localhost:9000
// @BasePath /
// @SecurityDefinitions BasicAuth
// @in header
// @name Authorization
func main() {

	r := gin.Default()
	controllers.RegisterRoutes(r)
	r.Run()
}
