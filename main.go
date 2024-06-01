package main

import (
	"github.com/gin-gonic/gin"
	"github.com/igordevopslabs/zapscan-integration/config"
	"github.com/igordevopslabs/zapscan-integration/migrations"
)

func init() {
	config.LoadEnvs()
	config.ConnectToDB()
	migrations.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.Run()
}
