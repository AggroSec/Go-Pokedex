package main

import (
	"fmt"
)

func commandInspect(conf *Config, param string) error {
	pokemon, exists := conf.pokedex[param]
	if !exists {
		fmt.Printf("You don't have a %s in your Pokedex. Try catching it first!\n", param)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, s := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}
