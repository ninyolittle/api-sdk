// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: other/amor_var/home.proto

package amor_var

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age          int64  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Address      string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	EmailAddress string `protobuf:"bytes,5,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	PhoneNumber  string `protobuf:"bytes,6,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_amor_var_home_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_other_amor_var_home_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_other_amor_var_home_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *User) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *User) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *User) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type Accommodation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address   string  `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Type      string  `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Rating    int32   `protobuf:"varint,5,opt,name=rating,proto3" json:"rating,omitempty"`
	Owner     *User   `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	Utilities []int32 `protobuf:"varint,7,rep,packed,name=utilities,proto3" json:"utilities,omitempty"`
	Rooms     []*Room `protobuf:"bytes,8,rep,name=rooms,proto3" json:"rooms,omitempty"`
	Latitude  float64 `protobuf:"fixed64,9,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,10,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Accommodation) Reset() {
	*x = Accommodation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_amor_var_home_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accommodation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accommodation) ProtoMessage() {}

func (x *Accommodation) ProtoReflect() protoreflect.Message {
	mi := &file_other_amor_var_home_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accommodation.ProtoReflect.Descriptor instead.
func (*Accommodation) Descriptor() ([]byte, []int) {
	return file_other_amor_var_home_proto_rawDescGZIP(), []int{1}
}

func (x *Accommodation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Accommodation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Accommodation) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Accommodation) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Accommodation) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Accommodation) GetOwner() *User {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *Accommodation) GetUtilities() []int32 {
	if x != nil {
		return x.Utilities
	}
	return nil
}

func (x *Accommodation) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

func (x *Accommodation) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Accommodation) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type LatLong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *LatLong) Reset() {
	*x = LatLong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_amor_var_home_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatLong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatLong) ProtoMessage() {}

func (x *LatLong) ProtoReflect() protoreflect.Message {
	mi := &file_other_amor_var_home_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatLong.ProtoReflect.Descriptor instead.
func (*LatLong) Descriptor() ([]byte, []int) {
	return file_other_amor_var_home_proto_rawDescGZIP(), []int{2}
}

func (x *LatLong) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *LatLong) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Occupied       int64               `protobuf:"varint,2,opt,name=occupied,proto3" json:"occupied,omitempty"`
	Capacity       int64               `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Type           string              `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Currency       string              `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	Name           string              `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Description    string              `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Prices         *Room_Prices        `protobuf:"bytes,9,opt,name=prices,proto3" json:"prices,omitempty"`
	Rating         int32               `protobuf:"varint,10,opt,name=rating,proto3" json:"rating,omitempty"`
	Utilities      []int32             `protobuf:"varint,11,rep,packed,name=utilities,proto3" json:"utilities,omitempty"`
	AddonUtilities []int32             `protobuf:"varint,12,rep,packed,name=addonUtilities,proto3" json:"addonUtilities,omitempty"`
	Reservations   []*Room_Reservation `protobuf:"bytes,13,rep,name=reservations,proto3" json:"reservations,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_amor_var_home_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_other_amor_var_home_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_other_amor_var_home_proto_rawDescGZIP(), []int{3}
}

func (x *Room) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Room) GetOccupied() int64 {
	if x != nil {
		return x.Occupied
	}
	return 0
}

func (x *Room) GetCapacity() int64 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *Room) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Room) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Room) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Room) GetPrices() *Room_Prices {
	if x != nil {
		return x.Prices
	}
	return nil
}

func (x *Room) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Room) GetUtilities() []int32 {
	if x != nil {
		return x.Utilities
	}
	return nil
}

func (x *Room) GetAddonUtilities() []int32 {
	if x != nil {
		return x.AddonUtilities
	}
	return nil
}

func (x *Room) GetReservations() []*Room_Reservation {
	if x != nil {
		return x.Reservations
	}
	return nil
}

