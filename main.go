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
		"map": {
			name:        "map",
			description: "Shows the 20 maps",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous 20 maps",
			callback:    commandMapB,
		},
	},
	url: map[string]string{
		"next" : "https://pokeapi.co/api/v2/location-area/",
		"previous" : "",
	},
}

func main() {
	startRepl(Config)
}
