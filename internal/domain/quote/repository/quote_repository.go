package repository

import (
	"math/rand"
	"quotesAPI/internal/domain/quote/entity"
	"quotesAPI/storage"
)

type QouteRepository interface {
	Create(quote entity.Quote) (*entity.Quote, error)
	GetRandom() (*entity.Quote, error)
	GetAll() ([]*entity.Quote, error)
	GetByAuthor(string) ([]*entity.Quote, error)
	Delete(id int) error
}

type quoteRepository struct {
	storage *storage.QuotesStorage
}

func NewQuoteRepository(storage *storage.QuotesStorage) QouteRepository {
	return &quoteRepository{storage: storage}
}

func (q *quoteRepository) Create(quote entity.Quote) (*entity.Quote, error) {
	createQuoteParams := storage.CreateQuoteParams{
		Author: quote.Author,
		Quote:  quote.Quote,
	}

	quoteModel, err := q.storage.Create(createQuoteParams)
	if err != nil {
		return nil, err
	}

	return &entity.Quote{
		ID:     quoteModel.ID,
		Author: quoteModel.Author,
		Quote:  quoteModel.Quote,
	}, err
}

func (q *quoteRepository) GetRandom() (*entity.Quote, error) {
	l, err := q.storage.Len()
	if err != nil {
		return nil, err
	}

	randID := rand.Intn(l)
	quoteModel, err := q.storage.GetByID(randID)

	if err != nil {
		return nil, err
	}

	return &entity.Quote{
		ID:     quoteModel.ID,
		Author: quoteModel.Author,
		Quote:  quoteModel.Quote,
	}, err
}

func (q *quoteRepository) GetAll() ([]*entity.Quote, error) {
	qouteModels, err := q.storage.GetAll()
	if err != nil {
		return nil, err
	}
	quotes := make([]*entity.Quote, len(qouteModels))

	for _, quoteModel := range qouteModels {
		quotes = append(quotes, &entity.Quote{
			ID:     quoteModel.ID,
			Author: quoteModel.Author,
			Quote:  quoteModel.Quote,
		})
	}

	return quotes, nil
}

func (q *quoteRepository) GetByAuthor(author string) ([]*entity.Quote, error) {
	qouteModels, err := q.storage.GetByAuthor(author)
	if err != nil {
		return nil, err
	}
	quotes := make([]*entity.Quote, len(qouteModels))

	for _, quoteModel := range qouteModels {
		quotes = append(quotes, &entity.Quote{
			ID:     quoteModel.ID,
			Author: quoteModel.Author,
			Quote:  quoteModel.Quote,
		})
	}

	return quotes, nil
}

func (q *quoteRepository) Delete(id int) error {
	return q.storage.Delete(id)
}
