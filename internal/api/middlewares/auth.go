package middlewares

import (
	"github.com/antonioalfa22/GoGin-API-REST-Template/pkg/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("authorization")
		if !crypto.ValidateToken(authorizationHeader) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		} else {
			c.Next()
		}
	}
}