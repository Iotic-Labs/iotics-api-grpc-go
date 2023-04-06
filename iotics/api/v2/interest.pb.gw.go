// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: iotics/api/v2/interest.proto

/*
Package ioticsapi is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package ioticsapi

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

var (
	filter_InterestAPI_FetchLastStored_0 = &utilities.DoubleArray{Encoding: map[string]int{"args": 0, "interest": 1, "followerTwinId": 2, "hostId": 3, "id": 4, "followedFeedId": 5, "twinId": 6}, Base: []int{1, 20, 1, 3, 19, 21, 1, 22, 0, 8, 7, 5, 0, 9, 12, 9, 0, 13, 15, 13, 0, 17, 17, 17, 0, 0, 0, 0, 0, 0}, Check: []int{0, 1, 2, 3, 1, 1, 4, 1, 7, 2, 10, 11, 12, 2, 14, 15, 16, 2, 18, 19, 20, 2, 22, 23, 24, 5, 5, 6, 6, 8}}
)

func request_InterestAPI_FetchLastStored_0(ctx context.Context, marshaler runtime.Marshaler, client InterestAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FetchLastStoredRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["args.interest.followerTwinId.hostId"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.followerTwinId.hostId")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.followerTwinId.hostId", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.followerTwinId.hostId", err)
	}

	val, ok = pathParams["args.interest.followerTwinId.id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.followerTwinId.id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.followerTwinId.id", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.followerTwinId.id", err)
	}

	val, ok = pathParams["args.interest.followedFeedId.hostId"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.followedFeedId.hostId")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.followedFeedId.hostId", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.followedFeedId.hostId", err)
	}

	val, ok = pathParams["args.interest.followedFeedId.twinId"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.followedFeedId.twinId")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.followedFeedId.twinId", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.followedFeedId.twinId", err)
	}

	val, ok = pathParams["args.interest.followedFeedId.id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.followedFeedId.id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.followedFeedId.id", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.followedFeedId.id", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_InterestAPI_FetchLastStored_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.FetchLastStored(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_InterestAPI_FetchLastStored_0(ctx context.Context, marshaler runtime.Marshaler, server InterestAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq FetchLastStoredRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["args.interest.followerTwinId.hostId"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.followerTwinId.hostId")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.followerTwinId.hostId", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.followerTwinId.hostId", err)
	}

	val, ok = pathParams["args.interest.followerTwinId.id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.followerTwinId.id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.followerTwinId.id", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.followerTwinId.id", err)
	}

	val, ok = pathParams["args.interest.followedFeedId.hostId"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.followedFeedId.hostId")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.followedFeedId.hostId", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.followedFeedId.hostId", err)
	}

	val, ok = pathParams["args.interest.followedFeedId.twinId"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.followedFeedId.twinId")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.followedFeedId.twinId", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.followedFeedId.twinId", err)
	}

	val, ok = pathParams["args.interest.followedFeedId.id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.followedFeedId.id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.followedFeedId.id", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.followedFeedId.id", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_InterestAPI_FetchLastStored_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.FetchLastStored(ctx, &protoReq)
	return msg, metadata, err

}

func request_InterestAPI_SendInputMessage_0(ctx context.Context, marshaler runtime.Marshaler, client InterestAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SendInputMessageRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["args.interest.senderTwinId.hostId"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.senderTwinId.hostId")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.senderTwinId.hostId", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.senderTwinId.hostId", err)
	}

	val, ok = pathParams["args.interest.senderTwinId.id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.senderTwinId.id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.senderTwinId.id", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.senderTwinId.id", err)
	}

	val, ok = pathParams["args.interest.destInputId.hostId"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.destInputId.hostId")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.destInputId.hostId", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.destInputId.hostId", err)
	}

	val, ok = pathParams["args.interest.destInputId.twinId"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.destInputId.twinId")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.destInputId.twinId", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.destInputId.twinId", err)
	}

	val, ok = pathParams["args.interest.destInputId.id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.destInputId.id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.destInputId.id", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.destInputId.id", err)
	}

	msg, err := client.SendInputMessage(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_InterestAPI_SendInputMessage_0(ctx context.Context, marshaler runtime.Marshaler, server InterestAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SendInputMessageRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["args.interest.senderTwinId.hostId"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.senderTwinId.hostId")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.senderTwinId.hostId", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.senderTwinId.hostId", err)
	}

	val, ok = pathParams["args.interest.senderTwinId.id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.senderTwinId.id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.senderTwinId.id", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.senderTwinId.id", err)
	}

	val, ok = pathParams["args.interest.destInputId.hostId"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.destInputId.hostId")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.destInputId.hostId", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.destInputId.hostId", err)
	}

	val, ok = pathParams["args.interest.destInputId.twinId"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.destInputId.twinId")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.destInputId.twinId", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.destInputId.twinId", err)
	}

	val, ok = pathParams["args.interest.destInputId.id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "args.interest.destInputId.id")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "args.interest.destInputId.id", val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "args.interest.destInputId.id", err)
	}

	msg, err := server.SendInputMessage(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterInterestAPIHandlerServer registers the http handlers for service InterestAPI to "mux".
// UnaryRPC     :call InterestAPIServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterInterestAPIHandlerFromEndpoint instead.
func RegisterInterestAPIHandlerServer(ctx context.Context, mux *runtime.ServeMux, server InterestAPIServer) error {

	mux.Handle("GET", pattern_InterestAPI_FetchLastStored_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/iotics.api.v2.InterestAPI/FetchLastStored", runtime.WithHTTPPathPattern("/api/v2/hosts/{args.interest.followerTwinId.hostId}/twins/{args.interest.followerTwinId.id}/interests/hosts/{args.interest.followedFeedId.hostId}/twins/{args.interest.followedFeedId.twinId}/feeds/{args.interest.followedFeedId.id}/samples/last"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_InterestAPI_FetchLastStored_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InterestAPI_FetchLastStored_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InterestAPI_SendInputMessage_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/iotics.api.v2.InterestAPI/SendInputMessage", runtime.WithHTTPPathPattern("/api/v2/hosts/{args.interest.senderTwinId.hostId}/twins/{args.interest.senderTwinId.id}/interests/hosts/{args.interest.destInputId.hostId}/twins/{args.interest.destInputId.twinId}/inputs/{args.interest.destInputId.id}/messages"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_InterestAPI_SendInputMessage_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InterestAPI_SendInputMessage_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterInterestAPIHandlerFromEndpoint is same as RegisterInterestAPIHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterInterestAPIHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterInterestAPIHandler(ctx, mux, conn)
}

// RegisterInterestAPIHandler registers the http handlers for service InterestAPI to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterInterestAPIHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterInterestAPIHandlerClient(ctx, mux, NewInterestAPIClient(conn))
}

// RegisterInterestAPIHandlerClient registers the http handlers for service InterestAPI
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "InterestAPIClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "InterestAPIClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "InterestAPIClient" to call the correct interceptors.
func RegisterInterestAPIHandlerClient(ctx context.Context, mux *runtime.ServeMux, client InterestAPIClient) error {

	mux.Handle("GET", pattern_InterestAPI_FetchLastStored_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/iotics.api.v2.InterestAPI/FetchLastStored", runtime.WithHTTPPathPattern("/api/v2/hosts/{args.interest.followerTwinId.hostId}/twins/{args.interest.followerTwinId.id}/interests/hosts/{args.interest.followedFeedId.hostId}/twins/{args.interest.followedFeedId.twinId}/feeds/{args.interest.followedFeedId.id}/samples/last"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_InterestAPI_FetchLastStored_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InterestAPI_FetchLastStored_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InterestAPI_SendInputMessage_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/iotics.api.v2.InterestAPI/SendInputMessage", runtime.WithHTTPPathPattern("/api/v2/hosts/{args.interest.senderTwinId.hostId}/twins/{args.interest.senderTwinId.id}/interests/hosts/{args.interest.destInputId.hostId}/twins/{args.interest.destInputId.twinId}/inputs/{args.interest.destInputId.id}/messages"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_InterestAPI_SendInputMessage_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InterestAPI_SendInputMessage_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_InterestAPI_FetchLastStored_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4, 1, 0, 4, 1, 5, 5, 2, 6, 2, 2, 1, 0, 4, 1, 5, 7, 2, 4, 1, 0, 4, 1, 5, 8, 2, 9, 1, 0, 4, 1, 5, 10, 2, 11, 2, 12}, []string{"api", "v2", "hosts", "args.interest.followerTwinId.hostId", "twins", "args.interest.followerTwinId.id", "interests", "args.interest.followedFeedId.hostId", "args.interest.followedFeedId.twinId", "feeds", "args.interest.followedFeedId.id", "samples", "last"}, ""))

	pattern_InterestAPI_SendInputMessage_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4, 1, 0, 4, 1, 5, 5, 2, 6, 2, 2, 1, 0, 4, 1, 5, 7, 2, 4, 1, 0, 4, 1, 5, 8, 2, 9, 1, 0, 4, 1, 5, 10, 2, 11}, []string{"api", "v2", "hosts", "args.interest.senderTwinId.hostId", "twins", "args.interest.senderTwinId.id", "interests", "args.interest.destInputId.hostId", "args.interest.destInputId.twinId", "inputs", "args.interest.destInputId.id", "messages"}, ""))
)

var (
	forward_InterestAPI_FetchLastStored_0 = runtime.ForwardResponseMessage

	forward_InterestAPI_SendInputMessage_0 = runtime.ForwardResponseMessage
)
