package main

import "fmt"

func callbackHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("")
	fmt.Println("Usage: <command> <arguments>")
	fmt.Println("Available commands:")

	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
