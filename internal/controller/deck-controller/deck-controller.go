package deckcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nunesyan/agendamento-nails/internal/database"
	handleerror "github.com/nunesyan/agendamento-nails/internal/pkg/handle-error"
	deckservice "github.com/nunesyan/agendamento-nails/internal/service/deck-service"
)

func ListDecks(ctx *gin.Context) {
	db := database.CallDatabase(ctx)
	defer database.CloseDatabase()

	service := deckservice.NewDeckService(db)

	result, err := service.ListDecks()
	if err != nil {
		handleerror.HandleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}
