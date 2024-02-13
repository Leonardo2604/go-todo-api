package repository

import (
	"time"

	"github.com/Leonardo2604/go-todo-api/internal/db"
	"github.com/Leonardo2604/go-todo-api/internal/entity"
)

type todo struct {
	Id        uint64
	Title     string
	IsDone    bool
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func GetAllTodos() *[]entity.Todo {
	db := db.GetDB()

	var r []todo

	todos := []entity.Todo{}

	db.Raw("SELECT id, title, is_done, created_at, updated_at, deleted_at from todos where deleted_at is null").Scan(&r)

	for _, t := range r {
		createdAt, createdAtErr := time.Parse("2006-01-02T15:04:05.999999-07:00", t.CreatedAt)
		updatedAt, updatedAtErr := time.Parse("2006-01-02T15:04:05.999999-07:00", t.UpdatedAt)
		deletedAt, _ := time.Parse("2006-01-02T15:04:05.999999-07:00", t.DeletedAt)

		if createdAtErr != nil || updatedAtErr != nil {
			continue
		}

		todo := entity.Todo{
			Id:        t.Id,
			Title:     t.Title,
			IsDone:    t.IsDone,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		}

		todos = append(todos, todo)
	}

	return &todos
}

type CreateTodoParams struct {
	Title  string
	IsDone bool
}

func CreateTodo(params *CreateTodoParams) *entity.Todo {
	db := db.GetDB()

	todo := struct {
		Id        uint64
		CreatedAt time.Time
		UpdatedAt time.Time
	}{}

	db.Raw("insert into todos (title, is_done) values (?, ?) returning id, created_at, updated_at", params.Title, params.IsDone).Scan(&todo)

	return &entity.Todo{
		Id:        todo.Id,
		Title:     params.Title,
		IsDone:    params.IsDone,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

type UpdateTodoParams struct {
	Title  string
	IsDone bool
}

func UpdateTodo(id *uint64, params *UpdateTodoParams) {
	db := db.GetDB()

	db.Exec("update todos set title = ?, is_done = ?, updated_at = now() where id = ? and deleted_at is null", params.Title, params.IsDone, id)
}

func FindTodo(id *uint64) *entity.Todo {
	db := db.GetDB()

	todo := entity.Todo{}

	db.Raw("SELECT id, title, is_done, created_at, updated_at, deleted_at from todos where id = ? and deleted_at is null", id).Scan(&todo)

	return &todo
}

func DeleteTodo(id *uint64) {
	db := db.GetDB()

	db.Exec("update todos set deleted_at = now() where id = ? and deleted_at is null", id)
}
