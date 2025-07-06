package mocks

import (
	"context"
)

// MockDB is a mock implementation of the db.DB interface
type MockDB struct {
	PingFunc  func(ctx context.Context) error
	CloseFunc func()
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
