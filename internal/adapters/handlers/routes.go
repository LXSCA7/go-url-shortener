package handlers

import "net/http"

func NewRouter(h *HTTPHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong\n"))
	})

	mux.HandleFunc("GET /{code}", h.Get)
	mux.HandleFunc("POST /api", h.Create)

	return mux
}
