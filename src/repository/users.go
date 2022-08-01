package repository

import (
	"database/sql"
	"fmt"
	"rscm/src/models"
)

type Users struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) GetUsers(user string) ([]models.User, error) {
	user = fmt.Sprintf("%%%s%%", user)

	query, error := repository.db.Query(
		"select user_id, username,name,email from users",
	)

	if error != nil {
		return nil, error
	}

	defer query.Close()

	var users []models.User

	for query.Next() {
		var user models.User

		if error = query.Scan(

			&user.User_id,
			&user.Username,
			&user.Name,
			&user.Email,
		); error != nil {
			return nil, error
		}

		users = append(users, user)
	}
	return users, nil

}
