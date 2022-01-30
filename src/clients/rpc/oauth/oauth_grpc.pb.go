// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package oauth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OauthProtoInterfaceClient is the client API for OauthProtoInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OauthProtoInterfaceClient interface {
	LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*JWTwRrefreshUuidResponse, error)
	ValidateRefreshToken(ctx context.Context, in *JWT, opts ...grpc.CallOption) (*JWTwRrefreshUuidResponse, error)
	RevokeUsersTokens(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*SvrMsg, error)
}

type oauthProtoInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewOauthProtoInterfaceClient(cc grpc.ClientConnInterface) OauthProtoInterfaceClient {
	return &oauthProtoInterfaceClient{cc}
}

func (c *oauthProtoInterfaceClient) LoginUser(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*JWTwRrefreshUuidResponse, error) {
	out := new(JWTwRrefreshUuidResponse)
	err := c.cc.Invoke(ctx, "/flydev_chat_oauth.OauthProtoInterface/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauthProtoInterfaceClient) ValidateRefreshToken(ctx context.Context, in *JWT, opts ...grpc.CallOption) (*JWTwRrefreshUuidResponse, error) {
	out := new(JWTwRrefreshUuidResponse)
	err := c.cc.Invoke(ctx, "/flydev_chat_oauth.OauthProtoInterface/ValidateRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauthProtoInterfaceClient) RevokeUsersTokens(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*SvrMsg, error) {
	out := new(SvrMsg)
	err := c.cc.Invoke(ctx, "/flydev_chat_oauth.OauthProtoInterface/RevokeUsersTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OauthProtoInterfaceServer is the server API for OauthProtoInterface service.
// All implementations must embed UnimplementedOauthProtoInterfaceServer
// for forward compatibility
type OauthProtoInterfaceServer interface {
	LoginUser(context.Context, *LoginRequest) (*JWTwRrefreshUuidResponse, error)
	ValidateRefreshToken(context.Context, *JWT) (*JWTwRrefreshUuidResponse, error)
	RevokeUsersTokens(context.Context, *Uuid) (*SvrMsg, error)
	mustEmbedUnimplementedOauthProtoInterfaceServer()
}

// UnimplementedOauthProtoInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedOauthProtoInterfaceServer struct {
}

func (UnimplementedOauthProtoInterfaceServer) LoginUser(context.Context, *LoginRequest) (*JWTwRrefreshUuidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedOauthProtoInterfaceServer) ValidateRefreshToken(context.Context, *JWT) (*JWTwRrefreshUuidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateRefreshToken not implemented")
}
func (UnimplementedOauthProtoInterfaceServer) RevokeUsersTokens(context.Context, *Uuid) (*SvrMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeUsersTokens not implemented")
}
func (UnimplementedOauthProtoInterfaceServer) mustEmbedUnimplementedOauthProtoInterfaceServer() {}

// UnsafeOauthProtoInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OauthProtoInterfaceServer will
// result in compilation errors.
type UnsafeOauthProtoInterfaceServer interface {
	mustEmbedUnimplementedOauthProtoInterfaceServer()
}

func RegisterOauthProtoInterfaceServer(s grpc.ServiceRegistrar, srv OauthProtoInterfaceServer) {
	s.RegisterService(&OauthProtoInterface_ServiceDesc, srv)
}

func _OauthProtoInterface_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OauthProtoInterfaceServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flydev_chat_oauth.OauthProtoInterface/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OauthProtoInterfaceServer).LoginUser(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OauthProtoInterface_ValidateRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JWT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OauthProtoInterfaceServer).ValidateRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flydev_chat_oauth.OauthProtoInterface/ValidateRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OauthProtoInterfaceServer).ValidateRefreshToken(ctx, req.(*JWT))
	}
	return interceptor(ctx, in, info, handler)
}

func _OauthProtoInterface_RevokeUsersTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Uuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OauthProtoInterfaceServer).RevokeUsersTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flydev_chat_oauth.OauthProtoInterface/RevokeUsersTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OauthProtoInterfaceServer).RevokeUsersTokens(ctx, req.(*Uuid))
	}
	return interceptor(ctx, in, info, handler)
}

// OauthProtoInterface_ServiceDesc is the grpc.ServiceDesc for OauthProtoInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OauthProtoInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flydev_chat_oauth.OauthProtoInterface",
	HandlerType: (*OauthProtoInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginUser",
			Handler:    _OauthProtoInterface_LoginUser_Handler,
		},
		{
			MethodName: "ValidateRefreshToken",
			Handler:    _OauthProtoInterface_ValidateRefreshToken_Handler,
		},
		{
			MethodName: "RevokeUsersTokens",
			Handler:    _OauthProtoInterface_RevokeUsersTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/clients/rpc/oauth/oauth.proto",
}
