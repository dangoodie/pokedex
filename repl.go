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
	nextLocationUrl *string
	prevLocationUrl *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
			err := command.callback(cfg)
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