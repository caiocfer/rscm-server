package routes

import (
	"rscm/src/router/routes"

	"github.com/gorilla/mux"
)

func GenerateRoutes() *mux.Router {
	router := mux.NewRouter()

	return routes.ConfigRoutes(router)
}
