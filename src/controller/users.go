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

func getRepository(w http.ResponseWriter) *repository.Users {

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return nil
	}

	defer db.Close()

	repositories := repository.NewUserRepo(db)
	return repositories

}

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

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, error := auth.GetUserID(r)
	parameters := mux.Vars(r)

	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
	}

	userId, error := strconv.ParseUint(parameters["userid"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()
	repository := repository.NewUserRepo(db)

	if error = repository.FollowUser(userId, followerId); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, error := auth.GetUserID(r)
	parameters := mux.Vars(r)

	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
	}

	userId, error := strconv.ParseUint(parameters["userid"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()
	repository := repository.NewUserRepo(db)

	if error = repository.UnfollowUser(userId, followerId); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func GetFollowedUserPosts(w http.ResponseWriter, r *http.Request) {
	userId, error := auth.GetUserID(r)
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
	posts, error := repositories.GetFollowedUserPosts(userId)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, posts)

}

func GetFollowing(w http.ResponseWriter, r *http.Request) {
	userId, error := auth.GetUserID(r)
	parameters := mux.Vars(r)

	var follow = false

	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
	}

	followerId, error := strconv.ParseUint(parameters["followerid"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := db.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer db.Close()
	repository := repository.NewUserRepo(db)

	if follow, error = repository.GetFollowing(userId, followerId); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if follow {
		responses.JSON(w, http.StatusOK, nil)
	} else {
		responses.JSON(w, http.StatusNoContent, nil)

	}

}
