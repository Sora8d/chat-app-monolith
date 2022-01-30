package users

import (
	"errors"

	"github.com/Sora8d/chat-app-monolith/src/domain/response"
	"github.com/Sora8d/common/logger"
	"github.com/Sora8d/common/server_message"
)

//Requests

type SearchContactRequest struct {
	Query        string   `json:"query"`
	ExcludeUuids []string `json:"exclude_uuids"`
}

type UserProfileUuid struct {
	Uuid        string `json:"uuid"`
	UserProfile `json:"user_profile"`
}

type UpdateUserRequest struct {
	UserProfileUuid `json:"content"`
	Partial         bool `json:"partial"`
}

// Responses
type UserMsgResponse struct {
	Users []User `json:"users"`
}

func (resp *UserMsgResponse) SetData(data interface{}) server_message.Svr_message {
	typedata, ok := data.([]User)
	if !ok {
		logger.Error("bad type ad UserMsgResponse SetData", errors.New("bad type"))
		return server_message.NewInternalError()
	}
	resp.Users = typedata
	return nil
}

type UserProfileMsgResponse struct {
	User []UserProfile `json:"user"`
}

func (resp UserProfileMsgResponse) SetData(data interface{}) (response.ResponseDataInterface, server_message.Svr_message) {
	typedata, ok := data.([]UserProfile)
	if !ok {
		logger.Error("bad type ad UserMsgResponse SetData", errors.New("bad type"))
		return nil, server_message.NewInternalError()
	}
	resp.User = typedata
	return resp, nil
}
