package order

import (
	"database/sql"
	"go-rest-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) CreateOrder(o types.Order) (int, error) {
	return 0, nil
}

func (s *Store) CreateOrderItem(oi types.OrderItem) error {
	return nil
}
