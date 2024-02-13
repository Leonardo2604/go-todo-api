package main

import (
	"fmt"

	"github.com/Leonardo2604/go-todo-api/internal/config"
	"github.com/Leonardo2604/go-todo-api/internal/db"
	"github.com/Leonardo2604/go-todo-api/internal/handler"
	"github.com/gin-gonic/gin"
)

func init() {
	var err error

	err = config.Init()

	if err != nil {
		fmt.Println("Failed to load settings.")
		return
	}

	err = db.Init()

	if err != nil {
		fmt.Println("Failed to connect database.")
		return
	}
}

func main() {
	c := config.GetConfig()
	r := gin.Default()

	r.GET("/todos", handler.GetAllTodos)
	r.GET("/todos/:id", handler.FindTodo)
	r.POST("/todos", handler.CreateTodo)
	r.PUT("/todos/:id", handler.UpdateTodo)
	r.DELETE("/todos/:id", handler.DeleteTodo)

	addr := fmt.Sprintf(":%d", c.GetServer().GetPort())

	r.Run(addr)
}
