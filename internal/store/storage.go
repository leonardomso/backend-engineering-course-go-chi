package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts    Posts
	Users    Users
	Comments Comments
}

type Posts interface {
	Create(context.Context) error
	Update(context.Context) error
	Delete(context.Context) error
}

type Users interface {
	Create(context.Context) error
	Update(context.Context) error
	Delete(context.Context) error
}

type Comments interface {
	Create(context.Context) error
	Delete(context.Context) error
	Update(context.Context) error
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostsStore{db: db},
		Users:    &UsersStore{db: db},
		Comments: &CommentsStore{db: db},
	}
}
