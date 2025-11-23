package store

import (
	"context"
	"database/sql"
)

type User struct {
	id         string
	username   string
	email      string
	password   string
	created_at string
}

type UserStore struct {
	db *sql.DB
}

// Ideally we would use a ORM library to handle the SQL queries.
// However, I'm doing this manually because the author is doing this manually for learning purposes.
func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO users (id, username, email, password, created_at) VALUES ($1, $2, $3, $4, $5)`

	err := s.db.QueryRowContext(ctx, query, user.id, user.username, user.email, user.password).Scan(&user.id, &user.created_at)

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
