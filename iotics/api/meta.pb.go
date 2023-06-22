// Copyright (c) 2019-2022 Iotic Labs Ltd. All rights reserved.

// Iotics Web protocol definitions (meta)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: iotics/api/meta.proto

package ioticsapi

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SparqlResultType defines the result content types for SPARQL requests. Note that applicable content types depend on
// the type of query.
type SparqlResultType int32

const (
	// Applicable to SELECT/ASK (SPARQL 1.1 Query Results JSON Format)
	SparqlResultType_SPARQL_JSON SparqlResultType = 0
	// Applicable to SELECT/ASK (SPARQL 1.1 Query Results XML Format)
	SparqlResultType_SPARQL_XML SparqlResultType = 1
	// Applicable to SELECT/ASK (SPARQL 1.1. Query Results CSV Format)
	SparqlResultType_SPARQL_CSV SparqlResultType = 2
	// Applicable to CONSTRUCT/DESCRIBE (Terse RDF Triple Language)
	SparqlResultType_RDF_TURTLE SparqlResultType = 3
	// Applicable to CONSTRUCT/DESCRIBE (RDF 1.1 XML)
	SparqlResultType_RDF_XML SparqlResultType = 4
	// Applicable to CONSTRUCT/DESCRIBE (RDF 1.1 N-Triples)
	SparqlResultType_RDF_NTRIPLES SparqlResultType = 5
)

// Enum value maps for SparqlResultType.
var (
	SparqlResultType_name = map[int32]string{
		0: "SPARQL_JSON",
		1: "SPARQL_XML",
		2: "SPARQL_CSV",
		3: "RDF_TURTLE",
		4: "RDF_XML",
		5: "RDF_NTRIPLES",
	}
	SparqlResultType_value = map[string]int32{
		"SPARQL_JSON":  0,
		"SPARQL_XML":   1,
		"SPARQL_CSV":   2,
		"RDF_TURTLE":   3,
		"RDF_XML":      4,
		"RDF_NTRIPLES": 5,
	}
)

func (x SparqlResultType) Enum() *SparqlResultType {
	p := new(SparqlResultType)
	*p = x
	return p
}

func (x SparqlResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SparqlResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_iotics_api_meta_proto_enumTypes[0].Descriptor()
}

func (SparqlResultType) Type() protoreflect.EnumType {
	return &file_iotics_api_meta_proto_enumTypes[0]
}

func (x SparqlResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SparqlResultType.Descriptor instead.
func (SparqlResultType) EnumDescriptor() ([]byte, []int) {
	return file_iotics_api_meta_proto_rawDescGZIP(), []int{0}
}

// ExplorerRequest - Deprecated. Use SparqlQueryRequest instead.
type ExplorerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Explorer request headers
	Headers *Headers `protobuf:"bytes,1,opt,name=headers,proto3" json:"headers,omitempty"`
	// Explorer request scope
	Scope Scope `protobuf:"varint,2,opt,name=scope,proto3,enum=iotics.api.Scope" json:"scope,omitempty"`
	// Explorer request payload
	Payload *ExplorerRequest_Payload `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ExplorerRequest) Reset() {
	*x = ExplorerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iotics_api_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExplorerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExplorerRequest) ProtoMessage() {}

func (x *ExplorerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iotics_api_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExplorerRequest.ProtoReflect.Descriptor instead.
func (*ExplorerRequest) Descriptor() ([]byte, []int) {
	return file_iotics_api_meta_proto_rawDescGZIP(), []int{0}
}

func (x *ExplorerRequest) GetHeaders() *Headers {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *ExplorerRequest) GetScope() Scope {
	if x != nil {
		return x.Scope
	}
	return Scope_GLOBAL
}

func (x *ExplorerRequest) GetPayload() *ExplorerRequest_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

// SparqlQueryRequest describes a SPARQL query.
type SparqlQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SPARQL query request headers
	Headers *Headers `protobuf:"bytes,1,opt,name=headers,proto3" json:"headers,omitempty"`
	// SPARQL query request scope
	Scope Scope `protobuf:"varint,2,opt,name=scope,proto3,enum=iotics.api.Scope" json:"scope,omitempty"`
	// SPARQL query request payload
	Payload *SparqlQueryRequest_Payload `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *SparqlQueryRequest) Reset() {
	*x = SparqlQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iotics_api_meta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SparqlQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SparqlQueryRequest) ProtoMessage() {}

