// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: iotics/api/search.proto

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

// SearchAPIClient is the client API for SearchAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchAPIClient interface {
	// Send a search request. Results are expected asynchronously. (local and remote)
	DispatchSearchRequest(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*DispatchSearchResponse, error)
	// Run a synchronous search based on a user timeout. (local and remote)
	SynchronousSearch(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (SearchAPI_SynchronousSearchClient, error)
	// Receive all search responses associated to a set of Search request for a given client application ID.
	ReceiveAllSearchResponses(ctx context.Context, in *SubscriptionHeaders, opts ...grpc.CallOption) (SearchAPI_ReceiveAllSearchResponsesClient, error)
}

type searchAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchAPIClient(cc grpc.ClientConnInterface) SearchAPIClient {
	return &searchAPIClient{cc}
}

func (c *searchAPIClient) DispatchSearchRequest(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*DispatchSearchResponse, error) {
	out := new(DispatchSearchResponse)
	err := c.cc.Invoke(ctx, "/iotics.api.SearchAPI/DispatchSearchRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchAPIClient) SynchronousSearch(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (SearchAPI_SynchronousSearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &SearchAPI_ServiceDesc.Streams[0], "/iotics.api.SearchAPI/SynchronousSearch", opts...)
	if err != nil {
		return nil, err
	}
	x := &searchAPISynchronousSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SearchAPI_SynchronousSearchClient interface {
	Recv() (*SearchResponse, error)
	grpc.ClientStream
}

type searchAPISynchronousSearchClient struct {
	grpc.ClientStream
}

func (x *searchAPISynchronousSearchClient) Recv() (*SearchResponse, error) {
	m := new(SearchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *searchAPIClient) ReceiveAllSearchResponses(ctx context.Context, in *SubscriptionHeaders, opts ...grpc.CallOption) (SearchAPI_ReceiveAllSearchResponsesClient, error) {
	stream, err := c.cc.NewStream(ctx, &SearchAPI_ServiceDesc.Streams[1], "/iotics.api.SearchAPI/ReceiveAllSearchResponses", opts...)
	if err != nil {
		return nil, err
	}
	x := &searchAPIReceiveAllSearchResponsesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SearchAPI_ReceiveAllSearchResponsesClient interface {
	Recv() (*SearchResponse, error)
	grpc.ClientStream
}

type searchAPIReceiveAllSearchResponsesClient struct {
	grpc.ClientStream
}

func (x *searchAPIReceiveAllSearchResponsesClient) Recv() (*SearchResponse, error) {
	m := new(SearchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SearchAPIServer is the server API for SearchAPI service.
// All implementations should embed UnimplementedSearchAPIServer
// for forward compatibility
type SearchAPIServer interface {
	// Send a search request. Results are expected asynchronously. (local and remote)
	DispatchSearchRequest(context.Context, *SearchRequest) (*DispatchSearchResponse, error)
	// Run a synchronous search based on a user timeout. (local and remote)
	SynchronousSearch(*SearchRequest, SearchAPI_SynchronousSearchServer) error
	// Receive all search responses associated to a set of Search request for a given client application ID.
	ReceiveAllSearchResponses(*SubscriptionHeaders, SearchAPI_ReceiveAllSearchResponsesServer) error
}

// UnimplementedSearchAPIServer should be embedded to have forward compatible implementations.
type UnimplementedSearchAPIServer struct {
}

func (UnimplementedSearchAPIServer) DispatchSearchRequest(context.Context, *SearchRequest) (*DispatchSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DispatchSearchRequest not implemented")
}
func (UnimplementedSearchAPIServer) SynchronousSearch(*SearchRequest, SearchAPI_SynchronousSearchServer) error {
	return status.Errorf(codes.Unimplemented, "method SynchronousSearch not implemented")
}
func (UnimplementedSearchAPIServer) ReceiveAllSearchResponses(*SubscriptionHeaders, SearchAPI_ReceiveAllSearchResponsesServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveAllSearchResponses not implemented")
}

// UnsafeSearchAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchAPIServer will
// result in compilation errors.
type UnsafeSearchAPIServer interface {
	mustEmbedUnimplementedSearchAPIServer()
}

func RegisterSearchAPIServer(s grpc.ServiceRegistrar, srv SearchAPIServer) {
	s.RegisterService(&SearchAPI_ServiceDesc, srv)
}

func _SearchAPI_DispatchSearchRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchAPIServer).DispatchSearchRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iotics.api.SearchAPI/DispatchSearchRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchAPIServer).DispatchSearchRequest(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchAPI_SynchronousSearch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SearchAPIServer).SynchronousSearch(m, &searchAPISynchronousSearchServer{stream})
}

type SearchAPI_SynchronousSearchServer interface {
	Send(*SearchResponse) error
	grpc.ServerStream
}

type searchAPISynchronousSearchServer struct {
	grpc.ServerStream
}

func (x *searchAPISynchronousSearchServer) Send(m *SearchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SearchAPI_ReceiveAllSearchResponses_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscriptionHeaders)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SearchAPIServer).ReceiveAllSearchResponses(m, &searchAPIReceiveAllSearchResponsesServer{stream})
}

type SearchAPI_ReceiveAllSearchResponsesServer interface {
	Send(*SearchResponse) error
	grpc.ServerStream
}

type searchAPIReceiveAllSearchResponsesServer struct {
	grpc.ServerStream
}

func (x *searchAPIReceiveAllSearchResponsesServer) Send(m *SearchResponse) error {
	return x.ServerStream.SendMsg(m)
}

// SearchAPI_ServiceDesc is the grpc.ServiceDesc for SearchAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iotics.api.SearchAPI",
	HandlerType: (*SearchAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DispatchSearchRequest",
			Handler:    _SearchAPI_DispatchSearchRequest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SynchronousSearch",
			Handler:       _SearchAPI_SynchronousSearch_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ReceiveAllSearchResponses",
			Handler:       _SearchAPI_ReceiveAllSearchResponses_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "iotics/api/search.proto",
}
