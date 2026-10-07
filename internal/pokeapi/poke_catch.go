package pokeapi


import (
	"io"
	"net/http"
	"encoding/json"
	"fmt"
)

func (c *Client) CatchPokemon(name *string) (PokemonResponse, error) {
	if name == nil {
		return PokemonResponse{}, fmt.Errorf("Please provide the pokemon name")
	}

	url := baseURL + "/pokemon/" + *name

	var body []byte

	if data, ok := c.cache.Get(url); ok {
		body = data
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonResponse{}, err
		}

		resp, err := c.httpClient.Do(req)

		if err != nil {
			return PokemonResponse{}, err
		}

		defer resp.Body.Close()
		
		body, err = io.ReadAll(resp.Body)
		
		if err != nil {
			return PokemonResponse{}, err
		}
		c.cache.Add(url, body)
	}
	pokeResp := PokemonResponse{}

	err := json.Unmarshal(body, &pokeResp)
		if err != nil {
			return PokemonResponse{}, err
	}

	return pokeResp, nil
}

