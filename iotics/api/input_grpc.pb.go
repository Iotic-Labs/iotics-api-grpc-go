// Copyright (c) 2019-2022 Iotic Labs Ltd. All rights reserved.

// Iotics Web protocol definitions (input)

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: iotics/api/input.proto

package ioticsapi

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

const (
	InputAPI_DeleteInput_FullMethodName          = "/iotics.api.InputAPI/DeleteInput"
	InputAPI_DescribeInput_FullMethodName        = "/iotics.api.InputAPI/DescribeInput"
	InputAPI_ReceiveInputMessages_FullMethodName = "/iotics.api.InputAPI/ReceiveInputMessages"
)

// InputAPIClient is the client API for InputAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InputAPIClient interface {
	// Deletes an input. (Idempotent)
	DeleteInput(ctx context.Context, in *DeleteInputRequest, opts ...grpc.CallOption) (*DeleteInputResponse, error)
	// Describes an input. (local and remote)
	DescribeInput(ctx context.Context, in *DescribeInputRequest, opts ...grpc.CallOption) (*DescribeInputResponse, error)
	// Receives input messages for a specific input.
	ReceiveInputMessages(ctx context.Context, in *ReceiveInputMessageRequest, opts ...grpc.CallOption) (InputAPI_ReceiveInputMessagesClient, error)
}

type inputAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewInputAPIClient(cc grpc.ClientConnInterface) InputAPIClient {
	return &inputAPIClient{cc}
}

func (c *inputAPIClient) DeleteInput(ctx context.Context, in *DeleteInputRequest, opts ...grpc.CallOption) (*DeleteInputResponse, error) {
	out := new(DeleteInputResponse)
	err := c.cc.Invoke(ctx, InputAPI_DeleteInput_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inputAPIClient) DescribeInput(ctx context.Context, in *DescribeInputRequest, opts ...grpc.CallOption) (*DescribeInputResponse, error) {
	out := new(DescribeInputResponse)
	err := c.cc.Invoke(ctx, InputAPI_DescribeInput_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inputAPIClient) ReceiveInputMessages(ctx context.Context, in *ReceiveInputMessageRequest, opts ...grpc.CallOption) (InputAPI_ReceiveInputMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &InputAPI_ServiceDesc.Streams[0], InputAPI_ReceiveInputMessages_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &inputAPIReceiveInputMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InputAPI_ReceiveInputMessagesClient interface {
	Recv() (*ReceiveInputMessageResponse, error)
	grpc.ClientStream
}

type inputAPIReceiveInputMessagesClient struct {
	grpc.ClientStream
}

func (x *inputAPIReceiveInputMessagesClient) Recv() (*ReceiveInputMessageResponse, error) {
	m := new(ReceiveInputMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// InputAPIServer is the server API for InputAPI service.
// All implementations should embed UnimplementedInputAPIServer
// for forward compatibility
type InputAPIServer interface {
	// Deletes an input. (Idempotent)
	DeleteInput(context.Context, *DeleteInputRequest) (*DeleteInputResponse, error)
	// Describes an input. (local and remote)
	DescribeInput(context.Context, *DescribeInputRequest) (*DescribeInputResponse, error)
	// Receives input messages for a specific input.
	ReceiveInputMessages(*ReceiveInputMessageRequest, InputAPI_ReceiveInputMessagesServer) error
}

// UnimplementedInputAPIServer should be embedded to have forward compatible implementations.
type UnimplementedInputAPIServer struct {
}

func (UnimplementedInputAPIServer) DeleteInput(context.Context, *DeleteInputRequest) (*DeleteInputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInput not implemented")
}
func (UnimplementedInputAPIServer) DescribeInput(context.Context, *DescribeInputRequest) (*DescribeInputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeInput not implemented")
}
func (UnimplementedInputAPIServer) ReceiveInputMessages(*ReceiveInputMessageRequest, InputAPI_ReceiveInputMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveInputMessages not implemented")
}

// UnsafeInputAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InputAPIServer will
// result in compilation errors.
type UnsafeInputAPIServer interface {
	mustEmbedUnimplementedInputAPIServer()
}

func RegisterInputAPIServer(s grpc.ServiceRegistrar, srv InputAPIServer) {
	s.RegisterService(&InputAPI_ServiceDesc, srv)
}

func _InputAPI_DeleteInput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InputAPIServer).DeleteInput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InputAPI_DeleteInput_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InputAPIServer).DeleteInput(ctx, req.(*DeleteInputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InputAPI_DescribeInput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeInputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InputAPIServer).DescribeInput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InputAPI_DescribeInput_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InputAPIServer).DescribeInput(ctx, req.(*DescribeInputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InputAPI_ReceiveInputMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveInputMessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InputAPIServer).ReceiveInputMessages(m, &inputAPIReceiveInputMessagesServer{stream})
}

type InputAPI_ReceiveInputMessagesServer interface {
	Send(*ReceiveInputMessageResponse) error
	grpc.ServerStream
}

type inputAPIReceiveInputMessagesServer struct {
	grpc.ServerStream
}

func (x *inputAPIReceiveInputMessagesServer) Send(m *ReceiveInputMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

// InputAPI_ServiceDesc is the grpc.ServiceDesc for InputAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InputAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iotics.api.InputAPI",
	HandlerType: (*InputAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteInput",
			Handler:    _InputAPI_DeleteInput_Handler,
		},
		{
			MethodName: "DescribeInput",
			Handler:    _InputAPI_DescribeInput_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveInputMessages",
			Handler:       _InputAPI_ReceiveInputMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "iotics/api/input.proto",
}
