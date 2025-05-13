package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"

	"backend/database"
	"backend/models"
)

func GetTasks(c *fiber.Ctx) error {
	tasks := []models.Task{}

	sqlQuery := `
		SELECT id, title, description, status, created_at, updated_at
		FROM tasks
		ORDER BY id
	`

	rows, err := database.DB.Query(context.Background(), sqlQuery)
	if err != nil {
		log.Printf("Ошибка при получении списка задач: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Не вышло получить список задач")
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)

		if err != nil {
			log.Printf("Ошибка при чтении задачи: %v", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Не вышло получить список задач")
		}

		tasks = append(tasks, task)
	}

	return c.JSON(fiber.Map{
		"tasks": tasks,
	})
}

func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	if task.Status == "" {
		task.Status = "new"
	} else if !models.ValidateStatus(task.Status) {
		return fiber.NewError(http.StatusBadRequest, fmt.Sprintf("Неверный статус. Возможные статусы: %v", models.ValidStatuses))
	}

	sqlQuery := `
		INSERT INTO tasks (title, description, status)
		VALUES ($1, $2, $3)
		RETURNING id, title, description, status, created_at, updated_at
	`

	err := database.DB.QueryRow(context.Background(), sqlQuery,
		task.Title,
		task.Description,
		task.Status,
	).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		log.Printf("Ошибка при создании задачи: %v", err)
		return fiber.NewError(http.StatusInternalServerError, "Не вышло создать задачу")
	}

	return c.Status(http.StatusCreated).JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	if task.Status != "" && !models.ValidateStatus(task.Status) {
		return fiber.NewError(http.StatusBadRequest, fmt.Sprintf("Неверный статус. Возможные статусы: %v", models.ValidStatuses))
	}

	sqlQuery := `
		UPDATE tasks
		SET title = COALESCE(NULLIF($1, ''), title),
		    description = COALESCE(NULLIF($2, ''), description),
		    status = COALESCE(NULLIF($3, ''), status),
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $4
		RETURNING id, title, description, status, created_at, updated_at
	`

	err := database.DB.QueryRow(context.Background(), sqlQuery,
		task.Title,
		task.Description,
		task.Status,
		id,
	).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		if err == pgx.ErrNoRows {
			return fiber.NewError(http.StatusNotFound, "Задача не найдена")
		}
		log.Printf("Ошибка при обновлении данных задачи: %v", err)
		return fiber.NewError(http.StatusInternalServerError, "Не вышло обновить данные задачи")
	}

	return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	sqlQuery := `DELETE FROM tasks WHERE id = $1`
	result, err := database.DB.Exec(context.Background(), sqlQuery, id)

	if err != nil {
		log.Printf("Ошибка при удалении задачи: %v", err)
		return fiber.NewError(http.StatusInternalServerError, "Не вышло удалить задачу")
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fiber.NewError(http.StatusNotFound, "Задача не найдена")
	}

	return c.SendStatus(http.StatusNoContent)
}
