package repository

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/LXSCA7/go-url-shortener/internal/core/domain"
	"github.com/LXSCA7/go-url-shortener/internal/core/ports"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{client: client}
}

var _ ports.CacheRepository = (*RedisRepository)(nil)

func (r *RedisRepository) Get(ctx context.Context, key string) (domain.Link, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return domain.Link{}, errors.New("cache miss")
		}
		return domain.Link{}, err
	}

	var link domain.Link
	if err := json.Unmarshal([]byte(val), &link); err != nil {
		return domain.Link{}, err
	}

	return link, nil
}

func (r *RedisRepository) Set(ctx context.Context, key string, link domain.Link, ttl time.Duration) error {
	data, err := json.Marshal(link)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, data, ttl).Err()
}
