package commands

func LoadAll(deps *Deps) []SlashCommand {
	return []SlashCommand{
		NewHelloCommand(deps),
		// Add other command constructors here
	}
}
