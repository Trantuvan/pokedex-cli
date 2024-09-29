package internal

import (
	"encoding/json"
	"io"
	"log"
	"time"
)

const (
	BaseURL      = "https://pokeapi.co/api/v2/"
	LocationArea = "location-area"
)

type pokeDex struct {
	Count    int64
	Next     string
	Previous *string //* allow nil string
	Results  []Result
}

type Result struct {
	Name string
	URL  string
}

func (c Client) GetLocationAreasPaginated(url string) pokeDex {
	caches := NewCache(time.Second)
	locationAreas := &pokeDex{}

	if location, ok := caches.Get(url); ok {
		errUnmarshal := json.Unmarshal(location, locationAreas)

		if errUnmarshal != nil {
			log.Fatal(errUnmarshal)
		}

		return *locationAreas
	}

	res, errGet := c.httpClient.Get(url)

	if errGet != nil {
		log.Fatal(errGet)
	}

	defer res.Body.Close()

	body, errRead := io.ReadAll(res.Body)

	// * filter response only show res with status code 2**
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody:%s\n", res.StatusCode, body)
	}

	if errRead != nil {
		log.Fatal(errRead)
	}

	caches.Add(url, body)
	errUnmarshal := json.Unmarshal(body, locationAreas)

	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}

	return *locationAreas
}
