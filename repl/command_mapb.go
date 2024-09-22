package repl

import (
	"errors"

	"github.com/trantuvan/pokedex-cli/internal"
)

func CommandMapb(urls *config) error {
	if urls.Prev == nil {
		return errors.New("cannot go prev")
	}

	locationAreas := internal.GetLocationAreasPaginated(*urls.Prev)
	urls.Next = locationAreas.Next
	urls.Prev = locationAreas.Previous
	return nil
}
