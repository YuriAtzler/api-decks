package deckcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nunesyan/agendamento-nails/internal/database"
	deckmodel "github.com/nunesyan/agendamento-nails/internal/model/deck-model"
	handleerror "github.com/nunesyan/agendamento-nails/internal/pkg/handle-error"
	deckservice "github.com/nunesyan/agendamento-nails/internal/service/deck-service"
)

func ListDecks(ctx *gin.Context) {
	db := database.Get()

	service := deckservice.NewDeckService(db)

	result, err := service.ListDecks()
	if err != nil {
		handleerror.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func CreateDeck(ctx *gin.Context) {
	db := database.Get()

	service := deckservice.NewDeckService(db)

	var deck deckmodel.CreateDeck
	if err := ctx.ShouldBindJSON(&deck); err != nil {
		handleerror.HandleError(ctx, &handleerror.HandleErrorModel{
			Status:  http.StatusBadRequest,
			Message: "Invalid request payload",
			Error:   err,
		})
		return
	}

	deckId, err := service.CreateDeck(deck)
	if err != nil {
		handleerror.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Deck created successfully",
		"id":      deckId,
	})
}
