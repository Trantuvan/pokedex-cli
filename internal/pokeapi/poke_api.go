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
	PokemonInfo           = "pokemon/"
	maximumBaseExperience = 200
	minimumBaseExperince  = 40
)

var catches userPokeDex = userPokeDex{}

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

func (c Client) Catch(url string) *catchPokemon {
	catch := &catchPokemon{}

	if pokemons, ok := c.cache.Get(url); ok {
		errUnmarshal := json.Unmarshal(pokemons, catch)

		if errUnmarshal != nil {
			log.Fatal(errUnmarshal)
		}

		return catch
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

	c.cache.Add(catch.Name, body)
	errUnmarshal := json.Unmarshal(body, catch)

	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}

	if percent := calcCatchPercentage(catch.BaseExperience); percent >= 68 {
		catches[catch.Name] = *catch
		return catch
	}

	return nil
}

func calcCatchPercentage(baseExperience float64) float64 {
	percent := (1 - (baseExperience-minimumBaseExperince)/(maximumBaseExperience-minimumBaseExperince)) * 100
	randomEffect := rand.Float64()
	catchPercent := percent + randomEffect
	return math.Min(catchPercent, 100)
}
