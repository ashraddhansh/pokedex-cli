package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func startRepl(cfg config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		commandName := cleanInput(userInput)[0]

		command, exists := cfg.registry[commandName]
		
		if exists {
			err := command.callback(cfg)
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
	callback    func(config) error
}


type config struct {
	registry map[string]cliCommand

}
