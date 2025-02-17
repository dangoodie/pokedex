package main

import (
	"errors"
	"fmt"
	"os"
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

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
