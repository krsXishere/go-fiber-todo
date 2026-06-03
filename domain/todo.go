package domain

import (
	"context"
	"database/sql"
)

type Todo struct {
	ID        string       `db:"id"`
	Title     string       `db:"title"`
	Subtitle  string       `db:"subtitle"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"update_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type TodoRepository interface {
	FindAll(ctx context.Context) ([]Todo, error)
	FindById(ctx context.Context, id string) (Todo, error)
	Save(ctx context.Context, t *Todo) error
	Update(ctx context.Context, t *Todo) error
	Delete(ctx context.Context, id string) error
}
