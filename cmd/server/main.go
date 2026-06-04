package main

import (
	"online-test/config"
	"online-test/pkg/migration"

	auth "online-test/internal/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	config.ConnectDB(cfg)

	// 🔥 RUN MIGRATION DI START
	migration.RunMigrations("./migrations")

	r := gin.Default()

	r.POST("/register", auth.RegisterHandler)
	r.POST("/login", auth.LoginHandler)

	r.Run(":8080")
}
