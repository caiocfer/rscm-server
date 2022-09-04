package auth

import (
	"errors"
	"fmt"
	"net/http"
	"rscm/src/config"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 720).Unix()
	permissions["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return " "
}

func ValidateToken(r *http.Request) error {
	tokenString := getToken(r)
	token, error := jwt.Parse(tokenString, returnVerificationKey)

	if error != nil {
		return error
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Invalid Token")
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signature methods! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

func GetUserID(r *http.Request) (uint64, error) {
	tokenString := getToken(r)
	token, error := jwt.Parse(tokenString, returnVerificationKey)

	if error != nil {
		return 0, error
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, error := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userID"]), 10, 64)
		if error != nil {
			return 0, error
		}

		return userID, nil
	}
	return 0, errors.New("Invalid Token")

}
