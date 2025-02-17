package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	// Get PokeMap from PokeApi
	pokeMap, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = pokeMap.Next
	cfg.prevLocationUrl = pokeMap.Previous

	// Print the locations
	for _, location := range pokeMap.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationUrl == nil {
		return errors.New("you're on the first page")
	}

	pokeMap, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = pokeMap.Next
	cfg.prevLocationUrl = pokeMap.Previous

	// Print the locations
	for _, location := range pokeMap.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}
