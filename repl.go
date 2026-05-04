package main

import (
	"strings"
	"fmt"
	"os"
	"net/http"
	"io"
	"encoding/json"
	"errors"
	)

type Location struct {
	Name string
	URL  string
}

type Locations struct {
	Count int
	Next  string
	Previous string
	Results []Location
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

func commandExit(c *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config) error {
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

func commandMap(c *Config) error {
	url := c.next
	if len(c.next) == 0 {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return errors.New(fmt.Sprintf("Response failed with status code %d", res.StatusCode))
	}

	locations := Locations{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return err
	}
	c.next = locations.Next
	c.previous = locations.Previous
	fmt.Println("")
	for i := range locations.Results {
		fmt.Println(locations.Results[i].Name)
	}
	fmt.Println("")
	return nil
}

func commandMapb(c *Config) error {
	url := c.previous
	if len(c.previous) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return errors.New(fmt.Sprintf("Response failed with status code %d", res.StatusCode))
	}

	locations := Locations{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return err
	}
	c.next = locations.Next
	c.previous = locations.Previous
	fmt.Println("")
	for i := range locations.Results {
		fmt.Println(locations.Results[i].Name)
	}
	fmt.Println("")
	return nil
}