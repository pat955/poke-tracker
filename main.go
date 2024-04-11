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
		x.Config.Command()
	}
}

type cliCommand struct {
	Name    string
	Desc    string
	Command func() error
	Config  *Config
}

type Config struct {
	Name    string
	Desc    string
	Command func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name: "Help",
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

		"map": {
			Name: "Map",
			Desc: "Map the next 20 area of pokemon",
			Command: func() error {
				call("")
				return nil
			},
			Config: &Config{
				Name: "Previous",
				Desc: "Get the previous 20 map areas",
				Command: func() error {
					fmt.Println("previous!")
					return nil
				},
			},
		},
	}
}

func call(s string) {
	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/%s/", s)
	fmt.Println(endpoint)

}
