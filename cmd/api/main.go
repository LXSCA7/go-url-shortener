package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/LXSCA7/go-url-shortener/internal/adapters/handlers"
	"github.com/LXSCA7/go-url-shortener/internal/adapters/idgen"
	"github.com/LXSCA7/go-url-shortener/internal/adapters/repository"
	"github.com/LXSCA7/go-url-shortener/internal/config"
	"github.com/LXSCA7/go-url-shortener/internal/core/services"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func main() {
	node, err := idgen.NewSnowflakeIDGen(1)
	if err != nil {
		panic(err)
	}

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.DBURL)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("banco de dados não está respondendo: %v", err)
	}

	redisOpt, err := redis.ParseURL(cfg.CacheURL)
	if err != nil {
		log.Fatalf("redis nao respondendo: %v", err)
	}

	redisOpt.PoolSize = 100
	redisOpt.MinIdleConns = 10

	redisClient := redis.NewClient(redisOpt)

	pgRepo := repository.NewPostgresRepository(pool)
	redisRepo := repository.NewRedisRepository(redisClient)

	repo := repository.NewCachedLinkRepository(pgRepo, redisRepo)
	svc := services.NewShortenerService(node, repo)
	handler := handlers.NewHTTPHandler(svc)
	mux := handlers.NewRouter(handler)

	fmt.Printf("🚀 api running at %s\n\n", cfg.ServerPort)

	srv := &http.Server{
		Addr:         cfg.ServerPort,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	srv.ListenAndServe()
}
