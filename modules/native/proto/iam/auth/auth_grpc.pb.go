// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: auth.proto

package auth

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

// IAMAuthServiceClient is the client API for IAMAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IAMAuthServiceClient interface {
	// OAuth. Create access token and refresh token using password
	CreateTokenWithPassword(ctx context.Context, in *CreateTokenWithPasswordRequest, opts ...grpc.CallOption) (*CreateTokenWithPasswordResponse, error)
	// OAuth. Creates new access token using refresh tokenna
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	// rpc VerifyResoureAccess(VerifyResourceAccessRequest) returns (VerifyResourceAccessResponse);
	// OAuth. Check if token is allowed to perform actions from the specified scopes
	CheckAccessWithToken(ctx context.Context, in *CheckAccessWithTokenRequest, opts ...grpc.CallOption) (*CheckAccessWithTokenResponse, error)
	// Basic Auth. Check if provided identity with proposed password is allowed to perform actions from the provided scopes
	CheckAccessWithPassword(ctx context.Context, in *CheckAccessWithPasswordRequest, opts ...grpc.CallOption) (*CheckAccessWithPasswordResponse, error)
}

type iAMAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIAMAuthServiceClient(cc grpc.ClientConnInterface) IAMAuthServiceClient {
	return &iAMAuthServiceClient{cc}
}

func (c *iAMAuthServiceClient) CreateTokenWithPassword(ctx context.Context, in *CreateTokenWithPasswordRequest, opts ...grpc.CallOption) (*CreateTokenWithPasswordResponse, error) {
	out := new(CreateTokenWithPasswordResponse)
	err := c.cc.Invoke(ctx, "/native_iam_auth.IAMAuthService/CreateTokenWithPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAuthServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, "/native_iam_auth.IAMAuthService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAuthServiceClient) CheckAccessWithToken(ctx context.Context, in *CheckAccessWithTokenRequest, opts ...grpc.CallOption) (*CheckAccessWithTokenResponse, error) {
	out := new(CheckAccessWithTokenResponse)
	err := c.cc.Invoke(ctx, "/native_iam_auth.IAMAuthService/CheckAccessWithToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMAuthServiceClient) CheckAccessWithPassword(ctx context.Context, in *CheckAccessWithPasswordRequest, opts ...grpc.CallOption) (*CheckAccessWithPasswordResponse, error) {
	out := new(CheckAccessWithPasswordResponse)
	err := c.cc.Invoke(ctx, "/native_iam_auth.IAMAuthService/CheckAccessWithPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IAMAuthServiceServer is the server API for IAMAuthService service.
// All implementations must embed UnimplementedIAMAuthServiceServer
// for forward compatibility
type IAMAuthServiceServer interface {
	// OAuth. Create access token and refresh token using password
	CreateTokenWithPassword(context.Context, *CreateTokenWithPasswordRequest) (*CreateTokenWithPasswordResponse, error)
	// OAuth. Creates new access token using refresh tokenna
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	// rpc VerifyResoureAccess(VerifyResourceAccessRequest) returns (VerifyResourceAccessResponse);
	// OAuth. Check if token is allowed to perform actions from the specified scopes
	CheckAccessWithToken(context.Context, *CheckAccessWithTokenRequest) (*CheckAccessWithTokenResponse, error)
	// Basic Auth. Check if provided identity with proposed password is allowed to perform actions from the provided scopes
	CheckAccessWithPassword(context.Context, *CheckAccessWithPasswordRequest) (*CheckAccessWithPasswordResponse, error)
	mustEmbedUnimplementedIAMAuthServiceServer()
}

// UnimplementedIAMAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIAMAuthServiceServer struct {
}

func (UnimplementedIAMAuthServiceServer) CreateTokenWithPassword(context.Context, *CreateTokenWithPasswordRequest) (*CreateTokenWithPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTokenWithPassword not implemented")
}
func (UnimplementedIAMAuthServiceServer) RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedIAMAuthServiceServer) CheckAccessWithToken(context.Context, *CheckAccessWithTokenRequest) (*CheckAccessWithTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAccessWithToken not implemented")
}
func (UnimplementedIAMAuthServiceServer) CheckAccessWithPassword(context.Context, *CheckAccessWithPasswordRequest) (*CheckAccessWithPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAccessWithPassword not implemented")
}
func (UnimplementedIAMAuthServiceServer) mustEmbedUnimplementedIAMAuthServiceServer() {}

// UnsafeIAMAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IAMAuthServiceServer will
// result in compilation errors.
type UnsafeIAMAuthServiceServer interface {
	mustEmbedUnimplementedIAMAuthServiceServer()
}

func RegisterIAMAuthServiceServer(s grpc.ServiceRegistrar, srv IAMAuthServiceServer) {
	s.RegisterService(&IAMAuthService_ServiceDesc, srv)
}

func _IAMAuthService_CreateTokenWithPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenWithPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAuthServiceServer).CreateTokenWithPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/native_iam_auth.IAMAuthService/CreateTokenWithPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAuthServiceServer).CreateTokenWithPassword(ctx, req.(*CreateTokenWithPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAuthService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAuthServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/native_iam_auth.IAMAuthService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAuthServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAuthService_CheckAccessWithToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAccessWithTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAuthServiceServer).CheckAccessWithToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/native_iam_auth.IAMAuthService/CheckAccessWithToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAuthServiceServer).CheckAccessWithToken(ctx, req.(*CheckAccessWithTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMAuthService_CheckAccessWithPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAccessWithPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMAuthServiceServer).CheckAccessWithPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/native_iam_auth.IAMAuthService/CheckAccessWithPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMAuthServiceServer).CheckAccessWithPassword(ctx, req.(*CheckAccessWithPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IAMAuthService_ServiceDesc is the grpc.ServiceDesc for IAMAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IAMAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "native_iam_auth.IAMAuthService",
	HandlerType: (*IAMAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTokenWithPassword",
			Handler:    _IAMAuthService_CreateTokenWithPassword_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _IAMAuthService_RefreshToken_Handler,
		},
		{
			MethodName: "CheckAccessWithToken",
			Handler:    _IAMAuthService_CheckAccessWithToken_Handler,
		},
		{
			MethodName: "CheckAccessWithPassword",
			Handler:    _IAMAuthService_CheckAccessWithPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
