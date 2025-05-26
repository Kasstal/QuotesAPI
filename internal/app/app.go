package app

import (
	"fmt"
	"quotesAPI/internal/domain/quote/repository"
	"quotesAPI/internal/domain/quote/service"
	"quotesAPI/storage"
)

type Application struct {
	QuoteService service.QuoteService
	Port         string
}

func NewApplication(port string) *Application {
	// Initialize storage
	storage := storage.NewQuotesStorage()

	// Initialize repository
	repo := repository.NewQuoteRepository(storage)

	// Initialize service
	quoteService := service.NewQuoteService(repo)

	return &Application{
		Port:         fmt.Sprintf(":%s", port),
		QuoteService: quoteService,
	}
}
