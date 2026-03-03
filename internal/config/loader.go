package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/fezcode/go-piml"
	"github.com/fezcode/atlas.deck/internal/model"
)

func LoadDeck() (*model.Deck, error) {
	// 1. Try local deck.piml
	if _, err := os.Stat("deck.piml"); err == nil {
		return loadFile("deck.piml")
	}

	// 2. Try global ~/.atlas/deck.piml
	home, err := os.UserHomeDir()
	if err == nil {
		globalPath := filepath.Join(home, ".atlas", "deck.piml")
		if _, err := os.Stat(globalPath); err == nil {
			return loadFile(globalPath)
		}
	}

	return nil, nil // No deck found
}

func loadFile(path string) (*model.Deck, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var deck model.Deck
	err = piml.Unmarshal(data, &deck)
	if err != nil {
		return nil, err
	}

	// Post-process: Trim quotes if piml.Unmarshal didn't
	deck.Name = strings.Trim(deck.Name, "\"")
	deck.Version = strings.Trim(deck.Version, "\"")
	for i := range deck.Pads {
		deck.Pads[i].Key = strings.Trim(deck.Pads[i].Key, "\"")
		deck.Pads[i].Label = strings.Trim(deck.Pads[i].Label, "\"")
		deck.Pads[i].Command = strings.Trim(deck.Pads[i].Command, "\"")
		deck.Pads[i].Color = strings.Trim(deck.Pads[i].Color, "\"")
	}

	return &deck, nil
}
