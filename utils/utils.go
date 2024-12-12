package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return errors.New("no body found for the request")
	}

	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		return err
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	_ = WriteJSON(w, statusCode, map[string]string{
		"error": err.Error(),
	})
}
