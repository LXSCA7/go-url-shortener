package main

import (
	"fmt"
	"net/http"

	"github.com/LXSCA7/go-url-shortener/internal/adapters/handlers"
	"github.com/LXSCA7/go-url-shortener/internal/adapters/idgen"
	"github.com/LXSCA7/go-url-shortener/internal/adapters/repository"
	"github.com/LXSCA7/go-url-shortener/internal/core/services"
)

func main() {
	node, err := idgen.NewSnowflakeIDGen(1)
	if err != nil {
		panic(err)
	}

	repo := repository.NewMemoryRepository()
	svc := services.NewShortenerService(node, repo)
	handler := handlers.NewHTTPHandler(svc)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong\n"))
	})

	mux.HandleFunc("GET /{code}", handler.Get)
	mux.HandleFunc("POST /api", handler.Create)

	port := ":8080"
	fmt.Printf("🚀 api running at %s\n\n", port)
	http.ListenAndServe(port, mux)
}
