package store

import (
	"context"
	"database/sql"
)

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context) error {
	return nil
}

func (s *UsersStore) Update(ctx context.Context) error {
	return nil
}

func (s *UsersStore) Delete(ctx context.Context) error {
	return nil
}
