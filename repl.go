package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("pokedex > ")
		scanner.Scan()

		// Read the input from the user and sanitize it
		unsanitizedInput := scanner.Text()
		sanitizedCommand := sanitizeInput(unsanitizedInput)
		if len(sanitizedCommand) == 0 {
			continue
		}

		// Parse commands
		requestedCommand := sanitizedCommand[0]

		validCommands := getCommands()

		command, ok := validCommands[requestedCommand]

		if !ok {
			fmt.Println("Invalid command: ", requestedCommand)
			continue
		}
		command.callback()
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func sanitizeInput(input string) []string {
	lowerCaseInput := strings.ToLower(input)
	words := strings.Fields(lowerCaseInput)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    callbackExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "List the areas on the map!",
			callback:    callbackAreaMap,
		},
	}
}
