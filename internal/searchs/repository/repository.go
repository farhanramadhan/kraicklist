package repository

import (
	"context"

	"challenge.haraj.com.sa/kraicklist/internal/searchs"
)

type SearchRepositoryInterface interface {
	Load(ctx context.Context, filepath string) error
	Search(ctx context.Context, query string) ([]searchs.Record, error)
}

type searchRepository struct {
	records []searchs.Record
}

func NewSearchRepository() *searchRepository {
	return &searchRepository{
		records: make([]searchs.Record, 0),
	}
}
