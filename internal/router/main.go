package router

import (
	"github.com/gin-gonic/gin"
	deckrouter "github.com/nunesyan/agendamento-nails/internal/router/deck-route"
)

func Initialize() {
	router := gin.Default()

	// Initialize Routes
	deckrouter.Routes(router)

	// Run the server
	router.Run()
}
