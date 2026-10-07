package main

import "fmt"


func commandExplore(cfg *config, areaName *string) error {
	nameResp, err := cfg.pokeapiClient.ListNames(areaName)
	if err != nil {
		return err
	}
	fmt.Println("Exploring", *areaName, "...")
	fmt.Println("Found Pokemon:")
	for _, poke := range nameResp.PokemonEncounters {
		fmt.Println("-" ,poke.Pokemon.Name)
	}
	return nil
}
