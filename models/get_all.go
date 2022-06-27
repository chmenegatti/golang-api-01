package models

import (
	"github.com/chmenegatti/golang-api-01/database"
)

func GetAll() (todos []Todo, err error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM todos`)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}
