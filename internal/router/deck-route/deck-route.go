package deckrouter

import (
	"github.com/gin-gonic/gin"
	deckcontroller "github.com/nunesyan/agendamento-nails/internal/controller/deck-controller"
)

func Routes(router *gin.Engine) {
	route := router.Group("/deck")
	{
		route.GET("", deckcontroller.ListDecks)
		route.POST("", deckcontroller.CreateDeck)
	}
}
