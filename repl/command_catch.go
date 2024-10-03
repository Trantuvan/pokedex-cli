package repl

import (
	"errors"
	"fmt"

	"github.com/trantuvan/pokedex-cli/internal/pokeapi"
)

func CommandCatch(url *config, filters ...string) error {
	if len(filters) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	catchUrl := fmt.Sprintf("%s%s%s", pokeapi.BaseURL, pokeapi.PokemonInfo, filters[0])
	fmt.Printf("Throwing a Pokeball at %s...\n", filters[0])

	if pokemon := url.httpClient.Catch(catchUrl); pokemon == nil {
		fmt.Printf("%s escaped!\n", filters[0])
	} else {
		url.caughtPokemons[pokemon.Name] = *pokemon
		fmt.Printf("%s was caught!\n", filters[0])
	}

	return nil
}
