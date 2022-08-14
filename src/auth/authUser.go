package auth

import (
	"net/http"
	"rscm/src/responses"
)

func AuthUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if error := ValidateToken(r); error != nil {
			responses.Error(w, http.StatusUnauthorized, error)
			return
		}

		next(w, r)
	}
}
