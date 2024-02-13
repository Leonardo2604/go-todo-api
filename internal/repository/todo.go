package repository

import "github.com/Leonardo2604/go-todo-api/internal/entity"

type CreateTodoParams struct {
	Title  string
	IsDone bool
}

type UpdateTodoParams struct {
	Title  string
	IsDone bool
}

type TodoRepository interface {
	GetAllTodos() []entity.Todo

	CreateTodo(params CreateTodoParams) entity.Todo

	UpdateTodo(id uint64, params UpdateTodoParams)

	FindTodo(id uint64) entity.Todo

	DeleteTodo(id uint64)
}
