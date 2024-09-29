package main

import (
	"time"

	"github.com/igga95/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caugthPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caugthPokemon: map[string]pokeapi.Pokemon{},
	}
	startRepl(&cfg)
}
