package handleerror

import "github.com/gin-gonic/gin"

type HandleErrorModel struct {
	Status  int
	Message string
	Error   error
}

func HandleError(ctx *gin.Context, e *HandleErrorModel) {
	ctx.JSON(e.Status, gin.H{
		"message": e.Message,
		"error":   e.Error.Error(),
	})
}
