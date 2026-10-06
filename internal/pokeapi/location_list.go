package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	var body []byte

	if data, ok := c.cache.Get(url); ok {
		body = data
	} else {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		c.cache.Add(url, body)
	}
	locationResp := LocationAreaResponse{}

	err := json.Unmarshal(body, &locationResp)
		if err != nil {
			return LocationAreaResponse{}, err
	}

	return locationResp, nil
}
