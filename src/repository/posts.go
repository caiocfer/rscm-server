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
