include .env

DB_URL=postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/${DB_NAME}?sslmode=disable

run:
	go run cmd/api/main.go

migrate-up:
	goose -dir ./internal/adapters/repository/migrations postgres "$(DB_URL)" up

migrate-down:
	goose -dir ./internal/adapters/repository/migrations postgres "$(DB_URL)" down