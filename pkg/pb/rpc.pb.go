// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: rpc.proto

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

type StatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interfaces []*Interface `protobuf:"bytes,1,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
}

func (x *StatusResp) Reset() {
	*x = StatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResp) ProtoMessage() {}

func (x *StatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResp.ProtoReflect.Descriptor instead.
func (*StatusResp) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *StatusResp) GetInterfaces() []*Interface {
	if x != nil {
		return x.Interfaces
	}
	return nil
}

type StatusParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intf string `protobuf:"bytes,1,opt,name=intf,proto3" json:"intf,omitempty"`
	Peer []byte `protobuf:"bytes,2,opt,name=peer,proto3" json:"peer,omitempty"`
}

func (x *StatusParams) Reset() {
	*x = StatusParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusParams) ProtoMessage() {}

func (x *StatusParams) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusParams.ProtoReflect.Descriptor instead.
func (*StatusParams) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *StatusParams) GetIntf() string {
	if x != nil {
		return x.Intf
	}
	return ""
}

func (x *StatusParams) GetPeer() []byte {
	if x != nil {
		return x.Peer
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
		mi := &file_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnWaitParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnWaitParams) ProtoMessage() {}

func (x *UnWaitParams) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[2]
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
	return file_rpc_proto_rawDescGZIP(), []int{2}
}

type StopParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopParams) Reset() {
	*x = StopParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopParams) ProtoMessage() {}

func (x *StopParams) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[3]
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
	return file_rpc_proto_rawDescGZIP(), []int{3}
}

type StreamEventsParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamEventsParams) Reset() {
	*x = StreamEventsParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamEventsParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamEventsParams) ProtoMessage() {}

func (x *StreamEventsParams) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[4]
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
	return file_rpc_proto_rawDescGZIP(), []int{4}
}

type SyncParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SyncParams) Reset() {
	*x = SyncParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncParams) ProtoMessage() {}

func (x *SyncParams) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[5]
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
	return file_rpc_proto_rawDescGZIP(), []int{5}
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
		mi := &file_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveInterfaceParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveInterfaceParams) ProtoMessage() {}

func (x *RemoveInterfaceParams) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[6]
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
	return file_rpc_proto_rawDescGZIP(), []int{6}
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
		mi := &file_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestartPeerParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestartPeerParams) ProtoMessage() {}

func (x *RestartPeerParams) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[7]
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
	return file_rpc_proto_rawDescGZIP(), []int{7}
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
		mi := &file_rpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterfaceConfigParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfaceConfigParams) ProtoMessage() {}

func (x *InterfaceConfigParams) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[8]
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
	return file_rpc_proto_rawDescGZIP(), []int{8}
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

type GetSignalingMessageParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Interface name
	Intf string `protobuf:"bytes,1,opt,name=intf,proto3" json:"intf,omitempty"`
	// Public key of peer
	Peer []byte `protobuf:"bytes,2,opt,name=peer,proto3" json:"peer,omitempty"`
}

func (x *GetSignalingMessageParams) Reset() {
	*x = GetSignalingMessageParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSignalingMessageParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSignalingMessageParams) ProtoMessage() {}

func (x *GetSignalingMessageParams) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSignalingMessageParams.ProtoReflect.Descriptor instead.
func (*GetSignalingMessageParams) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{9}
}

func (x *GetSignalingMessageParams) GetIntf() string {
	if x != nil {
		return x.Intf
	}
	return ""
}

func (x *GetSignalingMessageParams) GetPeer() []byte {
	if x != nil {
		return x.Peer
	}
	return nil
}

type GetSignalingMessageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Envelope *SignalingEnvelope `protobuf:"bytes,1,opt,name=envelope,proto3" json:"envelope,omitempty"`
}

func (x *GetSignalingMessageResp) Reset() {
	*x = GetSignalingMessageResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSignalingMessageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSignalingMessageResp) ProtoMessage() {}

func (x *GetSignalingMessageResp) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSignalingMessageResp.ProtoReflect.Descriptor instead.
func (*GetSignalingMessageResp) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{10}
}

func (x *GetSignalingMessageResp) GetEnvelope() *SignalingEnvelope {
	if x != nil {
		return x.Envelope
	}
	return nil
}

type PutSignalingMessageParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Envelope *SignalingEnvelope `protobuf:"bytes,1,opt,name=envelope,proto3" json:"envelope,omitempty"`
}

func (x *PutSignalingMessageParams) Reset() {
	*x = PutSignalingMessageParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutSignalingMessageParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutSignalingMessageParams) ProtoMessage() {}

func (x *PutSignalingMessageParams) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutSignalingMessageParams.ProtoReflect.Descriptor instead.
func (*PutSignalingMessageParams) Descriptor() ([]byte, []int) {
	return file_rpc_proto_rawDescGZIP(), []int{11}
}

func (x *PutSignalingMessageParams) GetEnvelope() *SignalingEnvelope {
	if x != nil {
		return x.Envelope
	}
	return nil
}

var File_rpc_proto protoreflect.FileDescriptor

var file_rpc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x77, 0x69, 0x63,
	0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0a,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x0a, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52,
	0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x0c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x6e, 0x74, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x74, 0x66, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x70,
	0x65, 0x65, 0x72, 0x22, 0x0e, 0x0a, 0x0c, 0x55, 0x6e, 0x57, 0x61, 0x69, 0x74, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x79, 0x6e, 0x63, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x2b, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x6e, 0x74, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e,
	0x74, 0x66, 0x22, 0x3b, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x65, 0x65,
	0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x74, 0x66, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x74, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x65, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x22,
	0x60, 0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x09,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x22, 0x43, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e,
	0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x6e, 0x74, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e,
	0x74, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x22, 0x4e, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x33, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x08, 0x65, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x22, 0x50, 0x0a, 0x19, 0x50, 0x75, 0x74, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x33, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x08,
	0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x32, 0x99, 0x01, 0x0a, 0x06, 0x53, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x39, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e,
	0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2b,
	0x0a, 0x06, 0x55, 0x6e, 0x57, 0x61, 0x69, 0x74, 0x12, 0x12, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e,
	0x55, 0x6e, 0x57, 0x61, 0x69, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e, 0x77,
	0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x27, 0x0a, 0x04, 0x53,
	0x74, 0x6f, 0x70, 0x12, 0x10, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x32, 0x8d, 0x04, 0x0a, 0x07, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x27, 0x0a, 0x04, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x10, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x79, 0x6e, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e, 0x77, 0x69, 0x63,
	0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x10, 0x2e, 0x77, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x1b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0b,
	0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x13, 0x53, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x40, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x77, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1d, 0x2e, 0x77,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x45, 0x0a,
	0x13, 0x50, 0x75, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x75, 0x74, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x32, 0x50, 0x0a, 0x17, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x35, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x65, 0x65, 0x72, 0x12, 0x17,
	0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x50, 0x65, 0x65,
	0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x32, 0xb1, 0x01, 0x0a, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x69, 0x6e, 0x67, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x57, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1f, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x6c, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x1d, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x13, 0x50, 0x75, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x2e, 0x77, 0x69, 0x63,
	0x65, 0x2e, 0x50, 0x75, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x0b, 0x2e, 0x77, 0x69,
	0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x72, 0x69,
	0x61, 0x73, 0x63, 0x2e, 0x65, 0x75, 0x2f, 0x77, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_proto_rawDescOnce sync.Once
	file_rpc_proto_rawDescData = file_rpc_proto_rawDesc
)

func file_rpc_proto_rawDescGZIP() []byte {
	file_rpc_proto_rawDescOnce.Do(func() {
		file_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_proto_rawDescData)
	})
	return file_rpc_proto_rawDescData
}

var file_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_rpc_proto_goTypes = []interface{}{
	(*StatusResp)(nil),                // 0: wice.StatusResp
	(*StatusParams)(nil),              // 1: wice.StatusParams
	(*UnWaitParams)(nil),              // 2: wice.UnWaitParams
	(*StopParams)(nil),                // 3: wice.StopParams
	(*StreamEventsParams)(nil),        // 4: wice.StreamEventsParams
	(*SyncParams)(nil),                // 5: wice.SyncParams
	(*RemoveInterfaceParams)(nil),     // 6: wice.RemoveInterfaceParams
	(*RestartPeerParams)(nil),         // 7: wice.RestartPeerParams
	(*InterfaceConfigParams)(nil),     // 8: wice.InterfaceConfigParams
	(*GetSignalingMessageParams)(nil), // 9: wice.GetSignalingMessageParams
	(*GetSignalingMessageResp)(nil),   // 10: wice.GetSignalingMessageResp
	(*PutSignalingMessageParams)(nil), // 11: wice.PutSignalingMessageParams
	(*Interface)(nil),                 // 12: wice.Interface
	(*InterfaceConfig)(nil),           // 13: wice.InterfaceConfig
	(*SignalingEnvelope)(nil),         // 14: wice.SignalingEnvelope
	(*Event)(nil),                     // 15: wice.Event
	(*Empty)(nil),                     // 16: wice.Empty
}
var file_rpc_proto_depIdxs = []int32{
	12, // 0: wice.StatusResp.interfaces:type_name -> wice.Interface
	13, // 1: wice.InterfaceConfigParams.interface:type_name -> wice.InterfaceConfig
	14, // 2: wice.GetSignalingMessageResp.envelope:type_name -> wice.SignalingEnvelope
	14, // 3: wice.PutSignalingMessageParams.envelope:type_name -> wice.SignalingEnvelope
	4,  // 4: wice.Socket.StreamEvents:input_type -> wice.StreamEventsParams
	2,  // 5: wice.Socket.UnWait:input_type -> wice.UnWaitParams
	3,  // 6: wice.Socket.Stop:input_type -> wice.StopParams
	5,  // 7: wice.Watcher.Sync:input_type -> wice.SyncParams
	1,  // 8: wice.Watcher.GetStatus:input_type -> wice.StatusParams
	6,  // 9: wice.Watcher.RemoveInterface:input_type -> wice.RemoveInterfaceParams
	8,  // 10: wice.Watcher.SyncInterfaceConfig:input_type -> wice.InterfaceConfigParams
	8,  // 11: wice.Watcher.AddInterfaceConfig:input_type -> wice.InterfaceConfigParams
	8,  // 12: wice.Watcher.SetInterfaceConfig:input_type -> wice.InterfaceConfigParams
	9,  // 13: wice.Watcher.GetSignalingMessage:input_type -> wice.GetSignalingMessageParams
	11, // 14: wice.Watcher.PutSignalingMessage:input_type -> wice.PutSignalingMessageParams
	7,  // 15: wice.EndpointDiscoverySocket.RestartPeer:input_type -> wice.RestartPeerParams
	9,  // 16: wice.SignalingSocket.GetSignalingMessage:input_type -> wice.GetSignalingMessageParams
	11, // 17: wice.SignalingSocket.PutSignalingMessage:input_type -> wice.PutSignalingMessageParams
	15, // 18: wice.Socket.StreamEvents:output_type -> wice.Event
	16, // 19: wice.Socket.UnWait:output_type -> wice.Empty
	16, // 20: wice.Socket.Stop:output_type -> wice.Empty
	16, // 21: wice.Watcher.Sync:output_type -> wice.Empty
	0,  // 22: wice.Watcher.GetStatus:output_type -> wice.StatusResp
	16, // 23: wice.Watcher.RemoveInterface:output_type -> wice.Empty
	16, // 24: wice.Watcher.SyncInterfaceConfig:output_type -> wice.Empty
	16, // 25: wice.Watcher.AddInterfaceConfig:output_type -> wice.Empty
	16, // 26: wice.Watcher.SetInterfaceConfig:output_type -> wice.Empty
	10, // 27: wice.Watcher.GetSignalingMessage:output_type -> wice.GetSignalingMessageResp
	16, // 28: wice.Watcher.PutSignalingMessage:output_type -> wice.Empty
	16, // 29: wice.EndpointDiscoverySocket.RestartPeer:output_type -> wice.Empty
	10, // 30: wice.SignalingSocket.GetSignalingMessage:output_type -> wice.GetSignalingMessageResp
	16, // 31: wice.SignalingSocket.PutSignalingMessage:output_type -> wice.Empty
	18, // [18:32] is the sub-list for method output_type
	4,  // [4:18] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_proto_init() }
func file_rpc_proto_init() {
	if File_rpc_proto != nil {
		return
	}
	file_common_proto_init()
	file_event_proto_init()
	file_interface_proto_init()
	file_config_proto_init()
	file_signaling_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResp); i {
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
		file_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusParams); i {
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
		file_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSignalingMessageParams); i {
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
		file_rpc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSignalingMessageResp); i {
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
		file_rpc_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutSignalingMessageParams); i {
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
			RawDescriptor: file_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_rpc_proto_goTypes,
		DependencyIndexes: file_rpc_proto_depIdxs,
		MessageInfos:      file_rpc_proto_msgTypes,
	}.Build()
	File_rpc_proto = out.File
	file_rpc_proto_rawDesc = nil
	file_rpc_proto_goTypes = nil
	file_rpc_proto_depIdxs = nil
}
