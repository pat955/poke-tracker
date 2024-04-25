package main

import (
	"errors"

	"github.com/pat955/pokedex/internal/pokeapi"
)

func dataTypeToAreaData(cache pokeapi.Cache, endpoint string) (AreaData, error) {
	var areaData AreaData
	d, err := checkAndCall(cache, endpoint, &areaData)
	if err != nil {
		return AreaData{}, err
	}
	data, ok := d.(*AreaData)
	if !ok {
		return AreaData{}, errors.New("conversion error: converting datatype to AreaData not working")
	}
	return *data, nil
}
