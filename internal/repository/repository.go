package repository

var Todo TodoRepository

func init() {
	Todo = gormTodoRepository{}
}
