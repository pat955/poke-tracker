package main

// TODO
// Change to cases instead of decr strings
// Add Save Game
//

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/pat955/pokedex/internal/pokeapi"
)

func main() {
	fmt.Println("Loading...")
	scanner := bufio.NewScanner(os.Stdin)
	pokeInventory := newPokedex()
	inventory := NewItemInventory()
	inventory.AddStarterItems()

	cache := pokeapi.NewCache(30 * time.Minute)
	commands := getCommands(cache, pokeInventory, inventory)

	for {
		fmt.Print(color.HiGreenString("PokeCLI >>> "))

		var cmd, args string

		if scanner.Scan() {
			line := strings.Split(scanner.Text(), " ")
			cmd = line[0]
			if len(line) != 1 {
				args = line[1]
			}
		}
		if cmd == "" {
			continue
		}
		command, ok := commands[cmd]
		if !ok {
			fmt.Println(errors.New("unknown command, type help for commands"))
			continue
		}
		err := command.Command(args)
		if err != nil {
			fmt.Println(err)
		}

	}
}

// Update the CLI to support the "up" arrow to cycle through previous commands
// Simulate battles between pokemon
// Add more unit tests
// Refactor your code to organize it better and make it more testable
// Keep pokemon in a "party" and allow them to level up
// Allow for pokemon that are caught to evolve after a set amount of time
// Persist a user's Pokedex to disk so they can save progress between sessions
// Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
// Random encounters with wild pokemon
// Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon
