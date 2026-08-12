package service

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/kaushiknaik9/improved-spoon/internal/model"
	"github.com/kaushiknaik9/improved-spoon/internal/repository"
)

type URLService struct {
	repo repository.URLRepository
}

func NewURLService(repo repository.URLRepository) *URLService {
	return &URLService{
		repo: repo,
	}
}

func (s *URLService) CreateShortURL(ctx context.Context, originalUrl string) (*model.URL, error) {
	if strings.TrimSpace(originalUrl) == "" {
		return nil, errors.New("url cant be empty")
	}
	parsedUrl, err := url.ParseRequestURI(originalUrl)
	if err != nil {
		return nil, errors.New("invalid url")
	}
	if parsedUrl.Host == "" || parsedUrl.Scheme == "" {
		return nil, errors.New("invalid url")
	}

	newUrl := &model.URL{
		ID:          time.Now().UnixNano(),
		ShortCode:   "temp-" + time.Now().Format("150405"),
		OriginalURL: originalUrl,
		CreatedAt:   time.Now(),
	}

	err = s.repo.Create(ctx, newUrl)
	if err != nil {
		return nil, err
	}
	return newUrl, nil
}

func (s *URLService) GetOriginalURL(ctx context.Context, shortCode string) (*model.URL, error) {
	if strings.TrimSpace(shortCode) == "" {
		return nil, errors.New("short code can't be empty")
	}
	return s.repo.GetByShortCode(ctx, shortCode)
}
