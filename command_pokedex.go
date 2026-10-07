package main

import "fmt"


func commandPokedex(cfg *config, _ *string) error {
	
	resp := cfg.caught
	
	if len(resp) == 0 {
		return fmt.Errorf("You haven't caught any pokemon yet!")
	}

	fmt.Println("Your Pokedex:")

	for _, poke := range (resp) {
		fmt.Println("-", poke.Name)
	}


	return nil
	
}
