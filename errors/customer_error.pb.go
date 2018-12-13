// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/customer_error.proto

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

// Set of errors that are related to requests dealing with Customer.
// Next id: 26
type CustomerErrorEnum_CustomerError int32

const (
	// Enum unspecified.
	CustomerErrorEnum_UNSPECIFIED CustomerErrorEnum_CustomerError = 0
	// The received error code is not known in this version.
	CustomerErrorEnum_UNKNOWN CustomerErrorEnum_CustomerError = 1
	// Customer status is not allowed to be changed from DRAFT and CLOSED.
	// Currency code and at least one of country code and time zone needs to be
	// set when status is changed to ENABLED.
	CustomerErrorEnum_STATUS_CHANGE_DISALLOWED CustomerErrorEnum_CustomerError = 2
	// CustomerService cannot get a customer that has not been fully set up.
	CustomerErrorEnum_ACCOUNT_NOT_SET_UP CustomerErrorEnum_CustomerError = 3
)

var CustomerErrorEnum_CustomerError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "STATUS_CHANGE_DISALLOWED",
	3: "ACCOUNT_NOT_SET_UP",
}

var CustomerErrorEnum_CustomerError_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"STATUS_CHANGE_DISALLOWED": 2,
	"ACCOUNT_NOT_SET_UP":       3,
}

func (x CustomerErrorEnum_CustomerError) String() string {
	return proto.EnumName(CustomerErrorEnum_CustomerError_name, int32(x))
}

func (CustomerErrorEnum_CustomerError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c9e66c3070369bbd, []int{0, 0}
}

// Container for enum describing possible customer errors.
type CustomerErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerErrorEnum) Reset()         { *m = CustomerErrorEnum{} }
func (m *CustomerErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CustomerErrorEnum) ProtoMessage()    {}
func (*CustomerErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9e66c3070369bbd, []int{0}
}

func (m *CustomerErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerErrorEnum.Unmarshal(m, b)
}
func (m *CustomerErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerErrorEnum.Marshal(b, m, deterministic)
}
func (m *CustomerErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerErrorEnum.Merge(m, src)
}
func (m *CustomerErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CustomerErrorEnum.Size(m)
}
func (m *CustomerErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.errors.CustomerErrorEnum_CustomerError", CustomerErrorEnum_CustomerError_name, CustomerErrorEnum_CustomerError_value)
	proto.RegisterType((*CustomerErrorEnum)(nil), "google.ads.googleads.v0.errors.CustomerErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/customer_error.proto", fileDescriptor_c9e66c3070369bbd)
}

var fileDescriptor_c9e66c3070369bbd = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x5d, 0x07, 0x0a, 0x19, 0x62, 0xcd, 0x41, 0x3c, 0xc8, 0x0e, 0xfd, 0x00, 0x69, 0x61,
	0x47, 0x4f, 0x59, 0x1b, 0x6b, 0x71, 0xa4, 0x85, 0x26, 0x1b, 0x48, 0x21, 0xd4, 0xb6, 0x04, 0x61,
	0x6d, 0x46, 0xb2, 0x0d, 0x3f, 0x8f, 0x47, 0x3f, 0x89, 0xf8, 0xa9, 0xa4, 0xcd, 0x36, 0xd8, 0x41,
	0x4f, 0x79, 0xf2, 0xf2, 0xfc, 0xde, 0x3f, 0x0f, 0x98, 0x49, 0xa5, 0xe4, 0xba, 0xf1, 0xcb, 0xda,
	0xf8, 0x56, 0xf6, 0x6a, 0x1f, 0xf8, 0x8d, 0xd6, 0x4a, 0x1b, 0xbf, 0xda, 0x99, 0xad, 0x6a, 0x1b,
	0x2d, 0x86, 0x3f, 0xda, 0x68, 0xb5, 0x55, 0x70, 0x6a, 0x9d, 0xa8, 0xac, 0x0d, 0x3a, 0x41, 0x68,
	0x1f, 0x20, 0x0b, 0x79, 0x1f, 0xe0, 0x36, 0x3c, 0x70, 0xa4, 0xaf, 0x90, 0x6e, 0xd7, 0x7a, 0x15,
	0xb8, 0x3e, 0x2b, 0xc2, 0x1b, 0x30, 0xe1, 0x34, 0xcf, 0x48, 0x98, 0x3c, 0x25, 0x24, 0x72, 0x2f,
	0xe0, 0x04, 0x5c, 0x71, 0xfa, 0x42, 0xd3, 0x15, 0x75, 0x47, 0xf0, 0x01, 0xdc, 0xe7, 0x0c, 0x33,
	0x9e, 0x8b, 0xf0, 0x19, 0xd3, 0x98, 0x88, 0x28, 0xc9, 0xf1, 0x62, 0x91, 0xae, 0x48, 0xe4, 0x3a,
	0xf0, 0x0e, 0x40, 0x1c, 0x86, 0x29, 0xa7, 0x4c, 0xd0, 0x94, 0x89, 0x9c, 0x30, 0xc1, 0x33, 0x77,
	0x3c, 0xff, 0x1e, 0x01, 0xaf, 0x52, 0x2d, 0xfa, 0x7f, 0xc1, 0x39, 0x3c, 0xdb, 0x24, 0xeb, 0x8f,
	0xca, 0x46, 0xaf, 0xd1, 0x81, 0x92, 0x6a, 0x5d, 0x76, 0x12, 0x29, 0x2d, 0x7d, 0xd9, 0x74, 0xc3,
	0xc9, 0xc7, 0x6c, 0x36, 0xef, 0xe6, 0xaf, 0xa8, 0x1e, 0xed, 0xf3, 0xe9, 0x8c, 0x63, 0x8c, 0xbf,
	0x9c, 0x69, 0x6c, 0x9b, 0xe1, 0xda, 0x20, 0x2b, 0x7b, 0xb5, 0x0c, 0xd0, 0x30, 0xd2, 0xfc, 0x1c,
	0x0d, 0x05, 0xae, 0x4d, 0x71, 0x32, 0x14, 0xcb, 0xa0, 0xb0, 0x86, 0xb7, 0xcb, 0x61, 0xf0, 0xec,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x98, 0x99, 0x68, 0x2b, 0xa2, 0x01, 0x00, 0x00,
}
