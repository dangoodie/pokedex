package main

import "fmt"

func commandPokedex(cfg *config, args *[]string) error {
	if len(cfg.userPokedex) == 0 {
		return fmt.Errorf("you must catch a pokemon first")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.userPokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	return nil
}