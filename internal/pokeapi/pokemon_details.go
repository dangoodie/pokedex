package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetPokemonDetails(pokemon *string) (PokemonDetails, error) {
	fullURL := ""
	if pokemon == nil {
		return PokemonDetails{}, fmt.Errorf("must have pokemon")
	} else {
		fullURL = BaseURL + "pokemon/" + *pokemon + "/"
	}

	var data []byte
	var found bool
	var err error

	data, found = c.cache.Get(&fullURL)
	if !found {
		// Make Get Request
		res, err := c.httpClient.Get(fullURL)
		if err != nil {
			return PokemonDetails{}, err
		}
		defer res.Body.Close()

		// Check if Not Found
		if res.StatusCode > 299 {
			return PokemonDetails{}, fmt.Errorf("error code: %s", res.Status)
		}

		// Read data from the response body
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return PokemonDetails{}, err
		}

		// Cache the response
		c.cache.Add(&fullURL, data)
	}

	var pokemonDetails PokemonDetails
	err = json.Unmarshal(data, &pokemonDetails)
	if err != nil {
		return PokemonDetails{}, err
	}

	return pokemonDetails, nil
}
