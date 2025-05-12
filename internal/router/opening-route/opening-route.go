package openingrouter

import (
	"github.com/gin-gonic/gin"
	openingcontroller "github.com/nunesyan/agendamento-nails/internal/controller/opening-controller"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// Show opening
		v1.GET("/opening/:id/por-id", openingcontroller.ShowOpening)

		// Create opening
		v1.POST("/opening", openingcontroller.CreateOpening)

		// Delete opening
		v1.DELETE("/opening", openingcontroller.DeleteOpening)

		// Update opening
		v1.PUT("/opening", openingcontroller.UpdateOpening)

		// Show all opening
		v1.GET("/openings", openingcontroller.ShowAllOpening)
	}
}
