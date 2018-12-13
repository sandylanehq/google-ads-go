// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/search_term_targeting_status.proto

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

// Indicates whether the search term is one of your targeted or excluded
// keywords.
type SearchTermTargetingStatusEnum_SearchTermTargetingStatus int32

const (
	// Not specified.
	SearchTermTargetingStatusEnum_UNSPECIFIED SearchTermTargetingStatusEnum_SearchTermTargetingStatus = 0
	// Used for return value only. Represents value unknown in this version.
	SearchTermTargetingStatusEnum_UNKNOWN SearchTermTargetingStatusEnum_SearchTermTargetingStatus = 1
	// Search term is added to targeted keywords.
	SearchTermTargetingStatusEnum_ADDED SearchTermTargetingStatusEnum_SearchTermTargetingStatus = 2
	// Search term matches a negative keyword.
	SearchTermTargetingStatusEnum_EXCLUDED SearchTermTargetingStatusEnum_SearchTermTargetingStatus = 3
	// Search term has been both added and excluded.
	SearchTermTargetingStatusEnum_ADDED_EXCLUDED SearchTermTargetingStatusEnum_SearchTermTargetingStatus = 4
	// Search term is neither targeted nor excluded.
	SearchTermTargetingStatusEnum_NONE SearchTermTargetingStatusEnum_SearchTermTargetingStatus = 5
)

var SearchTermTargetingStatusEnum_SearchTermTargetingStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ADDED",
	3: "EXCLUDED",
	4: "ADDED_EXCLUDED",
	5: "NONE",
}

var SearchTermTargetingStatusEnum_SearchTermTargetingStatus_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"ADDED":          2,
	"EXCLUDED":       3,
	"ADDED_EXCLUDED": 4,
	"NONE":           5,
}

func (x SearchTermTargetingStatusEnum_SearchTermTargetingStatus) String() string {
	return proto.EnumName(SearchTermTargetingStatusEnum_SearchTermTargetingStatus_name, int32(x))
}

func (SearchTermTargetingStatusEnum_SearchTermTargetingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7ac067e7cfba810f, []int{0, 0}
}

// Container for enum indicating whether a search term is one of your targeted
// or excluded keywords.
type SearchTermTargetingStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchTermTargetingStatusEnum) Reset()         { *m = SearchTermTargetingStatusEnum{} }
func (m *SearchTermTargetingStatusEnum) String() string { return proto.CompactTextString(m) }
func (*SearchTermTargetingStatusEnum) ProtoMessage()    {}
func (*SearchTermTargetingStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac067e7cfba810f, []int{0}
}

func (m *SearchTermTargetingStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchTermTargetingStatusEnum.Unmarshal(m, b)
}
func (m *SearchTermTargetingStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchTermTargetingStatusEnum.Marshal(b, m, deterministic)
}
func (m *SearchTermTargetingStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchTermTargetingStatusEnum.Merge(m, src)
}
func (m *SearchTermTargetingStatusEnum) XXX_Size() int {
	return xxx_messageInfo_SearchTermTargetingStatusEnum.Size(m)
}
func (m *SearchTermTargetingStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchTermTargetingStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SearchTermTargetingStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.SearchTermTargetingStatusEnum_SearchTermTargetingStatus", SearchTermTargetingStatusEnum_SearchTermTargetingStatus_name, SearchTermTargetingStatusEnum_SearchTermTargetingStatus_value)
	proto.RegisterType((*SearchTermTargetingStatusEnum)(nil), "google.ads.googleads.v0.enums.SearchTermTargetingStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/search_term_targeting_status.proto", fileDescriptor_7ac067e7cfba810f)
}

var fileDescriptor_7ac067e7cfba810f = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x4e, 0xbb, 0x30,
	0x1c, 0xc5, 0x7f, 0xec, 0xcf, 0xcf, 0xd9, 0x19, 0x25, 0xbd, 0xf3, 0x62, 0x26, 0xee, 0x01, 0x0a,
	0x89, 0x97, 0xde, 0x08, 0xa3, 0x2e, 0x8b, 0xa6, 0x23, 0x61, 0x4c, 0x63, 0x48, 0x48, 0x1d, 0x4d,
	0x5d, 0x32, 0x28, 0x69, 0xcb, 0xde, 0xc3, 0x57, 0xf0, 0xd2, 0x47, 0xf1, 0xc6, 0x57, 0x32, 0x14,
	0xe1, 0x0e, 0x6f, 0x9a, 0xd3, 0x9e, 0xf6, 0xd3, 0xf3, 0x3d, 0xe0, 0x8e, 0x0b, 0xc1, 0x0f, 0xcc,
	0xa1, 0x99, 0x72, 0x1a, 0x59, 0xab, 0xa3, 0xeb, 0xb0, 0xa2, 0xca, 0x95, 0xa3, 0x18, 0x95, 0xbb,
	0xb7, 0x54, 0x33, 0x99, 0xa7, 0x9a, 0x4a, 0xce, 0xf4, 0xbe, 0xe0, 0xa9, 0xd2, 0x54, 0x57, 0x0a,
	0x95, 0x52, 0x68, 0x01, 0x67, 0xcd, 0x33, 0x44, 0x33, 0x85, 0x3a, 0x02, 0x3a, 0xba, 0xc8, 0x10,
	0xe6, 0xef, 0x16, 0x98, 0x45, 0x86, 0xb2, 0x61, 0x32, 0xdf, 0xb4, 0x8c, 0xc8, 0x20, 0x70, 0x51,
	0xe5, 0xf3, 0x12, 0x5c, 0xf6, 0x5e, 0x80, 0x17, 0x60, 0x1a, 0x93, 0x28, 0xc4, 0x8b, 0xd5, 0xfd,
	0x0a, 0x07, 0xf6, 0x3f, 0x38, 0x05, 0x27, 0x31, 0x79, 0x20, 0xeb, 0x27, 0x62, 0x5b, 0xf0, 0x14,
	0x8c, 0xbd, 0x20, 0xc0, 0x81, 0x3d, 0x80, 0x67, 0x60, 0x82, 0x9f, 0x17, 0x8f, 0x71, 0xbd, 0x1b,
	0x42, 0x08, 0xce, 0x8d, 0x91, 0x76, 0x67, 0x23, 0x38, 0x01, 0x23, 0xb2, 0x26, 0xd8, 0x1e, 0xfb,
	0xdf, 0x16, 0xb8, 0xde, 0x89, 0x1c, 0xfd, 0x99, 0xdc, 0xbf, 0xea, 0x4d, 0x15, 0xd6, 0x83, 0x87,
	0xd6, 0x8b, 0xff, 0x0b, 0xe0, 0xe2, 0x40, 0x0b, 0x8e, 0x84, 0xe4, 0x0e, 0x67, 0x85, 0xa9, 0xa5,
	0x2d, 0xb3, 0xdc, 0xab, 0x9e, 0x6e, 0x6f, 0xcd, 0xfa, 0x31, 0x18, 0x2e, 0x3d, 0xef, 0x73, 0x30,
	0x5b, 0x36, 0x28, 0x2f, 0x53, 0xa8, 0x91, 0xb5, 0xda, 0xba, 0xa8, 0xae, 0x48, 0x7d, 0xb5, 0x7e,
	0xe2, 0x65, 0x2a, 0xe9, 0xfc, 0x64, 0xeb, 0x26, 0xc6, 0x7f, 0xfd, 0x6f, 0x3e, 0xbd, 0xf9, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x67, 0x8b, 0xe4, 0x66, 0xcf, 0x01, 0x00, 0x00,
}
