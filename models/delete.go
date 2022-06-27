package models

import "github.com/chmenegatti/golang-api-01/database"

func Delete(id int64) (int64, error) {
	conn, err := database.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id=$4`, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
