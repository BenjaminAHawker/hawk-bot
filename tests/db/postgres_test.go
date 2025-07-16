package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/BenjaminAHawker/hawk-bot/tests/mocks"
	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/require"
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

func TestUpsertUser_UsingMockDB(t *testing.T) {
	mock := &mocks.MockDB{
		UpsertUserFunc: func(ctx context.Context, user *discordgo.User) (int, error) {
			require.Equal(t, "testuser", user.Username)
			return 42, nil
		},
	}

	user := &discordgo.User{
		ID:       "123",
		Username: "testuser",
		Avatar:   "hash",
	}

	id, err := mock.UpsertUser(context.Background(), user)
	require.NoError(t, err)
	require.Equal(t, 42, id)
}
