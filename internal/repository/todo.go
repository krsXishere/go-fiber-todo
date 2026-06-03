package repository

import (
	"cmd/main/domain"
	"context"
)

type todoRepository struct{
	db *goqu.Database
}

func NewTodo(conn *sqlDB) domain.TodoRepository {
	return &todoRepository{
		db: goqu.New("default", conn)
	}
}

func (tr todoRepository) FindAll(ctx context.Context) ([]domain.Todo, error) {
	panic("")
}

func (tr todoRepository) FindById(ctx context.Context, id string) (domain.Todo, error) {
	panic("")
}

func (tr todoRepository) Save(ctx context.Context, t *domain.Todo) error {
	panic("")
}

func (tr todoRepository) Update(ctx context.Context, t *domain.Todo) error {
	panic("")
}

func (tr todoRepository) Delete(ctx context.Context, id string) error {
	panic("")
}
