package pokeAPI

import (
	"fmt"
	"io"
	"net/http"

	"github.com/AggroSec/Go-Pokedex/internal/pokecache"
)

func CreatePokeAPIRequest(url string, pokeCache pokecache.Cache) ([]byte, error) {
	if cachedResponse, ok := pokeCache.Get(url); ok {
		return cachedResponse, nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	pokeCache.Add(url, body)

	return body, nil
}
