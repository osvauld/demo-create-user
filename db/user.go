package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const createUser = `
INSERT INTO users (username, name, temp_password, type)
VALUES ($1, $2, $3, 'user')
`

type CreateUserParams struct {
	Username     string `json:"username"`
	Name         string `json:"name"`
	TempPassword string `json:"tempPassword"`
}

func CreateUser(db *sql.DB, arg CreateUserParams) (int64, error) {
	result, err := db.Exec(createUser, arg.Username, arg.Name, arg.TempPassword)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
