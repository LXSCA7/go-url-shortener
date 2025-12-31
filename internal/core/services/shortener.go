package services

import (
	"context"
	"time"

	"github.com/LXSCA7/go-url-shortener/internal/core/domain"
	"github.com/LXSCA7/go-url-shortener/internal/core/ports"
	"github.com/LXSCA7/go-url-shortener/pkg/base62"
)

type ShortenerService struct {
	idGen ports.IDGenerator
	repo  ports.LinkRepository
}

func NewShortenerService(idGen ports.IDGenerator, repo ports.LinkRepository) *ShortenerService {
	return &ShortenerService{
		idGen: idGen,
		repo:  repo,
	}
}

var _ ports.ShortenerService = (*ShortenerService)(nil)

func (s *ShortenerService) Shorten(ctx context.Context, originalURL string, alias string) (domain.Link, error) {
	id := s.idGen.Generate()
	code := alias
	if code == "" {
		code = base62.Encode(id)
	}

	l := domain.Link{
		Id:          id,
		ShortCode:   code,
		OriginalURL: originalURL,
		CreatedAt:   time.Now(),
		Visits:      0,
	}

	if err := s.repo.Save(ctx, l); err != nil {
		return domain.Link{}, err
	}

	return l, nil
}

func (s *ShortenerService) GetOriginalURL(ctx context.Context, code string) (string, error) {
	l, err := s.repo.GetByCode(ctx, code)
	if err != nil {
		return "", err
	}

	return l.OriginalURL, nil
}
