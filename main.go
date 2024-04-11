package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	commands := getCommands()

	for {
		fmt.Print(color.HiGreenString("Pokedex > "))

		var input string

		fmt.Scanln(&input)
		x := commands[input]
		x.Command()
	}
}

type cliCommand struct {
	Name    string
	Desc    string
	Command func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name: "help",
			Desc: "Get info about this cli and other commands",
			Command: func() error {
				fmt.Println("help received")
				return nil
			},
		},
		"exit": {
			Name: "Exit",
			Desc: "Exit command line",
			Command: func() error {
				fmt.Println("Exiting")
				os.Exit(0)
				return nil
			},
		},
	}
}
