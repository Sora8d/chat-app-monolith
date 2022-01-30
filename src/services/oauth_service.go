package services

import (
	"context"

	"github.com/Sora8d/chat-app-monolith/src/clients/rpc/oauth"
	"github.com/Sora8d/chat-app-monolith/src/domain/response"
	"github.com/Sora8d/chat-app-monolith/src/repository"
	"github.com/Sora8d/common/server_message"
)

type oauthService struct {
	oauthrepo repository.OauthRepositoryInterface
}

type OauthServiceInterface interface {
	LoginUser(*oauth.LoginRequest) (response.Response, []string)
	ValidateRefreshToken(request *oauth.JWT) (response.Response, []string)
	RevokeUsersTokens(request *oauth.Uuid) response.Response
}

func NewOauthService(oauthrepo repository.OauthRepositoryInterface) OauthServiceInterface {
	return &oauthService{oauthrepo: oauthrepo}
}

func (oauthsvs oauthService) LoginUser(request *oauth.LoginRequest) (response.Response, []string) {
	ctx := context.Background()
	response, err := oauthsvs.oauthrepo.LoginUser(ctx, request)
	if err != nil {
		return Response.CreateResponse(nil, err), nil
	}
	return Response.CreateResponse(response.Uuid, server_message.NewCustomMessage(int(response.Response.Status), response.Response.Message)), []string{response.AccessToken, response.RefreshToken}
}

func (oauthsvs oauthService) ValidateRefreshToken(request *oauth.JWT) (response.Response, []string) {
	ctx := context.Background()
	response, err := oauthsvs.oauthrepo.ValidateRefreshToken(ctx, request)
	if err != nil {
		return Response.CreateResponse(nil, err), nil
	}
	return Response.CreateResponse(response.Uuid, server_message.NewCustomMessage(int(response.Response.Status), response.Response.Message)), []string{response.AccessToken, response.RefreshToken}

}

func (oauthsvs oauthService) RevokeUsersTokens(request *oauth.Uuid) response.Response {
	ctx := context.Background()
	response_message := oauthsvs.oauthrepo.RevokeUsersTokens(ctx, request)
	return Response.CreateResponse(nil, response_message)
}
