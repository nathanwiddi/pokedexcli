package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationArea(url string) (LocationArea, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	if res.StatusCode > 299 {
		return LocationArea{}, fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
	}

	var locationArea LocationArea
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}
