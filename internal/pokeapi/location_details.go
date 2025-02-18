package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) GetLocationDetails(location *string) (LocationDetails, error) {
	fullURL := ""
	if location == nil {
		return LocationDetails{}, fmt.Errorf("must have location")
	} else {
		fullURL = BaseURL + "location-area/" + *location + "/"
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
			return LocationDetails{}, err
		}
		defer res.Body.Close()

		// Check if Not Found
		if res.StatusCode > 299 {
			return LocationDetails{}, fmt.Errorf("error code: %s", res.Status)
		}

		// Read data from the response body
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationDetails{}, err
		}

		// Cache the response
		c.cache.Add(&fullURL, data)
	}

	// Unmarshal JSON
	var locationDetails LocationDetails
	err = json.Unmarshal(data, &locationDetails)
	if err != nil {
		return LocationDetails{}, err
	}

	return locationDetails, nil
}
