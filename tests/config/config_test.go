package tests

import (
	"testing"

	"github.com/BenjaminAHawker/hawk-bot/internal/config"
	"github.com/BenjaminAHawker/hawk-bot/tests/testutil"
)

func TestLoadConfig(t *testing.T) {
	env := map[string]string{
		"DISCORD_TOKEN":   "test-token",
		"DB_HOST":         "localhost",
		"DB_PORT":         "5432",
		"DB_USER":         "user",
		"DB_PASSWORD":     "pass",
		"DB_NAME":         "testdb",
		"REQUEST_CHANNEL": "requests",
	}

	testutil.SetEnvVars(env)
	defer testutil.ClearEnvVars(getMapKeys(env))

	cfg := config.Load()

	if cfg.DiscordToken != "test-token" {
		t.Errorf("expected DISCORD_TOKEN to be 'test-token', got '%s'", cfg.DiscordToken)
	}
}

func getMapKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
