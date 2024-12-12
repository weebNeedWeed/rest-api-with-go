package user

import (
	"database/sql"
	"fmt"
	"go-rest-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	q := `SELECT ID, FirstName, LastName, Email, Password, CreatedAt FROM Users WHERE Email = ?`
	rows, err := s.db.Query(q, email)
	if err != nil {
		return nil, err
	}

	var u *types.User
	for rows.Next() {
		u, err = scanRowToUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u == nil {
		return nil, fmt.Errorf("no user found with given email %s", email)
	}

	return u, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (s *Store) CreateUser(user types.User) error {
	return nil
}

func scanRowToUser(rows *sql.Rows) (*types.User, error) {
	u := &types.User{}

	err := rows.Scan(&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.CreatedAt)
	if err != nil {
		return nil, err
	}

	return u, nil
}
