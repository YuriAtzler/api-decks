package router

import (
	"github.com/gin-gonic/gin"
	openingrouter "github.com/nunesyan/agendamento-nails/internal/router/opening-route"
)

func Initialize() {
	// Initialize Router
	router := gin.Default()
	// Initialize Routes
	openingrouter.InitializeRoutes(router)

	// Run the server
	router.Run() // listen and serve on 0.0.0.0:8080
}
