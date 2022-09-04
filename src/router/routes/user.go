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
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controller.CreateUser,
		RequerisAuthentication: false,
	},
	{
		URI:                    "/user",
		Method:                 http.MethodGet,
		Function:               controller.GetUserProfile,
		RequerisAuthentication: true,
	},
}
