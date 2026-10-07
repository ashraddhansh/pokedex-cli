package pokeapi

import (
	"io"
	"net/http"
	"encoding/json"
	"fmt"
)


func (c *Client) ListNames(areaName *string) (NameResponse, error) {
	if areaName == nil {
		return NameResponse{}, fmt.Errorf("Please add arguments")
	}

	url := baseURL + "/location-area/" + *areaName

	var body []byte

	if data, ok := c.cache.Get(url); ok {
		body = data
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return NameResponse{}, err
		}

		resp, err := c.httpClient.Do(req)

		if err != nil {
			return NameResponse{}, err
		}

		defer resp.Body.Close()
		
		body, err = io.ReadAll(resp.Body)
		
		if err != nil {
			return NameResponse{}, err
		}
		c.cache.Add(url, body)
	}
	nameResp := NameResponse{}

	err := json.Unmarshal(body, &nameResp)
		if err != nil {
			return NameResponse{}, err
	}

	return nameResp, nil
}

