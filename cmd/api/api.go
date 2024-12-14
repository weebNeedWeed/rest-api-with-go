package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"go-rest-api/service/cart"
	"go-rest-api/service/order"
	"go-rest-api/service/product"
	"go-rest-api/service/user"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr, db}
}

func (a *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(a.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(a.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(a.db)
	cartHandler := cart.NewHandler(orderStore, productStore)
	cartHandler.RegisterRoutes(subrouter)

	log.Println("server is listening on", a.addr)

	return http.ListenAndServe(a.addr, router)
}
