package services

import (
	"context"

	"github.com/Sora8d/chat-app-monolith/src/clients/rpc/messaging"
	"github.com/Sora8d/chat-app-monolith/src/domain/response"
	"github.com/Sora8d/chat-app-monolith/src/repository"
)

var Response response.Response

type messagingService struct {
	msg_repo repository.MessagingRepositoryInterface
}

type MessagingServiceInterface interface {
	CreateMessage(context.Context, *messaging.CreateMessageRequest) response.Response
	CreateConversation(context.Context, *messaging.Conversation) response.Response
	CreateUserConversation(context.Context, *messaging.CreateUserConversationRequest) response.Response
	KickUser(context.Context, *messaging.KickUserRequest) response.Response
	GetConversationsByUser(context.Context, *messaging.Uuid) response.Response
	GetMessagesByConversation(context.Context, *messaging.GetMessages) response.Response
	UpdateMessage(context.Context, *messaging.Message) response.Response
	UpdateConversationInfo(context.Context, *messaging.Conversation) response.Response
}

func NewMessagingService(msg_repo repository.MessagingRepositoryInterface) MessagingServiceInterface {
	return &messagingService{msg_repo: msg_repo}
}

//TODO: later create context

func (ms messagingService) CreateMessage(ctx context.Context, request *messaging.CreateMessageRequest) response.Response {
	return Response.CreateResponse(ms.msg_repo.CreateMessage(ctx, request))

}

func (ms messagingService) CreateConversation(ctx context.Context, request *messaging.Conversation) response.Response {
	return Response.CreateResponse(ms.msg_repo.CreateConversation(ctx, request))
}

func (ms messagingService) CreateUserConversation(ctx context.Context, request *messaging.CreateUserConversationRequest) response.Response {
	return Response.CreateResponse(nil, ms.msg_repo.CreateUserConversation(ctx, request))
}

func (ms messagingService) KickUser(ctx context.Context, request *messaging.KickUserRequest) response.Response {
	return Response.CreateResponse(nil, ms.msg_repo.KickUser(ctx, request))
}

func (ms messagingService) GetConversationsByUser(ctx context.Context, request *messaging.Uuid) response.Response {
	return Response.CreateResponse(ms.msg_repo.GetConversationsByUser(ctx, request))
}

func (ms messagingService) GetMessagesByConversation(ctx context.Context, request *messaging.GetMessages) response.Response {
	return Response.CreateResponse(ms.msg_repo.GetMessagesByConversation(ctx, request))
}

func (ms messagingService) UpdateMessage(ctx context.Context, request *messaging.Message) response.Response {
	return Response.CreateResponse(ms.msg_repo.UpdateMessage(ctx, request))
}

func (ms messagingService) UpdateConversationInfo(ctx context.Context, request *messaging.Conversation) response.Response {
	return Response.CreateResponse(ms.msg_repo.UpdateConversationInfo(ctx, request))
}
