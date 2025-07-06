package testutil

import "os"

// SetEnvVars sets environment variables for testing
func SetEnvVars(env map[string]string) {
	for key, value := range env {
		os.Setenv(key, value)
	}
}

// ClearEnvVars clears environment variables after testing
func ClearEnvVars(keys []string) {
	for _, key := range keys {
		os.Unsetenv(key)
	}
}

// GetMapKeys returns the keys of a map[string]string
func GetMapKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
