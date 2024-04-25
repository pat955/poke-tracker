package main

// TODO
// Maybe replace nickname= "" with nickname = name
import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Pokedex struct {
	CurrentPokemon map[string]PokemonData
	Storage        map[string]PokemonData
	Pokedex        map[string]PokemonData
}

func newPokedex() Pokedex {
	p := Pokedex{}
	p.CurrentPokemon = make(map[string]PokemonData, 6)
	p.Storage = make(map[string]PokemonData, 0)
	p.Pokedex = make(map[string]PokemonData, 0)
	return p
}
func (dex Pokedex) Inspect(pokemon_name string) error {
	data, ok := dex.CurrentPokemon[pokemon_name]
	if !ok {
		return errors.New(pokemon_name + " not in your inventory")
	}
	fmt.Println(color.RedString("-------Stats-------"))
	data.PrintBaseStats()
	fmt.Println(color.RedString("-------------------"))

	fmt.Println(color.HiYellowString("-------Types-------"))
	data.PrintTypes()
	fmt.Println(color.HiYellowString("-------------------"))
	return nil
}
func (dex Pokedex) PrintOutMyPokemon() {
	fmt.Println("Pokemon in pokedex:")
	for _, pokemon := range dex.CurrentPokemon {
		if pokemon.Nickname != pokemon.Name {
			fmt.Println("-", color.MagentaString(pokemon.Nickname), "("+color.CyanString((strings.Title(pokemon.Name)))+")")
		} else {
			fmt.Println("-", color.MagentaString(strings.Title(pokemon.Name)))
		}
	}
}
func (dex Pokedex) PrintOutPokedex() {
	for _, pokemon := range dex.Pokedex {
		fmt.Println("-", color.MagentaString(strings.Title(pokemon.Name)))
	}
}
func (dex Pokedex) PrintOutCurrentPokemon() {
	fmt.Println("-----HELD POKEMON-----")
	for _, pokemon := range dex.CurrentPokemon {
		if pokemon.Nickname != pokemon.Name {
			fmt.Println("-", color.HiCyanString(pokemon.Nickname), "("+color.MagentaString(strings.Title(pokemon.Name))+")")
			continue
		}
		fmt.Println("-", color.MagentaString(strings.Title(pokemon.Name)))
	}
}
func (dex Pokedex) Add(poke *PokemonData) {
	if len(dex.CurrentPokemon) >= 6 {
		dex.Storage[poke.Name] = *poke
		fmt.Println("Put", poke.Nickname, "in storage")
	} else {
		dex.CurrentPokemon[poke.Name] = *poke
	}
	dex.Pokedex[poke.Name] = *poke
}

type ItemInventory struct {
	Items []Item
}

type Item struct {
	Attributes []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"attributes"`
	BabyTriggerFor any `json:"baby_trigger_for"`
	Category       struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"category"`
	Cost          int `json:"cost"`
	EffectEntries []struct {
		Effect   string `json:"effect"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		ShortEffect string `json:"short_effect"`
	} `json:"effect_entries"`
	FlavorTextEntries []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Text         string `json:"text"`
		VersionGroup struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version_group"`
	} `json:"flavor_text_entries"`
	FlingEffect any `json:"fling_effect"`
	FlingPower  any `json:"fling_power"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	HeldByPokemon []any  `json:"held_by_pokemon"`
	ID            int    `json:"id"`
	Machines      []any  `json:"machines"`
	Name          string `json:"name"`
	Names         []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Sprites struct {
		Default string `json:"default"`
	} `json:"sprites"`
}

func (i *Item) GetID() int {
	return i.ID
}
func (i *Item) GetURL() string {
	return i.Category.URL
}
func (i *Item) PrintInfo() {
	fmt.Println("Nothing here, printing out item info")
}
