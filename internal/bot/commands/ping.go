package commands

import (
	"github.com/bwmarrin/discordgo"
)

// PingCommand handles the /ping slash command
type PingCommand struct{}

func NewPingCommand() SlashCommand {
	return &PingCommand{}
}

func (c *PingCommand) Command() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Replies with pong.",
	}
}

func (c *PingCommand) Handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}
