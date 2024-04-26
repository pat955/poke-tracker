package main

import (
	"errors"

	"github.com/pat955/pokedex/internal/pokeapi"
)

func dataToLocationData(cache pokeapi.Cache, endpoint string) (*LocationData, error) {
	var locData LocationData
	d, err := checkAndCall(cache, endpoint, &locData)
	if err != nil {
		return &LocationData{}, err
	}
	data, ok := d.(*LocationData)
	if !ok {
		return &LocationData{}, errors.New("conversion error: converting datatype to LocationData not working")
	}
	return data, nil
}
func dataToAreaData(cache pokeapi.Cache, endpoint string) (*AreaData, error) {
	var areaData AreaData
	d, err := checkAndCall(cache, endpoint, &areaData)
	if err != nil {
		return &AreaData{}, err
	}
	data, ok := d.(*AreaData)
	if !ok {
		return &AreaData{}, errors.New("conversion error: converting datatype to AreaData not working")
	}
	return data, nil
}

func dataToPokemonData(cache pokeapi.Cache, endpoint string) (*PokemonData, error) {
	var pokeDataHolder PokemonData
	d, err := checkAndCall(cache, endpoint, &pokeDataHolder)
	if err != nil {
		return &PokemonData{}, err
	}
	pokeData, ok := d.(*PokemonData)
	if !ok {
		return &PokemonData{}, errors.New("conversion error: converting datatype to AreaData not working")
	}
	return pokeData, nil
}
