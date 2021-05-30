package services

import (
	"context"

	"challenge.haraj.com.sa/kraicklist/internal/searchs"
	"challenge.haraj.com.sa/kraicklist/internal/searchs/repository"
)

type SearhServiceInterface interface {
	Load(ctx context.Context, filepath string) error
	Search(ctx context.Context, query string) ([]searchs.Record, error)
}

type searchService struct {
	searchRepo repository.SearchRepositoryInterface
}

func NewSearchService(searchRepo repository.SearchRepositoryInterface) *searchService {
	return &searchService{
		searchRepo: searchRepo,
	}
}

func (ss *searchService) Load(ctx context.Context, filepath string) error {
	err := ss.searchRepo.Load(ctx, filepath)

	if err != nil {
		return err
	}

	return nil
}

func (ss *searchService) Search(ctx context.Context, query string) ([]searchs.Record, error) {
	records, err := ss.searchRepo.Search(ctx, query)
	if err != nil {
		return nil, err
	}

	return records, nil
}
