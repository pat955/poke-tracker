package main

// TODO
// Add callnames, []string
// start command?

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/go-zoox/fetch"
	"github.com/pat955/pokedex/internal/pokeapi"
)

type cliCommand struct {
	Name    string
	Desc    string
	Command func(arguments string) error
}

// Get the names of all commands, execute with x.command(arg, arg, arg)
func getCommands(cache pokeapi.Cache, pokedex Pokedex, inventory ItemInventory) map[string]cliCommand {
	// 1 is the starting id in the api instead of 0.
	var currentArea string
	boldPrint := color.New(color.Bold).PrintlnFunc()
	currentLocationID := 1
	return map[string]cliCommand{
		"help": {
			Name: "Help",
			Desc: "Get the description of all available commands",
			Command: func(_ string) error {
				commandHelp()
				return nil
			},
		},
		"exit": {
			Name: "Exit",
			Desc: "Exit command line",
			Command: func(_ string) error {
				fmt.Println("See you next time!\nExiting...")
				os.Exit(0)
				return nil
			},
		},
		"map": {
			Name: "Map",
			Desc: "Get the next 2 location and their areas. The cyan name is the location. Explore the areas. eks: >>> explore eterna-city-west-gate",
			Command: func(_ string) error {
				boldPrint(color.GreenString("EXPLORABLE AREAS:"))
				for i := 0; i < 2; i++ {
					endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/location/%v", currentLocationID)
					var locData LocationData
					data, err := checkAndCall(cache, endpoint, &locData)
					if err != nil {
						return err
					}
					data.PrintInfo()
					fmt.Println()
					currentLocationID++
				}
				return nil
			},
		},
		"mapb": {
			Name: "Map Back",
			Desc: "Get the previous 2 locations",
			Command: func(_ string) error {
				if currentLocationID <= 3 {
					return errors.New("location error: You're at the start, unable to go back further. Type map to go forward")
				}
				currentLocationID -= 2
				boldPrint(color.GreenString("EXPLORABLE AREAS:"))

				// print backwards instead of 10 9 8 like it does now, maybe...
				for i := 0; i < 5; i++ {
					currentLocationID--
					endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/location/%v", currentLocationID)

					var locData LocationData
					data, err := checkAndCall(cache, endpoint, &locData)
					if err != nil {
						return err
					}
					data.PrintInfo()
					fmt.Println()
				}
				return nil
			},
		},
		"explore": {
			Name: "Explore Area",
			Desc: "explore current area, called with: >>> explore <areaName>",
			Command: func(areaName string) error {
				if areaName == "" {
					return errors.New("explore error: No location provided.\nUse map command to see accepted areas")
				}
				currentArea = areaName
				endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v/", areaName)
				data, err := dataTypeToAreaData(cache, endpoint)
				if err != nil {
					return err
				}
				data.Explored = true
				fmt.Println("Exploring", areaName, "...")
				fmt.Println("Found Pokemon: ")
				// add random chance so there are fewer pokemon
				for _, pokemon := range data.GetEncounters() {
					pokemondata, found := pokedex.Pokedex[pokemon.Name]
					if !found || pokemondata.AreaCaughtIn != areaName {
						fmt.Println("- " + pokemon.Name)
						continue
					}
					//fmt.Println(color.BlackString("- " + pokemon.Name))

				}
				return nil
			},
		},
		"catch": {
			// Add area checking so you cannot catch mewtwo in region 1...
			Name: "Catch Pokemon",
			Desc: "Catch pokemon using this command after exploring area",
			Command: func(pokemonName string) error {
				return commandCatch(cache, pokedex, currentArea, pokemonName)
			},
		},
		"inspect": {
			Name: "Inspect",
			Desc: "Inspect a pokemon or an item",
			Command: func(pokemonName string) error {
				return pokedex.Inspect(pokemonName)
			},
		},
		"pokedex": {
			Name: "Pokedex",
			Desc: "See all the pokemon you've caught so far",
			Command: func(_ string) error {
				pokedex.PrintOutMyPokemon()
				return nil
			},
		},
		"cache": {
			Name: "Check Cache",
			Desc: "Check Cache for debugging",
			Command: func(_ string) error {
				cache.Print()
				return nil
			},
		},
		"inventory": {
			Name: "Check Inventory",
			Desc: "Check inventory and use items",
			Command: func(_ string) error {
				inventory.PrintOutItems()
				return nil
			},
		},
		"buy": {
			Name: "Buy Items",
			Desc: "Buy items like pokeballs, moves and more",
			Command: func(_ string) error {
				return buyItems(cache, inventory)
			},
		},
	}
}

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!\nUsage: ")
	// placeholders, pokeapi.cache and pokedex{} not used at all
	for _, cmd := range getCommands(pokeapi.Cache{}, Pokedex{}, ItemInventory{}) {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Desc)
	}
	fmt.Println()
	return nil
}

func call(endpoint string) ([]byte, error) {
	response, err := fetch.Get(endpoint)
	if err != nil {
		return nil, err
	}
	if len(response.Body) < 9 {
		return nil, errors.New("404 Not Found: Check if you've typed in the name correctly")
	}
	return response.Body, nil
}
func checkAndCall(cache pokeapi.Cache, endpoint string, dataStruct pokeapi.DataTypes) (pokeapi.DataTypes, error) {
	data, found := cache.Get(endpoint)
	if found {
		return data, nil
	}
	bytes, err := call(endpoint)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bytes, &dataStruct)
	cache.Add(endpoint, dataStruct)
	return dataStruct, nil
}
