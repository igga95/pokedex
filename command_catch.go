package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	res, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	randNum := rand.Intn(res.BaseExperience)
	if randNum > int(0.6*float64(res.BaseExperience)) {
		return fmt.Errorf("failed to catch %s!\n", pokemonName)
	}

	cfg.caugthPokemon[pokemonName] = res
	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}
