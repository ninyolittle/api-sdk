// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.12
// source: amornotification/amornotification.proto

package amornotification

import (
	amor_var "github.com/Ninolito/api-sdk/other/amor_var"
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

type GetRealTimeNotificationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRealTimeNotificationsRequest) Reset() {
	*x = GetRealTimeNotificationsRequest{}
	mi := &file_amornotification_amornotification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRealTimeNotificationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRealTimeNotificationsRequest) ProtoMessage() {}

func (x *GetRealTimeNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_amornotification_amornotification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRealTimeNotificationsRequest.ProtoReflect.Descriptor instead.
func (*GetRealTimeNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_amornotification_amornotification_proto_rawDescGZIP(), []int{0}
}

type GetRealTimeNotificationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notification *amor_var.Notification `protobuf:"bytes,1,opt,name=Notification,proto3" json:"Notification,omitempty"`
}

func (x *GetRealTimeNotificationsResponse) Reset() {
	*x = GetRealTimeNotificationsResponse{}
	mi := &file_amornotification_amornotification_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRealTimeNotificationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRealTimeNotificationsResponse) ProtoMessage() {}

func (x *GetRealTimeNotificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_amornotification_amornotification_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRealTimeNotificationsResponse.ProtoReflect.Descriptor instead.
func (*GetRealTimeNotificationsResponse) Descriptor() ([]byte, []int) {
	return file_amornotification_amornotification_proto_rawDescGZIP(), []int{1}
}

func (x *GetRealTimeNotificationsResponse) GetNotification() *amor_var.Notification {
	if x != nil {
		return x.Notification
	}
	return nil
}

var File_amornotification_amornotification_proto protoreflect.FileDescriptor

var file_amornotification_amornotification_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x6d, 0x6f, 0x72, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x61, 0x6d, 0x6f, 0x72, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x21,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x61, 0x6d, 0x6f, 0x72, 0x5f, 0x76, 0x61, 0x72, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x21, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x5d, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x76, 0x61, 0x72, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x32, 0xc0, 0x01, 0x0a, 0x10, 0x41, 0x6d, 0x6f, 0x72, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xab, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x61,
	0x6d, 0x6f, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6d,
	0x6f, 0x72, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x69, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x61, 0x6d, 0x6f, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6d,
	0x6f, 0x72, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_amornotification_amornotification_proto_rawDescOnce sync.Once
	file_amornotification_amornotification_proto_rawDescData = file_amornotification_amornotification_proto_rawDesc
)

func file_amornotification_amornotification_proto_rawDescGZIP() []byte {
	file_amornotification_amornotification_proto_rawDescOnce.Do(func() {
		file_amornotification_amornotification_proto_rawDescData = protoimpl.X.CompressGZIP(file_amornotification_amornotification_proto_rawDescData)
	})
	return file_amornotification_amornotification_proto_rawDescData
}

var file_amornotification_amornotification_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_amornotification_amornotification_proto_goTypes = []any{
	(*GetRealTimeNotificationsRequest)(nil),  // 0: projectamor_api.amornotification.v1.GetRealTimeNotificationsRequest
	(*GetRealTimeNotificationsResponse)(nil), // 1: projectamor_api.amornotification.v1.GetRealTimeNotificationsResponse
	(*amor_var.Notification)(nil),            // 2: amorvar.Notification
}
var file_amornotification_amornotification_proto_depIdxs = []int32{
	2, // 0: projectamor_api.amornotification.v1.GetRealTimeNotificationsResponse.Notification:type_name -> amorvar.Notification
	0, // 1: projectamor_api.amornotification.v1.AmorNotification.GetRealTimeNotifications:input_type -> projectamor_api.amornotification.v1.GetRealTimeNotificationsRequest
	1, // 2: projectamor_api.amornotification.v1.AmorNotification.GetRealTimeNotifications:output_type -> projectamor_api.amornotification.v1.GetRealTimeNotificationsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_amornotification_amornotification_proto_init() }
func file_amornotification_amornotification_proto_init() {
	if File_amornotification_amornotification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_amornotification_amornotification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_amornotification_amornotification_proto_goTypes,
		DependencyIndexes: file_amornotification_amornotification_proto_depIdxs,
		MessageInfos:      file_amornotification_amornotification_proto_msgTypes,
	}.Build()
	File_amornotification_amornotification_proto = out.File
	file_amornotification_amornotification_proto_rawDesc = nil
	file_amornotification_amornotification_proto_goTypes = nil
	file_amornotification_amornotification_proto_depIdxs = nil
}
