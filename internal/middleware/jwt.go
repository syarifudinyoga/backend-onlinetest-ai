package middleware

import (
	"net/http"
	"strings"

	"online-test/config"
	pkgjwt "online-test/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "missing token",
				},
			)
			return
		}

		tokenString := strings.TrimPrefix(
			authHeader,
			"Bearer ",
		)

		cfg := config.LoadConfig()

		claims, err := pkgjwt.ParseToken(
			tokenString,
			cfg.JWTSecret,
		)

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "invalid token",
				},
			)
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
