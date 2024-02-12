package main

import (
	"github.com/Leonardo2604/go-todo-api/internal/config"
	"github.com/Leonardo2604/go-todo-api/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitializeConfig();

	r := gin.Default()

	r.GET("/todos", handler.GetAllTodos)
	r.GET("/todos/:id", handler.FindTodo)
	r.POST("/todos", handler.CreateTodo)
	r.PUT("/todos/:id", handler.UpdateTodo)
	r.DELETE("/todos/:id", handler.DeleteTodo)

	r.Run(":3000");
}