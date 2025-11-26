package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64
	Username  string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

type UserStore struct {
	db *sql.DB
}

// Ideally we would use a ORM library to handle the SQL queries.
// However, I'm doing this manually because the author is doing this manually for learning purposes.
func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`

	err := s.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) Update(ctx context.Context, user *User) error {
	return nil
}

func (s *UserStore) Delete(ctx context.Context, id int64) error {
	return nil
}
