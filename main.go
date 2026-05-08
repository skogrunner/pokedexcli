package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
	"github.com/skogrunner/pokedexcli/internal/pokecache"
	)

type Config struct {
	cache pokecache.Cache
	previous string
	next string
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*Config, []string) error
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	config := &Config {
    	cache : pokecache.NewCache(120 * time.Second),
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
			if len(userInput) == 0 {
    			processCommand(command.callback, config, []string{})
			} else {
    			processCommand(command.callback, config, userInput[1:])
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func processCommand(callback func(*Config, []string) error, values *Config, args []string) {
	err := callback(values, args)
	if err != nil {
		fmt.Println("user input failed. Error:", err)
	}
}

