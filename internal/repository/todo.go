package repository

import (
	"cmd/main/domain"
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type todoRepository struct {
	db *goqu.Database
}

func NewTodo(conn *sql.DB) domain.TodoRepository {
	return &todoRepository{
		db: goqu.New("default", conn),
	}
}

func (tr todoRepository) FindAll(ctx context.Context) (result []domain.Todo, err error) {
	dataset := tr.db.From("todos").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (tr todoRepository) FindById(ctx context.Context, id string) (result domain.Todo, err error) {
	dataset := tr.db.From("todos").Where(goqu.C("deleted_at").IsNull(), goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (tr todoRepository) Save(ctx context.Context, t *domain.Todo) error {
	executor := tr.db.Insert("todos").Rows(t).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (tr todoRepository) Update(ctx context.Context, t *domain.Todo) error {
	executor := tr.db.Update("todos").Where(goqu.C("id").Eq(t.ID)).Set(t).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (tr todoRepository) Delete(ctx context.Context, id string) error {
	executor := tr.db.Update("todos").
		Where(goqu.C("id").Eq(id)).
		Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
