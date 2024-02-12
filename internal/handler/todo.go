package handler

import (
	"net/http"
	"strconv"

	"github.com/Leonardo2604/go-todo-api/internal/repository"
	"github.com/gin-gonic/gin"
)

func GetAllTodos(ctx *gin.Context) {
	todos := repository.GetAllTodos();

	ctx.JSON(http.StatusOK, todos)
}

func FindTodo(ctx *gin.Context) {
	idParam := ctx.Param("id");

	id, err := strconv.ParseUint(idParam, 10, 64);

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid id",
		})

		return
	}

	todo := repository.FindTodo(&id);

	if todo.Id == 0 {
		ctx.JSON(404, gin.H{
			"message": "Todo not found",
		})

		return
	}

	ctx.JSON(http.StatusOK, todo)
}

func CreateTodo(ctx *gin.Context) {
	body := struct{
		Title string `json:"title"`
		IsDone bool `json:"is_done"`
	}{}

	ctx.Bind(&body)

	params := repository.CreateTodoParams{
		Title: body.Title,
		IsDone: body.IsDone,
	}

	todo := repository.CreateTodo(&params)

	ctx.JSON(http.StatusCreated, todo)
}

func UpdateTodo(ctx *gin.Context) {
	idParam := ctx.Param("id");

	id, err := strconv.ParseUint(idParam, 10, 64);

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid id",
		})

		return
	}

	body := struct{
		Title string `json:"title"`
		IsDone bool `json:"is_done"`
	}{}

	ctx.Bind(&body)

	params := repository.UpdateTodoParams{
		Title: body.Title,
		IsDone: body.IsDone,
	}

	repository.UpdateTodo(&id, &params);

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func DeleteTodo(ctx *gin.Context) {
	idParam := ctx.Param("id");

	id, err := strconv.ParseUint(idParam, 10, 64);

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid id",
		})

		return
	}

	repository.DeleteTodo(&id);

	ctx.JSON(http.StatusNoContent, gin.H{})
}
