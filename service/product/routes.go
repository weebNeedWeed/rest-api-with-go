package product

import (
	"github.com/gorilla/mux"
	"go-rest-api/service/auth"
	"go-rest-api/types"
	"go-rest-api/utils"
	"net/http"
)

type Handler struct {
	store     types.ProductStore
	userStore types.UserStore
}

func NewHandler(store types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{store, userStore}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/products", h.handleGetProducts).Methods("GET")
	r.HandleFunc("/products", auth.WithJWTAuth(h.handleCreateProduct, h.userStore)).Methods("POST")
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}

func (*Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	// TODO: implement this later on
	utils.WriteJSON(w, http.StatusOK, "ok")
}
