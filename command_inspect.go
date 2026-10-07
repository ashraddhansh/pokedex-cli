package main

import "fmt"


func commandInspect(cfg *config, name *string) error {
	resp, ok := cfg.caught[*name]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}
	fmt.Println("Name:", resp.Name)
	fmt.Println("Height:", resp.Height)
	fmt.Println("Weight:", resp.Weight)
	fmt.Println("Stats:")

	for _, poke := range resp.Stats {
		fmt.Println("  -", poke.Stat.Name)
	}

	fmt.Println("Types:")
	for _, poke := range resp.Types {
		fmt.Println("  -", poke.Type.Name)
	}


	return nil
}
