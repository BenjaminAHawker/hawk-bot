package commands

import (
	"github.com/bwmarrin/discordgo"
)

type HelloCommand struct {
	Deps *Deps
}

// responder defines the minimal interface we need to test this command
type responder interface {
	InteractionRespond(i *discordgo.Interaction, r *discordgo.InteractionResponse, options ...discordgo.RequestOption) error
}

func NewHelloCommand(deps *Deps) *HelloCommand {
	return &HelloCommand{Deps: deps}
}

func (c *HelloCommand) Command() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "hello",
		Description: "Replies with a friendly greeting",
	}
}

// Handler is the live Discord event hook
func (c *HelloCommand) Handler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	c.replyHello(s, i.Interaction)
}

// replyHello is testable by injecting a mock responder
func (c *HelloCommand) replyHello(s responder, i *discordgo.Interaction) {
	s.InteractionRespond(i, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Hello! Nice to meet you.",
		},
	})
}
