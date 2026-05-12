package main

import (
	"strings"
	"fmt"
	"os"
	"net/http"
	"io"
	"encoding/json"
	"errors"
	"time"
	"math/rand"
	"github.com/skogrunner/pokedexcli/internal/pokecache"
	)

type LocationArea struct {
	Name string
	URL  string
}

type LocationAreas struct {
	Count int
	Next  string
	Previous string
	Results []LocationArea
}

type PokemonEncounter struct {
	Name string
	URL  string
}

type PokemonEncounters struct {
	Pokemon PokemonEncounter
}

type PE struct {
	Pokemon_Encounters []PokemonEncounters
}

type PokemonStatName struct {
	Name   string
	URL    string
}

type PokemonStat struct {
	Base_Stat  int
	Effort     int
	Stat  PokemonStatName
}

type PokemonTypeName struct {
	Name  string
}

type PokemonType struct {
	Type PokemonTypeName
}

type Pokemon struct {
	Name string
	Base_Experience int
	Height int
	Weight int
	Stats  []PokemonStat
	Types  []PokemonType
}

func getCommands() map[string]cliCommand {
	return 	map[string]cliCommand {
		"help": {
			name:			"help",
			description:	"Display a help message",
			callback:		commandHelp,	
		},
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback:		commandExit,
		},
		"map": {
			name:			"map",
			description:	"Display next 20 location areas",
			callback:		commandMap,
		},
		"mapb": {
			name:			"mapb",
			description:	"Display previous 20 location areas",
			callback:		commandMapb,
		},
		"explore": {
			name:			"explore",
			description:	"Display all Pokemon in a location area",
			callback:		commandExplore,
		},
		"catch": {
			name:			"catch",
			description:	"Add Pokemon to user's Pokedex",
			callback:		commandCatch,
		},
		"inspect": {
			name:			"inspect",
			description:	"Display info for a Pokemon in the Pokedex",
			callback:		commandInspect,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text))) 
}

func commandExit(c *Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *Config, args []string) error {
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

func commandMap(c *Config, args []string) error {
	url := c.next
	if len(c.next) == 0 {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	var body []byte
    body, ok :=pokecache.Get(c.cache, url)
	if !ok {
    	res, err := http.Get(url)
	    if err != nil {
		    return err
	    }
	    body, err = io.ReadAll(res.Body)
	    defer res.Body.Close()
	    if err != nil {
		    return err
	    }
	    if res.StatusCode > 299 {
		    return errors.New(fmt.Sprintf("Response failed with status code %d", res.StatusCode))
	    }
		pokecache.Add(c.cache, url, body)
	} 
	locationAreas := LocationAreas{}
	err := json.Unmarshal(body, &locationAreas)
	if err != nil {
		return err
	}
	c.next = locationAreas.Next
	c.previous = locationAreas.Previous
	fmt.Println("")
	for i := range locationAreas.Results {
		fmt.Println(locationAreas.Results[i].Name)
	}
	fmt.Println("")
	return nil
}

func commandMapb(c *Config, args []string) error {
	url := c.previous
	if len(c.previous) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}

	var body []byte
    body, ok :=pokecache.Get(c.cache, url)
	if !ok {
    	res, err := http.Get(url)
	    if err != nil {
		    return err
	    }
	    body, err = io.ReadAll(res.Body)
	    defer res.Body.Close()
	    if err != nil {
		    return err
	    }
	    if res.StatusCode > 299 {
		    return errors.New(fmt.Sprintf("Response failed with status code %d", res.StatusCode))
	    }
		pokecache.Add(c.cache, url, body)
	} 
	locationAreas := LocationAreas{}
	err := json.Unmarshal(body, &locationAreas)
	if err != nil {
		return err
	}
	c.next = locationAreas.Next
	c.previous = locationAreas.Previous
	fmt.Println("")
	for i := range locationAreas.Results {
		fmt.Println(locationAreas.Results[i].Name)
	}
	fmt.Println("")
	return nil
}

func commandExplore(c *Config, args []string) error {
	if len(args) != 1 {
		fmt.Println("Usage is: EXPLORE <location area>")
		return nil
	}
	url := "https://pokeapi.co/api/v2/location-area/" + args[0] + "/"
	var body []byte
    body, ok :=pokecache.Get(c.cache, url)
	if !ok {
    	res, err := http.Get(url)
		if res.StatusCode == 404 {
			fmt.Println("Location area", args[0], "not found")
		}
	    if err != nil {
		    return err
	    }
	    body, err = io.ReadAll(res.Body)
	    defer res.Body.Close()
	    if err != nil {
		    return err
	    }
	    if res.StatusCode > 299 {
		    return errors.New(fmt.Sprintf("Response failed with status code %d", res.StatusCode))
	    }
		pokecache.Add(c.cache, url, body)
	}
	fmt.Println("Exploring", args[0]) 
	pe := PE{}
	err := json.Unmarshal(body, &pe)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for i := range len(pe.Pokemon_Encounters) {
		fmt.Println(" -", pe.Pokemon_Encounters[i].Pokemon.Name)
	}
	return nil
}

func commandCatch(c *Config, args []string) error {
	if len(args) != 1 {
		fmt.Println("Usage is: CATCH <Pokemon name>")
		return nil
	}
	_, ok := c.pokedex[args[0]]
	if ok {
		fmt.Println(args[0], "has already been caught.")
		return nil
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + args[0] + "/"
    var body []byte
    body, ok =pokecache.Get(c.cache, url)
	if !ok {
    	res, err := http.Get(url)
		if res.StatusCode == 404 {
			fmt.Println("Pokemon name", args[0], "not found")
		}
	    if err != nil {
		    return err
	    }
	    body, err = io.ReadAll(res.Body)
	    defer res.Body.Close()
	    if err != nil {
		    return err
	    }
	    if res.StatusCode > 299 {
		    return errors.New(fmt.Sprintf("Response failed with status code %d", res.StatusCode))
	    }
		pokecache.Add(c.cache, url, body)
	}
	pokemon := Pokemon{}
	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	rand.Seed(time.Now().UnixNano())
	if (400.0 - float64(pokemon.Base_Experience)) / 400.0 > rand.Float64() {
        fmt.Println(args[0], "was caught!")
		c.pokedex[args[0]] = pokemon
	} else {
		fmt.Println(args[0], "escaped!")
	}
	return nil
}

func commandInspect(c *Config, args []string) error {
	if len(args) != 1 {
		fmt.Println("Usage is: INSPECT <Pokemon name>")
		return nil
	}
	val, ok := c.pokedex[args[0]]
	if !ok {
		fmt.Println(args[0], "you have not caught that pokemon")
		return nil
	}
	fmt.Println("Name:", val.Name)
	fmt.Println("Height: ", val.Height)
	fmt.Println("Weight: ", val.Weight)
	fmt.Println("Stats:")
	for i := range len(val.Stats) {
		fmt.Printf("  -%s: %d\n", val.Stats[i].Stat.Name,val.Stats[i].Base_Stat)
	}
	fmt.Println("Types:")
	for j := range len(val.Types) {
		fmt.Println("  -", val.Types[j].Type.Name)
	}
	return nil
}