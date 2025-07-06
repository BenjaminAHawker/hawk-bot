package commands

import (
	"github.com/BenjaminAHawker/hawk-bot/internal/db"
)

type Deps struct {
	DB db.DB
	// Add other shared dependencies here, like loggers, config, etc.
}
