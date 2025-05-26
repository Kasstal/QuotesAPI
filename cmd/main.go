package main

import (
	"log"
	"os"
	"quotesAPI/internal/app"
)

func main() {
	app := app.NewApplication()

	if err := app.Run(); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}
