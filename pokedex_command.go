package main

import (
	"fmt"
)

func commandPokedex(conf *Config, param string) error {
	if len(conf.pokedex) == 0 {
		fmt.Println("Your Pokedex is empty. Try catching some Pokemon first!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range conf.pokedex {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
