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

func (repository Users) CreateUser(user models.User) (uint64, error) {
	statement, error := repository.db.Prepare("insert into users (username, name, email, password) VALUES (?, ?, ?, ?)")

	if error != nil {
		return 0, error
	}

	defer statement.Close()

	result, error := statement.Exec(user.Username, user.Name, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	lastIDInserted, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastIDInserted), nil

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

func (repository Users) GetUserByEmail(email string) (models.User, error) {
	query, error := repository.db.Query(
		"select user_id, password from users where email = ?",
		email)
	if error != nil {
		return models.User{}, error
	}

	defer query.Close()

	var user models.User

	if query.Next() {
		if error = query.Scan(&user.User_id, &user.Password); error != nil {
			return models.User{}, error
		}
	}

	return user, error
}

func (repository Users) GetUserProfile(userID uint64) (models.User, error) {
	query, error := repository.db.Query(
		"select user_id, username,name,email from users where user_id = ?",
		userID,
	)

	if error != nil {
		return models.User{}, error
	}

	defer query.Close()

	var user models.User

	for query.Next() {

		if error = query.Scan(

			&user.User_id,
			&user.Username,
			&user.Name,
			&user.Email,
		); error != nil {
			return models.User{}, error
		}

	}
	return user, nil
}

func (repository Users) GetSearchedUser(searchedUser string) ([]models.User, error) {
	searchedUser = fmt.Sprintf("%s%%", searchedUser)

	query, error := repository.db.Query(
		"select user_id,username,name,email from users where username like ? or name like ?",
		searchedUser, searchedUser,
	)

	if error != nil {
		return nil, error
	}

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
