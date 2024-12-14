package user

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"go-rest-api/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserServiceHandlers(t *testing.T) {
	mockStore := &mockUserStore{}
	handler := NewHandler(mockStore)

	t.Run("should fail if user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "user",
			LastName:  "123",
			Email:     "das",
			Password:  "123",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected error code %v, got %v", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should correctly register a user", func(t *testing.T) {
		mockStore := &mockUserStore{}
		handler := NewHandler(mockStore)

		payload := types.RegisterUserPayload{
			FirstName: "hello",
			LastName:  "world",
			Email:     "abc@xyz.com",
			Password:  "123",
		}
		marshalled, _ := json.Marshal(payload)

		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

type mockUserStore struct {
}

func (m *mockUserStore) GetUserByEmail(string) (*types.User, error) {
	return nil, nil
}
func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}
