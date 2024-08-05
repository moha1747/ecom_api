package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moha1747/ecom_api/types"
	"github.com/moha1747/ecom_api/utils"
)
type Handler struct {
	store *types.UserStore

}

func NewHandler() *Handler {
	return &Handler{}
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

	if err := json.NewDecoder(request.Body).Decode(payload); err != nil {
		utils.WriteError(response, http.StatusBadRequest, err)
	}
	// check if the user exists

		// if not, create the new user
}