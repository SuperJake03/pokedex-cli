package pokeapi

import (
	"io"
	"net/http"
	"time"

	"github.com/SuperJake03/pokedex-cli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(interval),
	}
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
