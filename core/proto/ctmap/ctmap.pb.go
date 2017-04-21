// Code generated by protoc-gen-go.
// source: ctmap.proto
// DO NOT EDIT!

/*
Package ctmap is a generated protocol buffer package.

It is generated from these files:
	ctmap.proto

It has these top-level messages:
	MapHead
	SignedMapHead
*/
package ctmap

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import signature "github.com/google/keytransparency/core/proto/signature"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// MapHead is the head node of the Merkle Tree as well as additional metadata
// for the tree.
type MapHead struct {
	// realm is the domain identifier for the transparent map.
	Realm string `protobuf:"bytes,1,opt,name=realm" json:"realm,omitempty"`
	// epoch is the epoch number of this map head.
	Epoch int64 `protobuf:"varint,2,opt,name=epoch" json:"epoch,omitempty"`
	// root is the value of the root node of the Merkle tree.
	Root []byte `protobuf:"bytes,3,opt,name=root,proto3" json:"root,omitempty"`
	// issue_time is the time when this epoch was created. Monotonically
	// increasing.
	IssueTime *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=issue_time,json=issueTime" json:"issue_time,omitempty"`
	// max_sequence_number contains the maximum mutation sequence number included
	// in this map head.
	MaxSequenceNumber uint64 `protobuf:"varint,5,opt,name=max_sequence_number,json=maxSequenceNumber" json:"max_sequence_number,omitempty"`
}

func (m *MapHead) Reset()                    { *m = MapHead{} }
func (m *MapHead) String() string            { return proto.CompactTextString(m) }
func (*MapHead) ProtoMessage()               {}
func (*MapHead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MapHead) GetRealm() string {
	if m != nil {
		return m.Realm
	}
	return ""
}

func (m *MapHead) GetEpoch() int64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *MapHead) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *MapHead) GetIssueTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.IssueTime
	}
	return nil
}

func (m *MapHead) GetMaxSequenceNumber() uint64 {
	if m != nil {
		return m.MaxSequenceNumber
	}
	return 0
}

