package pokeAPI

import (
	"fmt"
	"testing"
	"time"

	"github.com/AggroSec/Go-Pokedex/internal/pokecache"
)

func TestCreatePokeAPIRequest(t *testing.T) {
	cases := []struct {
		url     string
		command string
	}{
		{
			url:     "https://pokeapi.co/api/v2/location-area/",
			command: "map",
		},
	}

	cache := pokecache.NewCache(5 * time.Second)

	for _, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", c), func(t *testing.T) {
			results, err := CreatePokeAPIRequest(c.url, cache)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			if len(results) == 0 {
				t.Errorf("expected to find results")
				return
			}
		})
	}
}
