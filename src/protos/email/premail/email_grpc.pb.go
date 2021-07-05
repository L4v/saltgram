// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package premail

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

// EmailClient is the client API for Email service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailClient interface {
	SendActivation(ctx context.Context, in *SendActivationRequest, opts ...grpc.CallOption) (*SendActivationResponse, error)
	Activate(ctx context.Context, in *ActivateRequest, opts ...grpc.CallOption) (*ActivateResponse, error)
	RequestReset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error)
	ConfirmReset(ctx context.Context, in *ConfirmRequest, opts ...grpc.CallOption) (*ConfirmResponse, error)
}

type emailClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailClient(cc grpc.ClientConnInterface) EmailClient {
	return &emailClient{cc}
}

func (c *emailClient) SendActivation(ctx context.Context, in *SendActivationRequest, opts ...grpc.CallOption) (*SendActivationResponse, error) {
	out := new(SendActivationResponse)
	err := c.cc.Invoke(ctx, "/Email/SendActivation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailClient) Activate(ctx context.Context, in *ActivateRequest, opts ...grpc.CallOption) (*ActivateResponse, error) {
	out := new(ActivateResponse)
	err := c.cc.Invoke(ctx, "/Email/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailClient) RequestReset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error) {
	out := new(ResetResponse)
	err := c.cc.Invoke(ctx, "/Email/RequestReset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailClient) ConfirmReset(ctx context.Context, in *ConfirmRequest, opts ...grpc.CallOption) (*ConfirmResponse, error) {
	out := new(ConfirmResponse)
	err := c.cc.Invoke(ctx, "/Email/ConfirmReset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServer is the server API for Email service.
// All implementations must embed UnimplementedEmailServer
// for forward compatibility
type EmailServer interface {
	SendActivation(context.Context, *SendActivationRequest) (*SendActivationResponse, error)
	Activate(context.Context, *ActivateRequest) (*ActivateResponse, error)
	RequestReset(context.Context, *ResetRequest) (*ResetResponse, error)
	ConfirmReset(context.Context, *ConfirmRequest) (*ConfirmResponse, error)
	mustEmbedUnimplementedEmailServer()
}

// UnimplementedEmailServer must be embedded to have forward compatible implementations.
type UnimplementedEmailServer struct {
}

func (UnimplementedEmailServer) SendActivation(context.Context, *SendActivationRequest) (*SendActivationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendActivation not implemented")
}
func (UnimplementedEmailServer) Activate(context.Context, *ActivateRequest) (*ActivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
func (UnimplementedEmailServer) RequestReset(context.Context, *ResetRequest) (*ResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestReset not implemented")
}
func (UnimplementedEmailServer) ConfirmReset(context.Context, *ConfirmRequest) (*ConfirmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmReset not implemented")
}
func (UnimplementedEmailServer) mustEmbedUnimplementedEmailServer() {}

// UnsafeEmailServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailServer will
// result in compilation errors.
type UnsafeEmailServer interface {
	mustEmbedUnimplementedEmailServer()
}

func RegisterEmailServer(s grpc.ServiceRegistrar, srv EmailServer) {
	s.RegisterService(&Email_ServiceDesc, srv)
}

func _Email_SendActivation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendActivationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServer).SendActivation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Email/SendActivation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServer).SendActivation(ctx, req.(*SendActivationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Email_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Email/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServer).Activate(ctx, req.(*ActivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Email_RequestReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServer).RequestReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Email/RequestReset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServer).RequestReset(ctx, req.(*ResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Email_ConfirmReset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServer).ConfirmReset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Email/ConfirmReset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServer).ConfirmReset(ctx, req.(*ConfirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Email_ServiceDesc is the grpc.ServiceDesc for Email service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Email_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Email",
	HandlerType: (*EmailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendActivation",
			Handler:    _Email_SendActivation_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _Email_Activate_Handler,
		},
		{
			MethodName: "RequestReset",
			Handler:    _Email_RequestReset_Handler,
		},
		{
			MethodName: "ConfirmReset",
			Handler:    _Email_ConfirmReset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email/email.proto",
}
