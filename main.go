package main

// TODO
// Add Catch
// Add Save Game
//

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/pat955/pokedex/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	pokeInventory := newPokedex()
	commands := getCommands()
	cache := pokeapi.NewCache(500 * time.Second)
	for {
		fmt.Print(color.HiGreenString("Pokedex >>> "))

		var cmd, args string

		if scanner.Scan() {
			line := strings.Split(scanner.Text(), " ")
			cmd = line[0]
			if len(line) != 1 {
				args = line[1]
			}
			// fmt.Printf("Input was: %q\n", line)
		}
		if cmd == "" {
			continue
		}
		command, ok := commands[cmd]
		if !ok {
			fmt.Println(errors.New("unknown command, type help for commands"))
			continue
		}
		err := command.Command(cache, pokeInventory, args)
		if err != nil {
			fmt.Println(err)
		}

	}
}
