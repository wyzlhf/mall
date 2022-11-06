// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: service/order/rpc/order.proto

package order

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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    uint64 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Pid    uint64 `protobuf:"varint,2,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Amount uint64 `protobuf:"varint,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Status uint64 `protobuf:"varint,4,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_rpc_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_rpc_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_service_order_rpc_order_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *CreateRequest) GetPid() uint64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *CreateRequest) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateRequest) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_rpc_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_rpc_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_service_order_rpc_order_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid    uint64 `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Pid    uint64 `protobuf:"varint,3,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Amount uint64 `protobuf:"varint,4,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Status uint64 `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_rpc_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_rpc_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_service_order_rpc_order_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRequest) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *UpdateRequest) GetPid() uint64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *UpdateRequest) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *UpdateRequest) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_rpc_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_rpc_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_service_order_rpc_order_proto_rawDescGZIP(), []int{3}
}

type RemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveRequest) Reset() {
	*x = RemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_rpc_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRequest) ProtoMessage() {}

func (x *RemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_rpc_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRequest.ProtoReflect.Descriptor instead.
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return file_service_order_rpc_order_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveResponse) Reset() {
	*x = RemoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_rpc_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveResponse) ProtoMessage() {}

func (x *RemoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_rpc_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveResponse.ProtoReflect.Descriptor instead.
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return file_service_order_rpc_order_proto_rawDescGZIP(), []int{5}
}

type DetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DetailRequest) Reset() {
	*x = DetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_rpc_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailRequest) ProtoMessage() {}

func (x *DetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_rpc_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailRequest.ProtoReflect.Descriptor instead.
func (*DetailRequest) Descriptor() ([]byte, []int) {
	return file_service_order_rpc_order_proto_rawDescGZIP(), []int{6}
}

func (x *DetailRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid    uint64 `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Pid    uint64 `protobuf:"varint,3,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Amount uint64 `protobuf:"varint,4,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Status uint64 `protobuf:"varint,5,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *DetailResponse) Reset() {
	*x = DetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_rpc_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailResponse) ProtoMessage() {}

func (x *DetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_rpc_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailResponse.ProtoReflect.Descriptor instead.
func (*DetailResponse) Descriptor() ([]byte, []int) {
	return file_service_order_rpc_order_proto_rawDescGZIP(), []int{7}
}

func (x *DetailResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DetailResponse) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *DetailResponse) GetPid() uint64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *DetailResponse) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *DetailResponse) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_rpc_order_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_rpc_order_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_service_order_rpc_order_proto_rawDescGZIP(), []int{8}
}

func (x *ListRequest) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*DetailResponse `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_rpc_order_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_rpc_order_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_service_order_rpc_order_proto_rawDescGZIP(), []int{9}
}

func (x *ListResponse) GetData() []*DetailResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type PaidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PaidRequest) Reset() {
	*x = PaidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_rpc_order_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaidRequest) ProtoMessage() {}

func (x *PaidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_rpc_order_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaidRequest.ProtoReflect.Descriptor instead.
func (*PaidRequest) Descriptor() ([]byte, []int) {
	return file_service_order_rpc_order_proto_rawDescGZIP(), []int{10}
}

func (x *PaidRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PaidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PaidResponse) Reset() {
	*x = PaidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_order_rpc_order_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaidResponse) ProtoMessage() {}

func (x *PaidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_order_rpc_order_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaidResponse.ProtoReflect.Descriptor instead.
func (*PaidResponse) Descriptor() ([]byte, []int) {
	return file_service_order_rpc_order_proto_rawDescGZIP(), []int{11}
}

var File_service_order_rpc_order_proto protoreflect.FileDescriptor

var file_service_order_rpc_order_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x63, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x50, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x73, 0x0a,
	0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x55, 0x69, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x50,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x74, 0x0a, 0x0e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x50, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x50, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1f,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x39, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1d, 0x0a, 0x0b, 0x50, 0x61,
	0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x61, 0x69,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc5, 0x02, 0x0a, 0x05, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x14, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x04, 0x50, 0x61, 0x69, 0x64, 0x12, 0x12, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x50, 0x61, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_order_rpc_order_proto_rawDescOnce sync.Once
	file_service_order_rpc_order_proto_rawDescData = file_service_order_rpc_order_proto_rawDesc
)

func file_service_order_rpc_order_proto_rawDescGZIP() []byte {
	file_service_order_rpc_order_proto_rawDescOnce.Do(func() {
		file_service_order_rpc_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_order_rpc_order_proto_rawDescData)
	})
	return file_service_order_rpc_order_proto_rawDescData
}

var file_service_order_rpc_order_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_service_order_rpc_order_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),  // 0: order.CreateRequest
	(*CreateResponse)(nil), // 1: order.CreateResponse
	(*UpdateRequest)(nil),  // 2: order.UpdateRequest
	(*UpdateResponse)(nil), // 3: order.UpdateResponse
	(*RemoveRequest)(nil),  // 4: order.RemoveRequest
	(*RemoveResponse)(nil), // 5: order.RemoveResponse
	(*DetailRequest)(nil),  // 6: order.DetailRequest
	(*DetailResponse)(nil), // 7: order.DetailResponse
	(*ListRequest)(nil),    // 8: order.ListRequest
	(*ListResponse)(nil),   // 9: order.ListResponse
	(*PaidRequest)(nil),    // 10: order.PaidRequest
	(*PaidResponse)(nil),   // 11: order.PaidResponse
}
var file_service_order_rpc_order_proto_depIdxs = []int32{
	7,  // 0: order.ListResponse.data:type_name -> order.DetailResponse
	0,  // 1: order.Order.Create:input_type -> order.CreateRequest
	2,  // 2: order.Order.Update:input_type -> order.UpdateRequest
	4,  // 3: order.Order.Remove:input_type -> order.RemoveRequest
	6,  // 4: order.Order.Detail:input_type -> order.DetailRequest
	8,  // 5: order.Order.List:input_type -> order.ListRequest
	10, // 6: order.Order.Paid:input_type -> order.PaidRequest
	1,  // 7: order.Order.Create:output_type -> order.CreateResponse
	3,  // 8: order.Order.Update:output_type -> order.UpdateResponse
	5,  // 9: order.Order.Remove:output_type -> order.RemoveResponse
	7,  // 10: order.Order.Detail:output_type -> order.DetailResponse
	9,  // 11: order.Order.List:output_type -> order.ListResponse
	11, // 12: order.Order.Paid:output_type -> order.PaidResponse
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_service_order_rpc_order_proto_init() }
func file_service_order_rpc_order_proto_init() {
	if File_service_order_rpc_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_order_rpc_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_service_order_rpc_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_service_order_rpc_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_service_order_rpc_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_service_order_rpc_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRequest); i {
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
		file_service_order_rpc_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveResponse); i {
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
		file_service_order_rpc_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailRequest); i {
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
		file_service_order_rpc_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailResponse); i {
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
		file_service_order_rpc_order_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_service_order_rpc_order_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_service_order_rpc_order_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaidRequest); i {
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
		file_service_order_rpc_order_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaidResponse); i {
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
			RawDescriptor: file_service_order_rpc_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_order_rpc_order_proto_goTypes,
		DependencyIndexes: file_service_order_rpc_order_proto_depIdxs,
		MessageInfos:      file_service_order_rpc_order_proto_msgTypes,
	}.Build()
	File_service_order_rpc_order_proto = out.File
	file_service_order_rpc_order_proto_rawDesc = nil
	file_service_order_rpc_order_proto_goTypes = nil
	file_service_order_rpc_order_proto_depIdxs = nil
}
