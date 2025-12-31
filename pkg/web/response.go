package web

import (
	"encoding/json"
	"net/http"
)

// object needs to be a pointer
func DecodeJSON(object any, r *http.Request) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(object)
}

func EncodeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		return json.NewEncoder(w).Encode(data)
	}

	return nil
}

func Redirect(w http.ResponseWriter, r *http.Request, originalURL string) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
}

type ErrorPayload struct {
	Error string `json:"error"`
}

func EncodeError(w http.ResponseWriter, status int, message string) error {
	return EncodeJSON(w, status, ErrorPayload{Error: message})
}
