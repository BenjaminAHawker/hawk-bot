package commands

func LoadAll(deps *Deps) []SlashCommand {
	return []SlashCommand{
		NewPingCommand(),
		// Add other command constructors here
	}
}
