package main

import (
	"time"

	"github.com/dangoodie/pokedex/internal/pokeapi"
)

func main() {
	httpClientTimeout := 5 * time.Second
	cacheInterval := 1 * time.Minute
	pokeClient := pokeapi.NewClient(httpClientTimeout, cacheInterval)

	cfg := config{
		pokeapiClient: pokeClient,
		userPokedex: make(map[string]pokeapi.PokemonDetails),
	}

	startREPL(&cfg)
}
