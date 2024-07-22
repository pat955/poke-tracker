package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/pat955/pokedex/internal/pokeapi"
)

func buyItems(_ pokeapi.Cache, inventory ItemInventory) error {
	options := map[string]string{"1": "poke-ball", "2": "ultra-ball", "3": "somethingelse"}
	fmt.Println("You can buy:\n[1]Poke Balls\n[2]Ultra Balls\n[3]somethingelse")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		key := scanner.Text()
		itemName, found := options[key]
		if !found {
			return errors.New("item not available, check spelling")
		}
		fmt.Println("How many? (0-(get max amount)) ")
		if scanner.Scan() {
			amountInput := scanner.Text()
			amountInt, err := strconv.Atoi(amountInput)
			if err != nil {
				return err
				//errors.New("unable to convert", amountInput, "to integer")
			}
			fmt.Println("Adding item to inventory...")

			inventory.Items[itemName].AddItem(amountInt)
		}
	}
	fmt.Println("Your Inventory Now:")
	inventory.PrintOutItems()
	return nil
}
