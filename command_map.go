package main

import (
	"fmt"

	"github.com/cole-maxwell1/pokedexcli/internal"
)

func callbackAreaMap() error {
	pokeapiClient := pokeapi.NewClient()

	areas, err := pokeapiClient.GetLocationAreas()
	if err != nil {
		return err
	}

	for _, area := range areas.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}
