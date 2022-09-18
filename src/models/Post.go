package models

import (
	"errors"
	"strings"
)

type Post struct {
	PostId         uint64 `json:"post_id,omitempty"`
	AuthorId       uint64 `json:"author_id,omitempty"`
	AuthorUsername string `json:"author_username,omitempty"`
	Title          string `json:"title,omitempty"`
	Content        string `json:"content,omitempty"`
}

func (post *Post) validateFields() error {
	if post.Title == "" {
		return errors.New("Title can't be empty")
	}
	if post.Content == "" {
		return errors.New("Post content can't be empty")
	}

	return nil
}

func (post *Post) formatFields() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}

func (post *Post) Prepare() error {
	if error := post.validateFields(); error != nil {
		return error
	}

	post.formatFields()
	return nil
}
