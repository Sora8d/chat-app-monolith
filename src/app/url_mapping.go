package app

func mapUrls() {
	messageUrls()
	usersUrls()
	oauthUrls()
}

func messageUrls() {
	router.POST("/message", messaging_controller.CreateMessage)
	router.POST("/conversation", messaging_controller.CreateConversation)
	router.POST("/conversation/add_user", messaging_controller.CreateUserConversation)
	router.POST("/conversation/kick_user", messaging_controller.KickUser)
	router.GET("/conversation/:uuid", messaging_controller.GetConversationsByUser)
	router.GET("/message/:uuid", messaging_controller.GetMessagesByConversation)
	router.PUT("/message", messaging_controller.UpdateMessage)
	router.PUT("/conversation/info", messaging_controller.UpdateConversationInfo)
}

func usersUrls() {
	router.POST("/user", users_server.CreateUser)
	router.GET("/user", users_server.GetUserProfileByUuid)
	router.PUT("/user", users_server.UpdateUser)
	router.GET("/search", users_server.SearchContact)
}

func oauthUrls() {
	router.POST("/login", oauth_controller.LoginUser)
	router.POST("/token", oauth_controller.ValidateRefreshToken)
	router.POST("/token/revoke", oauth_controller.RevokeUsersTokens)
}
