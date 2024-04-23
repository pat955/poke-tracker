package main

import (
	"fmt"

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
	for _, pokemon := range dex.Pokemon {
		fmt.Println("-", color.MagentaString(pokemon.Name))
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
