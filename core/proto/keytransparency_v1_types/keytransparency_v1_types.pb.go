// Code generated by protoc-gen-go.
// source: keytransparency_v1_types.proto
// DO NOT EDIT!

/*
Package keytransparency_v1_types is a generated protocol buffer package.

Key Transparency Service

The Key Transparency Service API consists of a map of user names to public
keys. Each user name also has a history of public keys that have been
associated with it.

It is generated from these files:
	keytransparency_v1_types.proto

It has these top-level messages:
	Committed
	Profile
	EntryUpdate
	Entry
	PublicKey
	KeyValue
	SignedKV
	LeafProof
	GetEntryRequest
	GetEntryResponse
	ListEntryHistoryRequest
	ListEntryHistoryResponse
	UpdateEntryRequest
	UpdateEntryResponse
*/
package keytransparency_v1_types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ctmap "github.com/google/keytransparency/core/proto/ctmap"
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

// Committed represents the data committed to in a cryptographic commitment.
// commitment = HMAC_SHA512_256(key, data)
type Committed struct {
	// key is the 16 byte random commitment key.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// data is the data being committed to.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Committed) Reset()                    { *m = Committed{} }
func (m *Committed) String() string            { return proto.CompactTextString(m) }
func (*Committed) ProtoMessage()               {}
func (*Committed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Committed) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Committed) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Profile contains data hidden behind the cryptographic commitment.
type Profile struct {
	// Keys is a map of application IDs to keys.
	Keys map[string][]byte `protobuf:"bytes,1,rep,name=keys" json:"keys,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Profile) GetKeys() map[string][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

// EntryUpdate contains the user entry update(s).
type EntryUpdate struct {
	// update authorizes the change to profile.
	Update *SignedKV `protobuf:"bytes,2,opt,name=update" json:"update,omitempty"`
	// commitment contains the serialized Profile protobuf.
	Committed *Committed `protobuf:"bytes,3,opt,name=committed" json:"committed,omitempty"`
}

func (m *EntryUpdate) Reset()                    { *m = EntryUpdate{} }
func (m *EntryUpdate) String() string            { return proto.CompactTextString(m) }
func (*EntryUpdate) ProtoMessage()               {}
func (*EntryUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EntryUpdate) GetUpdate() *SignedKV {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *EntryUpdate) GetCommitted() *Committed {
	if m != nil {
		return m.Committed
	}
	return nil
}

// Entry contains a commitment to profile and a set of authorized update keys.
// Entry is placed in the verifiable map as leaf data.
type Entry struct {
	// commitment is a cryptographic commitment to arbitrary data.
	Commitment []byte `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	// authorized_keys is the set of keys allowed to sign updates for this entry.
	AuthorizedKeys []*PublicKey `protobuf:"bytes,2,rep,name=authorized_keys,json=authorizedKeys" json:"authorized_keys,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Entry) GetCommitment() []byte {
	if m != nil {
		return m.Commitment
	}
	return nil
}

func (m *Entry) GetAuthorizedKeys() []*PublicKey {
	if m != nil {
		return m.AuthorizedKeys
	}
	return nil
}

// PublicKey defines a key this domain uses to sign MapHeads with.
type PublicKey struct {
	// Key formats from Keyczar.
	//
	// Types that are valid to be assigned to KeyType:
	//	*PublicKey_Ed25519
	//	*PublicKey_RsaVerifyingSha256_3072
	//	*PublicKey_EcdsaVerifyingP256
	KeyType isPublicKey_KeyType `protobuf_oneof:"key_type"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isPublicKey_KeyType interface {
	isPublicKey_KeyType()
}

type PublicKey_Ed25519 struct {
	Ed25519 []byte `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof"`
}
type PublicKey_RsaVerifyingSha256_3072 struct {
	RsaVerifyingSha256_3072 []byte `protobuf:"bytes,2,opt,name=rsa_verifying_sha256_3072,json=rsaVerifyingSha2563072,proto3,oneof"`
}
type PublicKey_EcdsaVerifyingP256 struct {
	EcdsaVerifyingP256 []byte `protobuf:"bytes,3,opt,name=ecdsa_verifying_p256,json=ecdsaVerifyingP256,proto3,oneof"`
}

func (*PublicKey_Ed25519) isPublicKey_KeyType()                 {}
func (*PublicKey_RsaVerifyingSha256_3072) isPublicKey_KeyType() {}
func (*PublicKey_EcdsaVerifyingP256) isPublicKey_KeyType()      {}

func (m *PublicKey) GetKeyType() isPublicKey_KeyType {
	if m != nil {
		return m.KeyType
	}
	return nil
}

func (m *PublicKey) GetEd25519() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_Ed25519); ok {
		return x.Ed25519
	}
	return nil
}

