// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: socket.proto

package pb

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

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interfaces []*Interface `protobuf:"bytes,1,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetInterfaces() []*Interface {
	if x != nil {
		return x.Interfaces
	}
	return nil
}

type UnWaitParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnWaitParams) Reset() {
	*x = UnWaitParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnWaitParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnWaitParams) ProtoMessage() {}

func (x *UnWaitParams) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnWaitParams.ProtoReflect.Descriptor instead.
func (*UnWaitParams) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{1}
}

type StopParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopParams) Reset() {
	*x = StopParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopParams) ProtoMessage() {}

func (x *StopParams) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopParams.ProtoReflect.Descriptor instead.
func (*StopParams) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{2}
}

type StreamEventsParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamEventsParams) Reset() {
	*x = StreamEventsParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamEventsParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamEventsParams) ProtoMessage() {}

func (x *StreamEventsParams) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamEventsParams.ProtoReflect.Descriptor instead.
func (*StreamEventsParams) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{3}
}

type SyncParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SyncParams) Reset() {
	*x = SyncParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncParams) ProtoMessage() {}

func (x *SyncParams) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncParams.ProtoReflect.Descriptor instead.
func (*SyncParams) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{4}
}

type RemoveInterfaceParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intf string `protobuf:"bytes,1,opt,name=intf,proto3" json:"intf,omitempty"`
}

func (x *RemoveInterfaceParams) Reset() {
	*x = RemoveInterfaceParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveInterfaceParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveInterfaceParams) ProtoMessage() {}

func (x *RemoveInterfaceParams) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveInterfaceParams.ProtoReflect.Descriptor instead.
func (*RemoveInterfaceParams) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveInterfaceParams) GetIntf() string {
	if x != nil {
		return x.Intf
	}
	return ""
}

type RestartPeerParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intf string `protobuf:"bytes,1,opt,name=intf,proto3" json:"intf,omitempty"`
	Peer []byte `protobuf:"bytes,2,opt,name=peer,proto3" json:"peer,omitempty"`
}

func (x *RestartPeerParams) Reset() {
	*x = RestartPeerParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestartPeerParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestartPeerParams) ProtoMessage() {}

func (x *RestartPeerParams) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestartPeerParams.ProtoReflect.Descriptor instead.
func (*RestartPeerParams) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{6}
}

func (x *RestartPeerParams) GetIntf() string {
	if x != nil {
		return x.Intf
	}
	return ""
}

func (x *RestartPeerParams) GetPeer() []byte {
	if x != nil {
		return x.Peer
	}
	return nil
}

type InterfaceConfigParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Interface *InterfaceConfig `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
}

func (x *InterfaceConfigParams) Reset() {
	*x = InterfaceConfigParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_socket_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterfaceConfigParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfaceConfigParams) ProtoMessage() {}

func (x *InterfaceConfigParams) ProtoReflect() protoreflect.Message {
	mi := &file_socket_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfaceConfigParams.ProtoReflect.Descriptor instead.
func (*InterfaceConfigParams) Descriptor() ([]byte, []int) {
	return file_socket_proto_rawDescGZIP(), []int{7}
}

func (x *InterfaceConfigParams) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InterfaceConfigParams) GetInterface() *InterfaceConfig {
	if x != nil {
		return x.Interface
	}
	return nil
}

var File_socket_proto protoreflect.FileDescriptor

var file_socket_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x77, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x77,
	0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x0a, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x55, 0x6e, 0x57,
	0x61, 0x69, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x74, 0x6f,
	0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x0c, 0x0a,
	0x0a, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x2b, 0x0a, 0x15, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x74, 0x66, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x74, 0x66, 0x22, 0x3b, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x50, 0x65, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x6e, 0x74, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x74,
	0x66, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x70, 0x65, 0x65, 0x72, 0x22, 0x60, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x32, 0xa8, 0x04, 0x0a, 0x06, 0x53, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x27, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0a, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x0c, 0x2e, 0x77, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0c, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x77, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2b, 0x0a, 0x06, 0x55, 0x6e, 0x57, 0x61, 0x69, 0x74,
	0x12, 0x12, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x6e, 0x57, 0x61, 0x69, 0x74, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x10, 0x2e, 0x77, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e,
	0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x04,
	0x53, 0x79, 0x6e, 0x63, 0x12, 0x10, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x50, 0x65, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x50, 0x65, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e,
	0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0f,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12,
	0x1b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e, 0x77,
	0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x13, 0x53,
	0x79, 0x6e, 0x63, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x1b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x12, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x72, 0x69, 0x61, 0x73, 0x63, 0x2e, 0x65, 0x75, 0x2f, 0x77,
	0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_socket_proto_rawDescOnce sync.Once
	file_socket_proto_rawDescData = file_socket_proto_rawDesc
)

