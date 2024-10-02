package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"math"
	"math/rand"
)

const (
	BaseURL               = "https://pokeapi.co/api/v2/"
	LocationArea          = "location-area/"
	maximumBaseExperience = 635
	minimumBaseExperince  = 40
)

var Pokemons map[string]catchPokemon

func (c Client) GetLocationAreasPaginated(url string) locationArea {
	locationAreas := &locationArea{}

	if location, ok := c.cache.Get(url); ok {
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

	c.cache.Add(url, body)
	errUnmarshal := json.Unmarshal(body, locationAreas)

	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}

	return *locationAreas
}

func (c Client) ExplorePokemonInArea(url string) pokemonInArea {
	pokemonInArea := &pokemonInArea{}

	if pokemons, ok := c.cache.Get(url); ok {
		errUnmarshal := json.Unmarshal(pokemons, pokemonInArea)

		if errUnmarshal != nil {
			log.Fatal(errUnmarshal)
		}

		return *pokemonInArea
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

	c.cache.Add(url, body)
	errUnmarshal := json.Unmarshal(body, pokemonInArea)

	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}

	return *pokemonInArea
}

func (c Client) Catch(url string) map[string]catchPokemon {
	Pokemons = map[string]catchPokemon{}
	catch := &catchPokemon{}
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

	errUnmarshal := json.Unmarshal(body, catch)

	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}

	if percent := calcCatchPercentage(catch.BaseExperience); percent >= 30 {
		Pokemons[catch.Name] = *catch
	}

	return Pokemons
}

func calcCatchPercentage(baseExperience float64) float64 {
	percent := (1 - (baseExperience-minimumBaseExperince)/(maximumBaseExperience-minimumBaseExperince)) * 100
	randomEffect := rand.Intn(100)
	catchPercent := percent + float64(randomEffect)
	return math.Min(catchPercent, 100)
}
