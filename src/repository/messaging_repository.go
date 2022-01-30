package repository

import (
	"context"

	"github.com/Sora8d/chat-app-monolith/src/clients/proto_clients"
	"github.com/Sora8d/chat-app-monolith/src/clients/rpc/messaging"
	"github.com/Sora8d/common/logger"
	"github.com/Sora8d/common/server_message"
)

type messagingRepository struct {
}

type MessagingRepositoryInterface interface {
	CreateMessage(context.Context, *messaging.CreateMessageRequest) (*messaging.Uuid, server_message.Svr_message)
	CreateConversation(context.Context, *messaging.Conversation) (*messaging.Uuid, server_message.Svr_message)
	CreateUserConversation(context.Context, *messaging.CreateUserConversationRequest) server_message.Svr_message
	KickUser(context.Context, *messaging.KickUserRequest) server_message.Svr_message
	GetConversationsByUser(context.Context, *messaging.Uuid) ([]*messaging.ConversationAndParticipants, server_message.Svr_message)
	GetMessagesByConversation(context.Context, *messaging.GetMessages) ([]*messaging.Message, server_message.Svr_message)
	UpdateMessage(context.Context, *messaging.Message) (*messaging.Message, server_message.Svr_message)
	UpdateConversationInfo(context.Context, *messaging.Conversation) (*messaging.Conversation, server_message.Svr_message)
}

func GetMessagingRepository() MessagingRepositoryInterface {
	return &messagingRepository{}
}

func (mr messagingRepository) CreateMessage(ctx context.Context, in *messaging.CreateMessageRequest) (*messaging.Uuid, server_message.Svr_message) {
	client := proto_clients.GetMessagingClient()
	response, err := client.Client.CreateMessage(ctx, in)
	if err != nil {
		logger.Error("error in messaging_repository,", err)
		return nil, server_message.NewInternalError()
	}
	return response.Uuid, server_message.NewCustomMessage(int(response.Msg.Status), response.Msg.Message)
}

func (mr messagingRepository) CreateConversation(ctx context.Context, in *messaging.Conversation) (*messaging.Uuid, server_message.Svr_message) {
	client := proto_clients.GetMessagingClient()
	response, err := client.Client.CreateConversation(ctx, in)
	if err != nil {
		logger.Error("error in messaging_repository,", err)
		return nil, server_message.NewInternalError()
	}
	return response.Uuid, server_message.NewCustomMessage(int(response.Msg.Status), response.Msg.Message)
}

func (mr messagingRepository) CreateUserConversation(ctx context.Context, in *messaging.CreateUserConversationRequest) server_message.Svr_message {
	client := proto_clients.GetMessagingClient()
	response, err := client.Client.CreateUserConversation(ctx, in)
	if err != nil {
		logger.Error("error in messaging_repository,", err)
		return server_message.NewInternalError()
	}
	return server_message.NewCustomMessage(int(response.Status), response.Message)
}

func (mr messagingRepository) GetConversationsByUser(ctx context.Context, in *messaging.Uuid) ([]*messaging.ConversationAndParticipants, server_message.Svr_message) {
	client := proto_clients.GetMessagingClient()
	response, err := client.Client.GetConversationsByUser(ctx, in)
	if err != nil {
		logger.Error("error in messaging_repository,", err)
		return nil, server_message.NewInternalError()
	}
	return response.Conversations, server_message.NewCustomMessage(int(response.Msg.Status), response.Msg.Message)
}

func (mr messagingRepository) GetMessagesByConversation(ctx context.Context, in *messaging.GetMessages) ([]*messaging.Message, server_message.Svr_message) {
	client := proto_clients.GetMessagingClient()
	response, err := client.Client.GetMessagesByConversation(ctx, in)
	if err != nil {
		logger.Error("error in messaging_repository,", err)
		return nil, server_message.NewInternalError()
	}
	return response.Message, server_message.NewCustomMessage(int(response.Msg.Status), response.Msg.Message)
}

func (mr messagingRepository) UpdateMessage(ctx context.Context, in *messaging.Message) (*messaging.Message, server_message.Svr_message) {
	client := proto_clients.GetMessagingClient()
	response, err := client.Client.UpdateMessage(ctx, in)
	if err != nil {
		logger.Error("error in messaging_repository,", err)
		return nil, server_message.NewInternalError()
	}
	return response.Message, server_message.NewCustomMessage(int(response.Msg.Status), response.Msg.Message)
}

func (mr messagingRepository) UpdateConversationInfo(ctx context.Context, in *messaging.Conversation) (*messaging.Conversation, server_message.Svr_message) {
	client := proto_clients.GetMessagingClient()
	response, err := client.Client.UpdateConversationInfo(ctx, in)
	if err != nil {
		logger.Error("error in messaging_repository,", err)
		return nil, server_message.NewInternalError()
	}
	return response.Conversation, server_message.NewCustomMessage(int(response.Msg.Status), response.Msg.Message)
}

func (mr messagingRepository) KickUser(ctx context.Context, in *messaging.KickUserRequest) server_message.Svr_message {
	client := proto_clients.GetMessagingClient()
	response, err := client.Client.KickUser(ctx, in)
	if err != nil {
		logger.Error("error in kick user", err)
		return server_message.NewInternalError()
	}
	return server_message.NewCustomMessage(int(response.GetStatus()), response.GetMessage())
}
