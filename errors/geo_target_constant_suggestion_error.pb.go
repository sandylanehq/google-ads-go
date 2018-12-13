// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/geo_target_constant_suggestion_error.proto

package errors

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

// Enum describing possible geo target constant suggestion errors.
type GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError int32

const (
	// Enum unspecified.
	GeoTargetConstantSuggestionErrorEnum_UNSPECIFIED GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError = 0
	// The received error code is not known in this version.
	GeoTargetConstantSuggestionErrorEnum_UNKNOWN GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError = 1
	// A location name cannot be greater than 300 characters.
	GeoTargetConstantSuggestionErrorEnum_LOCATION_NAME_SIZE_LIMIT GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError = 2
	// At most 25 location names can be specified in a SuggestGeoTargetConstants
	// method.
	GeoTargetConstantSuggestionErrorEnum_LOCATION_NAME_LIMIT GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError = 3
	// The country code is invalid.
	GeoTargetConstantSuggestionErrorEnum_INVALID_COUNTRY_CODE GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError = 4
	// Geo target constant resource names or location names must be provided in
	// the request.
	GeoTargetConstantSuggestionErrorEnum_REQUEST_PARAMETERS_UNSET GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError = 5
)

var GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "LOCATION_NAME_SIZE_LIMIT",
	3: "LOCATION_NAME_LIMIT",
	4: "INVALID_COUNTRY_CODE",
	5: "REQUEST_PARAMETERS_UNSET",
}

var GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"LOCATION_NAME_SIZE_LIMIT": 2,
	"LOCATION_NAME_LIMIT":      3,
	"INVALID_COUNTRY_CODE":     4,
	"REQUEST_PARAMETERS_UNSET": 5,
}

func (x GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError) String() string {
	return proto.EnumName(GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError_name, int32(x))
}

func (GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0211167de73473bf, []int{0, 0}
}

// Container for enum describing possible geo target constant suggestion errors.
type GeoTargetConstantSuggestionErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeoTargetConstantSuggestionErrorEnum) Reset()         { *m = GeoTargetConstantSuggestionErrorEnum{} }
func (m *GeoTargetConstantSuggestionErrorEnum) String() string { return proto.CompactTextString(m) }
func (*GeoTargetConstantSuggestionErrorEnum) ProtoMessage()    {}
func (*GeoTargetConstantSuggestionErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_0211167de73473bf, []int{0}
}

func (m *GeoTargetConstantSuggestionErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeoTargetConstantSuggestionErrorEnum.Unmarshal(m, b)
}
func (m *GeoTargetConstantSuggestionErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeoTargetConstantSuggestionErrorEnum.Marshal(b, m, deterministic)
}
func (m *GeoTargetConstantSuggestionErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeoTargetConstantSuggestionErrorEnum.Merge(m, src)
}
func (m *GeoTargetConstantSuggestionErrorEnum) XXX_Size() int {
	return xxx_messageInfo_GeoTargetConstantSuggestionErrorEnum.Size(m)
}
func (m *GeoTargetConstantSuggestionErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_GeoTargetConstantSuggestionErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_GeoTargetConstantSuggestionErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.errors.GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError", GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError_name, GeoTargetConstantSuggestionErrorEnum_GeoTargetConstantSuggestionError_value)
	proto.RegisterType((*GeoTargetConstantSuggestionErrorEnum)(nil), "google.ads.googleads.v0.errors.GeoTargetConstantSuggestionErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/geo_target_constant_suggestion_error.proto", fileDescriptor_0211167de73473bf)
}

