package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dangoodie/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	userPokedex     map[string]pokeapi.PokemonDetails
	nextLocationUrl *string
	prevLocationUrl *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args *[]string) error
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		args := cleanInput(scanner.Text())
		if len(args) == 0 {
			continue
		}

		command, ok := getCommands()[args[0]]
		if ok {
			err := command.callback(cfg, &args)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
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
		"explore": {
			name: "explore",
			description: "Search a location for Pokemon\n" +
				"Usage: explore <location-name>",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Try to catch a pokemon\n" +
				"Usage: catch <pokemon-name>",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Inspect pokemon that you have caught\n" +
				"Usage: inspect <pokemon-name>",
			callback: commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List the caught pokemon in your pokedex",
			callback:    commandPokedex,
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
