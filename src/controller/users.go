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
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(bodyRequest, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare(); error != nil {
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
	user.User_id, error = repositories.CreateUser(user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, user)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	user := strings.ToLower(r.URL.Query().Get("user"))

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()
	repositories := repository.NewUserRepo(db)
	users, error := repositories.GetUsers(user)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, users)

}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userID, error := auth.GetUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repositories := repository.NewUserRepo(db)
	users, error := repositories.GetUserProfile(userID)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, users)

}

func GetSearchedUser(w http.ResponseWriter, r *http.Request) {
	searchedUser := strings.ToLower(r.URL.Query().Get("user"))
	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repository := repository.NewUserRepo(db)
	users, error := repository.GetSearchedUser(searchedUser)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)

	userID, error := strconv.ParseUint(ID["userid"], 10, 64)
	if error != nil {
		responses.JSON(w, http.StatusBadRequest, error)
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()

	repositories := repository.NewUserRepo(db)
	user, error := repositories.GetUserById(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, user)

}
