package ports

import (
	"context"
	"time"

	"github.com/LXSCA7/go-url-shortener/internal/core/domain"
)

type LinkRepository interface {
	Save(ctx context.Context, link domain.Link) error
	GetByCode(ctx context.Context, code string) (domain.Link, error)
}

type CacheRepository interface {
	Get(ctx context.Context, key string) (domain.Link, error)
	Set(ctx context.Context, key string, link domain.Link, ttl time.Duration) error
}
