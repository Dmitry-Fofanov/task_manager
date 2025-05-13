package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	tasks := app.Group("/tasks")
	tasks.Get("/", GetTasks)
	tasks.Post("/", CreateTask)
	tasks.Put("/:id", UpdateTask)
	tasks.Delete("/:id", DeleteTask)
}
