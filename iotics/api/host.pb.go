// Copyright (c) 2019-2022 Iotic Labs Ltd. All rights reserved.

// Iotics Web protocol definitions (hosts)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: iotics/api/host.proto

package ioticsapi

import (
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

// GetHostIDRequest: gets the local host twin's ID
type GetHostIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers *Headers `protobuf:"bytes,1,opt,name=headers,proto3" json:"headers,omitempty"`
}

func (x *GetHostIDRequest) Reset() {
	*x = GetHostIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iotics_api_host_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHostIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHostIDRequest) ProtoMessage() {}

func (x *GetHostIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iotics_api_host_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHostIDRequest.ProtoReflect.Descriptor instead.
func (*GetHostIDRequest) Descriptor() ([]byte, []int) {
	return file_iotics_api_host_proto_rawDescGZIP(), []int{0}
}

func (x *GetHostIDRequest) GetHeaders() *Headers {
	if x != nil {
		return x.Headers
	}
	return nil
}

// GetHostIDResponse: response containing the local host twin's ID
type GetHostIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers *Headers                   `protobuf:"bytes,1,opt,name=headers,proto3" json:"headers,omitempty"`
	Payload *GetHostIDResponse_Payload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *GetHostIDResponse) Reset() {
	*x = GetHostIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iotics_api_host_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHostIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHostIDResponse) ProtoMessage() {}

func (x *GetHostIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iotics_api_host_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHostIDResponse.ProtoReflect.Descriptor instead.
func (*GetHostIDResponse) Descriptor() ([]byte, []int) {
	return file_iotics_api_host_proto_rawDescGZIP(), []int{1}
}

func (x *GetHostIDResponse) GetHeaders() *Headers {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *GetHostIDResponse) GetPayload() *GetHostIDResponse_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

type GetHostIDResponse_Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostId string `protobuf:"bytes,1,opt,name=hostId,proto3" json:"hostId,omitempty"`
}

func (x *GetHostIDResponse_Payload) Reset() {
	*x = GetHostIDResponse_Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iotics_api_host_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHostIDResponse_Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHostIDResponse_Payload) ProtoMessage() {}

func (x *GetHostIDResponse_Payload) ProtoReflect() protoreflect.Message {
	mi := &file_iotics_api_host_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHostIDResponse_Payload.ProtoReflect.Descriptor instead.
func (*GetHostIDResponse_Payload) Descriptor() ([]byte, []int) {
	return file_iotics_api_host_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetHostIDResponse_Payload) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

var File_iotics_api_host_proto protoreflect.FileDescriptor

var file_iotics_api_host_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x6f, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x1a, 0x17, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2d, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x22,
	0xa6, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x3f, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x21, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x32, 0x55, 0x0a, 0x07, 0x48, 0x6f, 0x73, 0x74,
	0x41, 0x50, 0x49, 0x12, 0x4a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x44,
	0x12, 0x1c, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x48,
	0x6f, 0x73, 0x74, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x7d, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x42, 0x09, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x6f, 0x74, 0x69, 0x63,
	0x2d, 0x4c, 0x61, 0x62, 0x73, 0x2f, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x2d, 0x67, 0x6f, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x71, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73,
	0x2f, 0x61, 0x70, 0x69, 0x3b, 0x69, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x61, 0x70, 0x69, 0xa2, 0x02,
	0x03, 0x49, 0x41, 0x58, 0xaa, 0x02, 0x0a, 0x49, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x41, 0x70,
	0x69, 0xca, 0x02, 0x0a, 0x49, 0x6f, 0x74, 0x69, 0x63, 0x73, 0x5c, 0x41, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iotics_api_host_proto_rawDescOnce sync.Once
	file_iotics_api_host_proto_rawDescData = file_iotics_api_host_proto_rawDesc
)

func file_iotics_api_host_proto_rawDescGZIP() []byte {
	file_iotics_api_host_proto_rawDescOnce.Do(func() {
		file_iotics_api_host_proto_rawDescData = protoimpl.X.CompressGZIP(file_iotics_api_host_proto_rawDescData)
	})
	return file_iotics_api_host_proto_rawDescData
}

var file_iotics_api_host_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_iotics_api_host_proto_goTypes = []interface{}{
	(*GetHostIDRequest)(nil),          // 0: iotics.api.GetHostIDRequest
	(*GetHostIDResponse)(nil),         // 1: iotics.api.GetHostIDResponse
	(*GetHostIDResponse_Payload)(nil), // 2: iotics.api.GetHostIDResponse.Payload
	(*Headers)(nil),                   // 3: iotics.api.Headers
}
var file_iotics_api_host_proto_depIdxs = []int32{
	3, // 0: iotics.api.GetHostIDRequest.headers:type_name -> iotics.api.Headers
	3, // 1: iotics.api.GetHostIDResponse.headers:type_name -> iotics.api.Headers
	2, // 2: iotics.api.GetHostIDResponse.payload:type_name -> iotics.api.GetHostIDResponse.Payload
	0, // 3: iotics.api.HostAPI.GetHostID:input_type -> iotics.api.GetHostIDRequest
	1, // 4: iotics.api.HostAPI.GetHostID:output_type -> iotics.api.GetHostIDResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_iotics_api_host_proto_init() }
func file_iotics_api_host_proto_init() {
	if File_iotics_api_host_proto != nil {
		return
	}
	file_iotics_api_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_iotics_api_host_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHostIDRequest); i {
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
		file_iotics_api_host_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHostIDResponse); i {
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
		file_iotics_api_host_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHostIDResponse_Payload); i {
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
			RawDescriptor: file_iotics_api_host_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iotics_api_host_proto_goTypes,
		DependencyIndexes: file_iotics_api_host_proto_depIdxs,
		MessageInfos:      file_iotics_api_host_proto_msgTypes,
	}.Build()
	File_iotics_api_host_proto = out.File
	file_iotics_api_host_proto_rawDesc = nil
	file_iotics_api_host_proto_goTypes = nil
	file_iotics_api_host_proto_depIdxs = nil
}
