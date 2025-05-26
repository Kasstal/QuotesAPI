package service

import (
	"errors"
	"quotesAPI/internal/domain/quote/entity"
	"quotesAPI/internal/domain/quote/repository"
)

var (
	ErrQuoteNotFound  = errors.New("quote not found")
	ErrInvalidQuote   = errors.New("invalid quote data")
	ErrAuthorRequired = errors.New("author is required")
	ErrTextRequired   = errors.New("quote text is required")
	ErrTextTooLong    = errors.New("quote text is too long")
)

type QuoteService interface {
	CreateQuote(author, text string) (*entity.Quote, error)
	GetRandomQuote() (*entity.Quote, error)
	GetAllQuotes() ([]*entity.Quote, error)
	GetQuotesByAuthor(author string) ([]*entity.Quote, error)
	DeleteQuote(id int) error
}

type quoteService struct {
	repo repository.QouteRepository
}

func NewQuoteService(repo repository.QouteRepository) QuoteService {
	return &quoteService{repo: repo}
}

func (s *quoteService) CreateQuote(author, text string) (*entity.Quote, error) {
	// Input validation
	if author == "" {
		return nil, ErrAuthorRequired
	}
	if text == "" {
		return nil, ErrTextRequired
	}
	if len(text) > 500 {
		return nil, ErrTextTooLong
	}

	quote := entity.Quote{
		Author: author,
		Quote:  text,
	}

	return s.repo.Create(quote)
}

func (s *quoteService) GetRandomQuote() (*entity.Quote, error) {
	quote, err := s.repo.GetRandom()
	if err != nil {
		return nil, ErrQuoteNotFound
	}
	return quote, nil
}

func (s *quoteService) GetAllQuotes() ([]*entity.Quote, error) {
	quotes, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	if len(quotes) == 0 {
		return []*entity.Quote{}, nil
	}
	return quotes, nil
}

func (s *quoteService) GetQuotesByAuthor(author string) ([]*entity.Quote, error) {
	if author == "" {
		return nil, ErrAuthorRequired
	}

	quotes, err := s.repo.GetByAuthor(author)
	if err != nil {
		return nil, err
	}
	if len(quotes) == 0 {
		return []*entity.Quote{}, nil
	}
	return quotes, nil
}

func (s *quoteService) DeleteQuote(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return ErrQuoteNotFound
	}
	return nil
}
