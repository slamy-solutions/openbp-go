// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: token.proto

package token

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

// IAMTokenServiceClient is the client API for IAMTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IAMTokenServiceClient interface {
	// Create new token
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Get token data using token UUID (unique identifier)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Get token data using raw access/refresh token. Validates if token still exists in the system.
	RawGet(ctx context.Context, in *RawGetRequest, opts ...grpc.CallOption) (*RawGetResponse, error)
	// Delete token using token UUID (unique identifier)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Disable token using its unique identifier
	Disable(ctx context.Context, in *DisableRequest, opts ...grpc.CallOption) (*DisableResponse, error)
	// Validates token and gets its data
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
	// Validates refresh token and create new token based on it. New token will have same scopes
	Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error)
	// Returns list of tokens for specified identity
	GetTokensForIdentity(ctx context.Context, in *GetTokensForIdentityRequest, opts ...grpc.CallOption) (IAMTokenService_GetTokensForIdentityClient, error)
}

type iAMTokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIAMTokenServiceClient(cc grpc.ClientConnInterface) IAMTokenServiceClient {
	return &iAMTokenServiceClient{cc}
}

func (c *iAMTokenServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/native_iam_token.IAMTokenService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMTokenServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/native_iam_token.IAMTokenService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMTokenServiceClient) RawGet(ctx context.Context, in *RawGetRequest, opts ...grpc.CallOption) (*RawGetResponse, error) {
	out := new(RawGetResponse)
	err := c.cc.Invoke(ctx, "/native_iam_token.IAMTokenService/RawGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMTokenServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/native_iam_token.IAMTokenService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMTokenServiceClient) Disable(ctx context.Context, in *DisableRequest, opts ...grpc.CallOption) (*DisableResponse, error) {
	out := new(DisableResponse)
	err := c.cc.Invoke(ctx, "/native_iam_token.IAMTokenService/Disable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMTokenServiceClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/native_iam_token.IAMTokenService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMTokenServiceClient) Refresh(ctx context.Context, in *RefreshRequest, opts ...grpc.CallOption) (*RefreshResponse, error) {
	out := new(RefreshResponse)
	err := c.cc.Invoke(ctx, "/native_iam_token.IAMTokenService/Refresh", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMTokenServiceClient) GetTokensForIdentity(ctx context.Context, in *GetTokensForIdentityRequest, opts ...grpc.CallOption) (IAMTokenService_GetTokensForIdentityClient, error) {
	stream, err := c.cc.NewStream(ctx, &IAMTokenService_ServiceDesc.Streams[0], "/native_iam_token.IAMTokenService/GetTokensForIdentity", opts...)
	if err != nil {
		return nil, err
	}
	x := &iAMTokenServiceGetTokensForIdentityClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IAMTokenService_GetTokensForIdentityClient interface {
	Recv() (*GetTokensForIdentityResponse, error)
	grpc.ClientStream
}

type iAMTokenServiceGetTokensForIdentityClient struct {
	grpc.ClientStream
}

func (x *iAMTokenServiceGetTokensForIdentityClient) Recv() (*GetTokensForIdentityResponse, error) {
	m := new(GetTokensForIdentityResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IAMTokenServiceServer is the server API for IAMTokenService service.
// All implementations must embed UnimplementedIAMTokenServiceServer
// for forward compatibility
type IAMTokenServiceServer interface {
	// Create new token
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Get token data using token UUID (unique identifier)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Get token data using raw access/refresh token. Validates if token still exists in the system.
	RawGet(context.Context, *RawGetRequest) (*RawGetResponse, error)
	// Delete token using token UUID (unique identifier)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Disable token using its unique identifier
	Disable(context.Context, *DisableRequest) (*DisableResponse, error)
	// Validates token and gets its data
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
	// Validates refresh token and create new token based on it. New token will have same scopes
	Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error)
	// Returns list of tokens for specified identity
	GetTokensForIdentity(*GetTokensForIdentityRequest, IAMTokenService_GetTokensForIdentityServer) error
	mustEmbedUnimplementedIAMTokenServiceServer()
}

// UnimplementedIAMTokenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIAMTokenServiceServer struct {
}

func (UnimplementedIAMTokenServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedIAMTokenServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedIAMTokenServiceServer) RawGet(context.Context, *RawGetRequest) (*RawGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RawGet not implemented")
}
func (UnimplementedIAMTokenServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedIAMTokenServiceServer) Disable(context.Context, *DisableRequest) (*DisableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedIAMTokenServiceServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedIAMTokenServiceServer) Refresh(context.Context, *RefreshRequest) (*RefreshResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refresh not implemented")
}
func (UnimplementedIAMTokenServiceServer) GetTokensForIdentity(*GetTokensForIdentityRequest, IAMTokenService_GetTokensForIdentityServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTokensForIdentity not implemented")
}
func (UnimplementedIAMTokenServiceServer) mustEmbedUnimplementedIAMTokenServiceServer() {}

// UnsafeIAMTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IAMTokenServiceServer will
// result in compilation errors.
type UnsafeIAMTokenServiceServer interface {
	mustEmbedUnimplementedIAMTokenServiceServer()
}

func RegisterIAMTokenServiceServer(s grpc.ServiceRegistrar, srv IAMTokenServiceServer) {
	s.RegisterService(&IAMTokenService_ServiceDesc, srv)
}

func _IAMTokenService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMTokenServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/native_iam_token.IAMTokenService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMTokenServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMTokenService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMTokenServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/native_iam_token.IAMTokenService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMTokenServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMTokenService_RawGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMTokenServiceServer).RawGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/native_iam_token.IAMTokenService/RawGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMTokenServiceServer).RawGet(ctx, req.(*RawGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMTokenService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMTokenServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/native_iam_token.IAMTokenService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMTokenServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMTokenService_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMTokenServiceServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/native_iam_token.IAMTokenService/Disable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMTokenServiceServer).Disable(ctx, req.(*DisableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMTokenService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMTokenServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/native_iam_token.IAMTokenService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMTokenServiceServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMTokenService_Refresh_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMTokenServiceServer).Refresh(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/native_iam_token.IAMTokenService/Refresh",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMTokenServiceServer).Refresh(ctx, req.(*RefreshRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMTokenService_GetTokensForIdentity_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTokensForIdentityRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IAMTokenServiceServer).GetTokensForIdentity(m, &iAMTokenServiceGetTokensForIdentityServer{stream})
}

type IAMTokenService_GetTokensForIdentityServer interface {
	Send(*GetTokensForIdentityResponse) error
	grpc.ServerStream
}

type iAMTokenServiceGetTokensForIdentityServer struct {
	grpc.ServerStream
}

func (x *iAMTokenServiceGetTokensForIdentityServer) Send(m *GetTokensForIdentityResponse) error {
	return x.ServerStream.SendMsg(m)
}

// IAMTokenService_ServiceDesc is the grpc.ServiceDesc for IAMTokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IAMTokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "native_iam_token.IAMTokenService",
	HandlerType: (*IAMTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _IAMTokenService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _IAMTokenService_Get_Handler,
		},
		{
			MethodName: "RawGet",
			Handler:    _IAMTokenService_RawGet_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _IAMTokenService_Delete_Handler,
		},
		{
			MethodName: "Disable",
			Handler:    _IAMTokenService_Disable_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _IAMTokenService_Validate_Handler,
		},
		{
			MethodName: "Refresh",
			Handler:    _IAMTokenService_Refresh_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTokensForIdentity",
			Handler:       _IAMTokenService_GetTokensForIdentity_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "token.proto",
}
