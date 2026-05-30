package repository

import (
	"context"
	"time"

	"github.com/LXSCA7/go-url-shortener/internal/core/domain"
	"github.com/LXSCA7/go-url-shortener/internal/core/ports"
)

type CachedLinkRepository struct {
	dbRepo    ports.LinkRepository
	cacheRepo ports.CacheRepository
}

func NewCachedLinkRepository(dbRepo ports.LinkRepository, cacheRepo ports.CacheRepository) *CachedLinkRepository {
	return &CachedLinkRepository{dbRepo: dbRepo, cacheRepo: cacheRepo}
}

var _ ports.LinkRepository = (*CachedLinkRepository)(nil)

func (c *CachedLinkRepository) GetByCode(ctx context.Context, code string) (domain.Link, error) {
	link, err := c.cacheRepo.Get(ctx, code)
	if err == nil {
		return link, nil
	}

	link, err = c.dbRepo.GetByCode(ctx, code)
	if err != nil {
		return domain.Link{}, err
	}

	go func(l domain.Link) {
		bgCtx := context.Background()
		_ = c.cacheRepo.Set(bgCtx, code, l, 24*time.Hour)
	}(link)

	return link, nil
}

func (c *CachedLinkRepository) Save(ctx context.Context, link domain.Link) error {
	if err := c.dbRepo.Save(ctx, link); err != nil {
		return err
	}

	go func(l domain.Link) {
		bgCtx := context.Background()
		_ = c.cacheRepo.Set(bgCtx, l.ShortCode, l, 24*time.Hour)
	}(link)

	return nil
}
