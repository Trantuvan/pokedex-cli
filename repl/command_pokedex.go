package repl

import (
	"errors"
	"fmt"
)

func CommandPokedex(url *config) error {
	if len(url.caughtPokemons) == 0 {
		return errors.New("nothing to list")
	}

	fmt.Println("Your Pokedex:")

	for k := range url.caughtPokemons {
		fmt.Printf(" - %s\n", k)
	}
	return nil
}
