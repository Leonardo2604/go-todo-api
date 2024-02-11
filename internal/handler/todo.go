package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTodos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func FindTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func CreateTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "OK",
	})
}

func UpdateTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func DeleteTodo(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{})
}
