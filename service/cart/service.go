package cart

import (
	"fmt"
	"go-rest-api/types"
)

func getCartItemIDs(cartItems []types.CartItem) ([]int, error) {
	productIDs := make([]int, len(cartItems))

	for i, item := range cartItems {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for product with id %v", item.ProductID)
		}
		productIDs[i] = item.ProductID
	}

	return productIDs, nil
}

func (h *Handler) createOrder(ps []*types.Product, items []types.CartItem, userID int) (int, float64, error) {
	productMap := make(map[int]*types.Product)
	for _, p := range ps {
		productMap[p.ID] = p
	}

	// check if products are in stock
	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, err
	}

	// calculate the total price
	totalPrice := calculateTotalPrice(items, productMap)

	// this should be refactored into a join SQL statement
	// reduce the quantity of products in the database
	for _, i := range items {
		product := productMap[i.ProductID]
		product.Quantity -= i.Quantity // because of this

		h.productStore.UpdateProduct(*product)
	}

	// create order
	orderID, err := h.store.CreateOrder(types.Order{
		UserID:  userID,
		Total:   totalPrice,
		Status:  "test",
		Address: "test",
	})
	if err != nil {
		return 0, 0, err
	}

	// create order item
	for _, i := range items {
		err := h.store.CreateOrderItem(types.OrderItem{
			OrderID:   orderID,
			ProductID: i.ProductID,
			Quantity:  i.Quantity,
			Price:     productMap[i.ProductID].Price,
		})
		if err != nil {
			return 0, 0, err
		}
	}

	return orderID, totalPrice, nil
}

func calculateTotalPrice(items []types.CartItem, productMap map[int]*types.Product) float64 {
	var totalPrice float64 = 0

	for _, i := range items {
		p := float64(i.Quantity) * productMap[i.ProductID].Price
		totalPrice += p
	}

	return totalPrice
}

func checkIfCartIsInStock(items []types.CartItem, productMap map[int]*types.Product) error {
	if len(items) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range items {
		p, ok := productMap[item.ProductID]
		if !ok {
			return fmt.Errorf("product with id %d is invalid", item.ProductID)
		}

		if p.Quantity < item.Quantity {
			return fmt.Errorf("quantity of product with id %v is invalid", item.ProductID)
		}
	}

	return nil
}
