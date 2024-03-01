package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *PokeapiClient) GetLocationAreas() (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint

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

	return locationAreasData, nil

}
