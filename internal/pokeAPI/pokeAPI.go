package pokeAPI

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GetLocations struct {
	Count    int               `json:"count"`
	Next     *string           `json:"next"`
	Previous *string           `json:"previous"`
	Results  []LocationResults `json:"results"`
}

type LocationResults struct {
	Name *string `json:"name"`
	URL  *string `json:"url"`
}

func CreatePokeAPIRequest(url, command string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if command == "map" {
		results, err := handleMapCommand(body)
		if err != nil {
			return nil, err
		}
		return results, nil
	}

	return nil, errors.New("http command not found.")
}

func handleMapCommand(body []byte) ([]string, error) {
	results := []string{}

	var locations GetLocations
	if err := json.Unmarshal(body, &locations); err != nil {
		return nil, err
	}

	if locations.Next == nil {
		results = append(results, "null")
	} else {
		results = append(results, *locations.Next)
	}

	if locations.Previous == nil {
		results = append(results, "null")
	} else {
		results = append(results, *locations.Previous)
	}

	for _, location := range locations.Results {
		results = append(results, *location.Name)
	}

	return results, nil
}
