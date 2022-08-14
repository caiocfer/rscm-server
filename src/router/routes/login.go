package routes

import (
	"net/http"
	"rscm/src/controller"
)

var loginRoute = Route{
	URI:                    "/login",
	Method:                 http.MethodPost,
	Function:               controller.Login,
	RequerisAuthentication: false,
}
