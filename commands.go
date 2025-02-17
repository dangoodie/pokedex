package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	apiUrl = "https://pokeapi.co/api/v2/"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Use your map to find an area",
			callback:    commandMap,
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

func commandMap() error {
	// Make Get request
	res, err := http.Get(apiUrl + "location-area/")
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

	// Print the JSON data for now
	fmt.Printf("Count: %d\n", pokeMap.Count)
	if pokeMap.Next != nil {
		fmt.Printf("Next: %s\n", *pokeMap.Next)
	}
	if pokeMap.Previous != nil {
		fmt.Printf("Previous: %s\n", *pokeMap.Previous)
	}
	for i, result := range pokeMap.Results {
		fmt.Printf("%d: %s\n", i+1, result.Name)
	}

	return nil
}

func commandHelp() error {
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
