package commands

import (
	"github.com/bwmarrin/discordgo"
)

// SlashCommand is a reusable interface for all commands
type SlashCommand interface {
	Command() *discordgo.ApplicationCommand
	Handler(s *discordgo.Session, i *discordgo.InteractionCreate)
}
