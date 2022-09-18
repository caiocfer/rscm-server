package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rscm/src/auth"
	"rscm/src/db"
	"rscm/src/models"
	"rscm/src/repository"
	"rscm/src/responses"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, error := auth.GetUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
	}

	bodyRequest, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var post models.Post
	if error = json.Unmarshal(bodyRequest, &post); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	post.AuthorId = userID

	if error = post.Prepare(); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repositories := repository.NewPostRepo(db)
	post.PostId, error = repositories.CreateNewPost(post)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, post)

}
