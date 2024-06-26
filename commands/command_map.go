package commands

import (
	"errors"
	"fmt"

	"github.com/vbaxan-linkedin/pokedexcli/internal/pokeapi"
	"github.com/vbaxan-linkedin/pokedexcli/internal/pokecache"
)

func commandMap(config *pokeapi.Config, cache *pokecache.AppCache, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if len(config.Next) > 0 {
		url = config.Next
	}
	return requestLocationAreas(url, config, cache)
}

func commandMapB(config *pokeapi.Config, cache *pokecache.AppCache, args ...string) error {
	if len(config.Previous) > 0 {
		return errors.New("don't have what to go back to")
	}
	return requestLocationAreas(config.Previous, config, cache)
}

func requestLocationAreas(url string, config *pokeapi.Config, cache *pokecache.AppCache) error {
	response := pokeapi.LocationAreasResponse{}

	if cacheHit := tryHitCache(url, cache, &response); !cacheHit {
		fmt.Printf("Requesting %s...\n", url)
		body, err := pokeapi.SendGetRequest(url, &response)
		if err != nil {
			return err
		}
		cache.Cache.Add(url, &body)
	}

	updateConfigAndPrintResults(config, &response)
	return nil
}

func updateConfigAndPrintResults(config *pokeapi.Config, response *pokeapi.LocationAreasResponse) {
	config.UpdateFromResponse(response)
	fmt.Println("Received location areas:")
	for _, area := range response.Results {
		if len(area.Name) > 0 {
			fmt.Println(area.Name)
		}
	}
	fmt.Println()
}
