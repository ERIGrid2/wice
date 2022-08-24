// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: candidate.proto

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

// ICE Candidate types
// See: https://datatracker.ietf.org/doc/html/rfc8445#section-5.1.1
type Candidate_Type int32

const (
	Candidate_TYPE_UNSPECIFIED      Candidate_Type = 0
	Candidate_TYPE_HOST             Candidate_Type = 1
	Candidate_TYPE_SERVER_REFLEXIVE Candidate_Type = 2
	Candidate_TYPE_PEER_REFLEXIVE   Candidate_Type = 3
	Candidate_TYPE_RELAY            Candidate_Type = 4
)

// Enum value maps for Candidate_Type.
var (
	Candidate_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_HOST",
		2: "TYPE_SERVER_REFLEXIVE",
		3: "TYPE_PEER_REFLEXIVE",
		4: "TYPE_RELAY",
	}
	Candidate_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":      0,
		"TYPE_HOST":             1,
		"TYPE_SERVER_REFLEXIVE": 2,
		"TYPE_PEER_REFLEXIVE":   3,
		"TYPE_RELAY":            4,
	}
)

func (x Candidate_Type) Enum() *Candidate_Type {
	p := new(Candidate_Type)
	*p = x
	return p
}

func (x Candidate_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Candidate_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_candidate_proto_enumTypes[0].Descriptor()
}

func (Candidate_Type) Type() protoreflect.EnumType {
	return &file_candidate_proto_enumTypes[0]
}

