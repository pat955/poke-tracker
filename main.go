package main

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	commands := getCommands()

	for {
		fmt.Print(color.HiGreenString("Pokedex > "))

		var input string

		fmt.Scanln(&input)
		x, ok := commands[input]
		if !ok {
			fmt.Println(errors.New("unknown command, type help for commands"))
			continue
		}
		x.Command()

	}
}
