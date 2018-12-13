// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/resource_access_denied_error.proto

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

// Enum describing possible resource access denied errors.
type ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError int32

const (
	// Enum unspecified.
	ResourceAccessDeniedErrorEnum_UNSPECIFIED ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError = 0
	// The received error code is not known in this version.
	ResourceAccessDeniedErrorEnum_UNKNOWN ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError = 1
	// User did not have write access.
	ResourceAccessDeniedErrorEnum_WRITE_ACCESS_DENIED ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError = 3
)

var ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "WRITE_ACCESS_DENIED",
}

var ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"WRITE_ACCESS_DENIED": 3,
}

func (x ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError) String() string {
	return proto.EnumName(ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_name, int32(x))
}

func (ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_be307f994c248b0a, []int{0, 0}
}

// Container for enum describing possible resource access denied errors.
type ResourceAccessDeniedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourceAccessDeniedErrorEnum) Reset()         { *m = ResourceAccessDeniedErrorEnum{} }
func (m *ResourceAccessDeniedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ResourceAccessDeniedErrorEnum) ProtoMessage()    {}
func (*ResourceAccessDeniedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_be307f994c248b0a, []int{0}
}

func (m *ResourceAccessDeniedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceAccessDeniedErrorEnum.Unmarshal(m, b)
}
func (m *ResourceAccessDeniedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceAccessDeniedErrorEnum.Marshal(b, m, deterministic)
}
func (m *ResourceAccessDeniedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceAccessDeniedErrorEnum.Merge(m, src)
}
func (m *ResourceAccessDeniedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ResourceAccessDeniedErrorEnum.Size(m)
}
func (m *ResourceAccessDeniedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceAccessDeniedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceAccessDeniedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.errors.ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError", ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_name, ResourceAccessDeniedErrorEnum_ResourceAccessDeniedError_value)
	proto.RegisterType((*ResourceAccessDeniedErrorEnum)(nil), "google.ads.googleads.v0.errors.ResourceAccessDeniedErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/resource_access_denied_error.proto", fileDescriptor_be307f994c248b0a)
}

var fileDescriptor_be307f994c248b0a = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0xed, 0x06, 0x0a, 0xd9, 0xc1, 0x51, 0x0f, 0xe2, 0xc1, 0x1e, 0xfa, 0x00, 0x69, 0xc1,
	0xa3, 0xa7, 0xac, 0x8d, 0xa3, 0x08, 0xb1, 0xb4, 0x6e, 0x03, 0x29, 0x84, 0xda, 0x84, 0x30, 0xd8,
	0x9a, 0x91, 0xcf, 0xed, 0x81, 0x3c, 0xfa, 0x28, 0x9e, 0x7d, 0x20, 0x49, 0xb2, 0xed, 0xd6, 0x9d,
	0xfa, 0x87, 0xef, 0xd7, 0xdf, 0x3f, 0xdf, 0x87, 0x88, 0xd2, 0x5a, 0x6d, 0x64, 0xd2, 0x0a, 0x48,
	0x7c, 0xb4, 0xe9, 0x90, 0x26, 0xd2, 0x18, 0x6d, 0x20, 0x31, 0x12, 0xf4, 0xde, 0x74, 0x92, 0xb7,
	0x5d, 0x27, 0x01, 0xb8, 0x90, 0xfd, 0x5a, 0x0a, 0xee, 0xa6, 0x78, 0x67, 0xf4, 0x97, 0x0e, 0x23,
	0xff, 0x1f, 0x6e, 0x05, 0xe0, 0xb3, 0x02, 0x1f, 0x52, 0xec, 0x15, 0x31, 0xa0, 0xc7, 0xea, 0x68,
	0x21, 0x4e, 0x92, 0x3b, 0x07, 0xb5, 0x53, 0xda, 0xef, 0xb7, 0x71, 0x85, 0x1e, 0x06, 0x81, 0xf0,
	0x16, 0x4d, 0x16, 0xac, 0x2e, 0x69, 0x56, 0xbc, 0x14, 0x34, 0x9f, 0x5e, 0x85, 0x13, 0x74, 0xb3,
	0x60, 0xaf, 0xec, 0x6d, 0xc5, 0xa6, 0x41, 0x78, 0x8f, 0xee, 0x56, 0x55, 0xf1, 0x4e, 0x39, 0xc9,
	0x32, 0x5a, 0xd7, 0x3c, 0xa7, 0xcc, 0x52, 0xe3, 0xd9, 0x5f, 0x80, 0xe2, 0x4e, 0x6f, 0xf1, 0xe5,
	0xb7, 0xcd, 0xa2, 0xc1, 0xe2, 0xd2, 0xee, 0x56, 0x06, 0x1f, 0xf9, 0xd1, 0xa0, 0xf4, 0xa6, 0xed,
	0x15, 0xd6, 0x46, 0x25, 0x4a, 0xf6, 0x6e, 0xf3, 0xd3, 0xc1, 0x76, 0x6b, 0x18, 0xba, 0xdf, 0xb3,
	0xff, 0x7c, 0x8f, 0xc6, 0x73, 0x42, 0x7e, 0x46, 0xd1, 0xdc, 0xcb, 0x88, 0x00, 0xec, 0xa3, 0x4d,
	0xcb, 0x14, 0xbb, 0x4a, 0xf8, 0x3d, 0x01, 0x0d, 0x11, 0xd0, 0x9c, 0x81, 0x66, 0x99, 0x36, 0x1e,
	0xf8, 0xbc, 0x76, 0xc5, 0x4f, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x64, 0xcb, 0xa6, 0x9e, 0xb7,
	0x01, 0x00, 0x00,
}
