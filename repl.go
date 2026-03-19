package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Next     *string
	Previous *string
}

func cleanInput(text string) []string {
	input := strings.Fields(strings.ToLower(text))
	return input
}

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		clInput := cleanInput(input)
		commands := getCommands()
		command, ok := commands[clInput[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
