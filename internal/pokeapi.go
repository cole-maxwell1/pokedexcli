package pokeapi

import (
	"net/http"
	"time"
)

type PokeapiClient struct {
	httpClient http.Client
}

const baseUrl = "https://pokeapi.co/api/v2"

func NewClient() PokeapiClient {
	return PokeapiClient{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}


