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
	Pokemon map[string]PokemonData
}

func newPokedex() Pokedex {
	p := Pokedex{}
	p.Pokemon = make(map[string]PokemonData, 0)
	return p
}
func (dex Pokedex) Inspect(pokemon_name string) error {
	data, ok := dex.Pokemon[pokemon_name]
	if !ok {
		return errors.New("inspect error: " + pokemon_name + " not found with you")
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
	for _, pokemon := range dex.Pokemon {
		if pokemon.Nickname != "" {
			fmt.Println("-", color.MagentaString(pokemon.Nickname), "("+color.CyanString((strings.Title(pokemon.Name)))+")")
		} else {
			fmt.Println("-", color.MagentaString(strings.Title(pokemon.Name)))
		}
	}
}
func (dex Pokedex) Add(poke PokemonData) {
	dex.Pokemon[poke.Name] = poke
}

type ItemInventory struct {
	Items []Item
}

type Item struct {
	Name string
	Desc string
}
