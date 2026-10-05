package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg config) error {
	res, err := http.Get(cfg.url["next"])
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	var locationArea LocationAreaResponse
	if err = json.Unmarshal(body, &locationArea); err != nil {
		return err
	}

	for _, result := range locationArea.Results{
		fmt.Println(result.Name)
	}
	cfg.url["next"] = locationArea.Next
	cfg.url["previous"] = locationArea.Previous

	return nil
}



func commandMapB(cfg config) error {
	exists := cfg.url["previous"]
	if exists == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := http.Get(cfg.url["previous"])
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}

	var locationArea LocationAreaResponse
	if err = json.Unmarshal(body, &locationArea); err != nil {
		return err
	}

	for _, result := range locationArea.Results{
		fmt.Println(result.Name)
	}
	cfg.url["next"] = locationArea.Next
	cfg.url["previous"] = locationArea.Previous

	return nil
}
