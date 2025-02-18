package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) ListLocations(pageUrl *string) (LocationList, error) {
	fullURL := ""
	if pageUrl != nil {
		fullURL = *pageUrl
	} else {
		fullURL = BaseURL + "location-area/?offset=0&limit=20" // default query for page 1
	}

	var data []byte
	var found bool
	var err error

	// Check cache for hit
	data, found = c.cache.Get(&fullURL)
	if !found {
		// Make Get Request
		res, err := c.httpClient.Get(fullURL)
		if err != nil {
			return LocationList{}, err
		}
		defer res.Body.Close()

		// Check if Not Found
		if res.StatusCode > 299 {
			return LocationList{}, fmt.Errorf("error code: %s", res.Status)
		}

		// Read data from the response body
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationList{}, err
		}

		// Cache the response
		c.cache.Add(&fullURL, data)
	}

	// Unmarshal JSON
	var locationList LocationList
	err = json.Unmarshal(data, &locationList)
	if err != nil {
		return LocationList{}, err
	}

	return locationList, nil
}