func (x Candidate_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Candidate_Type.Descriptor instead.
func (Candidate_Type) EnumDescriptor() ([]byte, []int) {
	return file_candidate_proto_rawDescGZIP(), []int{1, 0}
}

type Candidate_NetworkType int32

const (
	Candidate_NETWORK_TYPE_UNSPECIFIED Candidate_NetworkType = 0
	Candidate_NETWORK_TYPE_UDP4        Candidate_NetworkType = 1
	Candidate_NETWORK_TYPE_UDP6        Candidate_NetworkType = 2
	Candidate_NETWORK_TYPE_TCP4        Candidate_NetworkType = 3
	Candidate_NETWORK_TYPE_TCP6        Candidate_NetworkType = 4
)

// Enum value maps for Candidate_NetworkType.
var (
	Candidate_NetworkType_name = map[int32]string{
		0: "NETWORK_TYPE_UNSPECIFIED",
		1: "NETWORK_TYPE_UDP4",
		2: "NETWORK_TYPE_UDP6",
		3: "NETWORK_TYPE_TCP4",
		4: "NETWORK_TYPE_TCP6",
	}
	Candidate_NetworkType_value = map[string]int32{
		"NETWORK_TYPE_UNSPECIFIED": 0,
		"NETWORK_TYPE_UDP4":        1,
		"NETWORK_TYPE_UDP6":        2,
		"NETWORK_TYPE_TCP4":        3,
		"NETWORK_TYPE_TCP6":        4,
	}
)

func (x Candidate_NetworkType) Enum() *Candidate_NetworkType {
	p := new(Candidate_NetworkType)
	*p = x
	return p
}

func (x Candidate_NetworkType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Candidate_NetworkType) Descriptor() protoreflect.EnumDescriptor {
	return file_candidate_proto_enumTypes[1].Descriptor()
}

func (Candidate_NetworkType) Type() protoreflect.EnumType {
	return &file_candidate_proto_enumTypes[1]
}

func (x Candidate_NetworkType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Candidate_NetworkType.Descriptor instead.
func (Candidate_NetworkType) EnumDescriptor() ([]byte, []int) {
	return file_candidate_proto_rawDescGZIP(), []int{1, 1}
}

// Type of TCP candidate
// See: https://datatracker.ietf.org/doc/html/rfc6544
type Candidate_TCPType int32

const (
	Candidate_TCP_TYPE_UNSPECIFIED       Candidate_TCPType = 0
	Candidate_TCP_TYPE_ACTIVE            Candidate_TCPType = 1
	Candidate_TCP_TYPE_PASSIVE           Candidate_TCPType = 2
	Candidate_TCP_TYPE_SIMULTANEOUS_OPEN Candidate_TCPType = 3
)

// Enum value maps for Candidate_TCPType.
var (
	Candidate_TCPType_name = map[int32]string{
		0: "TCP_TYPE_UNSPECIFIED",
		1: "TCP_TYPE_ACTIVE",
		2: "TCP_TYPE_PASSIVE",
		3: "TCP_TYPE_SIMULTANEOUS_OPEN",
	}
	Candidate_TCPType_value = map[string]int32{
		"TCP_TYPE_UNSPECIFIED":       0,
		"TCP_TYPE_ACTIVE":            1,
		"TCP_TYPE_PASSIVE":           2,
		"TCP_TYPE_SIMULTANEOUS_OPEN": 3,
	}
)

func (x Candidate_TCPType) Enum() *Candidate_TCPType {
	p := new(Candidate_TCPType)
	*p = x
	return p
}

func (x Candidate_TCPType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Candidate_TCPType) Descriptor() protoreflect.EnumDescriptor {
	return file_candidate_proto_enumTypes[2].Descriptor()
}

func (Candidate_TCPType) Type() protoreflect.EnumType {
	return &file_candidate_proto_enumTypes[2]
}

func (x Candidate_TCPType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Candidate_TCPType.Descriptor instead.
func (Candidate_TCPType) EnumDescriptor() ([]byte, []int) {
	return file_candidate_proto_rawDescGZIP(), []int{1, 2}
}

// The Related Address conveys transport addresses related to the candidate,
// useful for diagnostics and other purposes.
// See: https://datatracker.ietf.org/doc/html/rfc8839#section-5.1
type RelatedAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port    int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *RelatedAddress) Reset() {
	*x = RelatedAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_candidate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelatedAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelatedAddress) ProtoMessage() {}

func (x *RelatedAddress) ProtoReflect() protoreflect.Message {
	mi := &file_candidate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelatedAddress.ProtoReflect.Descriptor instead.
func (*RelatedAddress) Descriptor() ([]byte, []int) {
	return file_candidate_proto_rawDescGZIP(), []int{0}
}

func (x *RelatedAddress) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RelatedAddress) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

// An ICE Candidate contains a transport address for a candidate that can be used for connectivity checks.
// See: https://datatracker.ietf.org/doc/html/rfc8839#section-5.1
type Candidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of candidate
	Type        Candidate_Type        `protobuf:"varint,1,opt,name=type,proto3,enum=wice.Candidate_Type" json:"type,omitempty"`
	NetworkType Candidate_NetworkType `protobuf:"varint,2,opt,name=network_type,json=networkType,proto3,enum=wice.Candidate_NetworkType" json:"network_type,omitempty"`
	TcpType     Candidate_TCPType     `protobuf:"varint,3,opt,name=tcp_type,json=tcpType,proto3,enum=wice.Candidate_TCPType" json:"tcp_type,omitempty"`
	// An identifier that is equivalent for two candidates that are of the same type, share the same base, and come from the same STUN server.
	Foundation string `protobuf:"bytes,4,opt,name=foundation,proto3" json:"foundation,omitempty"`
	// A positive integer between 1 and 256 that identifies the specific component of the media stream for which this is a candidate.
	Component int32 `protobuf:"varint,5,opt,name=component,proto3" json:"component,omitempty"`
	// A positive integer between 1 and (2**31 - 1).
	Priority int32 `protobuf:"varint,6,opt,name=priority,proto3" json:"priority,omitempty"`
	// The IP address of the candidate.
	Address string `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	// The port of the candidate.
	Port int32 `protobuf:"varint,8,opt,name=port,proto3" json:"port,omitempty"`
	// The related address conveys transport addresses related to the candidate, useful for diagnostics and other purposes.
	RelatedAddress *RelatedAddress `protobuf:"bytes,9,opt,name=related_address,json=relatedAddress,proto3" json:"related_address,omitempty"`
}

func (x *Candidate) Reset() {
	*x = Candidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_candidate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candidate) ProtoMessage() {}

func (x *Candidate) ProtoReflect() protoreflect.Message {
	mi := &file_candidate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candidate.ProtoReflect.Descriptor instead.
func (*Candidate) Descriptor() ([]byte, []int) {
	return file_candidate_proto_rawDescGZIP(), []int{1}
}

func (x *Candidate) GetType() Candidate_Type {
	if x != nil {
		return x.Type
	}
	return Candidate_TYPE_UNSPECIFIED
}

func (x *Candidate) GetNetworkType() Candidate_NetworkType {
	if x != nil {
		return x.NetworkType
	}
	return Candidate_NETWORK_TYPE_UNSPECIFIED
}

func (x *Candidate) GetTcpType() Candidate_TCPType {
	if x != nil {
		return x.TcpType
	}
	return Candidate_TCP_TYPE_UNSPECIFIED
}

func (x *Candidate) GetFoundation() string {
	if x != nil {
		return x.Foundation
	}
	return ""
}

func (x *Candidate) GetComponent() int32 {
	if x != nil {
		return x.Component
	}
	return 0
}

func (x *Candidate) GetPriority() int32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Candidate) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Candidate) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Candidate) GetRelatedAddress() *RelatedAddress {
	if x != nil {
		return x.RelatedAddress
	}
	return nil
}

var File_candidate_proto protoreflect.FileDescriptor

var file_candidate_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x77, 0x69, 0x63, 0x65, 0x22, 0x3e, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0xdb, 0x05, 0x0a, 0x09, 0x43, 0x61, 0x6e, 0x64,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x3e, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x32, 0x0a, 0x08, 0x74, 0x63, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x54, 0x43, 0x50, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x74, 0x63, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3d, 0x0a, 0x0f, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x77, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x6f, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x48, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x46, 0x4c, 0x45, 0x58, 0x49, 0x56, 0x45,
	0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x45, 0x45, 0x52, 0x5f,
	0x52, 0x45, 0x46, 0x4c, 0x45, 0x58, 0x49, 0x56, 0x45, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x59, 0x10, 0x04, 0x22, 0x87, 0x01, 0x0a, 0x0b,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x4e,
	0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x45, 0x54,
	0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x44, 0x50, 0x34, 0x10, 0x01,
	0x12, 0x15, 0x0a, 0x11, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x44, 0x50, 0x36, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x4e, 0x45, 0x54, 0x57, 0x4f,
	0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x43, 0x50, 0x34, 0x10, 0x03, 0x12, 0x15,
	0x0a, 0x11, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54,
	0x43, 0x50, 0x36, 0x10, 0x04, 0x22, 0x6e, 0x0a, 0x07, 0x54, 0x43, 0x50, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x14, 0x54, 0x43, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x43,
	0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12,
	0x14, 0x0a, 0x10, 0x54, 0x43, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x53, 0x53,
	0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x43, 0x50, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x49, 0x4d, 0x55, 0x4c, 0x54, 0x41, 0x4e, 0x45, 0x4f, 0x55, 0x53, 0x5f, 0x4f,
	0x50, 0x45, 0x4e, 0x10, 0x03, 0x42, 0x16, 0x5a, 0x14, 0x72, 0x69, 0x61, 0x73, 0x63, 0x2e, 0x65,
	0x75, 0x2f, 0x77, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_candidate_proto_rawDescOnce sync.Once
	file_candidate_proto_rawDescData = file_candidate_proto_rawDesc
)

func file_candidate_proto_rawDescGZIP() []byte {
	file_candidate_proto_rawDescOnce.Do(func() {
		file_candidate_proto_rawDescData = protoimpl.X.CompressGZIP(file_candidate_proto_rawDescData)
	})
	return file_candidate_proto_rawDescData
}

var file_candidate_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_candidate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_candidate_proto_goTypes = []interface{}{
	(Candidate_Type)(0),        // 0: wice.Candidate.Type
	(Candidate_NetworkType)(0), // 1: wice.Candidate.NetworkType
	(Candidate_TCPType)(0),     // 2: wice.Candidate.TCPType
	(*RelatedAddress)(nil),     // 3: wice.RelatedAddress
	(*Candidate)(nil),          // 4: wice.Candidate
}
var file_candidate_proto_depIdxs = []int32{
	0, // 0: wice.Candidate.type:type_name -> wice.Candidate.Type
	1, // 1: wice.Candidate.network_type:type_name -> wice.Candidate.NetworkType
	2, // 2: wice.Candidate.tcp_type:type_name -> wice.Candidate.TCPType
	3, // 3: wice.Candidate.related_address:type_name -> wice.RelatedAddress
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_candidate_proto_init() }
func file_candidate_proto_init() {
	if File_candidate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_candidate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelatedAddress); i {
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
		file_candidate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Candidate); i {
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
			RawDescriptor: file_candidate_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_candidate_proto_goTypes,
		DependencyIndexes: file_candidate_proto_depIdxs,
		EnumInfos:         file_candidate_proto_enumTypes,
		MessageInfos:      file_candidate_proto_msgTypes,
	}.Build()
	File_candidate_proto = out.File
	file_candidate_proto_rawDesc = nil
	file_candidate_proto_goTypes = nil
	file_candidate_proto_depIdxs = nil
}
