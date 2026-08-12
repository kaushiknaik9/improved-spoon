package repository

import (
	"context"

	"github.com/kaushiknaik9/improved-spoon/internal/model"
)

type URLRepository interface {
	Create(ctx context.Context, url *model.URL) error
	GetByShortCode(ctx context.Context, shortCode string) (*model.URL, error)
}
