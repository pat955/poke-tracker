package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/pat955/pokedex/internal/pokeapi"
)

func commandCatch(cache pokeapi.Cache, player Profile, currentArea, pokemonName string) error {
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
	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v/", strings.ToLower(pokemonName))

	pokeData, err := dataToPokemonData(cache, endpoint)
	if err != nil {
		return err
	}
	if pokeData.AreaCaughtIn == currentArea {
		return errors.New("catch error: This pokemon already caught, escaped or killed in this area. Come back later")
	}
	pokeData.Nickname = pokeData.Name
	formattedName := color.HiCyanString(strings.Title(pokemonName))

	fmt.Println("Attempting to catch", formattedName, "...")
	catchLoop(pokeData, player, currentArea, formattedName)
	return nil
}

func catchLoop(pokeData *PokemonData, player Profile, currentArea, name string) {
	err := player.Inventory.Items["poke-ball"].UseItem(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	// pokeball chances
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
		player.Pokedex.Add(pokeData)
		player.Pokedex.PrintOutCurrentPokemon()
	} else {
		fmt.Println("Failed to catch", name+"!\nTry again? (y/n)")
		if scanner.Scan() {
			answer := scanner.Text()
			if answer == "y" {
				catchLoop(pokeData, player, currentArea, name)
			}
		}
	}
}
func capture() bool {
	boldPrint := color.New(color.Bold).PrintlnFunc()
	boldPrint(color.HiBlackString("*Poke ball thrown!*"))
	time.Sleep(300 * time.Millisecond)
	// 95% chance of success
	if chanceCheck(0.95) {
		boldPrint(color.HiBlackString("*click*"))
		time.Sleep(800 * time.Millisecond)

		if chanceCheck(0.90) {
			boldPrint(color.HiBlackString("*click*"))
			time.Sleep(800 * time.Millisecond)

			if chanceCheck(0.85) {
				boldPrint(color.HiBlackString("*click*"))
				time.Sleep(800 * time.Millisecond)

				return true
			}
		}
	}

	return false
}
func chanceCheck(precentage float64) bool {
	return precentage > rand.Float64()
}