func (x *SparqlQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iotics_api_meta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SparqlQueryRequest.ProtoReflect.Descriptor instead.
func (*SparqlQueryRequest) Descriptor() ([]byte, []int) {
	return file_iotics_api_meta_proto_rawDescGZIP(), []int{1}
}

func (x *SparqlQueryRequest) GetHeaders() *Headers {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *SparqlQueryRequest) GetScope() Scope {
	if x != nil {
		return x.Scope
	}
	return Scope_GLOBAL
}

func (x *SparqlQueryRequest) GetPayload() *SparqlQueryRequest_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

// SparqlQueryResponse is a part of a result for a SPARQL query request. Multiple chunks form a complete result. Related
// chunks can be identified by a combination of:
// - The hostId
// - Client reference (in headers, set by caller)
// - Chunk sequence number
type SparqlQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Headers for the query result. clientRef within can be used to identify which query the result applies to.
	Headers *Headers `protobuf:"bytes,1,opt,name=headers,proto3" json:"headers,omitempty"`
	// SPARQL query result chunk payload.
	Payload *SparqlQueryResponse_Payload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *SparqlQueryResponse) Reset() {
	*x = SparqlQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iotics_api_meta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SparqlQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SparqlQueryResponse) ProtoMessage() {}

func (x *SparqlQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iotics_api_meta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SparqlQueryResponse.ProtoReflect.Descriptor instead.
func (*SparqlQueryResponse) Descriptor() ([]byte, []int) {
	return file_iotics_api_meta_proto_rawDescGZIP(), []int{2}
}

func (x *SparqlQueryResponse) GetHeaders() *Headers {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *SparqlQueryResponse) GetPayload() *SparqlQueryResponse_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

// Performs a SPARQL update against custom metadata only.
type SparqlUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SPARQL update request headers
	Headers *Headers `protobuf:"bytes,1,opt,name=headers,proto3" json:"headers,omitempty"`
	// SPARQL update request payload.
	Payload *SparqlUpdateRequest_Payload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *SparqlUpdateRequest) Reset() {
	*x = SparqlUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iotics_api_meta_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SparqlUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SparqlUpdateRequest) ProtoMessage() {}

func (x *SparqlUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iotics_api_meta_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SparqlUpdateRequest.ProtoReflect.Descriptor instead.
func (*SparqlUpdateRequest) Descriptor() ([]byte, []int) {
	return file_iotics_api_meta_proto_rawDescGZIP(), []int{3}
}

func (x *SparqlUpdateRequest) GetHeaders() *Headers {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *SparqlUpdateRequest) GetPayload() *SparqlUpdateRequest_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

// Response of the SPARQL update request.
type SparqlUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SPARQL update response headers
	Headers *Headers `protobuf:"bytes,1,opt,name=headers,proto3" json:"headers,omitempty"`
}

func (x *SparqlUpdateResponse) Reset() {
	*x = SparqlUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iotics_api_meta_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SparqlUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SparqlUpdateResponse) ProtoMessage() {}

