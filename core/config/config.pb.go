// Code generated by protoc-gen-go.
// source: config.proto
// DO NOT EDIT!

/*
Package config is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	Config
	BlockSigner
	ValueTuple
	ValueSet
*/
package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bc "chain/protocol/bc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	Id                   string         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	IsSigner             bool           `protobuf:"varint,2,opt,name=is_signer,json=isSigner" json:"is_signer,omitempty"`
	IsGenerator          bool           `protobuf:"varint,3,opt,name=is_generator,json=isGenerator" json:"is_generator,omitempty"`
	BlockchainId         *bc.Hash       `protobuf:"bytes,4,opt,name=blockchain_id,json=blockchainId" json:"blockchain_id,omitempty"`
	GeneratorUrl         string         `protobuf:"bytes,5,opt,name=generator_url,json=generatorUrl" json:"generator_url,omitempty"`
	GeneratorAccessToken string         `protobuf:"bytes,6,opt,name=generator_access_token,json=generatorAccessToken" json:"generator_access_token,omitempty"`
	BlockHsmUrl          string         `protobuf:"bytes,7,opt,name=block_hsm_url,json=blockHsmUrl" json:"block_hsm_url,omitempty"`
	BlockHsmAccessToken  string         `protobuf:"bytes,8,opt,name=block_hsm_access_token,json=blockHsmAccessToken" json:"block_hsm_access_token,omitempty"`
	ConfiguredAt         uint64         `protobuf:"varint,9,opt,name=configured_at,json=configuredAt" json:"configured_at,omitempty"`
	BlockPub             []byte         `protobuf:"bytes,10,opt,name=block_pub,json=blockPub,proto3" json:"block_pub,omitempty"`
	Signers              []*BlockSigner `protobuf:"bytes,11,rep,name=signers" json:"signers,omitempty"`
	Quorum               uint32         `protobuf:"varint,12,opt,name=quorum" json:"quorum,omitempty"`
	MaxIssuanceWindowMs  uint64         `protobuf:"varint,13,opt,name=max_issuance_window_ms,json=maxIssuanceWindowMs" json:"max_issuance_window_ms,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Config) GetIsSigner() bool {
	if m != nil {
		return m.IsSigner
	}
	return false
}

func (m *Config) GetIsGenerator() bool {
	if m != nil {
		return m.IsGenerator
	}
	return false
}

func (m *Config) GetBlockchainId() *bc.Hash {
	if m != nil {
		return m.BlockchainId
	}
	return nil
}

func (m *Config) GetGeneratorUrl() string {
	if m != nil {
		return m.GeneratorUrl
	}
	return ""
}

func (m *Config) GetGeneratorAccessToken() string {
	if m != nil {
		return m.GeneratorAccessToken
	}
	return ""
}

func (m *Config) GetBlockHsmUrl() string {
	if m != nil {
		return m.BlockHsmUrl
	}
	return ""
}

func (m *Config) GetBlockHsmAccessToken() string {
	if m != nil {
		return m.BlockHsmAccessToken
	}
	return ""
}

func (m *Config) GetConfiguredAt() uint64 {
	if m != nil {
		return m.ConfiguredAt
	}
	return 0
}

func (m *Config) GetBlockPub() []byte {
	if m != nil {
		return m.BlockPub
	}
	return nil
}

func (m *Config) GetSigners() []*BlockSigner {
	if m != nil {
		return m.Signers
	}
	return nil
}

func (m *Config) GetQuorum() uint32 {
	if m != nil {
		return m.Quorum
	}
	return 0
}

func (m *Config) GetMaxIssuanceWindowMs() uint64 {
	if m != nil {
		return m.MaxIssuanceWindowMs
	}
	return 0
}

type BlockSigner struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	Pubkey      []byte `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Url         string `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
}

func (m *BlockSigner) Reset()                    { *m = BlockSigner{} }
func (m *BlockSigner) String() string            { return proto.CompactTextString(m) }
func (*BlockSigner) ProtoMessage()               {}
func (*BlockSigner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BlockSigner) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *BlockSigner) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *BlockSigner) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ValueTuple struct {
	// values is an ordered list of the values in the tuple
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *ValueTuple) Reset()                    { *m = ValueTuple{} }
func (m *ValueTuple) String() string            { return proto.CompactTextString(m) }
func (*ValueTuple) ProtoMessage()               {}
func (*ValueTuple) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ValueTuple) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type ValueSet struct {
	// tuples is an unordered list of the tuples in the set
	Tuples []*ValueTuple `protobuf:"bytes,1,rep,name=tuples" json:"tuples,omitempty"`
}

func (m *ValueSet) Reset()                    { *m = ValueSet{} }
func (m *ValueSet) String() string            { return proto.CompactTextString(m) }
func (*ValueSet) ProtoMessage()               {}
func (*ValueSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ValueSet) GetTuples() []*ValueTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "config.Config")
	proto.RegisterType((*BlockSigner)(nil), "config.BlockSigner")
	proto.RegisterType((*ValueTuple)(nil), "config.ValueTuple")
	proto.RegisterType((*ValueSet)(nil), "config.ValueSet")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0xdf, 0x6f, 0xd3, 0x3e,
	0x14, 0xc5, 0x95, 0x66, 0xdf, 0x2c, 0xbd, 0x49, 0xbf, 0x42, 0x2e, 0xaa, 0xac, 0xf1, 0x12, 0x32,
	0x1e, 0x22, 0xa4, 0xb5, 0x52, 0x87, 0x78, 0x1f, 0x3c, 0xb0, 0x3d, 0x20, 0x21, 0x6f, 0x80, 0xc4,
	0x8b, 0xe5, 0x38, 0xa6, 0xb5, 0x9a, 0x1f, 0x25, 0x8e, 0xd9, 0xf8, 0xd3, 0x79, 0x43, 0xbe, 0x49,
	0x9b, 0xee, 0x2d, 0xf7, 0x9e, 0x4f, 0x8e, 0xaf, 0x8f, 0x2f, 0xc4, 0xb2, 0xa9, 0x7f, 0xea, 0xcd,
	0x72, 0xdf, 0x36, 0x5d, 0x43, 0x82, 0xbe, 0xba, 0xb8, 0x90, 0x5b, 0xa1, 0xeb, 0x15, 0x36, 0x65,
	0x53, 0xae, 0x72, 0xb9, 0xca, 0x65, 0xcf, 0xa4, 0x7f, 0x7d, 0x08, 0x3e, 0x22, 0x46, 0xfe, 0x87,
	0x89, 0x2e, 0xa8, 0x97, 0x78, 0xd9, 0x94, 0x4d, 0x74, 0x41, 0x5e, 0xc1, 0x54, 0x1b, 0x6e, 0xf4,
	0xa6, 0x56, 0x2d, 0x9d, 0x24, 0x5e, 0x16, 0xb2, 0x50, 0x9b, 0x7b, 0xac, 0xc9, 0x6b, 0x88, 0xb5,
	0xe1, 0x1b, 0x55, 0xab, 0x56, 0x74, 0x4d, 0x4b, 0x7d, 0xd4, 0x23, 0x6d, 0x3e, 0x1d, 0x5a, 0xe4,
	0x0a, 0x66, 0x79, 0xd9, 0xc8, 0x1d, 0x9e, 0xce, 0x75, 0x41, 0xcf, 0x12, 0x2f, 0x8b, 0xd6, 0xe1,
	0x32, 0x97, 0xcb, 0x5b, 0x61, 0xb6, 0x2c, 0x1e, 0xe5, 0xbb, 0x82, 0x5c, 0xc2, 0xec, 0x68, 0xc7,
	0x6d, 0x5b, 0xd2, 0xff, 0x70, 0x92, 0xf8, 0xd8, 0xfc, 0xda, 0x96, 0xe4, 0x1d, 0x2c, 0x46, 0x48,
	0x48, 0xa9, 0x8c, 0xe1, 0x5d, 0xb3, 0x53, 0x35, 0x0d, 0x90, 0x7e, 0x79, 0x54, 0x6f, 0x50, 0x7c,
	0x70, 0x1a, 0x49, 0x87, 0x49, 0xf8, 0xd6, 0x54, 0x68, 0x7d, 0x8e, 0x70, 0x84, 0xcd, 0x5b, 0x53,
	0x39, 0xe7, 0x6b, 0x58, 0x8c, 0xcc, 0x33, 0xe7, 0x10, 0xe1, 0xf9, 0x01, 0x3e, 0x35, 0xbe, 0x84,
	0x59, 0x9f, 0xb1, 0x6d, 0x55, 0xc1, 0x45, 0x47, 0xa7, 0x89, 0x97, 0x9d, 0xb1, 0x78, 0x6c, 0xde,
	0x74, 0x2e, 0xc7, 0xde, 0x79, 0x6f, 0x73, 0x0a, 0x89, 0x97, 0xc5, 0x2c, 0xc4, 0xc6, 0x17, 0x9b,
	0x93, 0x2b, 0x38, 0xef, 0x13, 0x36, 0x34, 0x4a, 0xfc, 0x2c, 0x5a, 0xcf, 0x97, 0xc3, 0x1b, 0x7e,
	0x70, 0x48, 0x9f, 0x36, 0x3b, 0x30, 0x64, 0x01, 0xc1, 0x2f, 0xdb, 0xb4, 0xb6, 0xa2, 0x71, 0xe2,
	0x65, 0x33, 0x36, 0x54, 0x6e, 0xfa, 0x4a, 0x3c, 0x71, 0x6d, 0x8c, 0x15, 0xb5, 0x54, 0xfc, 0x51,
	0xd7, 0x45, 0xf3, 0xc8, 0x2b, 0x43, 0x67, 0x38, 0xd1, 0xbc, 0x12, 0x4f, 0x77, 0x83, 0xf8, 0x1d,
	0xb5, 0xcf, 0x26, 0xfd, 0x01, 0xd1, 0xc9, 0x21, 0xee, 0x49, 0x9f, 0xdd, 0xbb, 0xdf, 0x84, 0x48,
	0x9c, 0xdc, 0x77, 0x01, 0xc1, 0xde, 0xe6, 0x3b, 0xf5, 0x07, 0xf7, 0x21, 0x66, 0x43, 0x45, 0x5e,
	0x80, 0xef, 0x62, 0xf5, 0xf1, 0x0f, 0xf7, 0x99, 0xbe, 0x01, 0xf8, 0x26, 0x4a, 0xab, 0x1e, 0xec,
	0xbe, 0x54, 0xee, 0xbf, 0xdf, 0xae, 0x32, 0xd4, 0x4b, 0xfc, 0x6c, 0xca, 0x86, 0x2a, 0x7d, 0x0f,
	0x21, 0x52, 0xf7, 0xaa, 0x23, 0x6f, 0x21, 0xe8, 0x1c, 0xdc, 0x33, 0xd1, 0x9a, 0x1c, 0x82, 0x18,
	0x7d, 0xd8, 0x40, 0xe4, 0x01, 0x2e, 0xef, 0xf5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x12,
	0x09, 0x4c, 0xf0, 0x02, 0x00, 0x00,
}
