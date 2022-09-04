package routes

import (
	"net/http"
	"rscm/src/controller"
)

var statusRoute = []Route{
	{
		URI:                    "/.status",
		Method:                 http.MethodGet,
		Function:               controller.Ready,
		RequerisAuthentication: false,
	},
}
