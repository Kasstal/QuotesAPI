package main

import (
	"log"
	"os"
	"quotesAPI/internal/app"
	"quotesAPI/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	app := app.NewApplication(cfg.Port)

	if err := app.Run(); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
