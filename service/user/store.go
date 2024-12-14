package user

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

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	q := `SELECT id, firstName, lastName, email, password, createdAt FROM users WHERE Email = ?`
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

	return u, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	q := `SELECT id, firstName, lastName, email, password, createdAt FROM users WHERE id = ?`
	rows, err := s.db.Query(q, id)
	if err != nil {
		return nil, err
	}

	var u *types.User
	for rows.Next() {
		if u, err = scanRowToUser(rows); err != nil {
			return nil, err
		}
	}

	return u, nil
}
func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec(`INSERT INTO users (firstName, lastName, email, password)
		VALUES (?,?,?,?)`, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}

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
