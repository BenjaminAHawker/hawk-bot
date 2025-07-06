package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func NewSession(token string) (*discordgo.Session, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	// Example: You can set intents here if needed
	// dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages

	err = dg.Open()
	if err != nil {
		return nil, err
	}

	log.Println("Discord session started")
	return dg, nil
}
