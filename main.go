package main

import (
	"time"

	pokeapi "github.com/cole-maxwell1/pokedexcli/internal"
)

type State struct {
	pokeapiClient       pokeapi.PokeapiClient
	nextLocationUrl     *string
	previousLocationUrl *string
}

func main() {
	conf := State{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	startRepl(&conf)
}
