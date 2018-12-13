// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/feed_attribute_type.proto

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

// Possible data types for a feed attribute.
type FeedAttributeTypeEnum_FeedAttributeType int32

const (
	// Not specified.
	FeedAttributeTypeEnum_UNSPECIFIED FeedAttributeTypeEnum_FeedAttributeType = 0
	// Used for return value only. Represents value unknown in this version.
	FeedAttributeTypeEnum_UNKNOWN FeedAttributeTypeEnum_FeedAttributeType = 1
	// Int64.
	FeedAttributeTypeEnum_INT64 FeedAttributeTypeEnum_FeedAttributeType = 2
	// Double.
	FeedAttributeTypeEnum_DOUBLE FeedAttributeTypeEnum_FeedAttributeType = 3
	// String.
	FeedAttributeTypeEnum_STRING FeedAttributeTypeEnum_FeedAttributeType = 4
	// Boolean.
	FeedAttributeTypeEnum_BOOLEAN FeedAttributeTypeEnum_FeedAttributeType = 5
	// Url.
	FeedAttributeTypeEnum_URL FeedAttributeTypeEnum_FeedAttributeType = 6
	// Datetime.
	FeedAttributeTypeEnum_DATE_TIME FeedAttributeTypeEnum_FeedAttributeType = 7
	// Int64 list.
	FeedAttributeTypeEnum_INT64_LIST FeedAttributeTypeEnum_FeedAttributeType = 8
	// Double (8 bytes) list.
	FeedAttributeTypeEnum_DOUBLE_LIST FeedAttributeTypeEnum_FeedAttributeType = 9
	// String list.
	FeedAttributeTypeEnum_STRING_LIST FeedAttributeTypeEnum_FeedAttributeType = 10
	// Boolean list.
	FeedAttributeTypeEnum_BOOLEAN_LIST FeedAttributeTypeEnum_FeedAttributeType = 11
	// Url list.
	FeedAttributeTypeEnum_URL_LIST FeedAttributeTypeEnum_FeedAttributeType = 12
	// Datetime list.
	FeedAttributeTypeEnum_DATE_TIME_LIST FeedAttributeTypeEnum_FeedAttributeType = 13
	// Price.
	FeedAttributeTypeEnum_PRICE FeedAttributeTypeEnum_FeedAttributeType = 14
)

var FeedAttributeTypeEnum_FeedAttributeType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "INT64",
	3:  "DOUBLE",
	4:  "STRING",
	5:  "BOOLEAN",
	6:  "URL",
	7:  "DATE_TIME",
	8:  "INT64_LIST",
	9:  "DOUBLE_LIST",
	10: "STRING_LIST",
	11: "BOOLEAN_LIST",
	12: "URL_LIST",
	13: "DATE_TIME_LIST",
	14: "PRICE",
}

var FeedAttributeTypeEnum_FeedAttributeType_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"INT64":          2,
	"DOUBLE":         3,
	"STRING":         4,
	"BOOLEAN":        5,
	"URL":            6,
	"DATE_TIME":      7,
	"INT64_LIST":     8,
	"DOUBLE_LIST":    9,
	"STRING_LIST":    10,
	"BOOLEAN_LIST":   11,
	"URL_LIST":       12,
	"DATE_TIME_LIST": 13,
	"PRICE":          14,
}

func (x FeedAttributeTypeEnum_FeedAttributeType) String() string {
	return proto.EnumName(FeedAttributeTypeEnum_FeedAttributeType_name, int32(x))
}

func (FeedAttributeTypeEnum_FeedAttributeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d63f2bf2722bc1cb, []int{0, 0}
}

// Container for enum describing possible data types for a feed attribute.
type FeedAttributeTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedAttributeTypeEnum) Reset()         { *m = FeedAttributeTypeEnum{} }
func (m *FeedAttributeTypeEnum) String() string { return proto.CompactTextString(m) }
func (*FeedAttributeTypeEnum) ProtoMessage()    {}
func (*FeedAttributeTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_d63f2bf2722bc1cb, []int{0}
}

func (m *FeedAttributeTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedAttributeTypeEnum.Unmarshal(m, b)
}
func (m *FeedAttributeTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedAttributeTypeEnum.Marshal(b, m, deterministic)
}
func (m *FeedAttributeTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedAttributeTypeEnum.Merge(m, src)
}
func (m *FeedAttributeTypeEnum) XXX_Size() int {
	return xxx_messageInfo_FeedAttributeTypeEnum.Size(m)
}
func (m *FeedAttributeTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedAttributeTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedAttributeTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.FeedAttributeTypeEnum_FeedAttributeType", FeedAttributeTypeEnum_FeedAttributeType_name, FeedAttributeTypeEnum_FeedAttributeType_value)
	proto.RegisterType((*FeedAttributeTypeEnum)(nil), "google.ads.googleads.v0.enums.FeedAttributeTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/feed_attribute_type.proto", fileDescriptor_d63f2bf2722bc1cb)
}

