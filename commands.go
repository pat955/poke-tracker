package main

// TODO
// Add callnames, []string

import (
	"fmt"
	"os"
)

type cliCommand struct {
	Name    string
	Desc    string
	Command func() error
	Config  *Config
}

type Config struct {
	CallName string
	Name     string
	Desc     string
	Command  func() error
}

func getConfigs() map[string]*Config {
	return map[string]*Config{
		"prevMap": {
			CallName: "prev",
			Name:     "Previous",
			Desc:     "placeholder",
			Command: func() error {
				fmt.Println("prev map!")
				return nil
			},
		},
	}
}

func getCommands() map[string]cliCommand {
	configs := getConfigs()
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
			Config: configs["prevmap"],
		},
	}
}

func call(s string) {
	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/%s/", s)
	fmt.Println(endpoint)

}
