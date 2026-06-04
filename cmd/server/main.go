package main

import (
	"online-test/config"
	"online-test/pkg/migration"

	auth "online-test/internal/auth"
	"online-test/internal/middleware"
	question "online-test/internal/question"
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

	questionRepo := &question.Repository{}
	questionService := question.NewService(questionRepo)
	questionHandler := question.NewHandler(questionService)

	admin := r.Group("/api/admin")
	admin.Use(
		middleware.AuthMiddleware(),
		middleware.RoleGuard("admin"),
	)

	{
		admin.POST("/questions", questionHandler.CreateQuestion)
		admin.POST("/questions/:id/options", questionHandler.AddOptions)

		admin.GET("/questions", questionHandler.List)
		admin.GET("/questions/:id", questionHandler.Detail)
		admin.DELETE("/questions/:id", questionHandler.Delete)
	}

	r.Run(":8080")
}
