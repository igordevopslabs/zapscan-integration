package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/igordevopslabs/zapscan-integration/docs"
	"github.com/igordevopslabs/zapscan-integration/pkg/middleware"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine) {
	//Documentation
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.POST("/create", middleware.BasicAuth(), CreateSite)
	r.POST("/start", middleware.BasicAuth(), StartScan)
	r.GET("/list", middleware.BasicAuth(), ListScans)
	r.GET("/results/:scanId", middleware.BasicAuth(), GetScanResult)
}
