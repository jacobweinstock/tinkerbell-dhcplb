// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: api/v1/lbc.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type BackendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MacAddress string `protobuf:"bytes,1,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	Backend    string `protobuf:"bytes,2,opt,name=backend,proto3" json:"backend,omitempty"`
}

func (x *BackendRequest) Reset() {
	*x = BackendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_lbc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackendRequest) ProtoMessage() {}

func (x *BackendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_lbc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackendRequest.ProtoReflect.Descriptor instead.
func (*BackendRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_lbc_proto_rawDescGZIP(), []int{0}
}

func (x *BackendRequest) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *BackendRequest) GetBackend() string {
	if x != nil {
		return x.Backend
	}
	return ""
}

type BackendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BackendResponse) Reset() {
	*x = BackendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_lbc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackendResponse) ProtoMessage() {}

func (x *BackendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_lbc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackendResponse.ProtoReflect.Descriptor instead.
func (*BackendResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_lbc_proto_rawDescGZIP(), []int{1}
}

func (x *BackendResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_v1_lbc_proto protoreflect.FileDescriptor

var file_api_v1_lbc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x62, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x6c, 0x62, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x4b,
	0x0a, 0x0e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x42,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x49, 0x0a, 0x07, 0x42, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x12, 0x3e, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x6c, 0x62, 0x63,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6c, 0x62, 0x63, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x6c, 0x62, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_lbc_proto_rawDescOnce sync.Once
	file_api_v1_lbc_proto_rawDescData = file_api_v1_lbc_proto_rawDesc
)

func file_api_v1_lbc_proto_rawDescGZIP() []byte {
	file_api_v1_lbc_proto_rawDescOnce.Do(func() {
		file_api_v1_lbc_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_lbc_proto_rawDescData)
	})
	return file_api_v1_lbc_proto_rawDescData
}

var file_api_v1_lbc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v1_lbc_proto_goTypes = []interface{}{
	(*BackendRequest)(nil),  // 0: lbc.api.v1.BackendRequest
	(*BackendResponse)(nil), // 1: lbc.api.v1.BackendResponse
}
var file_api_v1_lbc_proto_depIdxs = []int32{
	0, // 0: lbc.api.v1.Backend.Set:input_type -> lbc.api.v1.BackendRequest
	1, // 1: lbc.api.v1.Backend.Set:output_type -> lbc.api.v1.BackendResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v1_lbc_proto_init() }
func file_api_v1_lbc_proto_init() {
	if File_api_v1_lbc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_lbc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackendRequest); i {
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
		file_api_v1_lbc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackendResponse); i {
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
			RawDescriptor: file_api_v1_lbc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_lbc_proto_goTypes,
		DependencyIndexes: file_api_v1_lbc_proto_depIdxs,
		MessageInfos:      file_api_v1_lbc_proto_msgTypes,
	}.Build()
	File_api_v1_lbc_proto = out.File
	file_api_v1_lbc_proto_rawDesc = nil
	file_api_v1_lbc_proto_goTypes = nil
	file_api_v1_lbc_proto_depIdxs = nil
}