var fileDescriptor_d63f2bf2722bc1cb = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcd, 0x6a, 0xab, 0x40,
	0x18, 0xbd, 0x9a, 0x9b, 0xbf, 0x2f, 0x3f, 0x77, 0xee, 0xc0, 0xbd, 0xbb, 0x2c, 0x9a, 0x07, 0x18,
	0x85, 0x96, 0x76, 0xd1, 0x95, 0x26, 0x93, 0x20, 0xb5, 0x2a, 0x46, 0x53, 0x28, 0x82, 0x98, 0x3a,
	0x95, 0x40, 0xa2, 0x12, 0x4d, 0x20, 0xfb, 0x3e, 0x49, 0x97, 0x7d, 0x91, 0x42, 0x1f, 0xa5, 0x4f,
	0x51, 0xc6, 0x49, 0xb2, 0x09, 0xed, 0x46, 0x8e, 0xe7, 0x7c, 0xe7, 0x7c, 0xc3, 0x77, 0xe0, 0x26,
	0xc9, 0xb2, 0x64, 0xc5, 0x94, 0x28, 0x2e, 0x14, 0x01, 0x39, 0xda, 0xa9, 0x0a, 0x4b, 0xb7, 0xeb,
	0x42, 0x79, 0x66, 0x2c, 0x0e, 0xa3, 0xb2, 0xdc, 0x2c, 0x17, 0xdb, 0x92, 0x85, 0xe5, 0x3e, 0x67,
	0x24, 0xdf, 0x64, 0x65, 0x86, 0x07, 0x62, 0x9a, 0x44, 0x71, 0x41, 0x4e, 0x46, 0xb2, 0x53, 0x49,
	0x65, 0x1c, 0xbe, 0xc8, 0xf0, 0x6f, 0xc2, 0x58, 0xac, 0x1d, 0xbd, 0xde, 0x3e, 0x67, 0x34, 0xdd,
	0xae, 0x87, 0x9f, 0x12, 0xfc, 0x3d, 0x53, 0xf0, 0x1f, 0xe8, 0xf8, 0xd6, 0xcc, 0xa1, 0x23, 0x63,
	0x62, 0xd0, 0x31, 0xfa, 0x85, 0x3b, 0xd0, 0xf4, 0xad, 0x3b, 0xcb, 0x7e, 0xb0, 0x90, 0x84, 0xdb,
	0x50, 0x37, 0x2c, 0xef, 0xfa, 0x0a, 0xc9, 0x18, 0xa0, 0x31, 0xb6, 0x7d, 0xdd, 0xa4, 0xa8, 0xc6,
	0xf1, 0xcc, 0x73, 0x0d, 0x6b, 0x8a, 0x7e, 0xf3, 0x79, 0xdd, 0xb6, 0x4d, 0xaa, 0x59, 0xa8, 0x8e,
	0x9b, 0x50, 0xf3, 0x5d, 0x13, 0x35, 0x70, 0x0f, 0xda, 0x63, 0xcd, 0xa3, 0xa1, 0x67, 0xdc, 0x53,
	0xd4, 0xc4, 0x7d, 0x80, 0x2a, 0x27, 0x34, 0x8d, 0x99, 0x87, 0x5a, 0x7c, 0xab, 0x08, 0x13, 0x44,
	0x9b, 0x13, 0x22, 0x51, 0x10, 0x80, 0x11, 0x74, 0x0f, 0xb1, 0x82, 0xe9, 0xe0, 0x2e, 0xb4, 0x7c,
	0xd7, 0x14, 0x7f, 0x5d, 0x8c, 0xa1, 0x7f, 0x5a, 0x20, 0xb8, 0x1e, 0x7f, 0xad, 0xe3, 0x1a, 0x23,
	0x8a, 0xfa, 0xfa, 0xbb, 0x04, 0x17, 0x4f, 0xd9, 0x9a, 0xfc, 0x78, 0x2c, 0xfd, 0xff, 0xd9, 0x3d,
	0x1c, 0x7e, 0x63, 0x47, 0x7a, 0xd4, 0x0f, 0xc6, 0x24, 0x5b, 0x45, 0x69, 0x42, 0xb2, 0x4d, 0xa2,
	0x24, 0x2c, 0xad, 0x1a, 0x38, 0xd6, 0x95, 0x2f, 0x8b, 0x6f, 0xda, 0xbb, 0xad, 0xbe, 0xaf, 0x72,
	0x6d, 0xaa, 0x69, 0x6f, 0xf2, 0x60, 0x2a, 0xa2, 0xb4, 0xb8, 0x20, 0x02, 0x72, 0x34, 0x57, 0x09,
	0x6f, 0xa5, 0xf8, 0x38, 0xea, 0x81, 0x16, 0x17, 0xc1, 0x49, 0x0f, 0xe6, 0x6a, 0x50, 0xe9, 0x8b,
	0x46, 0xb5, 0xf4, 0xf2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xb4, 0xcc, 0xb0, 0x31, 0x02, 0x00,
	0x00,
}
