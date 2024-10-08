package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/moha1747/ecom_api/service/auth"
	"github.com/moha1747/ecom_api/types"
	"github.com/moha1747/ecom_api/utils"
	"github.com/moha1747/ecom_api/config"

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
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
    var user = types.LoginUserPayload{}
    if err := utils.ParseJSON(r, &user); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }

    // Validate the payload
    if err := utils.Validate.Struct(user); err != nil {
        errors := err.(validator.ValidationErrors)
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
        return
    }

    // Check if user exists
    u, err := h.store.GetUserByEmail(user.Email)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email or password"))
        return
    }

    // Validate password
    if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
        return
    }

    // Generate JWT token
    secret := []byte(config.Envs.JWTSecret)
    token, err := auth.CreateJWT(secret, u.ID)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }

    // Return the token
    utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var user = types.ResgisterPayload{}
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// check if user exists
	_, err := h.store.GetUserByEmail(user.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Email))
		return
	}

	// hash password
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateUser(types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}