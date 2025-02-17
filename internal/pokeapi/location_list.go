package pokeapi

import (
	"encoding/json"
)

func (c *Client) ListLocations(pageUrl *string) (LocationList, error) {
	fullURL := ""
	if pageUrl != nil {
		fullURL = *pageUrl
	} else {
		fullURL = BaseURL + "location-area/?offset=0&limit=20" // default query for page 1
	}

	// Make Get Request
	res, err := c.httpClient.Get(fullURL)
	if err != nil {
		return LocationList{}, err
	}
	defer res.Body.Close()

	// Unmarshal JSON
	var locationList LocationList
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locationList)
	if err != nil {
		return LocationList{}, err
	}

	// Save the full URL for caching purposes
	locationList.URL = &fullURL

	return locationList, nil
}
