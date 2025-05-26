package service

import (
	"github.com/stretchr/testify/assert"
	"quotesAPI/internal/domain/quote/entity"
	"quotesAPI/internal/domain/quote/repository"
	"testing"
)

type mockRepo struct {
	quotes map[int]*entity.Quote
	nextID int
}

func newMockRepo() repository.QouteRepository {
	return &mockRepo{
		quotes: make(map[int]*entity.Quote),
		nextID: 1,
	}
}

func (m *mockRepo) Create(author string, text string) (*entity.Quote, error) {
	quote := &entity.Quote{
		Author: author,
		Quote:  text,
	}
	quote.ID = m.nextID
	m.quotes[quote.ID] = quote
	m.nextID++
	return quote, nil
}

func (m *mockRepo) GetRandom() (*entity.Quote, error) {
	for _, q := range m.quotes {
		return q, nil
	}
	return nil, ErrQuoteNotFound
}

func (m *mockRepo) GetAll() ([]*entity.Quote, error) {
	var quotes []*entity.Quote
	for _, q := range m.quotes {
		quotes = append(quotes, q)
	}
	return quotes, nil
}

func (m *mockRepo) GetByAuthor(author string) ([]*entity.Quote, error) {
	var quotes []*entity.Quote
	for _, q := range m.quotes {
		if q.Author == author {
			quotes = append(quotes, q)
		}
	}
	return quotes, nil
}

func (m *mockRepo) Delete(id int) error {
	if _, exists := m.quotes[id]; !exists {
		return ErrQuoteNotFound
	}
	delete(m.quotes, id)
	return nil
}

func TestQuoteService(t *testing.T) {
	repo := newMockRepo()
	s := NewQuoteService(repo)

	t.Run("CreateQuote", func(t *testing.T) {
		quote, err := s.CreateQuote("Author", "Text")
		assert.NoError(t, err)
		assert.Equal(t, "Author", quote.Author)
		assert.Equal(t, "Text", quote.Quote)
		assert.Equal(t, 1, quote.ID)
	})

	t.Run("CreateQuote_Validation", func(t *testing.T) {
		_, err := s.CreateQuote("", "Text")
		assert.Equal(t, ErrAuthorRequired, err)

		_, err = s.CreateQuote("Author", "")
		assert.Equal(t, ErrTextRequired, err)
	})

	t.Run("GetRandomQuote", func(t *testing.T) {
		_, err := s.GetRandomQuote()
		assert.NoError(t, err)
	})

	t.Run("GetAllQuotes", func(t *testing.T) {
		quotes, err := s.GetAllQuotes()
		assert.NoError(t, err)
		assert.Len(t, quotes, 1)
	})

	t.Run("GetQuotesByAuthor", func(t *testing.T) {
		quotes, err := s.GetQuotesByAuthor("Author")
		assert.NoError(t, err)
		assert.Len(t, quotes, 1)
	})

	t.Run("DeleteQuote", func(t *testing.T) {
		err := s.DeleteQuote(1)
		assert.NoError(t, err)
		err = s.DeleteQuote(999)
		assert.Equal(t, ErrQuoteNotFound, err)
	})
}
