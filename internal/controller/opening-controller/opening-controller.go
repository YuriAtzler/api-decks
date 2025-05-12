package openingcontroller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	userservice "github.com/nunesyan/agendamento-nails/internal/service/user-service"
)

// Show some controller
func ShowOpening(ctx *gin.Context) {

	// userId, userIdError := strconv.Atoi(ctx.Query("id"))
	// if userIdError != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"msg":   "Id o usuário inválido",
	// 		"error": "Id o usuário inválido",
	// 	})
	// 	return
	// }

	userId, userIdError := strconv.Atoi(ctx.Param("id"))
	if userIdError != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Id o usuário inválido",
			"error": "Id o usuário inválido",
		})
		return
	}

	user, err := userservice.GetUserById(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "deu erro filhao",
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "GET Opening",
		"user": user,
	})
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
