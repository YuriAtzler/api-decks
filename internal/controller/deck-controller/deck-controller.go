package deckcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nunesyan/agendamento-nails/internal/database"
	deckservice "github.com/nunesyan/agendamento-nails/internal/service/deck-service"
)

func ListDecks(ctx *gin.Context) {
	db := database.CallDatabase(ctx)
	defer database.CloseDatabase()

	service := deckservice.NewDeckService(db)

	result, err := service.ListDecks()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Ocorreu um erro ao buscar os decks",
			"error":   err,
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// Create controller
func CreateOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "POST Opening",
	})
}

// Delete controller
func DeleteOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Opening",
	})
}

// Update controller
func UpdateOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}

// Show all controllers
func ShowAllOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Openings",
	})
}
