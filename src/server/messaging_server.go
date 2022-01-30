package server

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/Sora8d/chat-app-monolith/src/clients/rpc/messaging"
	"github.com/Sora8d/chat-app-monolith/src/domain/response"
	"github.com/Sora8d/chat-app-monolith/src/services"
	"github.com/Sora8d/common/server_message"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

type messagingController struct {
	msg_svs services.MessagingServiceInterface
}

type MessagingControllerInterface interface {
	CreateMessage(c *gin.Context)
	CreateConversation(c *gin.Context)
	CreateUserConversation(c *gin.Context)
	KickUser(c *gin.Context)
	GetConversationsByUser(c *gin.Context)
	GetMessagesByConversation(c *gin.Context)
	UpdateMessage(c *gin.Context)
	UpdateConversationInfo(c *gin.Context)
}

func NewMessagingController(svs services.MessagingServiceInterface) MessagingControllerInterface {
	return &messagingController{msg_svs: svs}
}

func (mc messagingController) CreateMessage(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		aErr := server_message.NewInternalError()
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	new_request := messaging.CreateMessageRequest{}
	if err := protojson.Unmarshal(bytes, &new_request); err != nil {
		aErr := server_message.NewBadRequestError("invalid json")
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	ctx := appendHeaderAccessToken(c.Request.Header, context.Background())
	result_response_object := mc.msg_svs.CreateMessage(ctx, &new_request)

	c.JSON(result_response_object.Response.GetStatus(), result_response_object)
}

func (mc messagingController) CreateConversation(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		aErr := server_message.NewInternalError()
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	new_request := messaging.Conversation{}
	if err := protojson.Unmarshal(bytes, &new_request); err != nil {
		aErr := server_message.NewBadRequestError("invalid json")
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	ctx := appendHeaderAccessToken(c.Request.Header, context.Background())
	result_response_object := mc.msg_svs.CreateConversation(ctx, &new_request)

	c.JSON(result_response_object.Response.GetStatus(), result_response_object)
}

func (mc messagingController) CreateUserConversation(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		aErr := server_message.NewInternalError()
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	new_request := messaging.CreateUserConversationRequest{}
	if err := protojson.Unmarshal(bytes, &new_request); err != nil {
		aErr := server_message.NewBadRequestError("invalid json")
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	ctx := appendHeaderAccessToken(c.Request.Header, context.Background())
	result_response_object := mc.msg_svs.CreateUserConversation(ctx, &new_request)
	c.JSON(result_response_object.Response.GetStatus(), result_response_object)
}

func (mc messagingController) KickUser(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		aErr := server_message.NewInternalError()
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	new_request := messaging.KickUserRequest{}
	if err := protojson.Unmarshal(bytes, &new_request); err != nil {
		aErr := server_message.NewBadRequestError("invalid json")
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	ctx := appendHeaderAccessToken(c.Request.Header, context.Background())
	result_response_object := mc.msg_svs.KickUser(ctx, &new_request)
	c.JSON(result_response_object.Response.GetStatus(), result_response_object)
}

func (mc messagingController) GetConversationsByUser(c *gin.Context) {
	uuid, ok := c.Params.Get("uuid")
	if !ok {
		aErr := server_message.NewBadRequestError("no uuid provided in path")
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	new_request := messaging.Uuid{Uuid: uuid}

	ctx := appendHeaderAccessToken(c.Request.Header, context.Background())
	result_response_object := mc.msg_svs.GetConversationsByUser(ctx, &new_request)

	c.JSON(result_response_object.Response.GetStatus(), result_response_object)
}

func (mc messagingController) GetMessagesByConversation(c *gin.Context) {
	uuid, ok := c.Params.Get("uuid")
	if !ok {
		aErr := server_message.NewBadRequestError("no uuid provided in path")
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	var parsing_error error
	before_date_string, b_ok := c.GetQuery("before")
	var before_date float64
	if b_ok {
		if _, err := fmt.Sscan(before_date_string, &before_date); err != nil {
			parsing_error = err
		}
	}
	after_date_string, a_ok := c.GetQuery("after")
	var after_date float64
	if a_ok {
		if _, err := fmt.Sscan(after_date_string, &after_date); err != nil {
			parsing_error = err
		}
	}
	if parsing_error != nil || (a_ok && b_ok) {
		response := response.Response{}.CreateResponse(nil, server_message.NewBadRequestError("bad before/after parameters"))
		c.JSON(response.Response.GetStatus(), response)
		return
	}
	new_request := messaging.GetMessages{Uuid: &messaging.Uuid{Uuid: uuid}, BeforeDate: before_date, AfterDate: after_date}

	ctx := appendHeaderAccessToken(c.Request.Header, context.Background())
	result_response_object := mc.msg_svs.GetMessagesByConversation(ctx, &new_request)

	c.JSON(result_response_object.Response.GetStatus(), result_response_object)
}

func (mc messagingController) UpdateMessage(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		aErr := server_message.NewInternalError()
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	new_request := messaging.Message{}
	if err := protojson.Unmarshal(bytes, &new_request); err != nil {
		aErr := server_message.NewBadRequestError("invalid json")
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	ctx := appendHeaderAccessToken(c.Request.Header, context.Background())
	result_response_object := mc.msg_svs.UpdateMessage(ctx, &new_request)

	c.JSON(result_response_object.Response.GetStatus(), result_response_object)
}

func (mc messagingController) UpdateConversationInfo(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		aErr := server_message.NewInternalError()
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	new_request := messaging.Conversation{}
	if err := protojson.Unmarshal(bytes, &new_request); err != nil {
		aErr := server_message.NewBadRequestError("invalid json")
		c.JSON(aErr.GetStatus(), aErr)
		return
	}
	ctx := appendHeaderAccessToken(c.Request.Header, context.Background())
	result_response_object := mc.msg_svs.UpdateConversationInfo(ctx, &new_request)

	c.JSON(result_response_object.Response.GetStatus(), result_response_object)
}
