package main

import (
	"log"

	"github.com/Firuz43/order-service/internal/config"
	"github.com/Firuz43/order-service/internal/db"
)

func main() {
	// 1. Load configuration
	cfg := config.LoadConfig()

	// 2. Conenct to PostgreSQL
	dbpool, err := db.ConnectPostgres(cfg)

	if err != nil {
		log.Fatalf("Fatal: Could not initialize database %v", err)
	}

	// Clean up pool when main function exists
	defer dbpool.Close()

	log.Printf("Server starting on port %s", cfg.ServerPort)

	// Next: We will set up HTTP server router & route here
}
