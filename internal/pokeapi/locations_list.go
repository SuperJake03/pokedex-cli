package pokeapi

import (
	"encoding/json"
)

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocationsAreas(pageURL *string) (LocationAreasResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	data, err := c.getPokeApi(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	var locationAreas LocationAreasResponse
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return LocationAreasResponse{}, err
	}

	return locationAreas, nil
}
