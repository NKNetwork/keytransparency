// Code generated by protoc-gen-go.
// source: proto/security_e2ekeys/e2ekeytypes.proto
// DO NOT EDIT!

/*
Package security_e2ekeys is a generated protocol buffer package.

It is generated from these files:
	proto/security_e2ekeys/e2ekeytypes.proto

It has these top-level messages:
	GetEntryResponse
	Profile
	Entry
	PublicKey
	KeyValue
	SignedKV
	GetEntryRequest
	ListEntryHistoryRequest
	ListEntryHistoryResponse
	EntryUpdate
	UpdateEntryRequest
	UpdateEntryResponse
*/
package security_e2ekeys

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import security_ctmap "github.com/google/e2e-key-server/proto/security_ctmap"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// GetEntryResponse
type GetEntryResponse struct {
	// vrf is the output of VRF on user_id.
	Vrf []byte `protobuf:"bytes,1,opt,name=vrf,proto3" json:"vrf,omitempty"`
	// vrf_proof is the proof for VRF on user_id.
	VrfProof []byte `protobuf:"bytes,2,opt,name=vrf_proof,proto3" json:"vrf_proof,omitempty"`
	// TODO: Combine into Commitment datatype.
	// commitment_key connects the profile data to the commitment in leaf_proof.
	CommitmentKey []byte `protobuf:"bytes,3,opt,name=commitment_key,proto3" json:"commitment_key,omitempty"`
	// profile contains the public key data for this account.
	Profile []byte `protobuf:"bytes,4,opt,name=profile,proto3" json:"profile,omitempty"`
	// leaf_proof contains an Entry and an inclusion proof in the sparse merkle tree at end_epoch.
	LeafProof *security_ctmap.GetLeafResponse `protobuf:"bytes,5,opt,name=leaf_proof" json:"leaf_proof,omitempty"`
	// seh contains the signed epoch head for the sparse merkle tree.
	// seh is also stored in the append only log.
	Seh *security_ctmap.SignedEpochHead `protobuf:"bytes,6,opt,name=seh" json:"seh,omitempty"`
	// seh_sct is the SCT showing that seh was submitted to CT logs.
	// TODO: Support storing seh in multiple logs.
	SehSct []byte `protobuf:"bytes,7,opt,name=seh_sct,proto3" json:"seh_sct,omitempty"`
}

func (m *GetEntryResponse) Reset()                    { *m = GetEntryResponse{} }
func (m *GetEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*GetEntryResponse) ProtoMessage()               {}
func (*GetEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetEntryResponse) GetLeafProof() *security_ctmap.GetLeafResponse {
	if m != nil {
		return m.LeafProof
	}
	return nil
}

func (m *GetEntryResponse) GetSeh() *security_ctmap.SignedEpochHead {
	if m != nil {
		return m.Seh
	}
	return nil
}

