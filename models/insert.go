package models

import "github.com/chmenegatti/golang-api-01/database"

func Insert(todo Todo) (id int64, err error) {

	conn, err := database.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
