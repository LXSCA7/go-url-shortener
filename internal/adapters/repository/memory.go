package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/LXSCA7/go-url-shortener/internal/core/domain"
	"github.com/LXSCA7/go-url-shortener/internal/core/ports"
)

type MemoryRepository struct {
	mu    sync.RWMutex
	links map[string]domain.Link
}

var _ ports.LinkRepository = (*MemoryRepository)(nil)

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		links: make(map[string]domain.Link),
	}
}

func (m *MemoryRepository) GetByCode(ctx context.Context, code string) (domain.Link, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	link, exists := m.links[code]
	if !exists {
		return domain.Link{}, errors.New("not found")
	}

	return link, nil
}

func (m *MemoryRepository) Save(ctx context.Context, link domain.Link) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.links[link.ShortCode]; exists {
		return errors.New("link already exists")
	}

	m.links[link.ShortCode] = link

	return nil
}
