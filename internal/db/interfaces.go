package db

import "context"

// DB defines the methods your app expects from a database connection
type DB interface {
	Ping(ctx context.Context) error
	Close()
	// Add more methods here as needed (e.g. Exec, Query, etc.)
}
