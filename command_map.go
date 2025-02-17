package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dangoodie/pokedex/internal/pokeapi"
)

func commandMap(cfg *config) error {
	// Check cache first
	var pokeMap pokeapi.PokeMap
	val, found := cfg.pokecache.Get(cfg.nextLocationUrl)
	if found {
		if err := json.Unmarshal(val, &pokeMap); err != nil {
			return fmt.Errorf("failed to unmarshal cached data: %w", err)
		}
	} else {
		// Get PokeMap from PokeApi
		var err error
		pokeMap, err = cfg.pokeapiClient.ListLocations(cfg.nextLocationUrl)
		if err != nil {
			return err
		}

		// Store response in cache
		jsonData, err := json.Marshal(pokeMap)
		if err ==  nil {
			cfg.pokecache.Add(pokeMap.URL, jsonData)
		}
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
	// Stop if on first page
	if cfg.prevLocationUrl == nil {
		return errors.New("you're on the first page")
	}

	var pokeMap pokeapi.PokeMap
	val, found := cfg.pokecache.Get(cfg.prevLocationUrl)
	if found {
		if err := json.Unmarshal(val, &pokeMap); err != nil {
			return fmt.Errorf("failed to unmarshal cached data: %w", err)
		}
	} else {
		// Get PokeMap from PokeApi
		var err error
		pokeMap, err = cfg.pokeapiClient.ListLocations(cfg.prevLocationUrl)
		if err != nil {
			return err
		}

		// Store response in cache
		jsonData, err := json.Marshal(pokeMap)
		if err == nil {
			cfg.pokecache.Add(pokeMap.URL, jsonData)
		}
	}

	cfg.nextLocationUrl = pokeMap.Next
	cfg.prevLocationUrl = pokeMap.Previous

	// Print the locations
	for _, location := range pokeMap.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}