type Room_Reservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DateReserved string                       `protobuf:"bytes,2,opt,name=dateReserved,proto3" json:"dateReserved,omitempty"`
	RoomId       string                       `protobuf:"bytes,3,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Status       string                       `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	ReservedBy   *Room_Reservation_ReservedBy `protobuf:"bytes,5,opt,name=reservedBy,proto3" json:"reservedBy,omitempty"`
}

func (x *Room_Reservation) Reset() {
	*x = Room_Reservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_amor_var_home_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room_Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room_Reservation) ProtoMessage() {}

func (x *Room_Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_other_amor_var_home_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room_Reservation.ProtoReflect.Descriptor instead.
func (*Room_Reservation) Descriptor() ([]byte, []int) {
	return file_other_amor_var_home_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Room_Reservation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Room_Reservation) GetDateReserved() string {
	if x != nil {
		return x.DateReserved
	}
	return ""
}

func (x *Room_Reservation) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *Room_Reservation) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Room_Reservation) GetReservedBy() *Room_Reservation_ReservedBy {
	if x != nil {
		return x.ReservedBy
	}
	return nil
}

type Room_Prices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Daily   float64 `protobuf:"fixed64,1,opt,name=daily,proto3" json:"daily,omitempty"`
	Monthly float64 `protobuf:"fixed64,2,opt,name=monthly,proto3" json:"monthly,omitempty"`
	Annual  float64 `protobuf:"fixed64,3,opt,name=annual,proto3" json:"annual,omitempty"`
}

func (x *Room_Prices) Reset() {
	*x = Room_Prices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_amor_var_home_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room_Prices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room_Prices) ProtoMessage() {}

func (x *Room_Prices) ProtoReflect() protoreflect.Message {
	mi := &file_other_amor_var_home_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room_Prices.ProtoReflect.Descriptor instead.
func (*Room_Prices) Descriptor() ([]byte, []int) {
	return file_other_amor_var_home_proto_rawDescGZIP(), []int{3, 1}
}

func (x *Room_Prices) GetDaily() float64 {
	if x != nil {
		return x.Daily
	}
	return 0
}

func (x *Room_Prices) GetMonthly() float64 {
	if x != nil {
		return x.Monthly
	}
	return 0
}

func (x *Room_Prices) GetAnnual() float64 {
	if x != nil {
		return x.Annual
	}
	return 0
}

type Room_Reservation_ReservedBy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhotoUrl *string `protobuf:"bytes,3,opt,name=photoUrl,proto3,oneof" json:"photoUrl,omitempty"`
}

func (x *Room_Reservation_ReservedBy) Reset() {
	*x = Room_Reservation_ReservedBy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_amor_var_home_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room_Reservation_ReservedBy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room_Reservation_ReservedBy) ProtoMessage() {}

func (x *Room_Reservation_ReservedBy) ProtoReflect() protoreflect.Message {
	mi := &file_other_amor_var_home_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room_Reservation_ReservedBy.ProtoReflect.Descriptor instead.
func (*Room_Reservation_ReservedBy) Descriptor() ([]byte, []int) {
	return file_other_amor_var_home_proto_rawDescGZIP(), []int{3, 0, 0}
}

func (x *Room_Reservation_ReservedBy) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Room_Reservation_ReservedBy) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Room_Reservation_ReservedBy) GetPhotoUrl() string {
	if x != nil && x.PhotoUrl != nil {
		return *x.PhotoUrl
	}
	return ""
}

var File_other_amor_var_home_proto protoreflect.FileDescriptor

