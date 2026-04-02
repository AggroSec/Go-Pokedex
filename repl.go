package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/AggroSec/Go-Pokedex/internal/pokecache"
)

const (
	locationURL = "https://pokeapi.co/api/v2/location-area/"
	pokemonURL  = "https://pokeapi.co/api/v2/pokemon/"
)

type Config struct {
	Next      *string
	Previous  *string
	PokeCache pokecache.Cache
	pokedex   map[string]Pokemon
}

func cleanInput(text string) []string {
	input := strings.Fields(strings.ToLower(text))
	return input
}

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{
		Next:      nil,
		Previous:  nil,
		PokeCache: pokecache.NewCache(time.Second * 60),
		pokedex:   make(map[string]Pokemon),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		clInput := cleanInput(input)
		if len(clInput) == 0 {
			continue
		}
		commands := getCommands()
		command, ok := commands[clInput[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		parameter := ""
		if len(clInput) > 1 {
			parameter = clInput[1]
		}
		err := command.callback(cfg, parameter)
		if err != nil {
			fmt.Println(err)
		}
	}
}
