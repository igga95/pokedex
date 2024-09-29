package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	availableCommands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, v := range availableCommands {
		fmt.Printf(" - %s: %s\n", v.name, v.description)
	}
	fmt.Println("")
	return nil
}
