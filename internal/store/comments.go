package store

import (
	"context"
	"database/sql"
)

type CommentsStore struct {
	db *sql.DB
}

func (s *CommentsStore) Create(ctx context.Context) error {
	return nil
}

func (s *CommentsStore) Update(ctx context.Context) error {
	return nil
}

func (s *CommentsStore) Delete(ctx context.Context) error {
	return nil
}
