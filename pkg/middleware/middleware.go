package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/igordevopslabs/zapscan-integration/config"
)

var basicUser string
var basicPass string

func init() {
	config.LoadEnvs()
	basicUser = os.Getenv("BASIC_USER")
	basicPass = os.Getenv("BASIC_PASS")
}

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, hashAuth := c.Request.BasicAuth()
		if hashAuth && user == basicUser && pass == basicPass {
			c.Next()
			return
		}
		c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
