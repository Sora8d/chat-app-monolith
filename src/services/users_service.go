package services

import (
	"context"

	"github.com/Sora8d/chat-app-monolith/src/clients/rpc/users"
	"github.com/Sora8d/chat-app-monolith/src/domain/response"
	"github.com/Sora8d/chat-app-monolith/src/repository"
)

type usersService struct {
	users_repo repository.UsersRepositoryInterface
}

type UsersServiceInterface interface {
	CreateUser(context.Context, *users.RegisterUser) response.Response
	GetUserProfileByUuid(context.Context, *users.MultipleUuids) response.Response
	UpdateUser(context.Context, *users.UpdateUserRequest) response.Response
	SearchContact(ctx context.Context, request *users.SearchContactRequest) response.Response
}

func NewUsersService(users_repo repository.UsersRepositoryInterface) UsersServiceInterface {
	return &usersService{users_repo: users_repo}
}

func (us usersService) CreateUser(ctx context.Context, request *users.RegisterUser) response.Response {
	return Response.CreateResponse(nil, us.users_repo.CreateUser(ctx, request))
}

func (us usersService) GetUserProfileByUuid(ctx context.Context, request *users.MultipleUuids) response.Response {
	return Response.CreateResponse(us.users_repo.GetUserProfileByUuid(ctx, request))

}
func (us usersService) UpdateUser(ctx context.Context, request *users.UpdateUserRequest) response.Response {
	return Response.CreateResponse(us.users_repo.UpdateUser(ctx, request))

}

func (us usersService) SearchContact(ctx context.Context, request *users.SearchContactRequest) response.Response {
	return Response.CreateResponse(us.users_repo.SearchContact(ctx, request))
}
