package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
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

func (c *Client) getPokeApi(url string) ([]byte, error) {
	entry, ok := c.cache.Get(url)
	if ok {
		return entry, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, data)

	return data, nil
}
