package main

import (
	"fmt"

	"github.com/nathanwiddi/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *Config) error {
	url := ""
	if cfg.Next != "" {
		url = cfg.Next
	}

	locArea, err := pokeapi.GetLocationArea(url)
	if err != nil {
		return err
	}

	for _, loc := range locArea.Results {
		fmt.Println(loc.Name)
	}

	// Update the config with the new pagination URLs
	cfg.Next = locArea.Next

	// Handle the case when Previous is null
	if locArea.Previous != nil {
		if prev, ok := locArea.Previous.(string); ok {
			cfg.Previous = prev
		}
	} else {
		cfg.Previous = ""
	}

	return nil
}
