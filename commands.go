package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	apiUrl = "https://pokeapi.co/api/v2/"
)

type Config struct {
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Move forward on the map 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Move backward on the map 20 locations",
			callback:    commandMapb,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

type PokeMap struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *Config) error {
	// Make Get request
	fullURL := ""
	if cfg.Next != nil {
		fullURL = *cfg.Next
	} else {
		fullURL = apiUrl + "location-area/"
	}

	res, err := http.Get(fullURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Unmarshal the JSON
	var pokeMap PokeMap
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&pokeMap)
	if err != nil {
		return err
	}

	// Set the map configuration
	setMapConfig(&pokeMap, cfg)
	offset, err := getOffset(fullURL)
	if err != nil {
		return err
	}

	// Print the JSON data for now
	fmt.Printf("Count: %d\n", pokeMap.Count)
	if pokeMap.Next != nil {
		fmt.Printf("Next: %s\n", *pokeMap.Next)
	}
	if pokeMap.Previous != nil {
		fmt.Printf("Previous: %s\n", *pokeMap.Previous)
	}
	for i, result := range pokeMap.Results {
		fmt.Printf("%d: %s\n", i+1+offset, result.Name)
	}

	return nil
}

func commandMapb(cfg *Config) error {
	// Make Get request
	fullURL := ""
	if cfg.Previous != nil {
		fullURL = *cfg.Previous
	} else {
		fullURL = apiUrl + "location-area/"
	}

	res, err := http.Get(fullURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Unmarshal the JSON
	var pokeMap PokeMap
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&pokeMap)
	if err != nil {
		return err
	}

	// Set the map configuration
	setMapConfig(&pokeMap, cfg)
	offset, err := getOffset(fullURL)
	if err != nil {
		return err
	}

	// Print the JSON data for now
	fmt.Printf("Count: %d\n", pokeMap.Count)
	if pokeMap.Next != nil {
		fmt.Printf("Next: %s\n", *pokeMap.Next)
	}
	if pokeMap.Previous != nil {
		fmt.Printf("Previous: %s\n", *pokeMap.Previous)
	}
	for i, result := range pokeMap.Results {
		fmt.Printf("%d: %s\n", i+1+offset, result.Name)
	}

	return nil
}

func setMapConfig(pokeMap *PokeMap, cfg *Config) {
	if pokeMap.Next != nil {
		cfg.Next = pokeMap.Next
	}

	if pokeMap.Previous != nil {
		cfg.Previous = pokeMap.Previous
	}
}

func getOffset(fullURL string) (int, error) {
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		return 0, err
	}

	queries := parsedURL.Query()
	offsetStr := queries.Get("offset")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return 0, nil
	}

	return offset, nil
}

func commandHelp(cfg *Config) error {
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

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
