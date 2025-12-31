package handlers

import (
	"net/http"

	"github.com/LXSCA7/go-url-shortener/internal/core/ports"
	"github.com/LXSCA7/go-url-shortener/pkg/web"
)

type HTTPHandler struct {
	svc ports.ShortenerService
}

type CreateLinkRequest struct {
	OriginalURL string `json:"original_url"`
	ShortCode   string `json:"short_code,omitempty"`
}

func NewHTTPHandler(svc ports.ShortenerService) *HTTPHandler {
	return &HTTPHandler{
		svc: svc,
	}
}

func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body CreateLinkRequest
	if err := web.DecodeJSON(&body, r); err != nil {
		web.EncodeError(w, http.StatusBadRequest, err.Error())
		return
	}

	link, err := h.svc.Shorten(r.Context(), body.OriginalURL, body.ShortCode)
	if err != nil {
		web.EncodeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	web.EncodeJSON(w, http.StatusCreated, link)
}
