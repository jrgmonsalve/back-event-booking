package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jrgmonsalve/back-event-booking/utils"
)

func AuthMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	claims, err := utils.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.Set("userId", int64(claims["userId"].(float64)))

	c.Next()
}
