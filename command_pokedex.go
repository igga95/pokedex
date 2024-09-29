package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokemon in Pokedex:")
	for _, pokemon := range cfg.caugthPokemon {
		fmt.Printf(" - %s \n", pokemon.Name)
	}

	return nil
}
