package main

type Profile struct {
	Inventory   ItemInventory
	Pokedex     Pokedex
	CurrentArea string
}

func newProfile() Profile {
	i := NewItemInventory()
	dex := newPokedex()
	return Profile{Inventory: i, Pokedex: dex}
}
