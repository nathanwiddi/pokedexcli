package main

import (
	"fmt"

	"github.com/nathanwiddi/pokedexcli/internal/pokeapi"
)

func commandMapb(cfg *Config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locArea, err := pokeapi.GetLocationArea(cfg.Previous)
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