var file_other_amor_var_home_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x61, 0x6d, 0x6f, 0x72, 0x5f, 0x76, 0x61, 0x72,
	0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x6d, 0x6f,
	0x72, 0x76, 0x61, 0x72, 0x22, 0x9e, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x9b, 0x02, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x23, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x76, 0x61, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x75, 0x74, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x76, 0x61, 0x72, 0x2e, 0x52, 0x6f,
	0x6f, 0x6d, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x22, 0x43, 0x0a, 0x07, 0x4c, 0x61, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0xf3, 0x05, 0x0a, 0x04, 0x52, 0x6f, 0x6f,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x69, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x63, 0x63, 0x75, 0x70, 0x69, 0x65, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x76, 0x61, 0x72, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x2e, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x73, 0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x55, 0x74, 0x69, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0e, 0x61, 0x64, 0x64,
	0x6f, 0x6e, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x76, 0x61, 0x72, 0x2e, 0x52, 0x6f, 0x6f, 0x6d,
	0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x9f, 0x02, 0x0a, 0x0b, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x44,
	0x0a, 0x0a, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x42, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x6d, 0x6f, 0x72, 0x76, 0x61, 0x72, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x42, 0x79, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x64, 0x42, 0x79, 0x1a, 0x66, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x55, 0x72, 0x6c, 0x1a, 0x50, 0x0a, 0x06,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x6e, 0x75, 0x61, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6e, 0x6e, 0x75, 0x61, 0x6c, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x69, 0x6e,
	0x6f, 0x6c, 0x69, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x2f, 0x61, 0x6d, 0x6f, 0x72, 0x5f, 0x76, 0x61, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_other_amor_var_home_proto_rawDescOnce sync.Once
	file_other_amor_var_home_proto_rawDescData = file_other_amor_var_home_proto_rawDesc
)

func file_other_amor_var_home_proto_rawDescGZIP() []byte {
	file_other_amor_var_home_proto_rawDescOnce.Do(func() {
		file_other_amor_var_home_proto_rawDescData = protoimpl.X.CompressGZIP(file_other_amor_var_home_proto_rawDescData)
	})
	return file_other_amor_var_home_proto_rawDescData
}

var file_other_amor_var_home_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_other_amor_var_home_proto_goTypes = []interface{}{
	(*User)(nil),                        // 0: amorvar.User
	(*Accommodation)(nil),               // 1: amorvar.Accommodation
	(*LatLong)(nil),                     // 2: amorvar.LatLong
	(*Room)(nil),                        // 3: amorvar.Room
	(*Room_Reservation)(nil),            // 4: amorvar.Room.Reservation
	(*Room_Prices)(nil),                 // 5: amorvar.Room.Prices
	(*Room_Reservation_ReservedBy)(nil), // 6: amorvar.Room.Reservation.ReservedBy
}
var file_other_amor_var_home_proto_depIdxs = []int32{
	0, // 0: amorvar.Accommodation.owner:type_name -> amorvar.User
	3, // 1: amorvar.Accommodation.rooms:type_name -> amorvar.Room
	5, // 2: amorvar.Room.prices:type_name -> amorvar.Room.Prices
	4, // 3: amorvar.Room.reservations:type_name -> amorvar.Room.Reservation
	6, // 4: amorvar.Room.Reservation.reservedBy:type_name -> amorvar.Room.Reservation.ReservedBy
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_other_amor_var_home_proto_init() }
func file_other_amor_var_home_proto_init() {
	if File_other_amor_var_home_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_other_amor_var_home_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_other_amor_var_home_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Accommodation); i {
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
		file_other_amor_var_home_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatLong); i {
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
		file_other_amor_var_home_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
		file_other_amor_var_home_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room_Reservation); i {
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
		file_other_amor_var_home_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room_Prices); i {
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
		file_other_amor_var_home_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room_Reservation_ReservedBy); i {
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
	file_other_amor_var_home_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_other_amor_var_home_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_other_amor_var_home_proto_goTypes,
		DependencyIndexes: file_other_amor_var_home_proto_depIdxs,
		MessageInfos:      file_other_amor_var_home_proto_msgTypes,
	}.Build()
	File_other_amor_var_home_proto = out.File
	file_other_amor_var_home_proto_rawDesc = nil
	file_other_amor_var_home_proto_goTypes = nil
	file_other_amor_var_home_proto_depIdxs = nil
}