func (m *PublicKey) GetRsaVerifyingSha256_3072() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_RsaVerifyingSha256_3072); ok {
		return x.RsaVerifyingSha256_3072
	}
	return nil
}

func (m *PublicKey) GetEcdsaVerifyingP256() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_EcdsaVerifyingP256); ok {
		return x.EcdsaVerifyingP256
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PublicKey) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PublicKey_OneofMarshaler, _PublicKey_OneofUnmarshaler, _PublicKey_OneofSizer, []interface{}{
		(*PublicKey_Ed25519)(nil),
		(*PublicKey_RsaVerifyingSha256_3072)(nil),
		(*PublicKey_EcdsaVerifyingP256)(nil),
	}
}

func _PublicKey_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_3072:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RsaVerifyingSha256_3072)
	case *PublicKey_EcdsaVerifyingP256:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.EcdsaVerifyingP256)
	case nil:
	default:
		return fmt.Errorf("PublicKey.KeyType has unexpected type %T", x)
	}
	return nil
}

func _PublicKey_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PublicKey)
	switch tag {
	case 1: // key_type.ed25519
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_Ed25519{x}
		return true, err
	case 2: // key_type.rsa_verifying_sha256_3072
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_RsaVerifyingSha256_3072{x}
		return true, err
	case 3: // key_type.ecdsa_verifying_p256
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_EcdsaVerifyingP256{x}
		return true, err
	default:
		return false, nil
	}
}

func _PublicKey_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PublicKey)
	// key_type
	switch x := m.KeyType.(type) {
	case *PublicKey_Ed25519:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ed25519)))
		n += len(x.Ed25519)
	case *PublicKey_RsaVerifyingSha256_3072:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RsaVerifyingSha256_3072)))
		n += len(x.RsaVerifyingSha256_3072)
	case *PublicKey_EcdsaVerifyingP256:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.EcdsaVerifyingP256)))
		n += len(x.EcdsaVerifyingP256)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// KeyValue is a map entry.
