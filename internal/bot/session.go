package bot

import (
	"log"

	"github.com/BenjaminAHawker/hawk-bot/internal/bot/commands"
	"github.com/BenjaminAHawker/hawk-bot/internal/db"
	"github.com/bwmarrin/discordgo"
)

func NewSession(token string, dbConn db.DB) (*discordgo.Session, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	deps := &commands.Deps{DB: dbConn}
	cmds := commands.LoadAll(deps)

	// Handle interactions
	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		for _, cmd := range cmds {
			if i.ApplicationCommandData().Name == cmd.Command().Name {
				cmd.Handler(s, i)
				return
			}
		}
	})

	// Wait for the Ready event to get bot user ID
	ready := make(chan struct{})
	dg.AddHandlerOnce(func(s *discordgo.Session, r *discordgo.Ready) {
		close(ready)
	})

	if err := dg.Open(); err != nil {
		return nil, err
	}

	log.Println("Discord session started")
	<-ready // wait for ready event

	appID := dg.State.User.ID

	// Fetch bot guilds
	guilds, err := dg.UserGuilds(10, "", "", true)
	if err != nil {
		log.Println("Failed to fetch guilds, no slash commands will be registered")
		return dg, nil
	}
	if len(guilds) == 0 {
		log.Println("Bot is not in any guilds")
		return dg, nil
	}

	// Register commands to all guilds (or just first if you prefer)
	for _, guild := range guilds {
		for _, cmd := range cmds {
			_, err := dg.ApplicationCommandCreate(appID, guild.ID, cmd.Command())
			if err != nil {
				log.Printf("Failed to register command %s in guild %s: %v", cmd.Command().Name, guild.ID, err)
			}
		}
	}

	return dg, nil
}
