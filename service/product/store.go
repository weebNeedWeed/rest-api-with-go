package product

import (
	"database/sql"
	"fmt"
	"go-rest-api/types"
	"strings"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) GetProducts() ([]*types.Product, error) {
	q := "SELECT id,name,description,image,price,quantity,createdAt FROM products"
	rows, err := s.db.Query(q)
	if err != nil {
		return nil, err
	}

	pList := make([]*types.Product, 0)
	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		pList = append(pList, p)
	}

	return pList, nil
}

func (s *Store) GetProductsByIDs(ids []int) ([]*types.Product, error) {
	prep := strings.Repeat(",?", len(ids)-1)
	q := fmt.Sprintf(`SELECT id,name,description,image,price,quantity,createdAt FROM products
			WHERE id IN(?,%s)`, prep)

	// converts []int into []any
	idsAsAnys := make([]any, len(ids))
	for i, id := range ids {
		idsAsAnys[i] = id
	}

	rows, err := s.db.Query(q, idsAsAnys...)
	if err != nil {
		return nil, err
	}

	prods := make([]*types.Product, 0)
	for rows.Next() {
		p, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		prods = append(prods, p)
	}

	return prods, nil
}

func (s *Store) UpdateProduct(p types.Product) {

}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	p := new(types.Product)

	err := rows.Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Image,
		&p.Price,
		&p.Quantity,
		&p.CreatedAt)
	if err != nil {
		return nil, err
	}

	return p, nil
}
