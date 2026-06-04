package user

import (
	// "net/http"
	// "online-test/config"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {

	userID := c.GetString("user_id")
	role := c.GetString("role")

	c.JSON(200, gin.H{
		"user_id": userID,
		"role":    role,
	})
}
