package routes

import (
	"net/http"
	"rscm/src/controller"
)

var userRoute = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controller.GetUsers,
		RequerisAuthentication: true,
	},
}
