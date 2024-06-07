package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, hashAuth := c.Request.BasicAuth()
		if hashAuth && user == "admin" && pass == "password" {
			c.Next()
			return
		}
		c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
