package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/moha1747/ecom_api/service/auth"
	"github.com/moha1747/ecom_api/types"
	"github.com/moha1747/ecom_api/utils"
)
type Handler struct {
	store types.UserStore

}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")

}

func (h *Handler) handleLogin(response http.ResponseWriter, request *http.Request) {

}


func (h *Handler) handleRegister(response http.ResponseWriter, request *http.Request) {
	// get JSON payload
	var payload types.ResgisterPayload

	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		utils.WriteError(response, http.StatusBadRequest, err)
		return
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(response, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return 
	}
	// check if the user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(response, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(response, http.StatusInternalServerError, err)
		return
	}
	
		// if not, create the new user
		err = h.store.CreateUser(types.User{
			FIrstName: payload.FIrstName,
			LastName: payload.LastName,
			Email: payload.Email,
			Password: hashedPassword,
		})
		if err != nil {
			utils.WriteError(response, http.StatusInternalServerError, err)
			return
		}
		utils.WriteJSON(response, http.StatusCreated, nil)
}