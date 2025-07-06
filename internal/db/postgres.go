package db

import (
	"context"
	"fmt"
	"log"

	"github.com/BenjaminAHawker/hawk-bot/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Postgres is the concrete implementation of the DB interface
type Postgres struct {
	pool *pgxpool.Pool
}

// Connect initializes and returns a DB (implemented by Postgres)
func Connect(cfg *config.Config) DB {
	connString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	return &Postgres{pool: pool}
}

// Ping implements the DB interface
func (p *Postgres) Ping(ctx context.Context) error {
	return p.pool.Ping(ctx)
}

// Close implements the DB interface
func (p *Postgres) Close() {
	p.pool.Close()
}
