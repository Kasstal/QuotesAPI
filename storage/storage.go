package storage

import (
	"errors"
	"quotesAPI/storage/models"
	"sync"
)

/*
	type Storage[T any] interface {
		Create(T) (T, error)
		Get(int) (T, error)
		List() ([]T, error)
		Delete(int) error
	}
*/

var ErrQuoteNotFound = errors.New("quote not found")

type QuotesStorage struct {
	mu     sync.RWMutex
	quotes map[int]models.Quote
	nextID int
}

func NewQuotesStorage() *QuotesStorage {
	return &QuotesStorage{nextID: 1, quotes: map[int]models.Quote{}}
}

type CreateQuoteParams struct {
	Author string
	Quote  string
}

func (q *QuotesStorage) Len() (int, error) {
	l := len(q.quotes)
	if l == 0 {
		return 0, ErrQuoteNotFound
	}

	return l, nil
}

func (q *QuotesStorage) Create(params CreateQuoteParams) (*models.Quote, error) {
	quote := models.Quote{
		Author: params.Author,
		Quote:  params.Quote,
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	quote.ID = q.nextID
	q.quotes[q.nextID] = quote
	q.nextID++

	return &quote, nil
}

func (q *QuotesStorage) GetAll() ([]models.Quote, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	quotes := make([]models.Quote, 0, len(q.quotes))
	for _, qoute := range q.quotes {
		quotes = append(quotes, qoute)
	}

	return quotes, nil
}

func (q *QuotesStorage) GetByID(id int) (models.Quote, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	quote, ok := q.quotes[id]

	if !ok {
		return models.Quote{}, ErrQuoteNotFound
	}
	return quote, nil
}

func (q *QuotesStorage) GetByAuthor(author string) ([]models.Quote, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()

	var quotes []models.Quote
	for _, qoute := range q.quotes {
		if qoute.Author == author {
			quotes = append(quotes, qoute)
		}
	}

	return quotes, nil
}

func (q *QuotesStorage) Delete(id int) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	_, ok := q.quotes[id]
	if !ok {
		return ErrQuoteNotFound
	}
	delete(q.quotes, id)
	return nil
}
