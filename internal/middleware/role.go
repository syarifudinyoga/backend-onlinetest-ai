package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleGuard(
	roles ...string,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		role, exists := c.Get("role")

		if !exists {
			c.AbortWithStatusJSON(
				http.StatusForbidden,
				gin.H{
					"message": "role not found",
				},
			)
			return
		}

		userRole := role.(string)

		for _, r := range roles {

			if userRole == r {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(
			http.StatusForbidden,
			gin.H{
				"message": "access denied",
			},
		)
	}
}