// Profile contains data hidden behind the crypto comitment.
type Profile struct {
	// Keys is a map of appIds to keys.
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

// Entry contains a commitment to profile and a set of authorized update keys.
// Entry is placed in the verifiable map as leaf data.
type Entry struct {
	// commitment is a cryptographic commitment to arbitrary data.
	Commitment []byte `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	// authorized_keys is the set of keys allowed to sign updates for this entry.
	AuthorizedKeys []*PublicKey `protobuf:"bytes,2,rep,name=authorized_keys" json:"authorized_keys,omitempty"`
	// update_count prevents replay attacks. Monotonically increasing.
	UpdateCount uint64 `protobuf:"varint,3,opt,name=update_count" json:"update_count,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Entry) GetAuthorizedKeys() []*PublicKey {
	if m != nil {
		return m.AuthorizedKeys
	}
	return nil
}

// PublicKey defines a key this domain uses to sign EpochHeads with.
type PublicKey struct {
	// KeyFormats from Keyczar.
	//
	// Types that are valid to be assigned to KeyType:
	//	*PublicKey_Ed25519
	//	*PublicKey_RsaVerifyingSha256_2048
	//	*PublicKey_EcdsaVerifyingP256
	KeyType isPublicKey_KeyType `protobuf_oneof:"key_type"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isPublicKey_KeyType interface {
	isPublicKey_KeyType()
}

type PublicKey_Ed25519 struct {
	Ed25519 []byte `protobuf:"bytes,1,opt,name=ed25519,proto3,oneof"`
}
type PublicKey_RsaVerifyingSha256_2048 struct {
	RsaVerifyingSha256_2048 []byte `protobuf:"bytes,2,opt,name=rsa_verifying_sha256_2048,proto3,oneof"`
}
type PublicKey_EcdsaVerifyingP256 struct {
	EcdsaVerifyingP256 []byte `protobuf:"bytes,3,opt,name=ecdsa_verifying_p256,proto3,oneof"`
}

func (*PublicKey_Ed25519) isPublicKey_KeyType()                 {}
func (*PublicKey_RsaVerifyingSha256_2048) isPublicKey_KeyType() {}
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

func (m *PublicKey) GetRsaVerifyingSha256_2048() []byte {
	if x, ok := m.GetKeyType().(*PublicKey_RsaVerifyingSha256_2048); ok {
		return x.RsaVerifyingSha256_2048
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
		(*PublicKey_RsaVerifyingSha256_2048)(nil),
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
	case *PublicKey_RsaVerifyingSha256_2048:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.RsaVerifyingSha256_2048)
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
	case 2: // key_type.rsa_verifying_sha256_2048
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.KeyType = &PublicKey_RsaVerifyingSha256_2048{x}
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
	case *PublicKey_RsaVerifyingSha256_2048:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RsaVerifyingSha256_2048)))
		n += len(x.RsaVerifyingSha256_2048)
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
	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// SignedKV is a signed change to a map entry.
type SignedKV struct {
	// keyvalue is a serialized KeyValue.
	KeyValue []byte `protobuf:"bytes,1,opt,name=key_value,proto3" json:"key_value,omitempty"`
	// signatures on keyvalue. Must be signed by keys from both previous and
	// current epochs. The first proves ownership of new epoch key, and the
	// second proves the the correct owner is making this change.
	Signatures map[uint64][]byte `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"fixed64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *SignedKV) Reset()                    { *m = SignedKV{} }
func (m *SignedKV) String() string            { return proto.CompactTextString(m) }
func (*SignedKV) ProtoMessage()               {}
func (*SignedKV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SignedKV) GetSignatures() map[uint64][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// Get request for a user object.
type GetEntryRequest struct {
	// Last trusted epoch by the client.
	// int64 epoch_start = 3;
	// Absence of the epoch_end field indicates a request for the current value.
	EpochEnd int64 `protobuf:"varint,1,opt,name=epoch_end" json:"epoch_end,omitempty"`
	// User identifier. Most commonly an email address.
	UserId string `protobuf:"bytes,2,opt,name=user_id" json:"user_id,omitempty"`
}

func (m *GetEntryRequest) Reset()                    { *m = GetEntryRequest{} }
func (m *GetEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetEntryRequest) ProtoMessage()               {}
func (*GetEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

// Get a list of historical values for a user.
type ListEntryHistoryRequest struct {
	// The user identifier.
	UserId string `protobuf:"bytes,1,opt,name=user_id" json:"user_id,omitempty"`
	// from_epoch is the starting epcoh.
	StartEpoch int64 `protobuf:"varint,2,opt,name=start_epoch" json:"start_epoch,omitempty"`
	// The maximum number of entries to return.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size" json:"page_size,omitempty"`
}

func (m *ListEntryHistoryRequest) Reset()                    { *m = ListEntryHistoryRequest{} }
func (m *ListEntryHistoryRequest) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryRequest) ProtoMessage()               {}
func (*ListEntryHistoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// A paginated history of values for a user.
type ListEntryHistoryResponse struct {
	// The list of values this user_id has contained over time.
	Values []*GetEntryResponse `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	// The next time to query for pagination.
	NextEpoch int64 `protobuf:"varint,2,opt,name=next_epoch" json:"next_epoch,omitempty"`
}

func (m *ListEntryHistoryResponse) Reset()                    { *m = ListEntryHistoryResponse{} }
func (m *ListEntryHistoryResponse) String() string            { return proto.CompactTextString(m) }
func (*ListEntryHistoryResponse) ProtoMessage()               {}
func (*ListEntryHistoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ListEntryHistoryResponse) GetValues() []*GetEntryResponse {
	if m != nil {
		return m.Values
	}
	return nil
}

type EntryUpdate struct {
	// update authorizes the change to profile.
	Update *SignedKV `protobuf:"bytes,2,opt,name=update" json:"update,omitempty"`
	// profile is the serialized protobuf Profile.
	Profile []byte `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	// commitment_key is 16 random bytes.
	CommitmentKey []byte `protobuf:"bytes,4,opt,name=commitment_key,proto3" json:"commitment_key,omitempty"`
}

func (m *EntryUpdate) Reset()                    { *m = EntryUpdate{} }
func (m *EntryUpdate) String() string            { return proto.CompactTextString(m) }
func (*EntryUpdate) ProtoMessage()               {}
func (*EntryUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *EntryUpdate) GetUpdate() *SignedKV {
	if m != nil {
		return m.Update
	}
	return nil
}

// Update a user's profile.
type UpdateEntryRequest struct {
	// user_id specifies the id for the new account to be registered.
	UserId      string       `protobuf:"bytes,1,opt,name=user_id" json:"user_id,omitempty"`
	EntryUpdate *EntryUpdate `protobuf:"bytes,2,opt,name=entry_update" json:"entry_update,omitempty"`
}

func (m *UpdateEntryRequest) Reset()                    { *m = UpdateEntryRequest{} }
func (m *UpdateEntryRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryRequest) ProtoMessage()               {}
func (*UpdateEntryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UpdateEntryRequest) GetEntryUpdate() *EntryUpdate {
	if m != nil {
		return m.EntryUpdate
	}
	return nil
}

// UpdateEntryResponse contains a proof once the update has been included in
// the Merkel Tree.
type UpdateEntryResponse struct {
	Proof *GetEntryResponse `protobuf:"bytes,1,opt,name=proof" json:"proof,omitempty"`
}

func (m *UpdateEntryResponse) Reset()                    { *m = UpdateEntryResponse{} }
func (m *UpdateEntryResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateEntryResponse) ProtoMessage()               {}
func (*UpdateEntryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UpdateEntryResponse) GetProof() *GetEntryResponse {
	if m != nil {
		return m.Proof
	}
	return nil
}

func init() {
	proto.RegisterType((*GetEntryResponse)(nil), "security.e2ekeys.GetEntryResponse")
	proto.RegisterType((*Profile)(nil), "security.e2ekeys.Profile")
	proto.RegisterType((*Entry)(nil), "security.e2ekeys.Entry")
	proto.RegisterType((*PublicKey)(nil), "security.e2ekeys.PublicKey")
	proto.RegisterType((*KeyValue)(nil), "security.e2ekeys.KeyValue")
	proto.RegisterType((*SignedKV)(nil), "security.e2ekeys.SignedKV")
	proto.RegisterType((*GetEntryRequest)(nil), "security.e2ekeys.GetEntryRequest")
	proto.RegisterType((*ListEntryHistoryRequest)(nil), "security.e2ekeys.ListEntryHistoryRequest")
	proto.RegisterType((*ListEntryHistoryResponse)(nil), "security.e2ekeys.ListEntryHistoryResponse")
	proto.RegisterType((*EntryUpdate)(nil), "security.e2ekeys.EntryUpdate")
	proto.RegisterType((*UpdateEntryRequest)(nil), "security.e2ekeys.UpdateEntryRequest")
	proto.RegisterType((*UpdateEntryResponse)(nil), "security.e2ekeys.UpdateEntryResponse")
}

var fileDescriptor0 = []byte{
	// 654 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x6d, 0xea, 0x26, 0x69, 0x6e, 0xf2, 0x7d, 0x69, 0xdd, 0x0a, 0x4c, 0x11, 0x3f, 0x72, 0x25,
	0x54, 0x15, 0xea, 0x10, 0xb7, 0x41, 0x85, 0x05, 0x0b, 0xa4, 0x8a, 0x88, 0x76, 0x81, 0x8a, 0xe8,
	0x82, 0xcd, 0xc8, 0xb1, 0x6f, 0x1c, 0x8b, 0xc4, 0x63, 0x3c, 0xe3, 0x8a, 0xf0, 0x14, 0xbc, 0x19,
	0xaf, 0xc4, 0x9d, 0x99, 0xa4, 0xf9, 0xad, 0xc4, 0x2a, 0x99, 0x99, 0x33, 0xe7, 0x9c, 0x7b, 0xee,
	0x1d, 0xc3, 0x51, 0x96, 0x73, 0xc9, 0x5b, 0x02, 0xc3, 0x22, 0x4f, 0xe4, 0x98, 0xa1, 0x8f, 0xdf,
	0x71, 0x2c, 0x5a, 0xe6, 0x57, 0x8e, 0x33, 0x14, 0x9e, 0x86, 0xd8, 0x3b, 0x53, 0x8c, 0x37, 0xc1,
	0x1c, 0x7c, 0x8a, 0x13, 0x39, 0x28, 0x7a, 0x5e, 0xc8, 0x47, 0xad, 0x98, 0xf3, 0x78, 0x88, 0xea,
	0xda, 0x09, 0x9d, 0x9d, 0x08, 0xcc, 0x6f, 0x31, 0x6f, 0x2d, 0x91, 0x87, 0x72, 0x14, 0x64, 0x4b,
	0x4b, 0xc3, 0xee, 0xfe, 0x29, 0xc1, 0xce, 0x47, 0x94, 0x17, 0xa9, 0xcc, 0xc7, 0xd7, 0x28, 0x32,
	0x9e, 0x0a, 0xb4, 0xeb, 0x60, 0xdd, 0xe6, 0x7d, 0xa7, 0xf4, 0xbc, 0x74, 0xd4, 0xb0, 0x77, 0xa1,
	0x46, 0x0b, 0x46, 0x70, 0xde, 0x77, 0x36, 0xf5, 0xd6, 0x03, 0xf8, 0x9f, 0xb4, 0x47, 0x89, 0x1c,
	0x61, 0x2a, 0x19, 0xe9, 0x3a, 0x96, 0xde, 0x6f, 0x42, 0x95, 0x60, 0xfd, 0x64, 0x88, 0xce, 0x96,
	0xde, 0x38, 0x05, 0x18, 0x62, 0x30, 0xbd, 0x5c, 0xa6, 0xbd, 0xba, 0xff, 0xcc, 0x5b, 0x32, 0x42,
	0xf2, 0x57, 0x04, 0xba, 0x53, 0x7f, 0x05, 0x96, 0xc0, 0x81, 0x53, 0x59, 0x8f, 0xfe, 0x92, 0xc4,
	0x29, 0x46, 0x17, 0x19, 0x0f, 0x07, 0x5d, 0x0c, 0x22, 0xa5, 0x49, 0x68, 0x26, 0x42, 0xe9, 0x54,
	0x95, 0xa6, 0x9b, 0x40, 0xf5, 0xb3, 0x31, 0x61, 0xb7, 0x61, 0x4b, 0x05, 0x46, 0x85, 0x58, 0x44,
	0x75, 0xe8, 0x2d, 0x27, 0xe9, 0x4d, 0x80, 0xde, 0x25, 0x2d, 0x74, 0x04, 0x07, 0x2f, 0xa1, 0x76,
	0xb7, 0x50, 0x39, 0xa8, 0xe2, 0x54, 0x0e, 0x35, 0xfb, 0x3f, 0x28, 0xdf, 0x06, 0xc3, 0x02, 0x4d,
	0x06, 0xef, 0x36, 0xcf, 0x4b, 0x6e, 0x0c, 0x65, 0x03, 0xb4, 0x01, 0x66, 0x81, 0x4c, 0x72, 0x3b,
	0x83, 0x66, 0x50, 0xc8, 0x01, 0xcf, 0x93, 0x5f, 0x18, 0x31, 0xed, 0x63, 0x53, 0xfb, 0x78, 0xbc,
	0xc6, 0x47, 0xd1, 0x1b, 0x26, 0x21, 0x09, 0xdb, 0xfb, 0xd0, 0x28, 0xb2, 0x28, 0x90, 0xc8, 0x42,
	0x5e, 0x10, 0x97, 0x0a, 0x76, 0xcb, 0xe5, 0x50, 0x9b, 0x41, 0x76, 0xa1, 0x8a, 0x91, 0xdf, 0xe9,
	0xb4, 0xdf, 0x1a, 0xa5, 0xee, 0x86, 0x7d, 0x08, 0x8f, 0x72, 0x11, 0x30, 0x6a, 0x7d, 0xd2, 0x1f,
	0x27, 0x69, 0xcc, 0xc4, 0x20, 0xf0, 0x3b, 0x6f, 0x98, 0xff, 0xfa, 0xec, 0xdc, 0xf8, 0x25, 0xd0,
	0x53, 0xd8, 0xc7, 0x30, 0x5a, 0x80, 0x65, 0x04, 0x32, 0xbd, 0xeb, 0x6e, 0x7c, 0x00, 0xd8, 0x26,
	0x33, 0x4c, 0xcd, 0x9e, 0xfb, 0x02, 0xb6, 0x49, 0xea, 0x46, 0xd5, 0x3b, 0x9f, 0x42, 0x63, 0x29,
	0x05, 0xf7, 0x77, 0x09, 0xb6, 0x4d, 0x47, 0x2e, 0x6f, 0xd4, 0xa4, 0x28, 0x02, 0x73, 0x6e, 0xe0,
	0xef, 0x01, 0x04, 0x1d, 0x07, 0xb2, 0xc8, 0x71, 0x5a, 0xff, 0xf1, 0x6a, 0xfd, 0x53, 0x0a, 0xfd,
	0xc7, 0x80, 0x4d, 0x3b, 0xda, 0xd0, 0x5c, 0xda, 0x9a, 0xb7, 0x53, 0x59, 0xd7, 0x94, 0x0e, 0x34,
	0x67, 0x03, 0xfd, 0xa3, 0x40, 0x21, 0x95, 0x31, 0x54, 0x03, 0xc3, 0x30, 0x8d, 0xf4, 0x45, 0x4b,
	0x8d, 0x4d, 0x41, 0x8f, 0x85, 0x25, 0x91, 0xbe, 0x5a, 0x73, 0xaf, 0xe1, 0xe1, 0x55, 0x22, 0xcc,
	0xbd, 0x2e, 0xfd, 0xe1, 0xb3, 0xeb, 0x73, 0x58, 0x33, 0x0a, 0x7b, 0x50, 0x17, 0x32, 0xc8, 0x25,
	0xd3, 0xac, 0x9a, 0xc0, 0x52, 0x22, 0x59, 0x10, 0x23, 0x13, 0xd4, 0x6f, 0x9d, 0x69, 0xd9, 0xed,
	0x81, 0xb3, 0xca, 0x39, 0x99, 0x72, 0x1f, 0x2a, 0xda, 0xf9, 0x74, 0x3a, 0xdd, 0xd5, 0x54, 0x56,
	0xde, 0x25, 0x8d, 0x59, 0x8a, 0x3f, 0x17, 0x64, 0x49, 0xa3, 0xae, 0x41, 0x5f, 0xf5, 0xd4, 0xd8,
	0xc7, 0x50, 0x31, 0xf3, 0xa3, 0x8f, 0xeb, 0xfe, 0xc1, 0xfd, 0x61, 0xcf, 0x3f, 0x57, 0xeb, 0x9e,
	0x77, 0xad, 0x9f, 0xb1, 0xfb, 0x0d, 0x6c, 0x43, 0xbf, 0x90, 0xea, 0x4a, 0x2c, 0xa7, 0xd0, 0x40,
	0x05, 0x60, 0x0b, 0x0e, 0x9e, 0xac, 0x3a, 0x98, 0x33, 0xec, 0x76, 0x61, 0x6f, 0x81, 0x7b, 0x52,
	0x6a, 0x1b, 0xca, 0xe6, 0xa3, 0x51, 0xd2, 0x24, 0xff, 0x90, 0x4e, 0xaf, 0xa2, 0xbf, 0x68, 0xa7,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x07, 0x42, 0x53, 0x10, 0x5b, 0x05, 0x00, 0x00,
}
