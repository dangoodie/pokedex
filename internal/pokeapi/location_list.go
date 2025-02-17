package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListLocations(pageUrl *string) (PokeMap, error) {
	fullURL := ""
	if pageUrl != nil {
		fullURL = *pageUrl
	} else {
		fullURL = BaseURL + "location-area/?offset=0&limit=20" // default query for page 1
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

	// Save the full URL for caching purposes
	pokeMap.URL = &fullURL

	return pokeMap, nil
}
