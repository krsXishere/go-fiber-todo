package main

import (
	"cmd/main/internal/config"
	"cmd/main/internal/connection"
	"cmd/main/internal/repository"

	"github.com/gofiber/fiber/v3"
)

func main() {
	conf := config.Get()
	dbConn := connection.GetDatabase(conf.Database)
	app := fiber.New()
	todoRepository := repository.NewTodo(dbConn)
	app.Listen(conf.Server.Host + ":" + conf.Server.Port)
}