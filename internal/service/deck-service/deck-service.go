package deckservice

import (
	"encoding/json"
	"net/http"

	deckmodel "github.com/nunesyan/agendamento-nails/internal/model/deck-model"
	handleerror "github.com/nunesyan/agendamento-nails/internal/pkg/handle-error"
	"gorm.io/gorm"
)

type deckService struct {
	db *gorm.DB
}

func NewDeckService(db *gorm.DB) *deckService {
	db = db.Debug()
	return &deckService{db}
}

func (s *deckService) ListDecks() ([]deckmodel.ListDeck, *handleerror.HandleErrorModel) {
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
		return result, &handleerror.HandleErrorModel{
			Status:  http.StatusInternalServerError,
			Message: "Failed to execute query",
			Error:   err,
		}
	}

	if err := json.Unmarshal(response, &result); err != nil {
		return nil, &handleerror.HandleErrorModel{
			Status:  http.StatusInternalServerError,
			Message: "Failed to parse response",
			Error:   err,
		}
	}

	return result, nil
}

func (s *deckService) CreateDeck(deck deckmodel.CreateDeck) (int, *handleerror.HandleErrorModel) {
	var deckId int

	query := `INSERT INTO tb_deck (name) VALUES ($1) RETURNING id;`

	if err := s.db.Raw(query, deck.Name).Row().Scan(&deckId); err != nil {
		return deckId, &handleerror.HandleErrorModel{
			Status:  http.StatusInternalServerError,
			Message: "Failed to execute query",
			Error:   err,
		}
	}

	return deckId, nil
}
