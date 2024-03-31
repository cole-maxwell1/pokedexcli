package pokeapi

import (
	"net/http"
	"time"

	"github.com/cole-maxwell1/pokedexcli/pokecache"
)

type PokeapiClient struct {
	httpClient http.Client
	cache      pokecache.Cache
}

const baseUrl = "https://pokeapi.co/api/v2"

func NewClient(cacheLifetime time.Duration) PokeapiClient {
	return PokeapiClient{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(cacheLifetime),
	}
}
