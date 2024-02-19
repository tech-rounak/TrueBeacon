package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helper "github.com/tech-rounak/true-beacon/helpers"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusForbidden, gin.H{"error": "token required"})
			c.Abort()
			return
		}
		claims, err := helper.ValidateToken(clientToken)

		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("Email", claims.Name)
		c.Set("UserName", claims.UserName)

	}
}
