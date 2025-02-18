package pokeapi

import (
	"net/http"
	"time"

	"github.com/dangoodie/pokedex/internal/pokecache"
)

const (
	BaseURL = "https://pokeapi.co/api/v2/"
)

// Client
type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

// New Client
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}
