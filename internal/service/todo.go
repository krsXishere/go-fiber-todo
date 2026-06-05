package service

import (
	"cmd/main/domain"
	"cmd/main/dto"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type todoService struct {
	todoRepository domain.TodoRepository
}

func NewTodo(todoRepository domain.TodoRepository) domain.TodoService {
	return &todoService{
		todoRepository: todoRepository,
	}
}

func (t todoService) Index(ctx context.Context) ([]dto.TodoData, error) {
	todos, err := t.todoRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var todoData []dto.TodoData
	for _, v := range todos {
		todoData = append(todoData, dto.TodoData{
			ID:        v.ID,
			Title:     v.Title,
			Subtitle:  v.Subtitle,
			CreatedAt: v.CreatedAt.Time.GoString(),
			UpdatedAt: v.UpdatedAt.Time.GoString(),
		})
	}

	return todoData, nil
}

func (t todoService) Create(ctx context.Context, req dto.CreateTodoData) error {
	todo := domain.Todo{
		ID:       uuid.NewString(),
		Title:    req.Title,
		Subtitle: req.Subtitle,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
		UpdatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}

	return t.todoRepository.Save(ctx, &todo)
}

func (t todoService) Update(ctx context.Context, req dto.UpdateTodoData) error {
	persistedTodo, err := t.todoRepository.FindById(ctx, req.ID)
	if err != nil {
		return err
	}

	if persistedTodo.ID == "" {
		return errors.New("todo data not found")
	}
	
	persistedTodo.Title = req.Title
	persistedTodo.Subtitle = req.Subtitle
	persistedTodo.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	return t.todoRepository.Update(ctx, &persistedTodo)
}
