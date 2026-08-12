package postgres

import (
	"context"
	"database/sql"

	"github.com/kaushiknaik9/improved-spoon/internal/model"
	"github.com/kaushiknaik9/improved-spoon/internal/repository"
)

type URLRepository struct {
	db *sql.DB
}

var _ repository.URLRepository = (*URLRepository)(nil)

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{
		db: db,
	}
}

func (r *URLRepository) Create(ctx context.Context, url *model.URL) error {
	query := `
	INSERT INTO urls (
	id,
	short_code,
	original_url,
	created_at
	)
	VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.ExecContext(ctx, query, url.ID, url.ShortCode, url.OriginalURL, url.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
func (r *URLRepository) GetByShortCode(ctx context.Context, shortCode string) (*model.URL, error) {
	query := `
	SELECT 
	id,
	short_code,
	original_url,
	created_at
	FROM urls
	WHERE short_url = $1
	`
	var url model.URL
	err := r.db.QueryRowContext(ctx, query, shortCode).Scan(&url.ID, &url.ShortCode, &url.OriginalURL, &url.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &url, nil
}
