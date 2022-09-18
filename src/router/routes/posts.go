package routes

import (
	"net/http"
	"rscm/src/controller"
)

var postRoute = []Route{
	{
		URI:                    "/posts",
		Method:                 http.MethodPost,
		Function:               controller.CreatePost,
		RequerisAuthentication: true,
	},
}
