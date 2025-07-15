package config

// CF defines the methods your app expects from a configuration
type CF interface {
	Load() *Config
	getEnv(key string) string
}
