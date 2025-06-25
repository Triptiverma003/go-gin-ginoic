package middleware

import (
	"log"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/Triptiverma003/go-gin-auth/helper"
)

func Authenticate(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(401, gin.H{"error": "Authorization header not present."})
		c.Abort()
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(401, gin.H{"error": "Invalid token format. Expected 'Bearer <token>'."})
		c.Abort()
		return
	}

	token := parts[1]

	claims, msg := helper.ValidateToken(token)
	if msg != "" {
		c.JSON(401, gin.H{"error": msg})
		c.Abort()
		return
	}

	// Optional: Log and set user info in context
	log.Println("Authenticated email:", claims.Email)
	c.Set("email", claims.Email)
	c.Set("userId", claims.UserId)

	c.Next()
}
