package repl

import (
	"errors"
	"fmt"
)

func CommandMapb(urls *config) error {
	if urls.Prev == nil {
		return errors.New("cannot go prev")
	}

	locationAreas := urls.httpClient.GetLocationAreasPaginated(*urls.Prev)

	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}

	urls.Next = locationAreas.Next
	urls.Prev = locationAreas.Previous
	return nil
}
