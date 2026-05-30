package config

import (
	"fmt"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL      string
	ServerPort string
	CacheURL   string
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	serverPort := fmt.Sprintf(":%s", getVarWithFallback("HTTP_PORT", "8080"))

	dbUser := getVarWithFallback("DB_USER", "postgres")
	dbPass := getVarWithFallback("DB_PASSWORD", "postgres")
	dbHost := getVarWithFallback("DB_HOST", "localhost")
	dbPort := getVarWithFallback("DB_PORT", "5432")
	dbName := getVarWithFallback("DB_NAME", "shortener")

	pgURL := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(dbUser, dbPass),
		Host:   fmt.Sprintf("%s:%s", dbHost, dbPort),
		Path:   dbName,
	}

	pgQuery := pgURL.Query()
	pgQuery.Set("sslmode", "disable")
	// pgQuery.Set("pool_max_conns", "100")
	pgURL.RawQuery = pgQuery.Encode()

	redisHost := getVarWithFallback("REDIS_HOST", "localhost")
	redisPort := getVarWithFallback("REDIS_PORT", "6379")
	redisPass := getVarWithFallback("REDIS_PASSWORD", "")
	redisDB := getVarWithFallback("REDIS_DB", "0")

	redisURL := url.URL{
		Scheme: "redis",
		Host:   fmt.Sprintf("%s:%s", redisHost, redisPort),
		Path:   "/" + redisDB,
	}

	if redisPass != "" {
		redisURL.User = url.UserPassword("", redisPass)
	}

	return &Config{
		ServerPort: serverPort,
		DBURL:      pgURL.String(),
		CacheURL:   redisURL.String(),
	}, nil
}

func getVarWithFallback(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback
	}
	return value
}
