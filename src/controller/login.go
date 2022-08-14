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
	"rscm/src/security"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(body, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repositories := repository.NewUserRepo(db)
	checkUserDB, error := repositories.GetUserByEmail(user.Email)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error = security.CheckPassword(checkUserDB.Password, user.Password); error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	token, _ := auth.CreateToken(checkUserDB.User_id)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return

	}
	w.Write([]byte(token))

}
