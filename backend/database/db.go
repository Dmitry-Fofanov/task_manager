package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func Connect(ctx context.Context) error {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@db/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	var err error
	DB, err = pgx.Connect(ctx, connectionString)
	if err != nil {
		return err
	}

	return nil
}
