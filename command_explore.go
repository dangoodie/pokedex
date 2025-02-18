package main

import "fmt"

func commandExplore(cfg *config, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("must have a location")
	}
	
	// Print exploring area
	location := args[1]
	fmt.Printf("Exploring %s...\n", location)

	// Get location details
	locationDetails, err := cfg.pokeapiClient.GetLocationDetails(&location)
	if err != nil {
		return fmt.Errorf("error getting location details: %w", err)
	}


	// List pokemon in area
	pokemonEncounters := locationDetails.PokemonEncounters
	if len(pokemonEncounters) == 0 {
		fmt.Println("No pokemon found!")
		return nil
	}

	for _, pokemonEncounter := range pokemonEncounters {
		fmt.Printf(" - %s\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}