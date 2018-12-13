// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/feed_mapping_criterion_type.proto

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

// Possible placeholder types for a feed mapping.
type FeedMappingCriterionTypeEnum_FeedMappingCriterionType int32

const (
	// Not specified.
	FeedMappingCriterionTypeEnum_UNSPECIFIED FeedMappingCriterionTypeEnum_FeedMappingCriterionType = 0
	// Used for return value only. Represents value unknown in this version.
	FeedMappingCriterionTypeEnum_UNKNOWN FeedMappingCriterionTypeEnum_FeedMappingCriterionType = 1
	// Allows campaign targeting at locations within a location feed.
	FeedMappingCriterionTypeEnum_CAMPAIGN_LOCATION_TARGETS FeedMappingCriterionTypeEnum_FeedMappingCriterionType = 2
	// Allows url targeting for your dynamic search ads within a page feed.
	FeedMappingCriterionTypeEnum_DSA_PAGE_FEED FeedMappingCriterionTypeEnum_FeedMappingCriterionType = 3
)

var FeedMappingCriterionTypeEnum_FeedMappingCriterionType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CAMPAIGN_LOCATION_TARGETS",
	3: "DSA_PAGE_FEED",
}

var FeedMappingCriterionTypeEnum_FeedMappingCriterionType_value = map[string]int32{
	"UNSPECIFIED":               0,
	"UNKNOWN":                   1,
	"CAMPAIGN_LOCATION_TARGETS": 2,
	"DSA_PAGE_FEED":             3,
}

func (x FeedMappingCriterionTypeEnum_FeedMappingCriterionType) String() string {
	return proto.EnumName(FeedMappingCriterionTypeEnum_FeedMappingCriterionType_name, int32(x))
}

func (FeedMappingCriterionTypeEnum_FeedMappingCriterionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2bbba35bc711b98, []int{0, 0}
}

// Container for enum describing possible criterion types for a feed mapping.
type FeedMappingCriterionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedMappingCriterionTypeEnum) Reset()         { *m = FeedMappingCriterionTypeEnum{} }
func (m *FeedMappingCriterionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*FeedMappingCriterionTypeEnum) ProtoMessage()    {}
func (*FeedMappingCriterionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2bbba35bc711b98, []int{0}
}

func (m *FeedMappingCriterionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedMappingCriterionTypeEnum.Unmarshal(m, b)
}
func (m *FeedMappingCriterionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedMappingCriterionTypeEnum.Marshal(b, m, deterministic)
}
func (m *FeedMappingCriterionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedMappingCriterionTypeEnum.Merge(m, src)
}
func (m *FeedMappingCriterionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_FeedMappingCriterionTypeEnum.Size(m)
}
func (m *FeedMappingCriterionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedMappingCriterionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedMappingCriterionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.FeedMappingCriterionTypeEnum_FeedMappingCriterionType", FeedMappingCriterionTypeEnum_FeedMappingCriterionType_name, FeedMappingCriterionTypeEnum_FeedMappingCriterionType_value)
	proto.RegisterType((*FeedMappingCriterionTypeEnum)(nil), "google.ads.googleads.v0.enums.FeedMappingCriterionTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/feed_mapping_criterion_type.proto", fileDescriptor_c2bbba35bc711b98)
}

var fileDescriptor_c2bbba35bc711b98 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xf3, 0x30,
	0x18, 0x86, 0xff, 0x6e, 0xf0, 0x0b, 0x19, 0x62, 0xed, 0x91, 0x82, 0x3d, 0x70, 0x17, 0x90, 0x16,
	0x3c, 0xf4, 0x40, 0xb2, 0x2e, 0x2b, 0x45, 0x97, 0x15, 0xd7, 0x4d, 0x90, 0x42, 0xa8, 0x4b, 0x0c,
	0x95, 0x35, 0x09, 0x4d, 0x37, 0xd8, 0x2d, 0x78, 0x19, 0x1e, 0x7a, 0x29, 0x1e, 0x78, 0x4d, 0xd2,
	0xd6, 0xee, 0xac, 0x9e, 0x84, 0x17, 0x5e, 0xbe, 0x27, 0xdf, 0xf7, 0x80, 0x3b, 0xa1, 0x94, 0xd8,
	0x72, 0x2f, 0x63, 0xc6, 0x6b, 0x63, 0x9d, 0xf6, 0xbe, 0xc7, 0xe5, 0xae, 0x30, 0xde, 0x2b, 0xe7,
	0x8c, 0x16, 0x99, 0xd6, 0xb9, 0x14, 0x74, 0x53, 0xe6, 0x15, 0x2f, 0x73, 0x25, 0x69, 0x75, 0xd0,
	0x1c, 0xea, 0x52, 0x55, 0xca, 0x71, 0xdb, 0x29, 0x98, 0x31, 0x03, 0x8f, 0x00, 0xb8, 0xf7, 0x61,
	0x03, 0x18, 0xbf, 0x5b, 0xe0, 0x6a, 0xc6, 0x39, 0x9b, 0xb7, 0x8c, 0xa0, 0x43, 0x24, 0x07, 0xcd,
	0xb1, 0xdc, 0x15, 0xe3, 0x37, 0x70, 0xd1, 0xd7, 0x3b, 0x67, 0x60, 0xb4, 0x22, 0xcb, 0x18, 0x07,
	0xd1, 0x2c, 0xc2, 0x53, 0xfb, 0x9f, 0x33, 0x02, 0x27, 0x2b, 0x72, 0x4f, 0x16, 0x4f, 0xc4, 0xb6,
	0x1c, 0x17, 0x5c, 0x06, 0x68, 0x1e, 0xa3, 0x28, 0x24, 0xf4, 0x61, 0x11, 0xa0, 0x24, 0x5a, 0x10,
	0x9a, 0xa0, 0xc7, 0x10, 0x27, 0x4b, 0x7b, 0xe0, 0x9c, 0x83, 0xd3, 0xe9, 0x12, 0xd1, 0x18, 0x85,
	0x98, 0xce, 0x30, 0x9e, 0xda, 0xc3, 0xc9, 0xb7, 0x05, 0xae, 0x37, 0xaa, 0x80, 0x7f, 0xae, 0x3c,
	0x71, 0xfb, 0xf6, 0x89, 0xeb, 0x83, 0x63, 0xeb, 0x79, 0xf2, 0x3b, 0x2f, 0xd4, 0x36, 0x93, 0x02,
	0xaa, 0x52, 0x78, 0x82, 0xcb, 0x46, 0x47, 0xe7, 0x50, 0xe7, 0xa6, 0x47, 0xe9, 0x6d, 0xf3, 0x7e,
	0x0c, 0x86, 0x21, 0x42, 0x9f, 0x03, 0x37, 0x6c, 0x51, 0x88, 0x19, 0xd8, 0xc6, 0x3a, 0xad, 0x7d,
	0x58, 0xbb, 0x31, 0x5f, 0x5d, 0x9f, 0x22, 0x66, 0xd2, 0x63, 0x9f, 0xae, 0xfd, 0xb4, 0xe9, 0x5f,
	0xfe, 0x37, 0x9f, 0xde, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xc4, 0xa2, 0x02, 0xc6, 0x01,
	0x00, 0x00,
}
