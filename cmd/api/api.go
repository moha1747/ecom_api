package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moha1747/ecom_api/service/product"
	"github.com/moha1747/ecom_api/service/user"
)

type ApiServer struct {
	addr string
	db *sql.DB

}

func NewApiServer (addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db: db,
	}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	productStore := product.NewStore(s.db)
	productHandler :=product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subRouter)

	log.Println("Listening on ", s.addr)
	return http.ListenAndServe(s.addr, router)
}