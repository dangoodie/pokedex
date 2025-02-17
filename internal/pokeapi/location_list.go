package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListLocations(pageUrl *string) (PokeMap, error) {
	fullURL := ""
	if pageUrl != nil {
		fullURL = *pageUrl
	} else {
		fullURL = baseURL + "location-area/"
	}

	// Make Get Request
	res, err := c.httpClient.Get(fullURL)
	if err != nil {
		return PokeMap{}, err
	}
	defer res.Body.Close()

	// Unmarshal JSON
	var pokeMap PokeMap
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&pokeMap)
	if err != nil {
		return PokeMap{}, err
	}

	return pokeMap, nil
}
