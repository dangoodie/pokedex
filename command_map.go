package main

import (
	"errors"
	"fmt"

)

func commandMap(cfg *config) error {
	// Get PokeMap from PokeApi
	locationList, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = locationList.Next
	cfg.prevLocationUrl = locationList.Previous

	// Print the locations
	for _, location := range locationList.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	// Stop if on first page
	if cfg.prevLocationUrl == nil {
		return errors.New("you're on the first page")
	}

	// Get PokeMap from PokeApi
	locationList, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationUrl = locationList.Next
	cfg.prevLocationUrl = locationList.Previous

	// Print the locations
	for _, location := range locationList.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}
