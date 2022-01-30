package app

import (
	"github.com/Sora8d/chat-app-monolith/src/config"
	"github.com/Sora8d/chat-app-monolith/src/repository"
	"github.com/Sora8d/chat-app-monolith/src/server"
	"github.com/Sora8d/chat-app-monolith/src/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router               = gin.Default()
	messaging_controller server.MessagingControllerInterface
	users_server         server.UsersServerInterface
	oauth_controller     server.OauthControllerInterface
)

func StartApplication() {
	new_config := cors.DefaultConfig()
	new_config.AllowOrigins = append(new_config.AllowOrigins, "*")
	new_config.AllowMethods = append(new_config.AllowMethods, "OPTION")
	new_config.AllowHeaders = append(new_config.AllowHeaders, "access-token", "refresh-token")
	router.Use(cors.New(new_config))
	messaging_controller = server.NewMessagingController(services.NewMessagingService(repository.GetMessagingRepository()))
	users_server = server.NewUsersServer(services.NewUsersService(repository.GetUsersRepository()))
	oauth_controller = server.NewOauthController(services.NewOauthService(repository.GetOauthRepository()))
	mapUrls()
	router.Run(config.Config["ADDRESS"] + config.Config["PORT"])
}
