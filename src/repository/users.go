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

func (repository Users) GetUserById(userID uint64) (models.User, error) {
	query, error := repository.db.Query(
		"select user_id, username,name,email from users where user_id = ?",
		userID,
	)

	if error != nil {
		return models.User{}, error
	}

	var user models.User
	if query.Next() {
		if error = query.Scan(
			&user.User_id,
			&user.Username,
			&user.Name,
			&user.Email,
		); error != nil {
			return models.User{}, nil
		}
	}
	return user, error
}

func (repository Users) FollowUser(userId, followerId uint64) error {
	statement, error := repository.db.Prepare(`
	insert into followers (user_id, follower_id) values (?,?)
	`)

	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error = statement.Exec(userId, followerId); error != nil {
		return error
	}

	return nil
}

func (repository Users) UnfollowUser(userId, followerId uint64) error {
	statement, error := repository.db.Prepare(`
	delete from followers where user_id = ? and follower_id = ?
	`)

	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error = statement.Exec(userId, followerId); error != nil {
		return error
	}

	return nil
}

func (repository Users) GetFollowedUserPosts(userId uint64) ([]models.Post, error) {
	query, error := repository.db.Query(`
	select distinct posts.post_id, posts.author_id,users.username, posts.title, posts.content, posts.music_title, posts.music_link 
	from posts inner join users on users.user_id = posts.author_id inner join followers f on posts.author_id = f.user_id 
	where users.user_id = ? or f.follower_id = ? order by 1 desc
	`, userId, userId)

	if error != nil {
		return nil, error
	}

	defer query.Close()

	var posts []models.Post

	for query.Next() {
		var post models.Post

		if error = query.Scan(
			&post.PostId,
			&post.AuthorId,
			&post.AuthorUsername,
			&post.Title,
			&post.Content,
			&post.MusicTitle,
			&post.MusicLink,
		); error != nil {
			return []models.Post{}, error
		}

		posts = append(posts, post)
	}
	return posts, nil
}

func (repository Users) GetFollowing(userId, followerId uint64) (bool, error) {
	query, error := repository.db.Query(
		"select * from followers where user_id = ? and follower_id = ?",
		userId, followerId,
	)

	if error != nil {
		return false, error
	}
	following := 0
	for query.Next() {
		following++
	}

	if following == 0 {
		return false, nil
	}

	defer query.Close()

	return true, nil

}
