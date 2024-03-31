package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackAreaMap(state *State) error {

	areas, err := state.pokeapiClient.GetLocationAreas(state.nextLocationUrl)
	if err != nil {
		return err
	}

	state.nextLocationUrl = areas.Next
	state.previousLocationUrl = areas.Previous

	for _, area := range areas.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}

func callbackAreaMapBack(state *State) error {

	if state.previousLocationUrl == nil {
		return errors.New("you are on the first page of locations")
	}

	areas, err := state.pokeapiClient.GetLocationAreas(state.previousLocationUrl)
	if err != nil {
		log.Fatal(err)
	}

	state.nextLocationUrl = areas.Next
	state.previousLocationUrl = areas.Previous

	for _, area := range areas.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}