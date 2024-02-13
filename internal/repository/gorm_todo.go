package repository

import (
	"time"

	"github.com/Leonardo2604/go-todo-api/internal/db"
	"github.com/Leonardo2604/go-todo-api/internal/entity"
)

type gormTodoRepository struct{}

func (r gormTodoRepository) GetAllTodos() []entity.Todo {
	todos := []entity.Todo{}

	db.GetDB().Raw("SELECT id, title, is_done, created_at, updated_at, deleted_at from todos where deleted_at is null").Scan(&todos)

	return todos
}

func (r gormTodoRepository) CreateTodo(params CreateTodoParams) entity.Todo {
	todo := struct {
		Id        uint64
		CreatedAt time.Time
		UpdatedAt time.Time
	}{}

	db.GetDB().Raw("insert into todos (title, is_done) values (?, ?) returning id, created_at, updated_at", params.Title, params.IsDone).Scan(&todo)

	return entity.Todo{
		Id:        todo.Id,
		Title:     params.Title,
		IsDone:    params.IsDone,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

func (r gormTodoRepository) UpdateTodo(id uint64, params UpdateTodoParams) {
	db.GetDB().Exec("update todos set title = ?, is_done = ?, updated_at = now() where id = ? and deleted_at is null", params.Title, params.IsDone, id)
}

func (r gormTodoRepository) FindTodo(id uint64) entity.Todo {
	todo := entity.Todo{}

	db.GetDB().Raw("SELECT id, title, is_done, created_at, updated_at, deleted_at from todos where id = ? and deleted_at is null", id).Scan(&todo)

	return todo
}

func (r gormTodoRepository) DeleteTodo(id uint64) {
	db.GetDB().Exec("update todos set deleted_at = now() where id = ? and deleted_at is null", id)
}
