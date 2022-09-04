package routes

import (
	"net/http"
	"rscm/src/auth"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequerisAuthentication bool
}

func ConfigRoutes(r *mux.Router) *mux.Router {

	routes := userRoute
	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.RequerisAuthentication {
			r.HandleFunc(route.URI, auth.AuthUser(route.Function))
		} else {
			r.HandleFunc(route.URI, route.Function).Methods(route.Method)

		}
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r

}
