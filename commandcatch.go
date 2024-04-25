package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/pat955/pokedex/internal/pokeapi"
)

func commandCatch(cache pokeapi.Cache, pokedex Pokedex, currentArea, pokemonName string) error {
	if pokemonName == "" {
		return errors.New("catch error: No pokemon name provided")
	}
	areaData, err := dataToAreaData(cache, fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v/", currentArea))
	if err != nil {
		return err
	}
	if !areaData.CheckIfPokemonInArea(pokemonName) {
		return errors.New("Pokemon not found in your current area")
	}
	// check if already caught
	// several rounds of *click, *click*, italic *click* when caught with a timer to create suspense
	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v/", strings.ToLower(pokemonName))

	pokeData, err := dataToPokemonData(cache, endpoint)
	if err != nil {
		return err
	}
	pokeData.Nickname = pokeData.Name
	if pokeData.AreaCaughtIn != "" {
		return errors.New("catch error: This pokemon already caught, escaped or killed in this area. Come back later")
	}
	formattedName := color.HiCyanString(strings.Title(pokemonName))

	fmt.Println("Attempting to catch", formattedName, "...")
	catchLoop(pokeData, pokedex, currentArea, formattedName)
	return nil
}

func catchLoop(pokeData *PokemonData, pokedex Pokedex, currentArea, name string) {
	// pokeball chances
	// add countdown
	scanner := bufio.NewScanner(os.Stdin)
	caught := capture()
	if caught {
		pokeData.AreaCaughtIn = currentArea
		fmt.Println("You caught", name+"!\nGive", name, "a nickname? (y/n)")
		if scanner.Scan() {
			answer := scanner.Text()
			if answer == "y" {
				if scanner.Scan() {
					pokeData.Nickname = scanner.Text()
					fmt.Println("Nickname", color.HiMagentaString(pokeData.Nickname), "given to", name)
				}
			}
		}
		pokedex.Add(pokeData)
		pokedex.PrintOutCurrentPokemon()
	} else {
		fmt.Println("Failed to catch", name+"!\nTry again? (y/n)")
		if scanner.Scan() {
			answer := scanner.Text()
			if answer == "y" {
				catchLoop(pokeData, pokedex, currentArea, name)
			}

		}
	}
}
func capture() bool {
	boldPrint := color.New(color.Bold).PrintlnFunc()
	rand.Seed(time.Now().UnixMilli())

	if rand.Intn(1000) >= 100 {
		time.Sleep(500 * time.Millisecond)
		boldPrint(color.HiBlackString("*click*"))
		time.Sleep(1 * time.Second)

		if rand.Intn(1000) >= 100 {
			boldPrint(color.HiBlackString("*click*"))
			time.Sleep(1 * time.Second)

			if rand.Intn(1000) >= 200 {
				boldPrint(color.HiBlackString("*click*"))
				time.Sleep(1000 * time.Millisecond)

				return true
			}
			return false
		}
		return false
	}
	return false
}
