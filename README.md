# PokeTracker ![code coverage badge](https://github.com/pat955/pokedex/actions/workflows/ci.yml/badge.svg)
Pokemon and pokedex-tracking game in a command-line REPL. Uses [PokeApi](https://pokeapi.co/) Catch critters and explore the world of pokemon in a text based format.
### Why?
Honestly it seemed like a fun project, and it really was! Go is by far my favorite coding language, so that in combination with a fun concept was a total blast to build!

## Functionality 
* Explore sub-areas or entire locations
* Check pokemon in that region
* Catch pokemon and name them
* Inspect your pokemon
* Buy items and check inventory

## How to play
First clone the project.
```bash
git clone https://github.com/pat955/pokedex/
```

After that simply run to start:
```bash
./main.sh
```

You will hopefully see something like this:

```
Loading...
PokeCLI >>>
```
Type help to start your adventure!


```
PokeCLI >>> help
Welcome to my PokeCLI!

------------Available Commands------------
help  Get the description of all available commands
exit  Exit command line
map  Get the next 2 location and their areas. The cyan name is the location.
     Explore the areas. eks: >>> explore eterna-city-west-gate
mapb  Map Back. Get the previous 2 locations
explore <area_name>  Explore current area, called with: >>> explore <areaName>
explore-location <locationName>  explore an entire location rather than a small area
catch <pokemon_name>  Catch pokemon using this command after exploring area
inspect <pokemon_name>  Inspect a pokemon in your inventory
pokedex  See all the pokemon you've caught so far
cache  Check Cache for debugging reasons
inventory  Check inventory and use items
shop  Buy items like pokeballs, moves and more
PokeCLI >>> 
```

# Contributing
go version: 1.18
### Clone project
```bash
git clone https://github.com/pat955/poke-tracker
```
### Install go
Install go 1.18, should also work with the newest version but i have not tested it yet. Follow the instructions for your environment here: [go.dev/dl](https://go.dev/dl/)

### Run Scripts
Run the build script then the run script to update binary.
```bash
./scripts/buildprod.sh && ./scripts/main.sh
```

### Run the tests

```bash
go test ./...
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.

# Roadmap
- [ ] Better documentation
- [ ] History when pressing up
- [ ] Make a rarity and chance system
- [ ] Battles, trainers, quests
- [ ] Ability to have more than one of one pokemon, i.e. two pikachu
- [ ] CI (maybe even cd) pipeline
- [ ] SQL db
## 
- [x] Select regions instead of starting from the start
- [x] Items
- [x] Nicknames
- [x] Better Cache

