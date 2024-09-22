package repl

import (
	"fmt"

	"github.com/trantuvan/pokedex-cli/internal"
)

func commadMap(urls *config) error {
	if len(urls.Next) == 0 {
		urls.Next = fmt.Sprintf("%s%s", internal.BaseURL, internal.LocationArea)
	}
	locationAreas := internal.GetLocationAreasPaginated(urls.Next)
	urls.Next = locationAreas.Next
	urls.Prev = locationAreas.Previous
	return nil
}
