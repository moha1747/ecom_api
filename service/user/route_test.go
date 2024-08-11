package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/moha1747/ecom_api/types"
)

func TestHandler_RegisterRoutes(t *testing.T) {
	UserStore := &mockUserStore{}
	handler := NewHandler(UserStore)

	t.Run("Fail if user payload is invalid", func(t *testing.T) {
		// create a request
		payload := types.ResgisterPayload{
			FIrstName: "John",
			LastName: "Doe",
			Email: "123",
			Password: "password",
		}
		marshalled, _ := json.Marshal(payload)
		req, err  := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		// create a response
		res := httptest.NewRecorder()
		router := mux.NewRouter()
		
		// call the handler
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(res, req)
		// check the response
		if res.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadRequest, res.Code)
		}
	})

	t.Run("Should correctly register a user", func(t *testing.T)  {
				// create a request
		payload := types.ResgisterPayload{
			FIrstName: "John",
			LastName: "Doe",
			Email: "valid@mail.com",
			Password: "password",
		}
		marshalled, _ := json.Marshal(payload)
		req, err  := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		// create a response
		res := httptest.NewRecorder()
		router := mux.NewRouter()
		
		// call the handler
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(res, req)
		// check the response
		if res.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, res.Code)
		}
		
	})
}
// TODO: Implement the GetUserByEmail method
func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	// TODO: Implement the logic to retrieve a user by email from the mock store
	return nil, nil
}

// TODO: Implement the GetUserById method
func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	// TODO: Implement the logic to retrieve a user by ID from the mock store
	return nil, nil
}

// TODO: Implement the CreateUser method
func (m *mockUserStore) CreateUser(user types.User) error {
	// TODO: Implement the logic to create a user in the mock store
	return nil
}
type mockUserStore struct  {}
