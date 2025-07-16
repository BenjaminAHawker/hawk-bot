package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

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
	migrationsDir := filepath.Join("internal", "db", "migrations")

	// Ensure schema_migrations table exists
	_, err := p.pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			filename TEXT PRIMARY KEY,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`)
	if err != nil {
		return fmt.Errorf("failed to ensure schema_migrations table: %w", err)
	}

	// Read migration directory
	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("failed to read migrations dir: %w", err)
	}

	// Sort filenames alphabetically
	var files []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".sql") {
			files = append(files, entry.Name())
		}
	}
	sort.Strings(files)

	// Apply each migration if not yet applied
	for _, file := range files {
		applied, err := p.migrationApplied(ctx, file)
		if err != nil {
			return err
		}
		if applied {
			continue
		}

		content, err := os.ReadFile(filepath.Join(migrationsDir, file))
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", file, err)
		}

		log.Printf("Applying migration: %s", file)
		if _, err := p.pool.Exec(ctx, string(content)); err != nil {
			return fmt.Errorf("failed to apply migration %s: %w", file, err)
		}

		if _, err := p.pool.Exec(ctx, `INSERT INTO schema_migrations (filename) VALUES ($1)`, file); err != nil {
			return fmt.Errorf("failed to record migration %s: %w", file, err)
		}
	}

	log.Println("All migrations applied successfully.")
	return nil
}

func (p *Postgres) migrationApplied(ctx context.Context, filename string) (bool, error) {
	var exists bool
	err := p.pool.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE filename = $1)`, filename).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check migration %s: %w", filename, err)
	}
	return exists, nil
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