func file_socket_proto_rawDescGZIP() []byte {
	file_socket_proto_rawDescOnce.Do(func() {
		file_socket_proto_rawDescData = protoimpl.X.CompressGZIP(file_socket_proto_rawDescData)
	})
	return file_socket_proto_rawDescData
}

var file_socket_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_socket_proto_goTypes = []interface{}{
	(*Status)(nil),                // 0: wice.Status
	(*UnWaitParams)(nil),          // 1: wice.UnWaitParams
	(*StopParams)(nil),            // 2: wice.StopParams
	(*StreamEventsParams)(nil),    // 3: wice.StreamEventsParams
	(*SyncParams)(nil),            // 4: wice.SyncParams
	(*RemoveInterfaceParams)(nil), // 5: wice.RemoveInterfaceParams
	(*RestartPeerParams)(nil),     // 6: wice.RestartPeerParams
	(*InterfaceConfigParams)(nil), // 7: wice.InterfaceConfigParams
	(*Interface)(nil),             // 8: wice.Interface
	(*InterfaceConfig)(nil),       // 9: wice.InterfaceConfig
	(*Void)(nil),                  // 10: wice.Void
	(*Event)(nil),                 // 11: wice.Event
	(*Error)(nil),                 // 12: wice.Error
}
var file_socket_proto_depIdxs = []int32{
	8,  // 0: wice.Status.interfaces:type_name -> wice.Interface
	9,  // 1: wice.InterfaceConfigParams.interface:type_name -> wice.InterfaceConfig
	10, // 2: wice.Socket.GetStatus:input_type -> wice.Void
	3,  // 3: wice.Socket.StreamEvents:input_type -> wice.StreamEventsParams
	1,  // 4: wice.Socket.UnWait:input_type -> wice.UnWaitParams
	2,  // 5: wice.Socket.Stop:input_type -> wice.StopParams
	4,  // 6: wice.Socket.Sync:input_type -> wice.SyncParams
	6,  // 7: wice.Socket.RestartPeer:input_type -> wice.RestartPeerParams
	5,  // 8: wice.Socket.RemoveInterface:input_type -> wice.RemoveInterfaceParams
	7,  // 9: wice.Socket.SyncInterfaceConfig:input_type -> wice.InterfaceConfigParams
	7,  // 10: wice.Socket.AddInterfaceConfig:input_type -> wice.InterfaceConfigParams
	7,  // 11: wice.Socket.SetInterfaceConfig:input_type -> wice.InterfaceConfigParams
	0,  // 12: wice.Socket.GetStatus:output_type -> wice.Status
	11, // 13: wice.Socket.StreamEvents:output_type -> wice.Event
	12, // 14: wice.Socket.UnWait:output_type -> wice.Error
	12, // 15: wice.Socket.Stop:output_type -> wice.Error
	12, // 16: wice.Socket.Sync:output_type -> wice.Error
	12, // 17: wice.Socket.RestartPeer:output_type -> wice.Error
	12, // 18: wice.Socket.RemoveInterface:output_type -> wice.Error
	12, // 19: wice.Socket.SyncInterfaceConfig:output_type -> wice.Error
	12, // 20: wice.Socket.AddInterfaceConfig:output_type -> wice.Error
	12, // 21: wice.Socket.SetInterfaceConfig:output_type -> wice.Error
	12, // [12:22] is the sub-list for method output_type
	2,  // [2:12] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_socket_proto_init() }
func file_socket_proto_init() {
	if File_socket_proto != nil {
		return
	}
	file_common_proto_init()
	file_event_proto_init()
	file_interface_proto_init()
	file_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_socket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_socket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnWaitParams); i {
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
		file_socket_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopParams); i {
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
		file_socket_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamEventsParams); i {
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
		file_socket_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncParams); i {
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
		file_socket_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveInterfaceParams); i {
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
		file_socket_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestartPeerParams); i {
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
		file_socket_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterfaceConfigParams); i {
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
			RawDescriptor: file_socket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_socket_proto_goTypes,
		DependencyIndexes: file_socket_proto_depIdxs,
		MessageInfos:      file_socket_proto_msgTypes,
	}.Build()
	File_socket_proto = out.File
	file_socket_proto_rawDesc = nil
	file_socket_proto_goTypes = nil
	file_socket_proto_depIdxs = nil
}
