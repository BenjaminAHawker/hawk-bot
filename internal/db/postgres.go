package db

import (
	"context"
	"fmt"
	"log"

	"github.com/BenjaminAHawker/hawk-bot/internal/config"
	"github.com/bwmarrin/discordgo"
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

// Migrate applies the necessary migrations to the database
func (p *Postgres) Migrate(ctx context.Context) error {
	schemaSQL := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			discord_id VARCHAR UNIQUE NOT NULL,
			avatar VARCHAR,
			username VARCHAR
		);`,
		`CREATE TABLE IF NOT EXISTS request_types (
			id SERIAL PRIMARY KEY,
			description VARCHAR NOT NULL UNIQUE
		);`,
		`CREATE TABLE IF NOT EXISTS request_status_types (
			id INTEGER PRIMARY KEY,
			description VARCHAR NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS requests (
			id SERIAL PRIMARY KEY,
			request_type_id INTEGER NOT NULL REFERENCES request_types(id),
			user_requested INTEGER NOT NULL REFERENCES users(id),
			status_id INTEGER NOT NULL DEFAULT 0 REFERENCES request_status_types(id),
			processed_by INTEGER REFERENCES users(id)
		);`,
		`CREATE INDEX IF NOT EXISTS idx_requests_user_requested ON requests(user_requested);`,
	}

	for _, stmt := range schemaSQL {
		if _, err := p.pool.Exec(ctx, stmt); err != nil {
			return fmt.Errorf("migration failed: %w", err)
		}
	}

	// Seed request types
	requestTypes := []string{"audiobook", "tv", "movie", "anime", "comic", "manga", "game", "music", "ebook"}
	for _, desc := range requestTypes {
		_, err := p.pool.Exec(ctx,
			`INSERT INTO request_types (description) VALUES ($1) ON CONFLICT (description) DO NOTHING;`,
			desc)
		if err != nil {
			return fmt.Errorf("seeding request_types failed: %w", err)
		}
	}

	// Seed request status types
	statusTypes := map[int]string{
		0: "pending",
		1: "approved",
		2: "denied",
	}
	for id, desc := range statusTypes {
		_, err := p.pool.Exec(ctx,
			`INSERT INTO request_status_types (id, description) VALUES ($1, $2) ON CONFLICT (id) DO NOTHING;`,
			id, desc)
		if err != nil {
			return fmt.Errorf("seeding request_status_types failed: %w", err)
		}
	}

	log.Println("Database migrations and seed complete")
	return nil
}

// UpsertUser inserts or updates a user in the database
func (p *Postgres) UpsertUser(ctx context.Context, user *discordgo.User) (int, error) {
	query := `
		INSERT INTO users (discord_id, username, avatar)
		VALUES ($1, $2, $3)
		ON CONFLICT (discord_id) DO UPDATE 
		SET username = EXCLUDED.username,
		    avatar = EXCLUDED.avatar
		RETURNING id;
	`

	var id int
	err := p.pool.QueryRow(ctx, query, user.ID, user.Username, user.Avatar).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert user: %w", err)
	}

	return id, nil
}
