package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func InitDB() error {
	godotenv.Load()
	url := os.Getenv("DATABASE_URL")
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return err
	}
	DB = pool
	return nil
}
