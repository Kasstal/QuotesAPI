package http

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"quotesAPI/internal/domain/quote/entity"
	"quotesAPI/internal/domain/quote/service"
	"testing"
)

type mockService struct {
	service.QuoteService
}

func (m *mockService) CreateQuote(author, text string) (*entity.Quote, error) {
	return &entity.Quote{
		ID:     1,
		Author: author,
		Quote:  text,
	}, nil
}

func (m *mockService) GetRandomQuote() (*entity.Quote, error) {
	return &entity.Quote{
		ID:     1,
		Author: "Author",
		Quote:  "Text",
	}, nil
}

func (m *mockService) GetAllQuotes() ([]*entity.Quote, error) {
	return []*entity.Quote{
		{ID: 1, Author: "Author", Quote: "Text"},
	}, nil
}

func (m *mockService) GetQuotesByAuthor(author string) ([]*entity.Quote, error) {
	return []*entity.Quote{
		{ID: 1, Author: author, Quote: "Text"},
	}, nil
}

func (m *mockService) DeleteQuote(id int) error {
	return nil
}

func TestHandlers(t *testing.T) {
	router := mux.NewRouter()
	mockSvc := &mockService{}
	handler := NewQuoteHandler(mockSvc)

	router.HandleFunc("/quotes", handler.CreateQuote).Methods("POST")
	router.HandleFunc("/quotes", handler.GetQuotes).Methods("GET")
	router.HandleFunc("/quotes/random", handler.GetRandomQuote).Methods("GET")
	router.HandleFunc("/quotes/{id}", handler.DeleteQuote).Methods("DELETE")

	t.Run("CreateQuote", func(t *testing.T) {
		body := bytes.NewBufferString(`{"author":"Author","quote":"Text"}`)
		req := httptest.NewRequest("POST", "/quotes", body)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusCreated, res.Code)

		var quote entity.Quote
		err := json.NewDecoder(res.Body).Decode(&quote)
		assert.NoError(t, err)
		assert.Equal(t, "Author", quote.Author)
	})

	t.Run("GetQuotes", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/quotes", nil)
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)

		var quotes []entity.Quote
		err := json.NewDecoder(res.Body).Decode(&quotes)
		assert.NoError(t, err)
		assert.Len(t, quotes, 1)
	})

	t.Run("GetRandomQuote", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/quotes/random", nil)
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("DeleteQuote", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/quotes/1", nil)
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)
		assert.Equal(t, http.StatusNoContent, res.Code)
	})
}
