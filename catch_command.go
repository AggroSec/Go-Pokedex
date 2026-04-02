package main

import (
	"encoding/json"
	"fmt"

	"math/rand"

	"github.com/AggroSec/Go-Pokedex/internal/pokeAPI"
)

type Pokemon struct {
	Name           string  `json:"name"`
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type Types struct {
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
}

func commandCatch(conf *Config, param string) error {
	fullURL := pokemonURL + param

	data, err := pokeAPI.CreatePokeAPIRequest(fullURL, conf.PokeCache)
	if err != nil {
		return err
	}

	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	catchChance := 0.5
	baseExp := pokemon.BaseExperience

	switch {
	case baseExp <= 50:
		catchChance = 0.8
	case baseExp <= 100:
		catchChance = 0.6
	case baseExp <= 150:
		catchChance = 0.5
	case baseExp <= 200:
		catchChance = 0.4
	case baseExp <= 250:
		catchChance = 0.3
	default:
		catchChance = 0.15
	}

	roll := rand.Float64()
	if roll <= catchChance {
		fmt.Println(catchSuccessString(pokemon.Name))
		conf.pokedex[pokemon.Name] = pokemon
		fmt.Printf("%s has been added to your Pokedex! use 'inspect' to view it.\n", pokemon.Name)
	} else {
		fmt.Println(catchFailureString(pokemon.Name))
	}

	return nil
}

func catchSuccessString(pokename string) string {
	// returns a random success string to display when pokemon is caught.

	var catchSuccessMessages = []string{
		"You caught %s!",
		"Gotcha! %s was caught!",
		"All right! %s was caught!",
		"%s was successfully caught!",
		"Nice! %s was caught!",
	}

	return fmt.Sprintf(catchSuccessMessages[rand.Intn(len(catchSuccessMessages))], pokename)
}

func catchFailureString(pokename string) string {
	// returns a random failure string to display when pokemon escapes.

	var catchMissMessages = []string{
		"Oh no! %s broke free!",
		"%s wriggled out of the Poké Ball!",
		"%s was too strong and escaped!",
		"Aww... %s broke free!",
		"%s popped out of the ball!",
		"%s dodged the Poké Ball!",
	}

	return fmt.Sprintf(catchMissMessages[rand.Intn(len(catchMissMessages))], pokename)
}
