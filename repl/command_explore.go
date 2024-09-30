package repl

import (
	"fmt"

	"github.com/trantuvan/pokedex-cli/internal/pokeapi"
)

func CommandExplore(url *config, filter string) error {
	exploreUrl := fmt.Sprintf("%s%s%s", pokeapi.BaseURL, pokeapi.LocationArea, filter)
	pokemonsInArea := url.httpClient.ExplorePokemonInArea(exploreUrl)
	fmt.Printf("Exploring %s...\n", pokemonsInArea.LocationName)
	fmt.Println("Found Pokemon:")

	for _, pokemon := range pokemonsInArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
