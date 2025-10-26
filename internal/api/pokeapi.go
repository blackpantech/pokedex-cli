package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/blackpantech/pokedex-cli/internal/models"
)

const apiURL = "https://pokeapi.co/api/v2/pokemon/%s"

func GetPokemon(pokemonNameOrID string) (*models.Pokemon, error) {
	response, err := http.Get(fmt.Sprintf(apiURL, pokemonNameOrID))
	if err != nil {
		log.Fatalf("Error fetching data for this Pokemon: %s\n", err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body for this Pokemon: %s\n", err)
		return nil, err
	}

	var pokemon models.Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		log.Fatalf("Error unmarshalling response body for this Pokemon: %s\n", err)
		return nil, err
	}

	return &pokemon, nil
}
