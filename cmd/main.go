package main

import (
	"log"

	"github.com/yeboahd24/movie-reservation-system/config"
	"github.com/yeboahd24/movie-reservation-system/routes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := config.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := routes.SetupRouter(db)
	router.Run(":8080")
}
