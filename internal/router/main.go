package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nunesyan/agendamento-nails/internal/database"
	deckrouter "github.com/nunesyan/agendamento-nails/internal/router/deck-route"
)

func Initialize() {
	database.Initialize()
	defer database.Close()

	router := gin.Default()

	// Initialize Routes
	deckrouter.Routes(router)

	// Run the server
	router.Run()
}
