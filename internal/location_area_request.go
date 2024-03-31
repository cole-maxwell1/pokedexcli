package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *PokeapiClient) GetLocationAreas(url *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

	if url != nil {
		fullUrl = *url
	}

	// Check if request has be cached
	cachedResp, ok := c.cache.Get(fullUrl)

	if ok {
		locationAreasData := LocationAreasResponse{}
		fmt.Println("Cache hit!")
		err := json.Unmarshal(cachedResp, &locationAreasData)

		if err != nil {
			return LocationAreasResponse{}, err
		}
		return locationAreasData, nil
	}
	fmt.Println("Cache miss... making a request!")

	// Create request
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return LocationAreasResponse{}, fmt.Errorf("bad status code %d", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasData := LocationAreasResponse{}

	err = json.Unmarshal(data, &locationAreasData)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullUrl, data)

	return locationAreasData, nil

}
