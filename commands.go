package main

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	cliCommands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List the location areas on the map - 20 at a time. returns the next 20 every use",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "same as the map command, but goes back to the previous 20 location areas",
			callback:    commandMapb,
		},
	}
	return cliCommands
}
