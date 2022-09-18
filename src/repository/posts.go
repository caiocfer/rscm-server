package repository

import (
	"database/sql"
	"rscm/src/models"
)

type Posts struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repository Posts) CreateNewPost(post models.Post) (uint64, error) {
	statment, error := repository.db.Prepare("insert into posts(author_id,title,content) values (?,?,?)")

	if error != nil {
		return 0, error
	}

	defer statment.Close()

	result, error := statment.Exec(post.AuthorId, post.Title, post.Content)
	if error != nil {
		return 0, error
	}

	lastIDInserted, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastIDInserted), nil
}

func (repository Posts) GetUserPosts(userId uint64) ([]models.Post, error) {
	query, error := repository.db.Query(
		"select posts.post_id, posts.author_id,users.username, posts.title, posts.content from posts inner join users on posts.author_id=users.user_id where posts.author_id = ?",
		userId,
	)

	if error != nil {
		return []models.Post{}, error
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
		); error != nil {
			return []models.Post{}, error
		}

		posts = append(posts, post)
	}

	return posts, nil

}
