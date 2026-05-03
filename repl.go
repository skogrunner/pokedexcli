package main

import (
	"strings"
	"fmt"
	"os"
	)

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}

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
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text))) 
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for key := range commands {
		command := commands[key]
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}