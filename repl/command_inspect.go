package repl

import (
	"errors"
	"fmt"
)

func CommandInspect(url *config, filters ...string) error {
	if len(filters) != 1 {
		return errors.New("you've must enter a pokemon name")
	}

	pokemon, ok := url.caughtPokemons[filters[0]]

	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight)

	if len(pokemon.Stats) > 0 {
		fmt.Println("Stats:")
	}

	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStatus)
	}

	if len(pokemon.Types) > 0 {
		fmt.Println("Types:")
	}

	for _, ty := range pokemon.Types {
		fmt.Printf(" - %s\n", ty.Type.Name)
	}

	return nil
}
