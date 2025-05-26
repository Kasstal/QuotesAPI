package app

import (
	"quotesAPI/internal/domain/quote/repository"
	"quotesAPI/internal/domain/quote/service"
	"quotesAPI/storage"
)

type Application struct {
	QuoteService service.QuoteService
}

func NewApplication() *Application {
	// Initialize storage
	storage := storage.NewQuotesStorage()

	// Initialize repository
	repo := repository.NewQuoteRepository(storage)

	// Initialize service
	quoteService := service.NewQuoteService(repo)

	return &Application{
		QuoteService: quoteService,
	}
}
