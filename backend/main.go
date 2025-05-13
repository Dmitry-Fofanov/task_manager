package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"

	"backend/database"
	"backend/handlers"
)

func main() {
	ctx := context.Background()
	err := database.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer database.DB.Close(ctx)

	app := fiber.New()
	handlers.SetupRoutes(app)

	log.Fatal(app.Listen(":80"))
}
