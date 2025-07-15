package db

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

// DB defines the methods your app expects from a database connection
type DB interface {
	Ping(ctx context.Context) error
	Close()
	Migrate(ctx context.Context) error
	UpsertUser(ctx context.Context, user *discordgo.User) (int, error)
	// Add more methods here as needed (e.g. Exec, Query, etc.)
}
