package main

import (
	"fmt"

	"github.com/AggroSec/Go-Pokedex/internal/pokeAPI"
)

const (
	locationURL = "https://pokeapi.co/api/v2/location-area/"
)

func commandMap(conf *Config) error {
	var getURL string
	if conf.Next != nil {
		getURL = *conf.Next
	} else {
		getURL = locationURL
	}

	locationResults, err := pokeAPI.CreatePokeAPIRequest(getURL, "map")
	if err != nil {
		return err
	}

	if locationResults[0] != "null" {
		conf.Next = &locationResults[0]
	}
	if locationResults[1] != "null" {
		conf.Previous = &locationResults[1]
	}

	for _, location := range locationResults[2:] {
		fmt.Println(location)
	}

	return nil
}

func commandMapb(conf *Config) error {
	var getURL string
	if conf.Previous != nil {
		getURL = *conf.Previous
	} else {
		getURL = locationURL
	}

	locationResults, err := pokeAPI.CreatePokeAPIRequest(getURL, "map")
	if err != nil {
		return err
	}

	if locationResults[0] != "null" {
		conf.Next = &locationResults[0]
	}
	if locationResults[1] != "null" {
		conf.Previous = &locationResults[1]
	}

	for _, location := range locationResults[2:] {
		fmt.Println(location)
	}

	return nil
}
