package middlawares

import (
	"auth_in_go/utils"
	"github.com/gin-gonic/gin"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")
		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			return
		}

		claims, err := utils.ParseToken(cookie)
		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Set("role", claims.Role)
		c.Next()
	}
}