type KeyValue struct {
	// key contains the map entry key.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value contains the map entry value.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *KeyValue) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// SignedKV is a signed change to a map entry.
type SignedKV struct {
	// key_value is a serialized KeyValue.
	KeyValue *KeyValue `protobuf:"bytes,1,opt,name=key_value,json=keyValue" json:"key_value,omitempty"`
	// signatures on key_value. Must be signed by keys from both previous and
	// current epochs. The first proves ownership of new epoch key, and the
	// second proves that the correct owner is making this change.
	Signatures map[string]*signature.DigitallySigned `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// previous contains the hash of the previous entry that this mutation is
	// modifying creating a hash chain of all mutations. The hash used is
	// CommonJSON in "github.com/benlaurie/objecthash/go/objecthash".
	Previous []byte `protobuf:"bytes,3,opt,name=previous,proto3" json:"previous,omitempty"`
}

func (m *SignedKV) Reset()                    { *m = SignedKV{} }
func (m *SignedKV) String() string            { return proto.CompactTextString(m) }
func (*SignedKV) ProtoMessage()               {}
func (*SignedKV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SignedKV) GetKeyValue() *KeyValue {
	if m != nil {
		return m.KeyValue
	}
	return nil
}

func (m *SignedKV) GetSignatures() map[string]*signature.DigitallySigned {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *SignedKV) GetPrevious() []byte {
	if m != nil {
		return m.Previous
	}
	return nil
}

// LeafProof for a verifiable map leaf.
type LeafProof struct {
	// leaf_data contains an entry stored in the leaf node.
	LeafData []byte `protobuf:"bytes,1,opt,name=leaf_data,json=leafData,proto3" json:"leaf_data,omitempty"`
	// neighbors is a list of all the adjacent nodes along the path
	// from the deepest node to the head.
	Neighbors [][]byte `protobuf:"bytes,2,rep,name=neighbors,proto3" json:"neighbors,omitempty"`
}

func (m *LeafProof) Reset()                    { *m = LeafProof{} }
func (m *LeafProof) String() string            { return proto.CompactTextString(m) }
func (*LeafProof) ProtoMessage()               {}
func (*LeafProof) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *LeafProof) GetLeafData() []byte {
	if m != nil {
		return m.LeafData
	}
	return nil
}

func (m *LeafProof) GetNeighbors() [][]byte {
	if m != nil {
		return m.Neighbors
	}
	return nil
}

// GetEntryRequest for a user object.
type GetEntryRequest struct {
	// user_id is the user identifier. Most commonly an email address.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *GetEntryRequest) Reset()                    { *m = GetEntryRequest{} }
func (m *GetEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEntryRequest) ProtoMessage()               {}
func (*GetEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetEntryRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// GetEntryResponse returns a requested user entry.
type GetEntryResponse struct {
	// vrf is the output of VRF on user_id.
	Vrf []byte `protobuf:"bytes,1,opt,name=vrf,proto3" json:"vrf,omitempty"`
	// vrf_proof is the proof for VRF on user_id.
	VrfProof []byte `protobuf:"bytes,2,opt,name=vrf_proof,json=vrfProof,proto3" json:"vrf_proof,omitempty"`
	// committed contains the profile for this account and connects the data
	// in profile to the commitment in leaf_proof.
	Committed *Committed `protobuf:"bytes,3,opt,name=committed" json:"committed,omitempty"`
	// leaf_proof contains an Entry and an inclusion proof in the sparse Merkle
	// Tree.
	LeafProof *LeafProof `protobuf:"bytes,5,opt,name=leaf_proof,json=leafProof" json:"leaf_proof,omitempty"`
	// smh contains the signed map head for the sparse Merkle Tree.
	// smh is also stored in the append only log.
	Smh *ctmap.SignedMapHead `protobuf:"bytes,6,opt,name=smh" json:"smh,omitempty"`
	// smh_sct is the signed certificate timestamp (SCT) showing that SMH was
	// submitted to CT logs.
	// TODO: Support storing smh in multiple logs.
	SmhSct []byte `protobuf:"bytes,7,opt,name=smh_sct,json=smhSct,proto3" json:"smh_sct,omitempty"`
}

func (m *GetEntryResponse) Reset()                    { *m = GetEntryResponse{} }
func (m *GetEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEntryResponse) ProtoMessage()               {}
func (*GetEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetEntryResponse) GetVrf() []byte {
	if m != nil {
		return m.Vrf
	}
	return nil
}

func (m *GetEntryResponse) GetVrfProof() []byte {
	if m != nil {
		return m.VrfProof
	}
	return nil
}

func (m *GetEntryResponse) GetCommitted() *Committed {
	if m != nil {
		return m.Committed
	}
	return nil
}

func (m *GetEntryResponse) GetLeafProof() *LeafProof {
	if m != nil {
		return m.LeafProof
	}
	return nil
}

func (m *GetEntryResponse) GetSmh() *ctmap.SignedMapHead {
	if m != nil {
		return m.Smh
	}
	return nil
}

func (m *GetEntryResponse) GetSmhSct() []byte {
	if m != nil {
		return m.SmhSct
	}
	return nil
}

// ListEntryHistoryRequest gets a list of historical keys for a user.
type ListEntryHistoryRequest struct {
	// user_id is the user identifier.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// start is the starting epcoh.
	Start int64 `protobuf:"varint,2,opt,name=start" json:"start,omitempty"`
	// page_size is the maximum number of entries to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
}

func (m *ListEntryHistoryRequest) Reset()                    { *m = ListEntryHistoryRequest{} }
func (m *ListEntryHistoryRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryRequest) ProtoMessage()               {}
func (*ListEntryHistoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ListEntryHistoryRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ListEntryHistoryRequest) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ListEntryHistoryRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

// ListEntryHistoryResponse requests a paginated history of keys for a user.
type ListEntryHistoryResponse struct {
	// values represents the list of keys this user_id has contained over time.
	Values []*GetEntryResponse `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	// next_start is the next page token to query for pagination.
	// next_start is 0 when there are no more results to fetch.
	NextStart int64 `protobuf:"varint,2,opt,name=next_start,json=nextStart" json:"next_start,omitempty"`
}

func (m *ListEntryHistoryResponse) Reset()                    { *m = ListEntryHistoryResponse{} }
func (m *ListEntryHistoryResponse) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryResponse) ProtoMessage()               {}
func (*ListEntryHistoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ListEntryHistoryResponse) GetValues() []*GetEntryResponse {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *ListEntryHistoryResponse) GetNextStart() int64 {
	if m != nil {
		return m.NextStart
	}
	return 0
}

