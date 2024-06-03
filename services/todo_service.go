package services

import (
	"database/sql"

	"github.com/hypercompl3x/todo/model"
)

func GetTodo(db *sql.DB, id int) (model.Todo, error) {
	var todo model.Todo

	row := db.QueryRow("SELECT * FROM todo WHERE id = ?", id)

	err := row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.IsComplete)

	if err != nil {
		return model.Todo{}, err
	}

	return todo, nil
}

func GetTodos(db *sql.DB) ([]model.Todo, error) {
	rows, err := db.Query("SELECT * FROM todo")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []model.Todo

	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.IsComplete)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	err = rows.Err()

	if err != nil {
    return nil, err
  }

	return todos, nil
}

func CreateTodo(db *sql.DB, title string, description string) (*model.Todo, error) {
	result, err := db.Exec("INSERT INTO todo (title, description, isComplete) VALUES (?, ?, ?)", title, description, false)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	todo := model.Todo{
		Id:          int(id),
		Title:       title,
		Description: description,
		IsComplete:  false,
	}

	return &todo, nil
}

func UpdateTodo(db *sql.DB, id int, isComplete bool) error {
	_, err := db.Exec("UPDATE todo SET isComplete = ? WHERE id = ?", isComplete, id)

	if err != nil {
		return err
	}

	return nil
}