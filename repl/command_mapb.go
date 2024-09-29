package repl

import (
	"errors"
)

func CommandMapb(urls *config) error {
	if urls.Prev == nil {
		return errors.New("cannot go prev")
	}

	locationAreas := urls.httpClient.GetLocationAreasPaginated(*urls.Prev)
	urls.Next = locationAreas.Next
	urls.Prev = locationAreas.Previous
	return nil
}