// UpdateEntryRequest updates a user's profile.
type UpdateEntryRequest struct {
	// user_id specifies the id for the user who's profile is being updated.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// entry_update contains the user submitted update.
	EntryUpdate *EntryUpdate `protobuf:"bytes,2,opt,name=entry_update,json=entryUpdate" json:"entry_update,omitempty"`
}

func (m *UpdateEntryRequest) Reset()                    { *m = UpdateEntryRequest{} }
func (m *UpdateEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryRequest) ProtoMessage()               {}
func (*UpdateEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *UpdateEntryRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UpdateEntryRequest) GetEntryUpdate() *EntryUpdate {
	if m != nil {
		return m.EntryUpdate
	}
	return nil
}

// UpdateEntryResponse contains a proof once the update has been included in
// the Merkel Tree.
type UpdateEntryResponse struct {
	// proof contains a proof that the update has been included in the tree.
	Proof *GetEntryResponse `protobuf:"bytes,1,opt,name=proof" json:"proof,omitempty"`
}

func (m *UpdateEntryResponse) Reset()                    { *m = UpdateEntryResponse{} }
func (m *UpdateEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryResponse) ProtoMessage()               {}
func (*UpdateEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *UpdateEntryResponse) GetProof() *GetEntryResponse {
	if m != nil {
		return m.Proof
	}
	return nil
}

func init() {
	proto.RegisterType((*Committed)(nil), "keytransparency.v1.types.Committed")
	proto.RegisterType((*Profile)(nil), "keytransparency.v1.types.Profile")
	proto.RegisterType((*EntryUpdate)(nil), "keytransparency.v1.types.EntryUpdate")
	proto.RegisterType((*Entry)(nil), "keytransparency.v1.types.Entry")
	proto.RegisterType((*PublicKey)(nil), "keytransparency.v1.types.PublicKey")
	proto.RegisterType((*KeyValue)(nil), "keytransparency.v1.types.KeyValue")
	proto.RegisterType((*SignedKV)(nil), "keytransparency.v1.types.SignedKV")
	proto.RegisterType((*LeafProof)(nil), "keytransparency.v1.types.LeafProof")
	proto.RegisterType((*GetEntryRequest)(nil), "keytransparency.v1.types.GetEntryRequest")
	proto.RegisterType((*GetEntryResponse)(nil), "keytransparency.v1.types.GetEntryResponse")
	proto.RegisterType((*ListEntryHistoryRequest)(nil), "keytransparency.v1.types.ListEntryHistoryRequest")
	proto.RegisterType((*ListEntryHistoryResponse)(nil), "keytransparency.v1.types.ListEntryHistoryResponse")
	proto.RegisterType((*UpdateEntryRequest)(nil), "keytransparency.v1.types.UpdateEntryRequest")
	proto.RegisterType((*UpdateEntryResponse)(nil), "keytransparency.v1.types.UpdateEntryResponse")
}

