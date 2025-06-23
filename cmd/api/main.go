package main

import (
	"log"

	"github.com/abdansya/go-hexagonal/internal/bootstrap"
	"github.com/abdansya/go-hexagonal/pkg/config"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Bootstrap application
	app := bootstrap.App(cfg)

	// Start server
	port := cfg.AppPort
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