func (x *SparqlUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iotics_api_meta_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SparqlUpdateResponse.ProtoReflect.Descriptor instead.
func (*SparqlUpdateResponse) Descriptor() ([]byte, []int) {
	return file_iotics_api_meta_proto_rawDescGZIP(), []int{4}
}

func (x *SparqlUpdateResponse) GetHeaders() *Headers {
	if x != nil {
		return x.Headers
	}
	return nil
}

// Explorer request payload.
type ExplorerRequest_Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The desired result content type. Note that choosing an invalid result type for the type of query will result in
	// an error status reported in the response. (See SparqlResultType for valid content-query type combinations.)
	ResultContentType SparqlResultType `protobuf:"varint,1,opt,name=resultContentType,proto3,enum=iotics.api.SparqlResultType" json:"resultContentType,omitempty"`
	// keyword defines the search term associated to the explorer request.
	Keyword string `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *ExplorerRequest_Payload) Reset() {
	*x = ExplorerRequest_Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iotics_api_meta_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExplorerRequest_Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExplorerRequest_Payload) ProtoMessage() {}

func (x *ExplorerRequest_Payload) ProtoReflect() protoreflect.Message {
	mi := &file_iotics_api_meta_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExplorerRequest_Payload.ProtoReflect.Descriptor instead.
func (*ExplorerRequest_Payload) Descriptor() ([]byte, []int) {
	return file_iotics_api_meta_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ExplorerRequest_Payload) GetResultContentType() SparqlResultType {
	if x != nil {
		return x.ResultContentType
	}
	return SparqlResultType_SPARQL_JSON
}

func (x *ExplorerRequest_Payload) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

// SPARQL query request payload.
type SparqlQueryRequest_Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The desired result content type. Note that choosing an invalid result type for the type of query will result in
	// an error status reported in the response. (See SparqlResultType for valid content-query type combinations.)
	ResultContentType SparqlResultType `protobuf:"varint,1,opt,name=resultContentType,proto3,enum=iotics.api.SparqlResultType" json:"resultContentType,omitempty"`
	// A UTF8-encoded SPARQL 1.1 query
	Query []byte `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *SparqlQueryRequest_Payload) Reset() {
	*x = SparqlQueryRequest_Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iotics_api_meta_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SparqlQueryRequest_Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SparqlQueryRequest_Payload) ProtoMessage() {}

