package api

import (
	"cmd/main/domain"
	"cmd/main/dto"
	"cmd/main/internal/utility"
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
)

type todoApi struct {
	todoService domain.TodoService
}

func NewTodo(app *fiber.App, todoService domain.TodoService) {
	ta := todoApi{
		todoService: todoService,
	}

	app.Get("/todos", ta.Index)
	app.Post("/todos", ta.Create)
	app.Put("/todos/:id", ta.Update)
}

func (ta todoApi) Index(ctx fiber.Ctx) error {
	t, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := ta.todoService.Index(t)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ta todoApi) Create(ctx fiber.Ctx) error {
	t, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateTodoData
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := utility.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	err := ta.todoService.Create(t, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).
		JSON(dto.CreateResponseSuccess("todo created"))
}

func (ta todoApi) Update(ctx fiber.Ctx) error {
	t, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateTodoData
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := utility.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	req.ID = ctx.Params("id")

	err := ta.todoService.Update(t, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).
		JSON(dto.CreateResponseSuccess("todo updated"))
}
