package db

import (
	"context"
	"fmt"
	"log"

	"github.com/BenjaminAHawker/hawk-bot/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(cfg *config.Config) *pgxpool.Pool {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Failed to create DB pool: %v", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	log.Println("Connected to PostgreSQL")
	return pool
}