func (x *SparqlQueryRequest_Payload) ProtoReflect() protoreflect.Message {
	mi := &file_iotics_api_meta_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SparqlQueryRequest_Payload.ProtoReflect.Descriptor instead.
func (*SparqlQueryRequest_Payload) Descriptor() ([]byte, []int) {
	return file_iotics_api_meta_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SparqlQueryRequest_Payload) GetResultContentType() SparqlResultType {
	if x != nil {
		return x.ResultContentType
	}
	return SparqlResultType_SPARQL_JSON
}

func (x *SparqlQueryRequest_Payload) GetQuery() []byte {
	if x != nil {
		return x.Query
	}
	return nil
}

// Payload of the query result chunk
type SparqlQueryResponse_Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Result host identifier. Indicates from which host this result chunk came from. For a local result, this field
	// will be unset.
	HostId string `protobuf:"bytes,1,opt,name=hostId,proto3" json:"hostId,omitempty"`
	// Position of a chunk in result from a given host (and request). The first chunk has a sequence number of 0.
	SeqNum uint64 `protobuf:"varint,2,opt,name=seqNum,proto3" json:"seqNum,omitempty"`
	// Indicates whether this is the last chunk from a given host, for a specific request. Results for different
	// requests can be identified by setting a unique clientRef in the request headers.
	Last bool `protobuf:"varint,3,opt,name=last,proto3" json:"last,omitempty"`
	// Result error status (only applicable to local results). If set, this will
	// indicate a problem with running the query (e.g. invalid syntax or content type) as opposed to a more general
	// issue (in which case the standard gRPC error mechanism will be used and the stream terminated).
	Status *status.Status `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	// Content type of the result.
	ContentType SparqlResultType `protobuf:"varint,5,opt,name=contentType,proto3,enum=iotics.api.SparqlResultType" json:"contentType,omitempty"`
	// Query result chunk, encoded according to actualType.
	// Note that:
	// - The maximum size of each chunk is host-specific.
	ResultChunk []byte `protobuf:"bytes,6,opt,name=resultChunk,proto3" json:"resultChunk,omitempty"`
}

func (x *SparqlQueryResponse_Payload) Reset() {
	*x = SparqlQueryResponse_Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iotics_api_meta_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SparqlQueryResponse_Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SparqlQueryResponse_Payload) ProtoMessage() {}

func (x *SparqlQueryResponse_Payload) ProtoReflect() protoreflect.Message {
	mi := &file_iotics_api_meta_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SparqlQueryResponse_Payload.ProtoReflect.Descriptor instead.
func (*SparqlQueryResponse_Payload) Descriptor() ([]byte, []int) {
	return file_iotics_api_meta_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SparqlQueryResponse_Payload) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *SparqlQueryResponse_Payload) GetSeqNum() uint64 {
	if x != nil {
		return x.SeqNum
	}
	return 0
}

func (x *SparqlQueryResponse_Payload) GetLast() bool {
	if x != nil {
		return x.Last
	}
	return false
}

func (x *SparqlQueryResponse_Payload) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *SparqlQueryResponse_Payload) GetContentType() SparqlResultType {
	if x != nil {
		return x.ContentType
	}
	return SparqlResultType_SPARQL_JSON
}

func (x *SparqlQueryResponse_Payload) GetResultChunk() []byte {
	if x != nil {
		return x.ResultChunk
	}
	return nil
}

// SPARQL update request payload.
type SparqlUpdateRequest_Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A UTF8-encoded SPARQL 1.1 update
	Update []byte `protobuf:"bytes,1,opt,name=update,proto3" json:"update,omitempty"`
}

func (x *SparqlUpdateRequest_Payload) Reset() {
	*x = SparqlUpdateRequest_Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iotics_api_meta_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SparqlUpdateRequest_Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SparqlUpdateRequest_Payload) ProtoMessage() {}

func (x *SparqlUpdateRequest_Payload) ProtoReflect() protoreflect.Message {
	mi := &file_iotics_api_meta_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SparqlUpdateRequest_Payload.ProtoReflect.Descriptor instead.
func (*SparqlUpdateRequest_Payload) Descriptor() ([]byte, []int) {
	return file_iotics_api_meta_proto_rawDescGZIP(), []int{3, 0}
}

func (x *SparqlUpdateRequest_Payload) GetUpdate() []byte {
	if x != nil {
		return x.Update
	}
	return nil
}

var File_iotics_api_meta_proto protoreflect.FileDescriptor

var file_iotics_api_meta_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x69, 0x6f,
	0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x02, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x12, 0x3d, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x1a, 0x6f, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x4a, 0x0a, 0x11, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x71, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x11, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x9b, 0x02, 0x0a, 0x12, 0x53, 0x70, 0x61, 0x72, 0x71, 0x6c, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x12, 0x40, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x70, 0x61, 0x72, 0x71, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x1a, 0x6b, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x4a, 0x0a,
	0x11, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x71, 0x6c, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x11, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22,
	0xe5, 0x02, 0x0a, 0x13, 0x53, 0x70, 0x61, 0x72, 0x71, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x07, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x41, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x71, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0xdb, 0x01, 0x0a, 0x07, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73,
	0x65, 0x71, 0x4e, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x69, 0x6f, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x71, 0x6c, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0xaa, 0x01, 0x0a, 0x13, 0x53, 0x70, 0x61, 0x72,
	0x71, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2d, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x41,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x70, 0x61,
	0x72, 0x71, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x1a, 0x21, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x45, 0x0a, 0x14, 0x53, 0x70, 0x61, 0x72, 0x71, 0x6c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x2a, 0x72, 0x0a, 0x10, 0x53,
	0x70, 0x61, 0x72, 0x71, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x50, 0x41, 0x52, 0x51, 0x4c, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x50, 0x41, 0x52, 0x51, 0x4c, 0x5f, 0x58, 0x4d, 0x4c, 0x10, 0x01,
	0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x50, 0x41, 0x52, 0x51, 0x4c, 0x5f, 0x43, 0x53, 0x56, 0x10, 0x02,
	0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x44, 0x46, 0x5f, 0x54, 0x55, 0x52, 0x54, 0x4c, 0x45, 0x10, 0x03,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x44, 0x46, 0x5f, 0x58, 0x4d, 0x4c, 0x10, 0x04, 0x12, 0x10, 0x0a,
	0x0c, 0x52, 0x44, 0x46, 0x5f, 0x4e, 0x54, 0x52, 0x49, 0x50, 0x4c, 0x45, 0x53, 0x10, 0x05, 0x32,
	0x85, 0x02, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x61, 0x41, 0x50, 0x49, 0x12, 0x52, 0x0a, 0x0b, 0x53,
	0x70, 0x61, 0x72, 0x71, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1e, 0x2e, 0x69, 0x6f, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x71, 0x6c, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6f, 0x74,
	0x69, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x70, 0x61, 0x72, 0x71, 0x6c, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12,
	0x53, 0x0a, 0x0c, 0x53, 0x70, 0x61, 0x72, 0x71, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1f, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x70, 0x61,
	0x72, 0x71, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x70,
	0x61, 0x72, 0x71, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x70, 0x61, 0x72, 0x71, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x7d, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x49, 0x6f, 0x74, 0x69, 0x63, 0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x69, 0x6f,
	0x74, 0x69, 0x63, 0x2d, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x71, 0x61, 0x70,
	0x69, 0x2f, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x69, 0x6f, 0x74,
	0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0xa2, 0x02, 0x03, 0x49, 0x41, 0x58, 0xaa, 0x02, 0x0a, 0x49,
	0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x41, 0x70, 0x69, 0xca, 0x02, 0x0a, 0x49, 0x6f, 0x74, 0x69,
	0x63, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iotics_api_meta_proto_rawDescOnce sync.Once
	file_iotics_api_meta_proto_rawDescData = file_iotics_api_meta_proto_rawDesc
)

func file_iotics_api_meta_proto_rawDescGZIP() []byte {
	file_iotics_api_meta_proto_rawDescOnce.Do(func() {
		file_iotics_api_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_iotics_api_meta_proto_rawDescData)
	})
	return file_iotics_api_meta_proto_rawDescData
}

var file_iotics_api_meta_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_iotics_api_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_iotics_api_meta_proto_goTypes = []interface{}{
	(SparqlResultType)(0),               // 0: iotics.api.SparqlResultType
	(*ExplorerRequest)(nil),             // 1: iotics.api.ExplorerRequest
	(*SparqlQueryRequest)(nil),          // 2: iotics.api.SparqlQueryRequest
	(*SparqlQueryResponse)(nil),         // 3: iotics.api.SparqlQueryResponse
	(*SparqlUpdateRequest)(nil),         // 4: iotics.api.SparqlUpdateRequest
	(*SparqlUpdateResponse)(nil),        // 5: iotics.api.SparqlUpdateResponse
	(*ExplorerRequest_Payload)(nil),     // 6: iotics.api.ExplorerRequest.Payload
	(*SparqlQueryRequest_Payload)(nil),  // 7: iotics.api.SparqlQueryRequest.Payload
	(*SparqlQueryResponse_Payload)(nil), // 8: iotics.api.SparqlQueryResponse.Payload
	(*SparqlUpdateRequest_Payload)(nil), // 9: iotics.api.SparqlUpdateRequest.Payload
	(*Headers)(nil),                     // 10: iotics.api.Headers
	(Scope)(0),                          // 11: iotics.api.Scope
	(*status.Status)(nil),               // 12: google.rpc.Status
}
var file_iotics_api_meta_proto_depIdxs = []int32{
	10, // 0: iotics.api.ExplorerRequest.headers:type_name -> iotics.api.Headers
	11, // 1: iotics.api.ExplorerRequest.scope:type_name -> iotics.api.Scope
	6,  // 2: iotics.api.ExplorerRequest.payload:type_name -> iotics.api.ExplorerRequest.Payload
	10, // 3: iotics.api.SparqlQueryRequest.headers:type_name -> iotics.api.Headers
	11, // 4: iotics.api.SparqlQueryRequest.scope:type_name -> iotics.api.Scope
	7,  // 5: iotics.api.SparqlQueryRequest.payload:type_name -> iotics.api.SparqlQueryRequest.Payload
	10, // 6: iotics.api.SparqlQueryResponse.headers:type_name -> iotics.api.Headers
	8,  // 7: iotics.api.SparqlQueryResponse.payload:type_name -> iotics.api.SparqlQueryResponse.Payload
	10, // 8: iotics.api.SparqlUpdateRequest.headers:type_name -> iotics.api.Headers
	9,  // 9: iotics.api.SparqlUpdateRequest.payload:type_name -> iotics.api.SparqlUpdateRequest.Payload
	10, // 10: iotics.api.SparqlUpdateResponse.headers:type_name -> iotics.api.Headers
	0,  // 11: iotics.api.ExplorerRequest.Payload.resultContentType:type_name -> iotics.api.SparqlResultType
	0,  // 12: iotics.api.SparqlQueryRequest.Payload.resultContentType:type_name -> iotics.api.SparqlResultType
	12, // 13: iotics.api.SparqlQueryResponse.Payload.status:type_name -> google.rpc.Status
	0,  // 14: iotics.api.SparqlQueryResponse.Payload.contentType:type_name -> iotics.api.SparqlResultType
	2,  // 15: iotics.api.MetaAPI.SparqlQuery:input_type -> iotics.api.SparqlQueryRequest
	4,  // 16: iotics.api.MetaAPI.SparqlUpdate:input_type -> iotics.api.SparqlUpdateRequest
	1,  // 17: iotics.api.MetaAPI.ExplorerQuery:input_type -> iotics.api.ExplorerRequest
	3,  // 18: iotics.api.MetaAPI.SparqlQuery:output_type -> iotics.api.SparqlQueryResponse
	5,  // 19: iotics.api.MetaAPI.SparqlUpdate:output_type -> iotics.api.SparqlUpdateResponse
	3,  // 20: iotics.api.MetaAPI.ExplorerQuery:output_type -> iotics.api.SparqlQueryResponse
	18, // [18:21] is the sub-list for method output_type
	15, // [15:18] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_iotics_api_meta_proto_init() }
func file_iotics_api_meta_proto_init() {
	if File_iotics_api_meta_proto != nil {
		return
	}
	file_iotics_api_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_iotics_api_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExplorerRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iotics_api_meta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SparqlQueryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iotics_api_meta_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SparqlQueryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iotics_api_meta_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SparqlUpdateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iotics_api_meta_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SparqlUpdateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iotics_api_meta_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExplorerRequest_Payload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iotics_api_meta_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SparqlQueryRequest_Payload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iotics_api_meta_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SparqlQueryResponse_Payload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_iotics_api_meta_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SparqlUpdateRequest_Payload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_iotics_api_meta_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iotics_api_meta_proto_goTypes,
		DependencyIndexes: file_iotics_api_meta_proto_depIdxs,
		EnumInfos:         file_iotics_api_meta_proto_enumTypes,
		MessageInfos:      file_iotics_api_meta_proto_msgTypes,
	}.Build()
	File_iotics_api_meta_proto = out.File
	file_iotics_api_meta_proto_rawDesc = nil
	file_iotics_api_meta_proto_goTypes = nil
	file_iotics_api_meta_proto_depIdxs = nil
}
