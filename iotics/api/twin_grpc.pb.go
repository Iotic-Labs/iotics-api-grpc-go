// Copyright (c) 2019-2022 Iotic Labs Ltd. All rights reserved.

// Iotics Web protocol definitions (twins)

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: iotics/api/twin.proto

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
	TwinAPI_CreateTwin_FullMethodName   = "/iotics.api.TwinAPI/CreateTwin"
	TwinAPI_UpsertTwin_FullMethodName   = "/iotics.api.TwinAPI/UpsertTwin"
	TwinAPI_DeleteTwin_FullMethodName   = "/iotics.api.TwinAPI/DeleteTwin"
	TwinAPI_UpdateTwin_FullMethodName   = "/iotics.api.TwinAPI/UpdateTwin"
	TwinAPI_DescribeTwin_FullMethodName = "/iotics.api.TwinAPI/DescribeTwin"
	TwinAPI_ListAllTwins_FullMethodName = "/iotics.api.TwinAPI/ListAllTwins"
)

// TwinAPIClient is the client API for TwinAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TwinAPIClient interface {
	// CreateTwin creates a twin.
	CreateTwin(ctx context.Context, in *CreateTwinRequest, opts ...grpc.CallOption) (*CreateTwinResponse, error)
	// UpsertTwin creates or update a twin with its metadata + the twin feeds with their metadata.
	// The full state is applied (ie. if the operation succeeds the state of the twin/feeds will be the one
	// described in the payload)
	UpsertTwin(ctx context.Context, in *UpsertTwinRequest, opts ...grpc.CallOption) (*UpsertTwinResponse, error)
	// DeleteTwin deletes a twin.
	DeleteTwin(ctx context.Context, in *DeleteTwinRequest, opts ...grpc.CallOption) (*DeleteTwinResponse, error)
	// UpdateTwin updates a twin (partial update).
	UpdateTwin(ctx context.Context, in *UpdateTwinRequest, opts ...grpc.CallOption) (*UpdateTwinResponse, error)
	// Describes a twin. (local and remote)
	DescribeTwin(ctx context.Context, in *DescribeTwinRequest, opts ...grpc.CallOption) (*DescribeTwinResponse, error)
	// List all twins.
	ListAllTwins(ctx context.Context, in *ListAllTwinsRequest, opts ...grpc.CallOption) (*ListAllTwinsResponse, error)
}

type twinAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewTwinAPIClient(cc grpc.ClientConnInterface) TwinAPIClient {
	return &twinAPIClient{cc}
}

