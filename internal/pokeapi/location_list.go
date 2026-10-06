package pokeapi

import (
	"net/http"
	"io"
	"encoding/json"
	"fmt"
)

func ListLocations(url string) (LocationAreaResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return LocationAreaResponse{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return LocationAreaResponse{}, err
	}

	var locationArea LocationAreaResponse
	if err = json.Unmarshal(body, &locationArea); err != nil {
		return LocationAreaResponse{}, err
	}
	return locationArea, nil

}
