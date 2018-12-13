// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/bid_modifier_source.proto

package enums

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enum describing possible bid modifier sources.
type BidModifierSourceEnum_BidModifierSource int32

const (
	// Not specified.
	BidModifierSourceEnum_UNSPECIFIED BidModifierSourceEnum_BidModifierSource = 0
	// Used for return value only. Represents value unknown in this version.
	BidModifierSourceEnum_UNKNOWN BidModifierSourceEnum_BidModifierSource = 1
	// The bid modifier is specified at the campaign level, on the campaign
	// level criterion.
	BidModifierSourceEnum_CAMPAIGN BidModifierSourceEnum_BidModifierSource = 2
	// The bid modifier is specified (overridden) at the ad group level.
	BidModifierSourceEnum_AD_GROUP BidModifierSourceEnum_BidModifierSource = 3
)

var BidModifierSourceEnum_BidModifierSource_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CAMPAIGN",
	3: "AD_GROUP",
}

var BidModifierSourceEnum_BidModifierSource_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"CAMPAIGN":    2,
	"AD_GROUP":    3,
}

func (x BidModifierSourceEnum_BidModifierSource) String() string {
	return proto.EnumName(BidModifierSourceEnum_BidModifierSource_name, int32(x))
}

func (BidModifierSourceEnum_BidModifierSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_add6bce8ed25a001, []int{0, 0}
}

// Container for enum describing possible bid modifier sources.
type BidModifierSourceEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BidModifierSourceEnum) Reset()         { *m = BidModifierSourceEnum{} }
func (m *BidModifierSourceEnum) String() string { return proto.CompactTextString(m) }
func (*BidModifierSourceEnum) ProtoMessage()    {}
func (*BidModifierSourceEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_add6bce8ed25a001, []int{0}
}

func (m *BidModifierSourceEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidModifierSourceEnum.Unmarshal(m, b)
}
func (m *BidModifierSourceEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidModifierSourceEnum.Marshal(b, m, deterministic)
}
func (m *BidModifierSourceEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidModifierSourceEnum.Merge(m, src)
}
func (m *BidModifierSourceEnum) XXX_Size() int {
	return xxx_messageInfo_BidModifierSourceEnum.Size(m)
}
func (m *BidModifierSourceEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BidModifierSourceEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BidModifierSourceEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.BidModifierSourceEnum_BidModifierSource", BidModifierSourceEnum_BidModifierSource_name, BidModifierSourceEnum_BidModifierSource_value)
	proto.RegisterType((*BidModifierSourceEnum)(nil), "google.ads.googleads.v0.enums.BidModifierSourceEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/bid_modifier_source.proto", fileDescriptor_add6bce8ed25a001)
}

var fileDescriptor_add6bce8ed25a001 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4f, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xa4, 0xcc, 0x94, 0xf8, 0xdc, 0xfc, 0x94, 0xcc, 0xb4, 0xcc, 0xd4, 0xa2,
	0xf8, 0xe2, 0xfc, 0xd2, 0xa2, 0xe4, 0x54, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88,
	0x6a, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x46, 0xbd, 0x32, 0x03, 0x3d, 0xb0, 0x46, 0xa5, 0x34,
	0x2e, 0x51, 0xa7, 0xcc, 0x14, 0x5f, 0xa8, 0xd6, 0x60, 0xb0, 0x4e, 0xd7, 0xbc, 0xd2, 0x5c, 0x25,
	0x5f, 0x2e, 0x41, 0x0c, 0x09, 0x21, 0x7e, 0x2e, 0xee, 0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f,
	0x37, 0x4f, 0x57, 0x17, 0x01, 0x06, 0x21, 0x6e, 0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70,
	0x3f, 0x01, 0x46, 0x21, 0x1e, 0x2e, 0x0e, 0x67, 0x47, 0xdf, 0x00, 0x47, 0x4f, 0x77, 0x3f, 0x01,
	0x26, 0x10, 0xcf, 0xd1, 0x25, 0xde, 0x3d, 0xc8, 0x3f, 0x34, 0x40, 0x80, 0xd9, 0xe9, 0x38, 0x23,
	0x97, 0x62, 0x72, 0x7e, 0xae, 0x1e, 0x5e, 0xd7, 0x38, 0x89, 0x61, 0x58, 0x19, 0x00, 0xf2, 0x44,
	0x00, 0x63, 0x94, 0x13, 0x54, 0x63, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51, 0xba,
	0x7e, 0x7a, 0x6a, 0x1e, 0xd8, 0x8b, 0xb0, 0xf0, 0x28, 0xc8, 0x2c, 0xc6, 0x11, 0x3c, 0xd6, 0x60,
	0x72, 0x11, 0x13, 0xb3, 0xbb, 0xa3, 0xe3, 0x2a, 0x26, 0x59, 0x77, 0x88, 0x51, 0x8e, 0x29, 0xc5,
	0x7a, 0x10, 0x26, 0x88, 0x15, 0x66, 0xa0, 0x07, 0xf2, 0x77, 0xf1, 0x29, 0x98, 0x7c, 0x8c, 0x63,
	0x4a, 0x71, 0x0c, 0x5c, 0x3e, 0x26, 0xcc, 0x20, 0x06, 0x2c, 0x9f, 0xc4, 0x06, 0xb6, 0xd4, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xa8, 0xae, 0xc8, 0x92, 0x01, 0x00, 0x00,
}
