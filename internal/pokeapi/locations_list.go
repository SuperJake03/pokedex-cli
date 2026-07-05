package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type locationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocationsAreas(pageURL *string) (locationAreasResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreasResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreasResponse{}, nil
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreasResponse{}, err
	}

	var locationAreas locationAreasResponse
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return locationAreasResponse{}, err
	}
	return locationAreas, nil
}
