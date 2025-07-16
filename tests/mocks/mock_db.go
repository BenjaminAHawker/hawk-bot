package mocks

import (
	"context"

	"github.com/bwmarrin/discordgo"
)

// MockDB is a mock implementation of the db.DB interface
type MockDB struct {
	PingFunc       func(ctx context.Context) error
	CloseFunc      func()
	MigrateFunc    func(ctx context.Context) error
	UpsertUserFunc func(ctx context.Context, user *discordgo.User) (int, error)
}

// Ping calls the configured mock function or returns nil
func (m *MockDB) Ping(ctx context.Context) error {
	if m.PingFunc != nil {
		return m.PingFunc(ctx)
	}
	return nil
}

// Close calls the configured mock function if present
func (m *MockDB) Close() {
	if m.CloseFunc != nil {
		m.CloseFunc()
	}
}

func (m *MockDB) Migrate(ctx context.Context) error {
	if m.MigrateFunc != nil {
		return m.MigrateFunc(ctx)
	}
	return nil
}

// UpsertUser calls the configured mock function or returns a default value
func (m *MockDB) UpsertUser(ctx context.Context, user *discordgo.User) (int, error) {
	if m.UpsertUserFunc != nil {
		return m.UpsertUserFunc(ctx, user)
	}
	return 0, nil
}
