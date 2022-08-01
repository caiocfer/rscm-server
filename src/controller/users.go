package controller

import (
	"net/http"
	"rscm/src/db"
	"rscm/src/repository"
	"rscm/src/responses"
	"strings"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	user := strings.ToLower(r.URL.Query().Get("user"))

	db, error := db.Connect()
	if error != nil {
		//responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()
	repositories := repository.NewUserRepo(db)
	users, error := repositories.GetUsers(user)

	if error != nil {
		//responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, users)

}
