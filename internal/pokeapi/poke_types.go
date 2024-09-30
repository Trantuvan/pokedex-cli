package pokeapi

type locationArea struct {
	Count    int64
	Next     string
	Previous *string //* allow nil string
	Results  []Location
}

type Location struct {
	Name string
	URL  string
}

type pokemonInArea struct {
	Id                int                `json:"id"`
	GameIndex         int                `json:"game_index"`
	LocationName      string             `json:"name"`
	PokemonEncounters []pokemonEncounter `json:"pokemon_encounters"`
}

type pokemonEncounter struct {
	Pokemon pokemon `json:"pokemon"`
}

type pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
