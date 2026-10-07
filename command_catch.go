package main

import (
	"math/rand"
	"fmt"
)

func iscaught(baseExp int) bool {
	return rand.Intn(700) > baseExp
}


func commandCatch(cfg *config, name *string) error {
	pokeResp, err := cfg.pokeapiClient.CatchPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", *name)
	if pokeResp.Name == "" {
		return fmt.Errorf("Please enter the name correctly")
	}
	if iscaught(pokeResp.BaseExperience) {
		fmt.Println(pokeResp.Name, "was caught!")
		cfg.caught[pokeResp.Name] = pokeResp
	} else {
		fmt.Println(pokeResp.Name, "escaped!")
	}

	return nil
}
