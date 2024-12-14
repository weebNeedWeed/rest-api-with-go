package utils

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

var Validate = validator.New(validator.WithRequiredStructEnabled())

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return errors.New("no body found for the request")
	}

	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		return err
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if v != nil {
		err := json.NewEncoder(w).Encode(v)
		if err != nil {
			log.Printf("Error writing JSON response: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	WriteJSON(w, statusCode, map[string]string{
		"error": err.Error(),
	})
}
