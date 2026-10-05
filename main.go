package main

var Config = config{
	registry: map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays help",
			callback:    commandHelp,
		},
	},
}

func main() {
	startRepl(Config)
}
