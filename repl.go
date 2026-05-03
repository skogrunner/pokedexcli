package main

import (
	"strings"
	"fmt"
	"os"
	)

func getCommands() map[string]cliCommand {
	return 	map[string]cliCommand {
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		commandHelp,	
		},
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback:		commandExit,
		},
		"map": {
			name:			"map",
			description:	"Displays next 20 location areas",
			callback:		commandMap,
		},
		"mapb": {
			name:			"mapb",
			description:	"Displays previous 20 location areas",
			callback:		commandMapb,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text))) 
}

func commandExit(*Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*Config) error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for key := range commands {
		command := commands[key]
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println("")
	return nil
}

func commandMap(*Config) error {
	return nil
}

func commandMapb(*Config) error {
	return nil
}