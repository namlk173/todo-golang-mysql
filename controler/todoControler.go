package controler

import (
	"database/sql"
	"namlk/gomysql/entities"
	"namlk/gomysql/helper"
)

type TodoControler struct {
	DB *sql.DB
}

// Get all todos in database
func (t *TodoControler) GetAllTodos() []entities.Todo {
	todos := make([]entities.Todo, 0)
	stmt := "SELECT * FROM todo"
	rows, err := t.DB.Query(stmt)
	helper.ErrCheck(err)
	defer rows.Close()
	for rows.Next() {
		var (
			id        int
			name      string
			expired   bool
			completed bool
		)
		err = rows.Scan(&id, &name, &expired, &completed)
		helper.ErrCheck(err)
		todo := *entities.NewTodo(id, name, expired, completed)
		todos = append(todos, todo)
	}
	return todos
}

// Get todo by id
func (t *TodoControler) GetTodoById(todoId int) (entities.Todo, error) {
	stmt := "SELECT * FROM todo WHERE id = ?"
	row := t.DB.QueryRow(stmt, todoId)
	var (
		id        int
		name      string
		expired   bool
		completed bool
	)
	err := row.Scan(&id, &name, &expired, &completed)
	return *entities.NewTodo(id, name, expired, completed), err
}

// Add new todo into database
func (t *TodoControler) CreateNewTodo(name string, expired, completed bool) error {
	stmt := "INSERT INTO todo (name, expired, completed) VALUE (?, ?, ?)"
	_, err := t.DB.Exec(stmt, name, expired, completed)
	return err
}

// Update a todo into database
func (t *TodoControler) UpdateTodo(todo *entities.Todo) error {
	stmt := "UPDATE todo SET name = ?, expired = ?, completed = ? WHERE id = ?"
	_, err := t.DB.Exec(stmt, todo.GetName(), todo.GetExpired(), todo.GetCompleted(), todo.GetId())
	return err
}

// Delete a todo from database
func (t *TodoControler) DeleteTodo(id int) error {
	stmt := "DELETE FROM todo WHERE id = ?"
	_, err := t.DB.Exec(stmt, id)
	return err
}

// Search todos by expired and completed
func (t *TodoControler) SearchTodo(typeSearch string) []entities.Todo {
	todos := make([]entities.Todo, 0)
	var stmt string
	if typeSearch == "expired" {
		// Search todo that not completed and expired
		stmt = "SELECT * FROM todo WHERE expired = 1 AND completed = 0"
	} else {
		// Search todo that not completed and not expried
		stmt = "SELECT * FROM todo WHERE completed = 0 AND expired = 0"
	}
	rows, err := t.DB.Query(stmt)
	helper.ErrCheck(err)
	defer rows.Close()
	for rows.Next() {
		var (
			id                 int
			name               string
			expired, completed bool
		)
		err = rows.Scan(&id, &name, &expired, &completed)
		helper.ErrCheck(err)
		todo := *entities.NewTodo(id, name, expired, completed)
		todos = append(todos, todo)
	}
	return todos
}
