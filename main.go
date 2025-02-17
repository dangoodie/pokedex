package main

import (
	"time"

	"github.com/dangoodie/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := config{
		pokeapiClient: pokeClient,
	}

	startREPL(&cfg)
}
