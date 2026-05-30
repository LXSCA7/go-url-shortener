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
	// RedisURL   string
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	serverPort := fmt.Sprintf(":%s", getVarWithFallback("HTTP_PORT", "8080"))

	dbUser := getVarWithFallback("DB_USER", "postgres")
	dbPass := getVarWithFallback("DB_PASSWORD", "postgres")
	dbHost := getVarWithFallback("DB_HOST", "localhost")
	dbPort := getVarWithFallback("DB_PORT", "5432")
	dbName := getVarWithFallback("DB_NAME", "shortener")

	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(dbUser, dbPass),
		Host:   fmt.Sprintf("%s:%s", dbHost, dbPort),
		Path:   dbName,
	}

	q := u.Query()
	q.Set("sslmode", "disable")
	// q.Set("pool_max_conns", "100")
	u.RawQuery = q.Encode()

	return &Config{
		ServerPort: serverPort,
		DBURL:      u.String(),
	}, nil
}

func getVarWithFallback(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback
	}
	return value
}
