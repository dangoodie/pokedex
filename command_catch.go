package main

import (
	"fmt"
	"math/rand"

	"github.com/dangoodie/pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, args *[]string) error {
	if len(*args) < 2 {
		return fmt.Errorf("must have a pokemon")
	}
	//Print catching pokemon
	pokemonName := (*args)[1]
	fmt.Printf("Throwing a Pokeball at %s\n", pokemonName)

	// Get pokemon details
	pokemonDetails, err := cfg.pokeapiClient.GetPokemonDetails(&pokemonName)
	if err != nil {
		return fmt.Errorf("error getting pokemon details: %w", err)
	}

	// Try to catch pokemon
	caught := catchPokemon(&pokemonDetails)
	if caught {
		fmt.Printf("Caught %s\n", pokemonDetails.Name)
	} else {
		fmt.Printf("Didn't catch %s\n", pokemonDetails.Name)
	}

	return nil
}

func getCatchChance(baseExp int) float64 {
	// Constants
	maxBaseExp := 395.0 // Chansey's base experience
	k := 50.0           // Arbitrary balancing factor to avoid low catch rates

	// Calculate probability
	prob := 1 - float64(baseExp)/(maxBaseExp+k)

	// Ensure at least 5% chance
	if prob < 0.05 {
		prob = 0.05
	}
	return prob
}

func catchPokemon(pokemon *pokeapi.PokemonDetails) bool {
	catchChance := getCatchChance(pokemon.BaseExperience)
	roll := rand.Float64()

	return roll < catchChance
}