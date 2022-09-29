package routes

import (
	"net/http"
	"rscm/src/controller"
)

var followRoute = []Route{
	{
		URI:                    "/users/follow/{userid}",
		Method:                 http.MethodPost,
		Function:               controller.FollowUser,
		RequerisAuthentication: true,
	},
	{
		URI:                    "/users/unfollow/{userid}",
		Method:                 http.MethodPost,
		Function:               controller.UnfollowUser,
		RequerisAuthentication: true,
	},
	{
		URI:                    "/users/follow/{followerid}",
		Method:                 http.MethodGet,
		Function:               controller.GetFollowing,
		RequerisAuthentication: true,
	},
}
