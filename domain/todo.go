package domain

import (
	"cmd/main/dto"
	"context"
	"database/sql"
)

type Todo struct {
	ID        string       `db:"id"`
	Title     string       `db:"title"`
	Subtitle  string       `db:"subtitle"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

type TodoRepository interface {
	FindAll(ctx context.Context) ([]Todo, error)
	FindById(ctx context.Context, id string) (Todo, error)
	Save(ctx context.Context, t *Todo) error
	Update(ctx context.Context, t *Todo) error
	Delete(ctx context.Context, id string) error
}

type TodoService interface {
	Index(ctx context.Context) ([]dto.Todo, error)
	Create(ctx context.Context, req dto.CreateTodo) error
	Update(ctx context.Context, req dto.UpdateTodo) error
}
