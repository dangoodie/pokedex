package main

import "fmt"

func commandInspect(cfg *config, args *[]string) error {
	if len(*args) < 2 {
		return fmt.Errorf("must have a pokemon name")
	}
	pokemonName := (*args)[1]
	pokemonDetails, found := cfg.userPokedex[pokemonName]
	if !found {
		return fmt.Errorf("you must catch a pokemon before inspecting")
	}

	fmt.Printf("Name: %s\n", pokemonDetails.Name)
	fmt.Printf("Height: %d\n", pokemonDetails.Height)
	fmt.Printf("Weight: %d\n", pokemonDetails.Weight)
	fmt.Printf("Stats:\n")
	for _, s := range pokemonDetails.Stats {
		fmt.Printf("  - %s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemonDetails.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}