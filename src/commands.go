package main

// TODO
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
func getCommands(cache pokeapi.Cache, player Profile) map[string]cliCommand {
	// 1 is the starting id in the api instead of 0.
	var currentArea string
	boldPrint := color.New(color.Bold).PrintlnFunc()
	currentLocationID := 1

	return map[string]cliCommand{
		"help": {
			Name: "help",
			Desc: "Get the description of all available commands",
			Command: func(_ string) error {
				commandHelp()
				return nil
			},
		},
		"exit": {
			Name: "exit",
			Desc: "Exit command line",
			Command: func(_ string) error {
				fmt.Println("See you next time!\nExiting...")
				os.Exit(0)
				return nil
			},
		},
		"map": {
			Name: "map",
			Desc: "Get the next 2 location and their areas. The cyan name is the location.\nExplore the areas. eks: >>> explore eterna-city-west-gate",
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
			Name: "mapb",
			Desc: "Map Back. Get the previous 2 locations",
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
			Name: "explore <area_name>",
			Desc: "Explore current area, called with: >>> explore <areaName>",
			Command: func(areaName string) error {
				if areaName == "" {
					return errors.New("explore error: No area name provided.\nUse map command to see accepted areas")
				}
				currentArea = areaName
				endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v/", areaName)
				data, err := dataToAreaData(cache, endpoint)
				if err != nil {
					return err
				}
				data.Explored = true
				fmt.Println("Exploring", areaName, "...")
				fmt.Println("Found Pokemon: ")
				// add random chance so there are fewer pokemon
				for _, pokemon := range data.GetEncounters() {
					pokemondata, found := player.Pokedex.Pokedex[pokemon.Name]
					if !found || pokemondata.AreaCaughtIn != areaName {
						fmt.Println("- " + pokemon.Name)
						continue
					}
					//fmt.Println(color.BlackString("- " + pokemon.Name))

				}
				return nil
			},
		},
		"explore-location": {
			Name: "explore-location <locationName>",
			Desc: "explore an entire location rather than a small area",
			Command: func(locationName string) error {
				if locationName == "" {
					return errors.New("explore error: No location provided.\nUse map command to see accepted locations")
				}
				fmt.Println("Starting to explore", locationName, "...")
				endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/location/%v/", locationName)
				locData, err := dataToLocationData(cache, endpoint)
				if err != nil {
					return err
				}
				for i, area := range locData.Areas {
					fmt.Println(i, area)
				}
				return nil
			},
		},
		"catch": {
			// Add area checking so you cannot catch mewtwo in region 1...
			Name: "catch <pokemon_name>",
			Desc: "Catch pokemon using this command after exploring area",
			Command: func(pokemonName string) error {
				return commandCatch(cache, player, currentArea, pokemonName)
			},
		},
		"inspect": {
			Name: "inspect <pokemon_name>",
			Desc: "Inspect a pokemon in your inventory",
			Command: func(pokemonName string) error {
				return player.Pokedex.Inspect(pokemonName)
			},
		},
		"pokedex": {
			Name: "pokedex",
			Desc: "See all the pokemon you've caught so far",
			Command: func(_ string) error {
				player.Pokedex.PrintOutMyPokemon()
				return nil
			},
		},
		"cache": {
			Name: "cache",
			Desc: "Check Cache for debugging reasons",
			Command: func(_ string) error {
				cache.Print()
				return nil
			},
		},
		"inventory": {
			Name: "inventory",
			Desc: "Check inventory and use items",
			Command: func(_ string) error {
				player.Inventory.PrintOutItems()
				return nil
			},
		},
		"shop": {
			Name: "shop",
			Desc: "Buy items like pokeballs, moves and more",
			Command: func(_ string) error {
				return buyItems(cache, player.Inventory)
			},
		},
	}
}

func commandHelp() error {
	boldRed := color.New(color.Bold, color.FgRed).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	fmt.Println(bold("Welcome to my PokeCLI!\n"))
	fmt.Println(("------------Available Commands------------"))
	// placeholders, pokeapi.cache and pokedex{} not used at all
	// add order
	for _, cmd := range getCommands(pokeapi.Cache{}, Profile{}) {
		fmt.Printf("%s  %s\n", boldRed(cmd.Name), cmd.Desc)
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
