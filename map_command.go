package main

import (
	"encoding/json"
	"fmt"

	"github.com/AggroSec/Go-Pokedex/internal/pokeAPI"
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

func commandMap(conf *Config, param string) error {
	var getURL string
	if conf.Next != nil {
		getURL = *conf.Next
	} else {
		getURL = locationURL
	}

	data, err := pokeAPI.CreatePokeAPIRequest(getURL, conf.PokeCache)
	if err != nil {
		return err
	}

	var results GetLocations
	err = json.Unmarshal(data, &results)
	if err != nil {
		return err
	}

	if results.Next != nil {
		conf.Next = results.Next
	}
	if results.Previous != nil {
		conf.Previous = results.Previous
	}

	for _, location := range results.Results {
		fmt.Printf(" - %s\n", *location.Name)
	}

	return nil
}

func commandMapb(conf *Config, param string) error {
	var getURL string
	if conf.Previous != nil {
		getURL = *conf.Previous
	} else {
		getURL = locationURL
	}

	data, err := pokeAPI.CreatePokeAPIRequest(getURL, conf.PokeCache)
	if err != nil {
		return err
	}

	var results GetLocations
	err = json.Unmarshal(data, &results)
	if err != nil {
		return err
	}

	if results.Next != nil {
		conf.Next = results.Next
	}
	if results.Previous != nil {
		conf.Previous = results.Previous
	}

	for _, location := range results.Results {
		fmt.Printf(" - %s\n", *location.Name)
	}

	return nil
}
