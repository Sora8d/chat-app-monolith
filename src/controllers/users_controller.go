package controllers

import (
	"encoding/json"
	"errors"

	"github.com/Sora8d/chat-app-monolith/src/domain/common"
	"github.com/Sora8d/chat-app-monolith/src/domain/response"
	"github.com/Sora8d/chat-app-monolith/src/domain/users_api/users"
	"github.com/Sora8d/chat-app-monolith/src/services/users_service"
	"github.com/Sora8d/common/logger"
	"github.com/Sora8d/common/server_message"
)

type userController struct {
	svc users_service.UsersServiceInterface
}

func ScanBytes(info []byte, object interface{}) server_message.Svr_message {
	err := json.Unmarshal(info, object)
	if err != nil {
		return server_message.NewBadRequestError("bad json")
	} else {
		return nil
	}
}

type UserControllerInterface interface {
	//BODY
	CreateUser(ctx common.AuthInfo, registerReq []byte) server_message.Svr_message
	UpdateUser(ctx common.AuthInfo, updateReq []byte) (response.ResponseDataInterface, server_message.Svr_message)
	//UpdateActive(ctx context.Context, req *pb.UpdateActiveRequest) server_message.Svr_message
	//DeleteUserByUuid(ctx context.Context, uuid *pb.Uuid) server_message.Svr_message

	//GET
	SearchContact(ctx common.AuthInfoJTW, query string, excludeUuids []string) (response.ResponseDataInterface, server_message.Svr_message)
	GetUserProfileByUuid(ctx common.AuthInfoJTW, uuids []string) (response.ResponseDataInterface, server_message.Svr_message)

	//NOT REST
	//	GetUserByUuid(ctx common.AuthInfoJTW, uuid common.MultipleUuids) (response.ResponseDataInterface, server_message.Svr_message)
	//	UserLogin(ctx common.AuthInfoJTW, u *users.User) (response.ResponseDataInterface, server_message.Svr_message)
}

func (us userController) CreateUser(ctx common.AuthInfo, registerReq []byte) server_message.Svr_message {
	var userProfile users.RegisterUser
	aErr := ScanBytes(registerReq, userProfile)
	if aErr != nil {
		return aErr
	}
	aErr = us.svc.CreateUser(userProfile)
	return aErr
}

func (us userController) UpdateUser(ctx common.AuthInfo, updateReq []byte) (response.ResponseDataInterface, server_message.Svr_message) {
	at_uuid := ctx.Uuid
	if at_uuid == "" {
		logger.Error("error getting uuid", errors.New("error in UpdateUser, uuid is nil"))
		return nil, server_message.NewInternalError()
	}
	var updateObject users.UpdateUserRequest
	aErr := ScanBytes(updateReq, updateObject)
	if aErr != nil {
		return nil, aErr
	}
	if updateObject.UserProfile.Uuid != at_uuid {
		return nil, MessageBadPermission()
	}

	resp_profile, err := us.svc.UpdateUserProfile(at_uuid, updateObject.UserProfile, updateObject.Partial)
	if aErr != nil {
		return nil, err
	}

	result, aErr := users.UserProfileMsgResponse{}.SetData(resp_profile)
	if aErr != nil {
		return nil, err
	}
	return result, nil
}

//later update server_messages
func MessageBadPermission() server_message.Svr_message {
	return server_message.NewCustomMessage(403, "")
}
