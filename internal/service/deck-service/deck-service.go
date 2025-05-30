package deckservice

import (
	"encoding/json"

	deckmodel "github.com/nunesyan/agendamento-nails/internal/model/deck-model"
	"gorm.io/gorm"
)

type deckService struct {
	db *gorm.DB
}

func NewDeckService(db *gorm.DB) *deckService {
	db = db.Debug()
	return &deckService{db}
}

func (s *deckService) ListDecks() ([]deckmodel.ListDeck, error) {
	var (
		response []byte
		result   []deckmodel.ListDeck
	)

	query := `
		SELECT COALESCE(jsonb_agg(
			jsonb_build_object(
			'id', id,
			'name', name
			)
		), '[]'::jsonb)
		FROM tb_deck;
	`

	if err := s.db.Raw(query).Row().Scan(&response); err != nil {
		return result, err
	}

	if err := json.Unmarshal(response, &result); err != nil {
		return nil, err
	}

	return result, nil
}
