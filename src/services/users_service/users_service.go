package users_service

import (
	"github.com/Sora8d/chat-app-monolith/src/domain/users_api/users"
	"github.com/Sora8d/chat-app-monolith/src/repository/users_repository/db"
	"github.com/Sora8d/common/server_message"
)

type UsersServiceInterface interface {
	CreateUser(users.RegisterUser) server_message.Svr_message
	LoginUser(users.User) (*users.User, server_message.Svr_message)
	GetUser([]string) (users.UserSlice, server_message.Svr_message)
	GetUserProfile([]string) (users.UserProfileSlice, server_message.Svr_message)
	DeleteUser(string) server_message.Svr_message
	UpdateUserProfileActive(string, bool) server_message.Svr_message

	SearchContact(string, string, []string) (users.UserProfileSlice, server_message.Svr_message)
	UpdateUserProfile(string, users.UserProfile, bool) (*users.UserProfile, server_message.Svr_message)
}

type userService struct {
	dbRepo db.UserDbRepository
}

func NewUsersService(dbrepo db.UserDbRepository) UsersServiceInterface {
	return userService{dbRepo: dbrepo}
}

func (us userService) CreateUser(uc users.RegisterUser) server_message.Svr_message {
	uc.LoginInfo.LoginPassword = users.HashPassword(uc.LoginInfo.LoginPassword)
	aErr := us.dbRepo.CreateUser(uc)
	if aErr != nil {
		return aErr
	}
	return nil
}

func (us userService) LoginUser(u users.User) (*users.User, server_message.Svr_message) {
	u.LoginPassword = users.HashPassword(u.LoginPassword)
	res_user, aerr := us.dbRepo.LoginUser(u)
	if aerr != nil {
		return nil, aerr
	}
	return res_user, nil
}

func (us userService) GetUser(uuids []string) (users.UserSlice, server_message.Svr_message) {
	users, aErr := us.dbRepo.GetUserByUuid(uuids)
	if aErr != nil {
		return nil, aErr
	}
	return users, nil
}

func (us userService) GetUserProfile(uuid []string) (users.UserProfileSlice, server_message.Svr_message) {
	user, aErr := us.dbRepo.GetUserProfileById(uuid)
	if aErr != nil {
		return nil, aErr
	}
	return user, nil
}

func (us userService) UpdateUserProfile(uuid string, updates users.UserProfile, partial bool) (*users.UserProfile, server_message.Svr_message) {
	if partial {
		profiles, aErr := us.dbRepo.GetUserProfileById([]string{uuid})
		if aErr != nil {
			return nil, aErr
		}
		profile_with_information := profiles[0]
		if updates.Phone != "" {
			profile_with_information.Phone = updates.Phone
		}
		if updates.FirstName != "" {
			profile_with_information.FirstName = updates.FirstName
		}
		if updates.LastName != "" {
			profile_with_information.LastName = updates.LastName
		}
		if updates.UserName != "" {
			profile_with_information.UserName = updates.UserName
		}
		if updates.AvatarUrl != "" {
			profile_with_information.AvatarUrl = updates.AvatarUrl
		}
		if updates.Description != "" {
			profile_with_information.Description = updates.Description
		}
		users, aErr := us.dbRepo.UpdateUserProfile(uuid, *profile_with_information)
		if aErr != nil {
			return nil, aErr
		}
		return users, nil
	}
	if updates.Phone == "" {
		return nil, server_message.NewBadRequestError("the request is marked as no_partial, but updating will bring a nil phone value")
	}
	users, aErr := us.dbRepo.UpdateUserProfile(uuid, updates)
	if aErr != nil {
		return nil, aErr
	}
	return users, nil
}

func (us userService) UpdateUserProfileActive(uuid string, active bool) server_message.Svr_message {
	aErr := us.dbRepo.UpdateUserProfileActive(uuid, active)
	if aErr != nil {
		return aErr
	}
	return nil
}

func (us userService) DeleteUser(uuid string) server_message.Svr_message {
	if aErr := us.dbRepo.DeleteUser(uuid); aErr != nil {
		return aErr
	}
	return nil
}

func (us userService) SearchContact(query, user_uuid string, exclude_uuids []string) (users.UserProfileSlice, server_message.Svr_message) {
	exclude_uuids = append(exclude_uuids, user_uuid)
	profile, aErr := us.dbRepo.SearchContact(query, exclude_uuids)
	if aErr != nil {
		return nil, aErr
	}
	return profile, nil
}
