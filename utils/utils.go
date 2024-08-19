package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	validate "github.com/go-playground/validator/v10"
)

var Validate = validate.New()

func ParseJSON(request *http.Request, payload any) error { 
	if request.Body == nil {
	return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(request.Body).Decode(payload)
}

func WriteJSON(response http.ResponseWriter, status int, payload any) error {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)

	return json.NewEncoder(response).Encode(payload)
}

// payload will be of type error,  consistent response error for frontend {
func WriteError(response http.ResponseWriter, status int, err error) {
	WriteJSON(response, status, map[string]string{"error": err.Error()})
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")
	
	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}