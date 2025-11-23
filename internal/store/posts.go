package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type Post struct {
	id         int64
	title      string
	content    string
	user_id    string
	tags       []string
	created_at string
	updated_at string
}

type PostStore struct {
	db *sql.DB
}

// Ideally we would use a ORM library to handle the SQL queries.
// However, I'm doing this manually because the author is doing this manually for learning purposes.
func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `INSERT INTO posts (title, content, user_id, tags, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`

	err := s.db.QueryRowContext(ctx, query, post.title, post.content, post.user_id, pq.Array(post.tags), post.created_at, post.updated_at).Scan(&post.id, &post.created_at, &post.updated_at)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) Update(ctx context.Context, post *Post) error {
	return nil
}

func (s *PostStore) Delete(ctx context.Context, id int64) error {
	return nil
}
