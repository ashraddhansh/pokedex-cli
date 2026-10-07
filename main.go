package main

import (
	"time"

	"github.com/ashraddhansh/pokedexcli/internal/pokeapi"
)

func main(){
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		commands: getCommands(),
		pokeapiClient: pokeClient,
		caught: make(map[string]pokeapi.PokemonResponse),
	}

	startRepl(cfg)
}
