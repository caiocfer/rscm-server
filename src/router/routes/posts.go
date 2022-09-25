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
	{
		URI:                    "/posts",
		Method:                 http.MethodGet,
		Function:               controller.GetFollowedUserPosts,
		RequerisAuthentication: true,
	},
	{
		URI:                    "/posts/{userID}",
		Method:                 http.MethodGet,
		Function:               controller.GetUserPosts,
		RequerisAuthentication: true,
	},
}
