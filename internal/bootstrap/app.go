package bootstrap

import (
	"github.com/gofiber/fiber/v2"

	"github.com/abdansya/go-hexagonal/internal/adapter/handler"
	"github.com/abdansya/go-hexagonal/internal/adapter/repository"
	"github.com/abdansya/go-hexagonal/internal/adapter/router"
	"github.com/abdansya/go-hexagonal/internal/core/service"
	"github.com/abdansya/go-hexagonal/pkg/config"
	"github.com/abdansya/go-hexagonal/pkg/database"
)

func App(cfg *config.Config) *fiber.App {
	// Database connection
	db := database.ConnectDB(cfg)

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)

	// Setup Fiber app
	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	// Setup routes
	router.SetupRoutes(app, userHandler, cfg)

	return app
}
