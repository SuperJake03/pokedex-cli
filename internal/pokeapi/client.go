package pokeapi

import (
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
