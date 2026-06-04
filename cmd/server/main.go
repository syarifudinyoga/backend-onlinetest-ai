package main

import (
	"online-test/config"
	"online-test/pkg/migration"

	auth "online-test/internal/auth"
	"online-test/internal/middleware"
	user "online-test/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	config.ConnectDB(cfg)

	// 🔥 RUN MIGRATION DI START
	migration.RunMigrations("./migrations")

	r := gin.Default()

	r.POST("/login", auth.LoginHandler)
	r.POST("/register", auth.RegisterHandler)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/profile", user.Profile)
	}

	r.Run(":8080")
}
