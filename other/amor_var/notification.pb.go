// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: other/amor_var/notification.proto

package amor_var

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Notification struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp string                 `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message   string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// Valid Values: READ, UNREAD, SEEN_BUT_UNREAD
	Status string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Type   string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	// Types that are valid to be assigned to Details:
	//
	//	*Notification_NewRoomReservationOption_
	//	*Notification_CancelRoomReservation_
	Details       isNotification_Details `protobuf_oneof:"details"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Notification) Reset() {
	*x = Notification{}
	mi := &file_other_amor_var_notification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_other_amor_var_notification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_other_amor_var_notification_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notification) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *Notification) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Notification) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Notification) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Notification) GetDetails() isNotification_Details {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *Notification) GetNewRoomReservationOption() *Notification_NewRoomReservationOption {
	if x != nil {
		if x, ok := x.Details.(*Notification_NewRoomReservationOption_); ok {
			return x.NewRoomReservationOption
		}
	}
	return nil
}

func (x *Notification) GetCancelRoomReservation() *Notification_CancelRoomReservation {
	if x != nil {
		if x, ok := x.Details.(*Notification_CancelRoomReservation_); ok {
			return x.CancelRoomReservation
		}
	}
	return nil
}

type isNotification_Details interface {
	isNotification_Details()
}

type Notification_NewRoomReservationOption_ struct {
	NewRoomReservationOption *Notification_NewRoomReservationOption `protobuf:"bytes,6,opt,name=newRoomReservationOption,proto3,oneof"`
}

type Notification_CancelRoomReservation_ struct {
	CancelRoomReservation *Notification_CancelRoomReservation `protobuf:"bytes,7,opt,name=cancelRoomReservation,proto3,oneof"`
}

func (*Notification_NewRoomReservationOption_) isNotification_Details() {}

func (*Notification_CancelRoomReservation_) isNotification_Details() {}

type Notification_NewRoomReservationOption struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RoomId        string                 `protobuf:"bytes,1,opt,name=roomId,proto3" json:"roomId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Notification_NewRoomReservationOption) Reset() {
	*x = Notification_NewRoomReservationOption{}
	mi := &file_other_amor_var_notification_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification_NewRoomReservationOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification_NewRoomReservationOption) ProtoMessage() {}

func (x *Notification_NewRoomReservationOption) ProtoReflect() protoreflect.Message {
	mi := &file_other_amor_var_notification_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification_NewRoomReservationOption.ProtoReflect.Descriptor instead.
func (*Notification_NewRoomReservationOption) Descriptor() ([]byte, []int) {
	return file_other_amor_var_notification_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Notification_NewRoomReservationOption) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type Notification_CancelRoomReservation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReservationId string                 `protobuf:"bytes,2,opt,name=reservationId,proto3" json:"reservationId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Notification_CancelRoomReservation) Reset() {
	*x = Notification_CancelRoomReservation{}
	mi := &file_other_amor_var_notification_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification_CancelRoomReservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification_CancelRoomReservation) ProtoMessage() {}

func (x *Notification_CancelRoomReservation) ProtoReflect() protoreflect.Message {
	mi := &file_other_amor_var_notification_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification_CancelRoomReservation.ProtoReflect.Descriptor instead.
func (*Notification_CancelRoomReservation) Descriptor() ([]byte, []int) {
	return file_other_amor_var_notification_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Notification_CancelRoomReservation) GetReservationId() string {
	if x != nil {
		return x.ReservationId
	}
	return ""
}

var File_other_amor_var_notification_proto protoreflect.FileDescriptor

var file_other_amor_var_notification_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x61, 0x6d, 0x6f, 0x72, 0x5f, 0x76, 0x61, 0x72,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x6d, 0x6f, 0x72, 0x76, 0x61, 0x72, 0x22, 0xd3, 0x03, 0x0a,
	0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x6c, 0x0a, 0x18, 0x6e, 0x65, 0x77, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x76, 0x61, 0x72, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x18, 0x6e, 0x65, 0x77, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x63, 0x0a, 0x15, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x76, 0x61, 0x72, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x6f, 0x6f, 0x6d,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x15, 0x63,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x32, 0x0a, 0x18, 0x4e, 0x65, 0x77, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x1a, 0x3d, 0x0a, 0x15, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4e, 0x69, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x61, 0x6d, 0x6f, 0x72, 0x5f, 0x76, 0x61, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_other_amor_var_notification_proto_rawDescOnce sync.Once
	file_other_amor_var_notification_proto_rawDescData []byte
)

func file_other_amor_var_notification_proto_rawDescGZIP() []byte {
	file_other_amor_var_notification_proto_rawDescOnce.Do(func() {
		file_other_amor_var_notification_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_other_amor_var_notification_proto_rawDesc), len(file_other_amor_var_notification_proto_rawDesc)))
	})
	return file_other_amor_var_notification_proto_rawDescData
}

var file_other_amor_var_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_other_amor_var_notification_proto_goTypes = []any{
	(*Notification)(nil),                          // 0: amorvar.Notification
	(*Notification_NewRoomReservationOption)(nil), // 1: amorvar.Notification.NewRoomReservationOption
	(*Notification_CancelRoomReservation)(nil),    // 2: amorvar.Notification.CancelRoomReservation
}
var file_other_amor_var_notification_proto_depIdxs = []int32{
	1, // 0: amorvar.Notification.newRoomReservationOption:type_name -> amorvar.Notification.NewRoomReservationOption
	2, // 1: amorvar.Notification.cancelRoomReservation:type_name -> amorvar.Notification.CancelRoomReservation
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_other_amor_var_notification_proto_init() }
func file_other_amor_var_notification_proto_init() {
	if File_other_amor_var_notification_proto != nil {
		return
	}
	file_other_amor_var_notification_proto_msgTypes[0].OneofWrappers = []any{
		(*Notification_NewRoomReservationOption_)(nil),
		(*Notification_CancelRoomReservation_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_other_amor_var_notification_proto_rawDesc), len(file_other_amor_var_notification_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_other_amor_var_notification_proto_goTypes,
		DependencyIndexes: file_other_amor_var_notification_proto_depIdxs,
		MessageInfos:      file_other_amor_var_notification_proto_msgTypes,
	}.Build()
	File_other_amor_var_notification_proto = out.File
	file_other_amor_var_notification_proto_goTypes = nil
	file_other_amor_var_notification_proto_depIdxs = nil
}
