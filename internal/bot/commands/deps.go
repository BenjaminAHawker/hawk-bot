package commands

import (
	"github.com/BenjaminAHawker/hawk-bot/internal/config"
	"github.com/BenjaminAHawker/hawk-bot/internal/db"
	"github.com/bwmarrin/discordgo"
)

type Deps struct {
	Session *discordgo.Session
	DB      db.DB
	CONFIG  config.CF
	// Add other shared dependencies here, like loggers, config, etc.
}
