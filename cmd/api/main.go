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
	mux := handlers.NewRouter(handler)

	port := ":8080"
	fmt.Printf("🚀 api running at %s\n\n", port)
	http.ListenAndServe(port, mux)
}
