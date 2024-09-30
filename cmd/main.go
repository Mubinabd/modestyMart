package main

import (
	"github.com/Mubinabd/modestyMart/internal/app"
	"github.com/Mubinabd/modestyMart/internal/pkg/config"
	_ "github.com/lib/pq"
)

func main() {
	// Load configuration from environment variables or a file.
	cfg := config.Load()

	// Run the application with the loaded configuration.
	app.Run(&cfg)
}
