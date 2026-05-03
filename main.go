package main

import (
	"fmt"
	"bufio"
	"os"
	)

type Config struct {
	previous string
	next string
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*Config) error
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	config := &Config {
		previous: "",
		next: "",
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		ui := scanner.Text()
		userInput := cleanInput(ui)
		if len(userInput) == 0 {
			fmt.Println("no command specified")
			continue
		}
		if command, ok := commands[userInput[0]]; ok {
			processCommand(config, command.callback)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func processCommand(values *Config, callback func(*Config) error) {
	err := callback(values)
	if err != nil {
		fmt.Println("user input failed. Error:", err)
	}
}

