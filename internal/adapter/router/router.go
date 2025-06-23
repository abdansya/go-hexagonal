package router

import (
	"github.com/abdansya/go-hexagonal/internal/adapter/handler"
	"github.com/abdansya/go-hexagonal/pkg/config"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler, cfg *config.Config) {
	router := app.Group("/api/v1")

	// Public routes (e.g., login, register) can be added here
	// router.Post("/login", loginHandler)

	// JWT Middleware for protected routes
	router.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cfg.JWTSecret)},
	}))

	// User routes (protected)
	router.Get("/users", userHandler.List)
	router.Post("/users", userHandler.Create)
	router.Get("/users/:id", userHandler.GetByID)
	router.Put("/users/:id", userHandler.Update)
	router.Delete("/users/:id", userHandler.Delete)

	// Add more route groups here as your app grows
	// Example:
	// products := app.Group("/products")
	// products.Post("/", productHandler.CreateProduct)
	// products.Get("/", productHandler.ListProducts)
	// etc.
}
