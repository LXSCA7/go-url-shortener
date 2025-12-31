package ports

import (
	"context"

	"github.com/LXSCA7/go-url-shortener/internal/core/domain"
)

type ShortenerService interface {
	Shorten(ctx context.Context, originalURL, alias string) (domain.Link, error)
	GetOriginalURL(ctx context.Context, code string) (string, error)
}
