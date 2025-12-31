package ports

import (
	"context"

	"github.com/LXSCA7/go-url-shortener/internal/core/domain"
)

type LinkRepository interface {
	Save(ctx context.Context, link domain.Link) error
	GetByCode(ctx context.Context, code string) (domain.Link, error)
}
