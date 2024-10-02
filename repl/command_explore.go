package repl

import (
	"errors"
	"fmt"

	"github.com/trantuvan/pokedex-cli/internal/pokeapi"
)

func CommandExplore(url *config, filters ...string) error {
	if len(filters) != 1 {
		return errors.New("you must provide a location name")
	}

	exploreUrl := fmt.Sprintf("%s%s%s", pokeapi.BaseURL, pokeapi.LocationArea, filters[0])
	pokemonsInArea := url.httpClient.ExplorePokemonInArea(exploreUrl)
	fmt.Printf("Exploring %s...\n", pokemonsInArea.LocationName)
	fmt.Println("Found Pokemon:")

	for _, pokemon := range pokemonsInArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
