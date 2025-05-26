package http

import (
	"github.com/gorilla/mux"
	"quotesAPI/internal/domain/quote/service"
)

func SetupRoutes(router *mux.Router, quoteService service.QuoteService) {
	handler := NewQuoteHandler(quoteService)
	router.Use(LoggingMiddleware)
	router.HandleFunc("/quotes", handler.CreateQuote).Methods("POST")
	router.HandleFunc("/quotes", handler.GetQuotes).Methods("GET")
	router.HandleFunc("/quotes/random", handler.GetRandomQuote).Methods("GET")
	router.HandleFunc("/quotes/{id}", handler.DeleteQuote).Methods("DELETE")
}
