# PokeTracker

## Description
Pokemon/pokedex game in a command-line REPL. Uses [PokeApi](https://pokeapi.co/)

### Functionality 
* Explore areas
* Check pokemon in that region
* Catch pokemon and name them
* Inspect your pokemon
* Buy items and check inventory


## How to play
First clone the project.
```
git clone https://github.com/pat955/pokedex/
```

After that simply run to start:
```
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
```

## Plans
* Better documentation
* History when pressing up
* Make a rarity and chance system
* Battles, trainers
* Ability to have more than one of one pokemon, i.e. two pikachu
* ~~Select regions instead of starting from the start~~ (Done!)
* ~~Items~~ (Done!)
* ~~Nicknames~~ (Done!)
* ~~Better Cache~~ (Done!)

