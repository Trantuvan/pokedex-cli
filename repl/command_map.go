package repl

import (
	"fmt"

	"github.com/trantuvan/pokedex-cli/internal/pokeapi"
)

func commadMap(urls *config) error {
	if len(urls.Next) == 0 {
		urls.Next = fmt.Sprintf("%s%s", pokeapi.BaseURL, pokeapi.LocationArea)
	}
	locationAreas := urls.httpClient.GetLocationAreasPaginated(urls.Next)

	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}

	urls.Next = locationAreas.Next
	urls.Prev = locationAreas.Previous
	return nil
}
