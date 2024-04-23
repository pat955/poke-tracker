package main

// TODO
// Add callnames, []string
// find out why id 21 is empty

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/go-zoox/fetch"
	"github.com/pat955/pokedex/internal/pokeapi"
)

type cliCommand struct {
	Name    string
	Desc    string
	Command func(pokeapi.Cache, string) error
	Config  *Config
}

type Config struct {
	CallName string
	PrevURL  string
}

func getCommands() map[string]cliCommand {
	currentLocationID := 0
	return map[string]cliCommand{
		"help": {
			Name: "Help",
			Desc: "Get info about this cli and other commands",
			Command: func(c pokeapi.Cache, s string) error {
				commandHelp(&Config{})
				return nil
			},
		},
		"exit": {
			Name: "Exit",
			Desc: "Exit command line",
			Command: func(c pokeapi.Cache, s string) error {
				fmt.Println("Exiting")
				os.Exit(0)
				return nil
			},
		},

		"map": {
			Name: "Map",
			Desc: "Map of the next 10 area of pokemon",
			Command: func(c pokeapi.Cache, s string) error {
				for i := 0; i < 10; i++ {
					currentLocationID++
					bytes := call(fmt.Sprintf("https://pokeapi.co/api/v2/location/%v", currentLocationID), c)

					l := LocationData{}
					json.Unmarshal(bytes, &l)
					fmt.Println(l.String())

				}
				return nil
			},
		},
		"mapb": {
			Name: "Map Back",
			Desc: "Get the previous 10 areas",
			Command: func(c pokeapi.Cache, s string) error {
				if currentLocationID == 0 {
					return errors.New("you're at the start, cannot go further back. type map to continue")
				}
				for i := 0; i < 10; i++ {
					currentLocationID--
					bytes := call(fmt.Sprintf("https://pokeapi.co/api/v2/location/%v", currentLocationID), c)

					l := LocationData{}
					json.Unmarshal(bytes, &l)
					fmt.Println(l.String())
				}
				return nil
			},
		},
		"explore": {
			Name: "Explore Area",
			Desc: "explore current area, called with: >>> explore <area_name>",
			Command: func(c pokeapi.Cache, area_name string) error {
				fmt.Println("Exploring", area_name, "...")
				bytes := call(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v-area/", area_name), c)
				area := AreaData{}
				json.Unmarshal(bytes, &area)
				fmt.Println("Found Pokemon: ")
				for _, pokemon := range area.GetPokemonEncounters() {
					fmt.Println("-", pokemon.Name)
				}
				return nil
			},
		},
	}
}

func call(endpoint string, c pokeapi.Cache) []byte {
	bytes, found := c.Get(endpoint)
	if found {
		return bytes
	}

	response, err := fetch.Get(endpoint)
	if err != nil {
		panic(err)
	}

	c.Add(endpoint, response.Body)
	return response.Body
}
func commandHelp(cfg *Config) error {
	fmt.Println("\nWelcome to the Pokedex!\nUsage: ")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Desc)
	}
	fmt.Println()
	return nil
}
