package server

import (
	"io/ioutil"

	"github.com/Sora8d/chat-app-monolith/src/clients/rpc/oauth"
	"github.com/Sora8d/chat-app-monolith/src/domain/response"
	"github.com/Sora8d/chat-app-monolith/src/services"
	"github.com/Sora8d/common/server_message"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

type oauthController struct {
	oauthsvs services.OauthServiceInterface
}

type OauthControllerInterface interface {
	LoginUser(*gin.Context)
	ValidateRefreshToken(c *gin.Context)
	RevokeUsersTokens(c *gin.Context)
}

func NewOauthController(oauthsvs services.OauthServiceInterface) OauthControllerInterface {
	return &oauthController{oauthsvs: oauthsvs}
}

func (oauthctrl oauthController) LoginUser(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		aErr := server_message.NewInternalError()
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	new_request := oauth.LoginRequest{}
	if err := protojson.Unmarshal(bytes, &new_request); err != nil {
		aErr := server_message.NewBadRequestError("invalid json")
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	result_response_object, headers := oauthctrl.oauthsvs.LoginUser(&new_request)
	if headers != nil {
		c.Header("access-control-expose-headers", "access-token, refresh-token")
		c.Header("access-token", headers[0])
		c.Header("refresh-token", headers[1])
	}
	c.JSON(result_response_object.Response.GetStatus(), result_response_object)
}

func (oauthctrl oauthController) ValidateRefreshToken(c *gin.Context) {
	token := c.GetHeader("refresh-token")
	if token == "" {
		response := response.Response{}.CreateResponse(nil, server_message.NewBadRequestError("nil refresh-token"))
		c.JSON(response.Response.GetStatus(), response)
		return
	}
	new_request := oauth.JWT{Jwt: token}
	result_response_object, headers := oauthctrl.oauthsvs.ValidateRefreshToken(&new_request)
	if headers != nil {
		c.Header("access-control-expose-headers", "access-token, refresh-token")
		c.Header("access-token", headers[0])
		c.Header("refresh-token", headers[1])
	}
	c.JSON(result_response_object.Response.GetStatus(), result_response_object)
}

func (oauthctrl oauthController) RevokeUsersTokens(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		aErr := server_message.NewInternalError()
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	new_request := oauth.Uuid{}
	if err := protojson.Unmarshal(bytes, &new_request); err != nil {
		aErr := server_message.NewBadRequestError("invalid json")
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	result_response_object := oauthctrl.oauthsvs.RevokeUsersTokens(&new_request)
	c.JSON(result_response_object.Response.GetStatus(), result_response_object)
}
