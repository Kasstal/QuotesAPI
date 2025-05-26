package repository

import (
	"log"
	"math/rand"
	"quotesAPI/internal/domain/quote/entity"
	"quotesAPI/storage"
)

type QouteRepository interface {
	Create(author string, text string) (*entity.Quote, error)
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

func (r *quoteRepository) Create(author string, text string) (*entity.Quote, error) {

	createParams := storage.CreateQuoteParams{
		Author: author,
		Quote:  text,
	}

	quoteModel, err := r.storage.Create(createParams)
	log.Println(quoteModel)
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

	randID := rand.Intn(l) + 1
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

func (r *quoteRepository) GetAll() ([]*entity.Quote, error) {
	qouteModels, err := r.storage.GetAll()
	if err != nil {
		return nil, err
	}
	quotes := make([]*entity.Quote, 0, len(qouteModels))

	for _, quoteModel := range qouteModels {
		quotes = append(quotes, &entity.Quote{
			ID:     quoteModel.ID,
			Author: quoteModel.Author,
			Quote:  quoteModel.Quote,
		})
	}

	return quotes, nil
}

func (r *quoteRepository) GetByAuthor(author string) ([]*entity.Quote, error) {
	qouteModels, err := r.storage.GetByAuthor(author)
	if err != nil {
		return nil, err
	}
	quotes := make([]*entity.Quote, 0, len(qouteModels))

	for _, quoteModel := range qouteModels {
		quotes = append(quotes, &entity.Quote{
			ID:     quoteModel.ID,
			Author: quoteModel.Author,
			Quote:  quoteModel.Quote,
		})
	}

	return quotes, nil
}

func (r *quoteRepository) Delete(id int) error {
	return r.storage.Delete(id)
}
