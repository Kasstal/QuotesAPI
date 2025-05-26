package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"quotesAPI/internal/domain/quote/entity"
	"quotesAPI/internal/domain/quote/service"
	"strconv"
)

type QuoteHandler struct {
	service service.QuoteService
}

func NewQuoteHandler(service service.QuoteService) *QuoteHandler {
	return &QuoteHandler{service: service}
}

type createQuoteRequest struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

func (h *QuoteHandler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var req createQuoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	quote, err := h.service.CreateQuote(req.Author, req.Quote)
	if err != nil {
		switch err {
		case service.ErrAuthorRequired, service.ErrTextRequired, service.ErrTextTooLong:
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "Failed to create quote", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(quote)
}

func (h *QuoteHandler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")

	var quotes []*entity.Quote
	var err error

	if author != "" {
		quotes, err = h.service.GetQuotesByAuthor(author)
	} else {
		quotes, err = h.service.GetAllQuotes()
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "Failed to get quotes", http.StatusInternalServerError)
		return
	}

	if quotes == nil {
		quotes = []*entity.Quote{}
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

func (h *QuoteHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := h.service.GetRandomQuote()
	if err != nil {
		if err == service.ErrQuoteNotFound {
			w.WriteHeader(http.StatusNotFound)
			http.Error(w, "No quotes available", http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "Failed to get random quote", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quote)
}

func (h *QuoteHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {

		http.Error(w, "Invalid quote ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteQuote(id)
	if err != nil {
		if err == service.ErrQuoteNotFound {

			http.Error(w, "Quote not found", http.StatusNotFound)
		} else {

			http.Error(w, "Failed to delete quote", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
