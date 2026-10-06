package main

import (
	"fmt"
	"github.com/ashraddhansh/pokedexcli/internal/pokeapi"
)


func commandMap(cfg config) error {

	res, err := pokeapi.ListLocations(cfg.url["next"])

	if err != nil {
		return err
	}

	cfg.url["next"] = res.Next
	cfg.url["previous"] = res.Previous

	for _, result := range res.Results{
		fmt.Println(result.Name)
	}

	return nil
}


func commandMapB(cfg config) error {
	exists := cfg.url["previous"]
	if exists == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := pokeapi.ListLocations(cfg.url["previous"])

	if err != nil {
		return err
	}

	cfg.url["next"] = res.Next
	cfg.url["previous"] = res.Previous

	for _, result := range res.Results{
		fmt.Println(result.Name)
	}

	return nil
}
