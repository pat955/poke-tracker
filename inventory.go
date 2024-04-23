package main

import (
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
