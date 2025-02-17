package main

import (
	"time"

	"github.com/dangoodie/pokedex/internal/pokeapi"
	"github.com/dangoodie/pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cache := pokecache.NewCache(1 * time.Minute)
	cfg := config{
		pokeapiClient: pokeClient,
		pokecache: cache,
	}

	startREPL(&cfg)
}
