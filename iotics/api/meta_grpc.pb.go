// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: iotics/api/meta.proto

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

// MetaAPIClient is the client API for MetaAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetaAPIClient interface {
	// SparqlQuery performs a SPARQL 1.1 query against the Federated Knowledge Graph of the Iotics network to which this
	// host belongs. The result is returned as a sequence of chunks. Note that:
	// - Result chunks MIGHT arrive out of order and it is the client's responsibility to re-assemble them.
	// - This RPC is currently in alpha!
	SparqlQuery(ctx context.Context, in *SparqlQueryRequest, opts ...grpc.CallOption) (MetaAPI_SparqlQueryClient, error)
	// SparqlUpdate performs a SPARQL 1.1 update. When performing an update, the update query must contain a reference to
	// one of the following graph IRIs:
	//  1. http://data.iotics.com/graph#custom-public (aka custom public graph) - All metadata written to this graph will be
	//     visible during SPARQL queries both with local & global scope (and thus, the Iotics network).
	SparqlUpdate(ctx context.Context, in *SparqlUpdateRequest, opts ...grpc.CallOption) (*SparqlUpdateResponse, error)
	// ExplorerQuery - Deprecated - use SparqlQuery instead.
	ExplorerQuery(ctx context.Context, in *ExplorerRequest, opts ...grpc.CallOption) (MetaAPI_ExplorerQueryClient, error)
}

type metaAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewMetaAPIClient(cc grpc.ClientConnInterface) MetaAPIClient {
	return &metaAPIClient{cc}
}

func (c *metaAPIClient) SparqlQuery(ctx context.Context, in *SparqlQueryRequest, opts ...grpc.CallOption) (MetaAPI_SparqlQueryClient, error) {
	stream, err := c.cc.NewStream(ctx, &MetaAPI_ServiceDesc.Streams[0], "/iotics.api.MetaAPI/SparqlQuery", opts...)
	if err != nil {
		return nil, err
	}
	x := &metaAPISparqlQueryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MetaAPI_SparqlQueryClient interface {
	Recv() (*SparqlQueryResponse, error)
	grpc.ClientStream
}

type metaAPISparqlQueryClient struct {
	grpc.ClientStream
}

func (x *metaAPISparqlQueryClient) Recv() (*SparqlQueryResponse, error) {
	m := new(SparqlQueryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *metaAPIClient) SparqlUpdate(ctx context.Context, in *SparqlUpdateRequest, opts ...grpc.CallOption) (*SparqlUpdateResponse, error) {
	out := new(SparqlUpdateResponse)
	err := c.cc.Invoke(ctx, "/iotics.api.MetaAPI/SparqlUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metaAPIClient) ExplorerQuery(ctx context.Context, in *ExplorerRequest, opts ...grpc.CallOption) (MetaAPI_ExplorerQueryClient, error) {
	stream, err := c.cc.NewStream(ctx, &MetaAPI_ServiceDesc.Streams[1], "/iotics.api.MetaAPI/ExplorerQuery", opts...)
	if err != nil {
		return nil, err
	}
	x := &metaAPIExplorerQueryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MetaAPI_ExplorerQueryClient interface {
	Recv() (*ExplorerQueryResponse, error)
	grpc.ClientStream
}

type metaAPIExplorerQueryClient struct {
	grpc.ClientStream
}

func (x *metaAPIExplorerQueryClient) Recv() (*ExplorerQueryResponse, error) {
	m := new(ExplorerQueryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetaAPIServer is the server API for MetaAPI service.
// All implementations should embed UnimplementedMetaAPIServer
// for forward compatibility
type MetaAPIServer interface {
	// SparqlQuery performs a SPARQL 1.1 query against the Federated Knowledge Graph of the Iotics network to which this
	// host belongs. The result is returned as a sequence of chunks. Note that:
	// - Result chunks MIGHT arrive out of order and it is the client's responsibility to re-assemble them.
	// - This RPC is currently in alpha!
	SparqlQuery(*SparqlQueryRequest, MetaAPI_SparqlQueryServer) error
	// SparqlUpdate performs a SPARQL 1.1 update. When performing an update, the update query must contain a reference to
	// one of the following graph IRIs:
	//  1. http://data.iotics.com/graph#custom-public (aka custom public graph) - All metadata written to this graph will be
	//     visible during SPARQL queries both with local & global scope (and thus, the Iotics network).
	SparqlUpdate(context.Context, *SparqlUpdateRequest) (*SparqlUpdateResponse, error)
	// ExplorerQuery - Deprecated - use SparqlQuery instead.
	ExplorerQuery(*ExplorerRequest, MetaAPI_ExplorerQueryServer) error
}

// UnimplementedMetaAPIServer should be embedded to have forward compatible implementations.
type UnimplementedMetaAPIServer struct {
}

func (UnimplementedMetaAPIServer) SparqlQuery(*SparqlQueryRequest, MetaAPI_SparqlQueryServer) error {
	return status.Errorf(codes.Unimplemented, "method SparqlQuery not implemented")
}
func (UnimplementedMetaAPIServer) SparqlUpdate(context.Context, *SparqlUpdateRequest) (*SparqlUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SparqlUpdate not implemented")
}
func (UnimplementedMetaAPIServer) ExplorerQuery(*ExplorerRequest, MetaAPI_ExplorerQueryServer) error {
	return status.Errorf(codes.Unimplemented, "method ExplorerQuery not implemented")
}

// UnsafeMetaAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetaAPIServer will
// result in compilation errors.
type UnsafeMetaAPIServer interface {
	mustEmbedUnimplementedMetaAPIServer()
}

func RegisterMetaAPIServer(s grpc.ServiceRegistrar, srv MetaAPIServer) {
	s.RegisterService(&MetaAPI_ServiceDesc, srv)
}

func _MetaAPI_SparqlQuery_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SparqlQueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetaAPIServer).SparqlQuery(m, &metaAPISparqlQueryServer{stream})
}

type MetaAPI_SparqlQueryServer interface {
	Send(*SparqlQueryResponse) error
	grpc.ServerStream
}

type metaAPISparqlQueryServer struct {
	grpc.ServerStream
}

func (x *metaAPISparqlQueryServer) Send(m *SparqlQueryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MetaAPI_SparqlUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SparqlUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetaAPIServer).SparqlUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iotics.api.MetaAPI/SparqlUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetaAPIServer).SparqlUpdate(ctx, req.(*SparqlUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetaAPI_ExplorerQuery_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExplorerRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetaAPIServer).ExplorerQuery(m, &metaAPIExplorerQueryServer{stream})
}

type MetaAPI_ExplorerQueryServer interface {
	Send(*ExplorerQueryResponse) error
	grpc.ServerStream
}

type metaAPIExplorerQueryServer struct {
	grpc.ServerStream
}

func (x *metaAPIExplorerQueryServer) Send(m *ExplorerQueryResponse) error {
	return x.ServerStream.SendMsg(m)
}

// MetaAPI_ServiceDesc is the grpc.ServiceDesc for MetaAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetaAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iotics.api.MetaAPI",
	HandlerType: (*MetaAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SparqlUpdate",
			Handler:    _MetaAPI_SparqlUpdate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SparqlQuery",
			Handler:       _MetaAPI_SparqlQuery_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ExplorerQuery",
			Handler:       _MetaAPI_ExplorerQuery_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "iotics/api/meta.proto",
}