func (c *twinAPIClient) CreateTwin(ctx context.Context, in *CreateTwinRequest, opts ...grpc.CallOption) (*CreateTwinResponse, error) {
	out := new(CreateTwinResponse)
	err := c.cc.Invoke(ctx, TwinAPI_CreateTwin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twinAPIClient) UpsertTwin(ctx context.Context, in *UpsertTwinRequest, opts ...grpc.CallOption) (*UpsertTwinResponse, error) {
	out := new(UpsertTwinResponse)
	err := c.cc.Invoke(ctx, TwinAPI_UpsertTwin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twinAPIClient) DeleteTwin(ctx context.Context, in *DeleteTwinRequest, opts ...grpc.CallOption) (*DeleteTwinResponse, error) {
	out := new(DeleteTwinResponse)
	err := c.cc.Invoke(ctx, TwinAPI_DeleteTwin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twinAPIClient) UpdateTwin(ctx context.Context, in *UpdateTwinRequest, opts ...grpc.CallOption) (*UpdateTwinResponse, error) {
	out := new(UpdateTwinResponse)
	err := c.cc.Invoke(ctx, TwinAPI_UpdateTwin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twinAPIClient) DescribeTwin(ctx context.Context, in *DescribeTwinRequest, opts ...grpc.CallOption) (*DescribeTwinResponse, error) {
	out := new(DescribeTwinResponse)
	err := c.cc.Invoke(ctx, TwinAPI_DescribeTwin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twinAPIClient) ListAllTwins(ctx context.Context, in *ListAllTwinsRequest, opts ...grpc.CallOption) (*ListAllTwinsResponse, error) {
	out := new(ListAllTwinsResponse)
	err := c.cc.Invoke(ctx, TwinAPI_ListAllTwins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TwinAPIServer is the server API for TwinAPI service.
// All implementations should embed UnimplementedTwinAPIServer
// for forward compatibility
type TwinAPIServer interface {
	// CreateTwin creates a twin.
	CreateTwin(context.Context, *CreateTwinRequest) (*CreateTwinResponse, error)
	// UpsertTwin creates or update a twin with its metadata + the twin feeds with their metadata.
	// The full state is applied (ie. if the operation succeeds the state of the twin/feeds will be the one
	// described in the payload)
	UpsertTwin(context.Context, *UpsertTwinRequest) (*UpsertTwinResponse, error)
	// DeleteTwin deletes a twin.
	DeleteTwin(context.Context, *DeleteTwinRequest) (*DeleteTwinResponse, error)
	// UpdateTwin updates a twin (partial update).
	UpdateTwin(context.Context, *UpdateTwinRequest) (*UpdateTwinResponse, error)
	// Describes a twin. (local and remote)
	DescribeTwin(context.Context, *DescribeTwinRequest) (*DescribeTwinResponse, error)
	// List all twins.
	ListAllTwins(context.Context, *ListAllTwinsRequest) (*ListAllTwinsResponse, error)
}

// UnimplementedTwinAPIServer should be embedded to have forward compatible implementations.
type UnimplementedTwinAPIServer struct {
}

func (UnimplementedTwinAPIServer) CreateTwin(context.Context, *CreateTwinRequest) (*CreateTwinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTwin not implemented")
}
func (UnimplementedTwinAPIServer) UpsertTwin(context.Context, *UpsertTwinRequest) (*UpsertTwinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertTwin not implemented")
}
func (UnimplementedTwinAPIServer) DeleteTwin(context.Context, *DeleteTwinRequest) (*DeleteTwinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTwin not implemented")
}
func (UnimplementedTwinAPIServer) UpdateTwin(context.Context, *UpdateTwinRequest) (*UpdateTwinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTwin not implemented")
}
func (UnimplementedTwinAPIServer) DescribeTwin(context.Context, *DescribeTwinRequest) (*DescribeTwinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTwin not implemented")
}
func (UnimplementedTwinAPIServer) ListAllTwins(context.Context, *ListAllTwinsRequest) (*ListAllTwinsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllTwins not implemented")
}

// UnsafeTwinAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TwinAPIServer will
// result in compilation errors.
type UnsafeTwinAPIServer interface {
	mustEmbedUnimplementedTwinAPIServer()
}

func RegisterTwinAPIServer(s grpc.ServiceRegistrar, srv TwinAPIServer) {
	s.RegisterService(&TwinAPI_ServiceDesc, srv)
}

func _TwinAPI_CreateTwin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTwinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwinAPIServer).CreateTwin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwinAPI_CreateTwin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwinAPIServer).CreateTwin(ctx, req.(*CreateTwinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwinAPI_UpsertTwin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertTwinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwinAPIServer).UpsertTwin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwinAPI_UpsertTwin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwinAPIServer).UpsertTwin(ctx, req.(*UpsertTwinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwinAPI_DeleteTwin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTwinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwinAPIServer).DeleteTwin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwinAPI_DeleteTwin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwinAPIServer).DeleteTwin(ctx, req.(*DeleteTwinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwinAPI_UpdateTwin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTwinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwinAPIServer).UpdateTwin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwinAPI_UpdateTwin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwinAPIServer).UpdateTwin(ctx, req.(*UpdateTwinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwinAPI_DescribeTwin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTwinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwinAPIServer).DescribeTwin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwinAPI_DescribeTwin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwinAPIServer).DescribeTwin(ctx, req.(*DescribeTwinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwinAPI_ListAllTwins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllTwinsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwinAPIServer).ListAllTwins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TwinAPI_ListAllTwins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwinAPIServer).ListAllTwins(ctx, req.(*ListAllTwinsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TwinAPI_ServiceDesc is the grpc.ServiceDesc for TwinAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TwinAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iotics.api.TwinAPI",
	HandlerType: (*TwinAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTwin",
			Handler:    _TwinAPI_CreateTwin_Handler,
		},
		{
			MethodName: "UpsertTwin",
			Handler:    _TwinAPI_UpsertTwin_Handler,
		},
		{
			MethodName: "DeleteTwin",
			Handler:    _TwinAPI_DeleteTwin_Handler,
		},
		{
			MethodName: "UpdateTwin",
			Handler:    _TwinAPI_UpdateTwin_Handler,
		},
		{
			MethodName: "DescribeTwin",
			Handler:    _TwinAPI_DescribeTwin_Handler,
		},
		{
			MethodName: "ListAllTwins",
			Handler:    _TwinAPI_ListAllTwins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iotics/api/twin.proto",
}