func init() { proto.RegisterFile("keytransparency_v1_types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 832 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdf, 0x6f, 0x1b, 0x45,
	0x10, 0xae, 0xed, 0xda, 0xf1, 0x8d, 0x23, 0x52, 0x2d, 0x11, 0x35, 0x06, 0xaa, 0xea, 0x10, 0xa8,
	0x2a, 0xd2, 0xb9, 0xb9, 0xca, 0x29, 0x14, 0x89, 0x42, 0x28, 0xc5, 0x28, 0x41, 0x8a, 0xd6, 0x22,
	0x88, 0xa7, 0xd3, 0xe6, 0x6e, 0x7c, 0x5e, 0xc5, 0xf7, 0x83, 0xdd, 0xbd, 0x83, 0x8b, 0xc4, 0x13,
	0x4f, 0x48, 0x3c, 0xf3, 0x17, 0xf0, 0x87, 0xa2, 0xdd, 0xbd, 0xb3, 0x9d, 0x10, 0x13, 0xaa, 0xbe,
	0x9c, 0x76, 0x67, 0xe7, 0x9b, 0xf9, 0xe6, 0x9b, 0xd9, 0x3d, 0x78, 0x70, 0x81, 0x95, 0x12, 0x2c,
	0x95, 0x39, 0x13, 0x98, 0x86, 0x55, 0x50, 0x1e, 0x04, 0xaa, 0xca, 0x51, 0x7a, 0xb9, 0xc8, 0x54,
	0x46, 0x86, 0xd7, 0xce, 0xbd, 0xf2, 0xc0, 0x33, 0xe7, 0xa3, 0x2f, 0x62, 0xae, 0x16, 0xc5, 0xb9,
	0x17, 0x66, 0xc9, 0x38, 0xce, 0xb2, 0x78, 0x89, 0xe3, 0x6b, 0xbe, 0xe3, 0x30, 0x13, 0x38, 0x36,
	0x71, 0xc6, 0xa1, 0x4a, 0x58, 0x6e, 0xbf, 0x36, 0xf2, 0xe8, 0xd5, 0x6b, 0xe1, 0x25, 0x8f, 0x53,
	0xa6, 0x0a, 0x81, 0xeb, 0x95, 0x8d, 0xe3, 0x1e, 0x80, 0xf3, 0x75, 0x96, 0x24, 0x5c, 0x29, 0x8c,
	0xc8, 0x3d, 0xe8, 0x5c, 0x60, 0x35, 0x6c, 0x3d, 0x6c, 0x3d, 0xda, 0xa5, 0x7a, 0x49, 0x08, 0xdc,
	0x8d, 0x98, 0x62, 0xc3, 0xb6, 0x31, 0x99, 0xb5, 0xfb, 0x7b, 0x0b, 0x76, 0x4e, 0x45, 0x36, 0xe7,
	0x4b, 0x24, 0x2f, 0xe0, 0xee, 0x05, 0x56, 0x72, 0xd8, 0x7a, 0xd8, 0x79, 0x34, 0xf0, 0x3f, 0xf1,
	0xb6, 0xd5, 0xeb, 0xd5, 0x00, 0xef, 0x18, 0x2b, 0xf9, 0x4d, 0xaa, 0x44, 0x45, 0x0d, 0x70, 0xf4,
	0x0c, 0x9c, 0x95, 0x69, 0x33, 0xbf, 0x63, 0xf3, 0xef, 0x43, 0xb7, 0x64, 0xcb, 0x02, 0x6b, 0x02,
	0x76, 0xf3, 0xbc, 0xfd, 0x69, 0xcb, 0xfd, 0xb3, 0x05, 0x03, 0x83, 0xfa, 0x21, 0x8f, 0x98, 0x42,
	0xf2, 0x1c, 0x7a, 0x85, 0x59, 0x19, 0xd7, 0x81, 0xef, 0x6e, 0xe7, 0x32, 0xe3, 0x71, 0x8a, 0xd1,
	0xf1, 0x19, 0xad, 0x11, 0xe4, 0x2b, 0x70, 0xc2, 0x46, 0x84, 0x61, 0xc7, 0xc0, 0x3f, 0xdc, 0x0e,
	0x5f, 0xe9, 0x45, 0xd7, 0x28, 0xb7, 0x80, 0xae, 0xad, 0xe1, 0x01, 0x80, 0xb5, 0x26, 0x98, 0xaa,
	0x5a, 0xca, 0x0d, 0x0b, 0x39, 0x81, 0x3d, 0x56, 0xa8, 0x45, 0x26, 0xf8, 0x25, 0x46, 0x81, 0x11,
	0xaf, 0x6d, 0xc4, 0xfb, 0x8f, 0x8c, 0xa7, 0xc5, 0xf9, 0x92, 0x87, 0xc7, 0x58, 0xd1, 0xb7, 0xd6,
	0x58, 0x2d, 0x9b, 0xfb, 0x77, 0x0b, 0x9c, 0xd5, 0x29, 0x19, 0xc1, 0x0e, 0x46, 0xfe, 0x64, 0x72,
	0xf0, 0x99, 0x4d, 0x3c, 0xbd, 0x43, 0x1b, 0x03, 0xf9, 0x1c, 0xde, 0x15, 0x92, 0x05, 0x25, 0x0a,
	0x3e, 0xaf, 0x78, 0x1a, 0x07, 0x72, 0xc1, 0xfc, 0xc9, 0x61, 0xf0, 0xf4, 0xc9, 0x33, 0xdf, 0xaa,
	0x3b, 0xbd, 0x43, 0xdf, 0x11, 0x92, 0x9d, 0x35, 0x1e, 0x33, 0xe3, 0xa0, 0xcf, 0x89, 0x0f, 0xfb,
	0x18, 0x46, 0x57, 0xe0, 0xb9, 0x3f, 0x39, 0x34, 0x5a, 0x69, 0x1c, 0x31, 0xa7, 0x2b, 0xe4, 0xa9,
	0x3f, 0x39, 0x3c, 0x02, 0xe8, 0x5f, 0x60, 0x65, 0xae, 0x83, 0xeb, 0x43, 0xff, 0x18, 0xab, 0x33,
	0xdd, 0xbc, 0x1b, 0x86, 0xec, 0xc6, 0x26, 0xbb, 0x7f, 0xb5, 0xa1, 0xdf, 0x74, 0x8a, 0xbc, 0x00,
	0x47, 0x07, 0xb3, 0x6e, 0xad, 0xdb, 0x1a, 0xdc, 0xe4, 0xa2, 0x9a, 0x81, 0xcd, 0x4a, 0x01, 0x56,
	0xa3, 0xdf, 0x28, 0xee, 0xdf, 0x3e, 0x22, 0x66, 0x61, 0x41, 0x76, 0x6a, 0x37, 0xa2, 0x90, 0x11,
	0xf4, 0x73, 0x81, 0x25, 0xcf, 0x0a, 0x69, 0x95, 0xa0, 0xab, 0xfd, 0xe8, 0x27, 0xd8, 0xbb, 0x06,
	0xbd, 0x61, 0xba, 0x9f, 0x6c, 0x16, 0x3e, 0xf0, 0x47, 0xde, 0xfa, 0x76, 0xbe, 0xe4, 0x31, 0x57,
	0x6c, 0xb9, 0xac, 0x2c, 0x93, 0xcd, 0xc9, 0x7f, 0x05, 0xce, 0x09, 0xb2, 0xf9, 0xa9, 0xc8, 0xb2,
	0x39, 0x79, 0x0f, 0x9c, 0x25, 0xb2, 0x79, 0x60, 0x6e, 0xa9, 0xd5, 0xb4, 0xaf, 0x0d, 0x2f, 0x99,
	0x62, 0xe4, 0x7d, 0x70, 0x52, 0xe4, 0xf1, 0xe2, 0x3c, 0x13, 0xb6, 0xe6, 0x5d, 0xba, 0x36, 0xb8,
	0x8f, 0x61, 0xef, 0x5b, 0x54, 0xb6, 0x2c, 0xfc, 0xb9, 0x40, 0xa9, 0xc8, 0x7d, 0xd8, 0x29, 0x24,
	0x8a, 0x80, 0x47, 0x35, 0xcd, 0x9e, 0xde, 0x7e, 0x17, 0xb9, 0x7f, 0xb4, 0xe1, 0xde, 0xda, 0x59,
	0xe6, 0x59, 0x2a, 0x4d, 0x27, 0x4b, 0x31, 0x6f, 0x3a, 0x59, 0x0a, 0xc3, 0xa6, 0x14, 0xf3, 0x20,
	0xd7, 0xd4, 0xea, 0x6e, 0xf6, 0x4b, 0x51, 0x53, 0x7d, 0xf3, 0x5b, 0x46, 0x8e, 0x00, 0x4c, 0xb5,
	0x36, 0x41, 0xf7, 0xb6, 0x18, 0x2b, 0x99, 0xa8, 0x11, 0xc9, 0xd2, 0xf8, 0x18, 0x3a, 0x32, 0x59,
	0x0c, 0x7b, 0x06, 0xbc, 0xef, 0xd9, 0x47, 0xd5, 0xaa, 0xfc, 0x3d, 0xcb, 0xa7, 0xc8, 0x22, 0xaa,
	0x1d, 0xb4, 0x16, 0x32, 0x59, 0x04, 0x32, 0x54, 0xc3, 0x1d, 0x53, 0x49, 0x4f, 0x26, 0x8b, 0x59,
	0xa8, 0x5c, 0x84, 0xfb, 0x27, 0x5c, 0x5a, 0x2d, 0xa6, 0x5c, 0xaa, 0xec, 0x76, 0xfd, 0xf4, 0x88,
	0x4b, 0xc5, 0x84, 0x32, 0xa2, 0x74, 0xa8, 0xdd, 0x68, 0xb9, 0x72, 0x16, 0x63, 0x20, 0xf9, 0x25,
	0x1a, 0x45, 0xba, 0xb4, 0xaf, 0x0d, 0x33, 0x7e, 0x89, 0xee, 0x6f, 0x30, 0xfc, 0x77, 0x9a, 0x5a,
	0xf9, 0x23, 0xe8, 0x99, 0x79, 0x68, 0x1e, 0xde, 0xc7, 0xdb, 0x35, 0xb8, 0xde, 0x35, 0x5a, 0x23,
	0xc9, 0x07, 0x00, 0x29, 0xfe, 0xaa, 0x82, 0x4d, 0x5e, 0x8e, 0xb6, 0xcc, 0xb4, 0xc1, 0xfd, 0x05,
	0x88, 0x7d, 0x59, 0xff, 0xd7, 0x80, 0x90, 0x29, 0xec, 0xa2, 0x76, 0x0c, 0xae, 0x3c, 0xc2, 0x1f,
	0x6d, 0xe7, 0xb5, 0xf1, 0x76, 0xd3, 0x01, 0xae, 0x37, 0xee, 0x8f, 0xf0, 0xf6, 0x95, 0xc4, 0x75,
	0xc9, 0x5f, 0x42, 0xd7, 0x76, 0xdd, 0xde, 0xfe, 0xd7, 0xa9, 0xd8, 0x02, 0xcf, 0x7b, 0xe6, 0x8f,
	0xf7, 0xf4, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xbb, 0xb4, 0x6e, 0xb5, 0x07, 0x00, 0x00,
}