var fileDescriptor_0211167de73473bf = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x86, 0xbf, 0xb4, 0x9f, 0x0a, 0xd3, 0x85, 0x21, 0x0a, 0x76, 0x21, 0x45, 0x8a, 0x6e, 0x27,
	0x01, 0x97, 0xae, 0xa6, 0xc9, 0x58, 0x06, 0xdb, 0x49, 0xcc, 0x4f, 0xc5, 0x12, 0x18, 0x62, 0x33,
	0x0c, 0x85, 0x36, 0x53, 0x32, 0x69, 0x2f, 0xc8, 0xa5, 0x1b, 0xef, 0xc3, 0x2b, 0x10, 0xbc, 0x19,
	0x49, 0xa6, 0x2d, 0xb8, 0xd0, 0xae, 0xf2, 0xc2, 0x79, 0xf2, 0x1c, 0xe6, 0x3d, 0x80, 0x08, 0x29,
	0xc5, 0x82, 0xdb, 0x59, 0xae, 0x6c, 0x1d, 0xeb, 0xb4, 0x71, 0x6c, 0x5e, 0x96, 0xb2, 0x54, 0xb6,
	0xe0, 0x92, 0x55, 0x59, 0x29, 0x78, 0xc5, 0x66, 0xb2, 0x50, 0x55, 0x56, 0x54, 0x4c, 0xad, 0x85,
	0xe0, 0xaa, 0x9a, 0xcb, 0x82, 0x35, 0x14, 0x5c, 0x95, 0xb2, 0x92, 0x56, 0x4f, 0xff, 0x0f, 0xb3,
	0x5c, 0xc1, 0xbd, 0x0a, 0x6e, 0x1c, 0xa8, 0x55, 0xfd, 0x4f, 0x03, 0x5c, 0x0f, 0xb9, 0x8c, 0x1b,
	0x9b, 0xbb, 0x95, 0x45, 0x7b, 0x17, 0xae, 0x29, 0x5c, 0xac, 0x97, 0xfd, 0x77, 0x03, 0x5c, 0x1d,
	0x02, 0xad, 0x53, 0xd0, 0x49, 0x68, 0x14, 0x60, 0x97, 0xdc, 0x13, 0xec, 0x99, 0xff, 0xac, 0x0e,
	0x38, 0x49, 0xe8, 0x03, 0xf5, 0x9f, 0xa8, 0x69, 0x58, 0x97, 0xa0, 0x3b, 0xf2, 0x5d, 0x14, 0x13,
	0x9f, 0x32, 0x8a, 0xc6, 0x98, 0x45, 0x64, 0x8a, 0xd9, 0x88, 0x8c, 0x49, 0x6c, 0xb6, 0xac, 0x0b,
	0x70, 0xf6, 0x73, 0xaa, 0x07, 0x6d, 0xab, 0x0b, 0xce, 0x09, 0x9d, 0xa0, 0x11, 0xf1, 0x98, 0xeb,
	0x27, 0x34, 0x0e, 0x9f, 0x99, 0xeb, 0x7b, 0xd8, 0xfc, 0x5f, 0x0b, 0x43, 0xfc, 0x98, 0xe0, 0x28,
	0x66, 0x01, 0x0a, 0xd1, 0x18, 0xc7, 0x38, 0x8c, 0x58, 0x42, 0x23, 0x1c, 0x9b, 0x47, 0x83, 0x2f,
	0x03, 0xf4, 0x67, 0x72, 0x09, 0xff, 0x6e, 0x60, 0x70, 0x73, 0xe8, 0x55, 0x41, 0x5d, 0x64, 0x60,
	0x4c, 0xbd, 0xad, 0x48, 0xc8, 0x45, 0x56, 0x08, 0x28, 0x4b, 0x61, 0x0b, 0x5e, 0x34, 0x35, 0xef,
	0xae, 0xb4, 0x9a, 0xab, 0xdf, 0x8e, 0x76, 0xa7, 0x3f, 0xaf, 0xad, 0xf6, 0x10, 0xa1, 0xb7, 0x56,
	0x6f, 0xa8, 0x65, 0x28, 0x57, 0x50, 0xc7, 0x3a, 0x4d, 0x1c, 0xd8, 0xac, 0x54, 0x1f, 0x3b, 0x20,
	0x45, 0xb9, 0x4a, 0xf7, 0x40, 0x3a, 0x71, 0x52, 0x0d, 0xbc, 0x1c, 0x37, 0x8b, 0x6f, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x71, 0xd6, 0x36, 0xfe, 0x2c, 0x02, 0x00, 0x00,
}
