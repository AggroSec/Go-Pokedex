package main

import (
	"fmt"

	"github.com/AggroSec/Go-Pokedex/internal/pokeAPI"
)

const (
	locationURL = "https://pokeapi.co/api/v2/location-area/"
)

func commandMap(conf *Config) error {
	locationResults := pokeAPI.CreatePokeAPIRequest(locationURL, "map")
	fmt.Println(locationResults)
	return nil
}

func commandMapb(conf *Config) error {
	return nil
}
