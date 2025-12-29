package web

import (
	"encoding/json"
	"net/http"
)

func EncodeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		return json.NewEncoder(w).Encode(data)
	}

	return nil
}

type ErrorPayload struct {
	Error string `json:"error"`
}

func EncodeError(w http.ResponseWriter, status int, message string) error {
	return EncodeJSON(w, status, ErrorPayload{Error: message})
}
