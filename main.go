package main

import (
	"cmd/main/internal/api"
	"cmd/main/internal/config"
	"cmd/main/internal/connection"
	"cmd/main/internal/repository"
	"cmd/main/internal/service"

	"github.com/gofiber/fiber/v3"
)

func main() {
	conf := config.Get()
	dbConn := connection.GetDatabase(conf.Database)
	app := fiber.New()

	// repositories
	todoRepository := repository.NewTodo(dbConn)

	// services
	todoService := service.NewTodo(todoRepository)

	// apis
	api.NewTodo(app, todoService)
	app.Listen(conf.Server.Host + ":" + conf.Server.Port)
}
