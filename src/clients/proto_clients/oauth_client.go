package proto_clients

import (
	"fmt"

	proto_oauth "github.com/Sora8d/chat-app-monolith/src/clients/rpc/oauth"
	"github.com/Sora8d/chat-app-monolith/src/config"
	"github.com/Sora8d/common/logger"
	"google.golang.org/grpc"
)

var oauth_proto_client oauthProtoClient

type oauthProtoClient struct {
	Client proto_oauth.OauthProtoInterfaceClient
	Conn   *grpc.ClientConn
}

func (lpc *oauthProtoClient) setConnection(c *grpc.ClientConn) {
	lpc.Conn = c
}
func (lpc *oauthProtoClient) setClient(c proto_oauth.OauthProtoInterfaceClient) {
	lpc.Client = c
}
func init() {
	logger.Info(fmt.Sprintf("connecting to oauth service with address: %s", config.Config["OAUTH_ADDRESS"]))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	connection, err := grpc.Dial(config.Config["OAUTH_ADDRESS"], opts...)
	if err != nil {
		logger.Error("unable to connect to oauth_api", err)
		panic(err)
	}

	oauth_proto_client.setConnection(connection)
	oauth_proto_client.setClient(proto_oauth.NewOauthProtoInterfaceClient(connection))
}

func GetOauthClient() *oauthProtoClient {
	return &oauth_proto_client
}
