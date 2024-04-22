package main

// TODO
// Add callnames, []string

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-zoox/fetch"
	"github.com/pat955/pokedex/internal/pokeapi"
)

type cliCommand struct {
	Name    string
	Desc    string
	Command func(pokeapi.Cache) error
	Config  *Config
}

type Config struct {
	prevURL    string
	currentURL string
}

func getCommands() map[string]cliCommand {
	currentLocationID := 0
	return map[string]cliCommand{
		"help": {
			Name: "Help",
			Desc: "Get info about this cli and other commands",
			Command: func(c pokeapi.Cache) error {
				commandHelp(&Config{})
				return nil
			},
		},
		"exit": {
			Name: "Exit",
			Desc: "Exit command line",
			Command: func(pokeapi.Cache) error {
				fmt.Println("Exiting")
				os.Exit(0)
				return nil
			},
		},

		"map": {
			Name: "Map",
			Desc: "Map of the next 20 area of pokemon",
			Command: func(c pokeapi.Cache) error {
				for i := 0; i < 20; i++ {
					currentLocationID++
					bytes := call(fmt.Sprintf("https://pokeapi.co/api/v2/location/%v", currentLocationID), c)

					l := LocationData{}
					json.Unmarshal(bytes, &l)
					fmt.Println(l.String())

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

type LocationData struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Region struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"region"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	Areas []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"areas"`
}

func (l LocationData) String() string {
	return fmt.Sprint(l.ID, " "+l.Name+" ", l.Region.Name+" ")
}
func (l LocationData) GetUrl() string {
	return l.Region.URL
}
func (l LocationData) GetId() int {
	return l.ID
}

func commandHelp(cfg *Config) error {
	fmt.Println("\nWelcome to the Pokedex!\nUsage:\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Desc)
	}
	fmt.Println()
	return nil
}
