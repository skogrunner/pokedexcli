package main

import (
	"fmt"
	"bufio"
	"os"
	)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
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
			processCommand(command.callback)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func processCommand(callback func() error) {
	err := callback()
	if err != nil {
		fmt.Println("user input failed. Error:", err)
	}
}

