package main

// TODO
// Add callnames, []string
// find out why id 21 is empty

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/go-zoox/fetch"
	"github.com/pat955/pokedex/internal/pokeapi"
)

type cliCommand struct {
	Name    string
	Desc    string
	Command func(pokeapi.Cache, Pokedex, string) error
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
			Command: func(c pokeapi.Cache, i Pokedex, s string) error {
				commandHelp()
				return nil
			},
		},
		"exit": {
			Name: "Exit",
			Desc: "Exit command line",
			Command: func(c pokeapi.Cache, i Pokedex, s string) error {
				fmt.Println("Exiting")
				os.Exit(0)
				return nil
			},
		},

		"map": {
			Name: "Map",
			Desc: "Map of the next 10 area of pokemon",
			Command: func(c pokeapi.Cache, i Pokedex, s string) error {
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
			Command: func(c pokeapi.Cache, i Pokedex, s string) error {
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
			Desc: "explore current area, called with: >>> explore <areaName>",
			Command: func(c pokeapi.Cache, i Pokedex, areaName string) error {
				if areaName == "" {
					return errors.New("explore error: No location provided.\nUse map command to see accepted areas")
				}
				fmt.Println("Exploring", areaName, "...")
				bytes := call(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v-area/", areaName), c)
				area := AreaData{}
				json.Unmarshal(bytes, &area)
				fmt.Println("Found Pokemon: ")
				for _, pokemon := range area.GetPokemonEncounters() {
					_, ok := i.Pokemon[pokemon.Name]
					if !ok {
						fmt.Println("-", pokemon.Name)
					}
				}
				return nil
			},
		},
		"catch": {
			Name: "Catch Pokemon",
			Desc: "Catch pokemon using this command after exploring area",
			Command: func(c pokeapi.Cache, i Pokedex, pokemonName string) error {
				if pokemonName == "" {
					return errors.New("catch error: No pokemon name provided.\nUse explore command to see pokemon in your area")
				}
				// check if already caught
				// several rounds of *click, *click*, italic *click* when caught with a timer to create suspense
				fmt.Println("Attempting to catch", pokemonName, "...")

				bytes := call(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v/", strings.ToLower(pokemonName)), c)
				pokemondata := &PokemonData{}
				fmt.Println("!" + pokemondata.Nickname + "!")
				json.Unmarshal(bytes, &pokemondata)

				rand.Seed(time.Now().UnixMilli())
				if rand.Intn(1000) >= 500 {
					fmt.Println("You caught", pokemonName+"!")
					fmt.Println("Give", pokemonName, "a nickname? (y/n)")
					scanner := bufio.NewScanner(os.Stdin)
					if scanner.Scan() {
						answer := scanner.Text()
						if answer == "y" {
							if scanner.Scan() {
								pokemondata.Nickname = scanner.Text()
								fmt.Println(&pokemondata.Nickname)
							}

						}

					}
					i.Add(*pokemondata)
				} else {
					fmt.Println("Failed to catch", pokemonName+"!")
				}
				i.PrintOutMyPokemon()
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
func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!\nUsage: ")
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Desc)
	}
	fmt.Println()
	return nil
}
