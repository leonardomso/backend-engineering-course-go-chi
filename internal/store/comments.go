package store

import (
	"context"
	"database/sql"
)

type CommentStore struct {
	db *sql.DB
}

func (s *CommentStore) Create(ctx context.Context) error {
	return nil
}

func (s *CommentStore) Update(ctx context.Context) error {
	return nil
}

func (s *CommentStore) Delete(ctx context.Context) error {
	return nil
}
