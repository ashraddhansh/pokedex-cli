package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/ashraddhansh/pokedexcli/internal/pokeapi"
)


type config struct {
	commands	map[string]cliCommand
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string


}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		var arg *string
		if len(words) > 1 {
			arg = &words[1]
		}

		commandName := words[0]

		command, exists := cfg.commands[commandName]

		if exists{
			err := command.callback(cfg, arg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}


	}
}

func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	trimmed := strings.Fields(lowercase)
	return trimmed
}


type cliCommand struct {
	name        string
	description string
	callback    func(*config, *string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore <area-name>",
			description: "Explore pokemons in particular location areas enters in argument",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
	}
}
