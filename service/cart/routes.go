package cart

import (
	"fmt"
	"go-rest-api/types"
	"go-rest-api/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{store, productStore}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/cart/checkout", h.handleCheckout).Methods("POST")
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	userID := 0
	var payload types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload"))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errs))
			return
		}
		return
	}

	productIDs, err := getCartItemIDs(payload.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	ps, err := h.productStore.GetProductsByIDs(productIDs)

	orderID, totalPrice, err := h.createOrder(ps, payload.Items, userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"totalPrice": totalPrice,
		"orderID":    orderID,
	})
}
