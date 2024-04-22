package main

// TODO
// Add callnames, []string

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-zoox/fetch"
)

type data interface {
	String() string
	GetUrl() string
	GetId() int
}
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
	currentLocationID := 0
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
			Desc: "Map of the next 20 area of pokemon",
			Command: func() error {
				for i := 0; i < 20; i++ {
					currentLocationID++
					Location := LocationData{}
					call(fmt.Sprintf("location/%v", currentLocationID), &Location)
					fmt.Println(Location.String())
				}
				return nil
			},
			Config: configs["prevmap"],
		},
	}
}

func call(str string, dataLocation data) {
	endpoint := fmt.Sprintf("https://pokeapi.co/api/v2/%s/", str)
	response, err := fetch.Get(endpoint)
	if err != nil {
		panic(err)
	}
	responsejson, err := response.JSON()
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(responsejson), &dataLocation)

}

type LocationData struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Region struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"region"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	Areas []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"areas"`
}

func (l LocationData) String() string {
	return fmt.Sprint(l.ID, " "+l.Name+" ", l.Region.Name+" ")
}

func (l LocationData) GetUrl() string {
	return l.Region.URL
}
func (l LocationData) GetId() int {
	return l.ID
}
