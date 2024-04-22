package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/pat955/pokedex/internal/pokeapi"
)

func main() {
	commands := getCommands()
	cache := pokeapi.NewCache(100 * time.Second)
	for {
		fmt.Print(color.HiGreenString("Pokedex > "))

		var input string

		fmt.Scanln(&input)
		command, ok := commands[input]
		if !ok {
			fmt.Println(errors.New("unknown command, type help for commands"))
			continue
		}
		command.Command(cache)

	}
}
