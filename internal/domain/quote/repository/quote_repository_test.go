package repository

import (
	"github.com/stretchr/testify/assert"
	"quotesAPI/internal/domain/quote/entity"
	"quotesAPI/storage"
	"testing"
)

func TestQuoteRepository(t *testing.T) {
	store := storage.NewQuotesStorage()
	repo := NewQuoteRepository(store)

	testQuote := entity.Quote{
		Author: "Test Author",
		Quote:  "Test Quote",
	}

	t.Run("Create", func(t *testing.T) {
		created, err := repo.Create(testQuote.Author, testQuote.Quote)
		assert.NoError(t, err)
		assert.Equal(t, testQuote.Author, created.Author)
		assert.Equal(t, testQuote.Quote, created.Quote)
		assert.NotZero(t, created.ID)
	})

	t.Run("GetRandom", func(t *testing.T) {
		// First create a quote to ensure there's one available
		_, err := repo.Create(testQuote.Author, testQuote.Quote)
		assert.NoError(t, err)

		quote, err := repo.GetRandom()
		assert.NoError(t, err)
		assert.NotNil(t, quote)
	})

	t.Run("GetAll", func(t *testing.T) {
		// Clear storage first
		store := storage.NewQuotesStorage()
		repo := NewQuoteRepository(store)

		// Create test data
		_, err := repo.Create(testQuote.Author, testQuote.Quote)
		assert.NoError(t, err)

		quotes, err := repo.GetAll()
		assert.NoError(t, err)
		assert.Len(t, quotes, 1)
		assert.Equal(t, testQuote.Author, quotes[0].Author)
	})

	t.Run("GetByAuthor", func(t *testing.T) {
		// Clear storage first
		store := storage.NewQuotesStorage()
		repo := NewQuoteRepository(store)

		// Create test data
		_, err := repo.Create(testQuote.Author, testQuote.Quote)
		assert.NoError(t, err)

		quotes, err := repo.GetByAuthor(testQuote.Author)
		assert.NoError(t, err)
		assert.Len(t, quotes, 1)
		assert.Equal(t, testQuote.Author, quotes[0].Author)
	})

	t.Run("Delete", func(t *testing.T) {
		// Clear storage first
		store := storage.NewQuotesStorage()
		repo := NewQuoteRepository(store)

		// Create test data
		created, err := repo.Create(testQuote.Author, testQuote.Quote)
		assert.NoError(t, err)

		// Test successful delete
		err = repo.Delete(created.ID)
		assert.NoError(t, err)

		// Test not found
		err = repo.Delete(999)
		assert.Equal(t, storage.ErrQuoteNotFound, err)
	})

}
