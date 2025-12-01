package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("Resource not found")
)

type Storage struct {
	Posts    Posts
	Users    Users
	Comments Comments
}

type Posts interface {
	GetByID(context.Context, int64) (*Post, error)
	Create(context.Context, *Post) error
	Update(context.Context, *Post) error
	Delete(context.Context, int64) error
}

type Users interface {
	Create(context.Context, *User) error
	Update(context.Context, *User) error
	Delete(context.Context, int64) error
}

type Comments interface {
	Create(context.Context) error
	Delete(context.Context) error
	Update(context.Context) error
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db: db},
		Users:    &UserStore{db: db},
		Comments: &CommentStore{db: db},
	}
}