// SignedMapHead represents a signed state of the Merkel Tree.
type SignedMapHead struct {
	// map_head contains the head node of the Merkle Tree along with other
	// metadata.
	MapHead *MapHead `protobuf:"bytes,1,opt,name=map_head,json=mapHead" json:"map_head,omitempty"`
	// signatures is a set of map_head signatures. Each signature is identified by
	// the first 64 bits of the public key that verifies it.
	Signatures map[string]*signature.DigitallySigned `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SignedMapHead) Reset()                    { *m = SignedMapHead{} }
func (m *SignedMapHead) String() string            { return proto.CompactTextString(m) }
func (*SignedMapHead) ProtoMessage()               {}
func (*SignedMapHead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignedMapHead) GetMapHead() *MapHead {
	if m != nil {
		return m.MapHead
	}
	return nil
}

func (m *SignedMapHead) GetSignatures() map[string]*signature.DigitallySigned {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func init() {
	proto.RegisterType((*MapHead)(nil), "ctmap.MapHead")
	proto.RegisterType((*SignedMapHead)(nil), "ctmap.SignedMapHead")
}

func init() { proto.RegisterFile("ctmap.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbd, 0x52, 0xeb, 0x30,
	0x10, 0x85, 0x47, 0x71, 0x7c, 0x73, 0x23, 0xdf, 0x1f, 0x10, 0x14, 0x1e, 0x37, 0x78, 0x32, 0x14,
	0xa6, 0xb1, 0x99, 0xd0, 0x00, 0x75, 0x60, 0x68, 0xa0, 0x50, 0x68, 0xa8, 0x3c, 0xb2, 0xb3, 0x38,
	0x9e, 0x58, 0x96, 0x91, 0x65, 0x26, 0x7e, 0x2f, 0xde, 0x86, 0x97, 0x61, 0x2c, 0x29, 0x04, 0xe8,
	0xf6, 0xec, 0x7e, 0x5a, 0x9d, 0x3d, 0xd8, 0xcb, 0x15, 0x67, 0x4d, 0xdc, 0x48, 0xa1, 0x04, 0x71,
	0xb5, 0x08, 0x4e, 0x0a, 0x21, 0x8a, 0x0a, 0x12, 0xdd, 0xcc, 0xba, 0xe7, 0x44, 0x95, 0x1c, 0x5a,
	0xc5, 0xb8, 0xe5, 0x82, 0xdb, 0xa2, 0x54, 0xeb, 0x2e, 0x8b, 0x73, 0xc1, 0x13, 0xcb, 0x6e, 0xa0,
	0x57, 0x92, 0xd5, 0x6d, 0xc3, 0x24, 0xd4, 0x79, 0x9f, 0xe4, 0x42, 0xda, 0x05, 0x49, 0x5b, 0x16,
	0x35, 0x53, 0x9d, 0x84, 0x7d, 0x65, 0xf6, 0xcc, 0xde, 0x10, 0x9e, 0xdc, 0xb3, 0xe6, 0x0e, 0xd8,
	0x8a, 0x1c, 0x63, 0x57, 0x02, 0xab, 0xb8, 0x8f, 0x42, 0x14, 0x4d, 0xa9, 0x11, 0x43, 0x17, 0x1a,
	0x91, 0xaf, 0xfd, 0x51, 0x88, 0x22, 0x87, 0x1a, 0x41, 0x08, 0x1e, 0x4b, 0x21, 0x94, 0xef, 0x84,
	0x28, 0xfa, 0x43, 0x75, 0x4d, 0xae, 0x30, 0x2e, 0xdb, 0xb6, 0x83, 0x74, 0x30, 0xeb, 0x8f, 0x43,
	0x14, 0x79, 0xf3, 0x20, 0x36, 0xee, 0xe2, 0xdd, 0x25, 0xf1, 0xe3, 0xee, 0x12, 0x3a, 0xd5, 0xf4,
	0xa0, 0x49, 0x8c, 0x8f, 0x38, 0xdb, 0xa6, 0x2d, 0xbc, 0x74, 0x50, 0xe7, 0x90, 0xd6, 0x1d, 0xcf,
	0x40, 0xfa, 0x6e, 0x88, 0xa2, 0x31, 0x3d, 0xe4, 0x6c, 0xbb, 0xb4, 0x93, 0x07, 0x3d, 0x98, 0xbd,
	0x23, 0xfc, 0x77, 0x59, 0x16, 0x35, 0xac, 0x76, 0xe6, 0xcf, 0xf0, 0x6f, 0xce, 0x9a, 0x74, 0x0d,
	0x6c, 0xa5, 0xfd, 0x7b, 0xf3, 0x7f, 0xb1, 0x09, 0xd6, 0x12, 0x74, 0xc2, 0x2d, 0xba, 0xc0, 0xf8,
	0x33, 0x86, 0xd6, 0x1f, 0x85, 0x4e, 0xe4, 0xcd, 0x4f, 0x2d, 0xfc, 0x6d, 0xa9, 0x56, 0x06, 0xbb,
	0xa9, 0x95, 0xec, 0xe9, 0x97, 0x77, 0xc1, 0x13, 0xfe, 0xff, 0x63, 0x4c, 0x0e, 0xb0, 0xb3, 0x81,
	0xde, 0xc6, 0x37, 0x94, 0xe4, 0x1c, 0xbb, 0xaf, 0xac, 0xea, 0x40, 0x87, 0x37, 0xa4, 0xb1, 0xcf,
	0x7f, 0x51, 0x16, 0xa5, 0x62, 0x55, 0xd5, 0x9b, 0x2f, 0xa9, 0x01, 0xaf, 0x47, 0x97, 0x28, 0xfb,
	0xa5, 0xc3, 0xba, 0xf8, 0x08, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xc8, 0x6d, 0x81, 0x1a, 0x02, 0x00,
	0x00,
}
