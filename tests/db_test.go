package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/BenjaminAHawker/hawk-bot/tests/mocks"
)

func TestMockDBPingSuccess(t *testing.T) {
	db := &mocks.MockDB{
		PingFunc: func(ctx context.Context) error {
			return nil // simulate success
		},
	}

	err := db.Ping(context.Background())
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestMockDBPingFailure(t *testing.T) {
	db := &mocks.MockDB{
		PingFunc: func(ctx context.Context) error {
			return errors.New("mock failure")
		},
	}

	err := db.Ping(context.Background())
	if err == nil {
		t.Error("expected error, got nil")
	}
}
