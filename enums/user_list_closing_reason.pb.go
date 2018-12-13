// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/user_list_closing_reason.proto

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

// Enum describing possible user list closing reasons.
type UserListClosingReasonEnum_UserListClosingReason int32

const (
	// Not specified.
	UserListClosingReasonEnum_UNSPECIFIED UserListClosingReasonEnum_UserListClosingReason = 0
	// Used for return value only. Represents value unknown in this version.
	UserListClosingReasonEnum_UNKNOWN UserListClosingReasonEnum_UserListClosingReason = 1
	// The userlist was closed because of not being used for over one year.
	UserListClosingReasonEnum_UNUSED UserListClosingReasonEnum_UserListClosingReason = 2
)

var UserListClosingReasonEnum_UserListClosingReason_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "UNUSED",
}

var UserListClosingReasonEnum_UserListClosingReason_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"UNUSED":      2,
}

func (x UserListClosingReasonEnum_UserListClosingReason) String() string {
	return proto.EnumName(UserListClosingReasonEnum_UserListClosingReason_name, int32(x))
}

func (UserListClosingReasonEnum_UserListClosingReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5cc89bf8e48e4df0, []int{0, 0}
}

// Indicates the reason why the userlist was closed.
// This enum is only used when a list is auto-closed by the system.
type UserListClosingReasonEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListClosingReasonEnum) Reset()         { *m = UserListClosingReasonEnum{} }
func (m *UserListClosingReasonEnum) String() string { return proto.CompactTextString(m) }
func (*UserListClosingReasonEnum) ProtoMessage()    {}
func (*UserListClosingReasonEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cc89bf8e48e4df0, []int{0}
}

func (m *UserListClosingReasonEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListClosingReasonEnum.Unmarshal(m, b)
}
func (m *UserListClosingReasonEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListClosingReasonEnum.Marshal(b, m, deterministic)
}
func (m *UserListClosingReasonEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListClosingReasonEnum.Merge(m, src)
}
func (m *UserListClosingReasonEnum) XXX_Size() int {
	return xxx_messageInfo_UserListClosingReasonEnum.Size(m)
}
func (m *UserListClosingReasonEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListClosingReasonEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListClosingReasonEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.UserListClosingReasonEnum_UserListClosingReason", UserListClosingReasonEnum_UserListClosingReason_name, UserListClosingReasonEnum_UserListClosingReason_value)
	proto.RegisterType((*UserListClosingReasonEnum)(nil), "google.ads.googleads.v0.enums.UserListClosingReasonEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/user_list_closing_reason.proto", fileDescriptor_5cc89bf8e48e4df0)
}

var fileDescriptor_5cc89bf8e48e4df0 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xd2, 0xe2, 0xd4, 0xa2, 0xf8, 0x9c, 0xcc, 0xe2, 0x92, 0xf8, 0xe4, 0x9c,
	0xfc, 0xe2, 0xcc, 0xbc, 0xf4, 0xf8, 0xa2, 0xd4, 0xc4, 0xe2, 0xfc, 0x3c, 0xbd, 0x82, 0xa2, 0xfc,
	0x92, 0x7c, 0x21, 0x59, 0x88, 0x16, 0xbd, 0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x6e, 0xbd, 0x32, 0x03,
	0x3d, 0xb0, 0x6e, 0xa5, 0x38, 0x2e, 0xc9, 0xd0, 0xe2, 0xd4, 0x22, 0x9f, 0xcc, 0xe2, 0x12, 0x67,
	0x88, 0xf6, 0x20, 0xb0, 0x6e, 0xd7, 0xbc, 0xd2, 0x5c, 0x25, 0x47, 0x2e, 0x51, 0xac, 0x92, 0x42,
	0xfc, 0x5c, 0xdc, 0xa1, 0x7e, 0xc1, 0x01, 0xae, 0xce, 0x9e, 0x6e, 0x9e, 0xae, 0x2e, 0x02, 0x0c,
	0x42, 0xdc, 0x5c, 0xec, 0xa1, 0x7e, 0xde, 0x7e, 0xfe, 0xe1, 0x7e, 0x02, 0x8c, 0x42, 0x5c, 0x5c,
	0x6c, 0xa1, 0x7e, 0xa1, 0xc1, 0xae, 0x2e, 0x02, 0x4c, 0x4e, 0xa7, 0x19, 0xb9, 0x14, 0x93, 0xf3,
	0x73, 0xf5, 0xf0, 0xba, 0xc2, 0x49, 0x0a, 0xab, 0x35, 0x01, 0x20, 0x0f, 0x04, 0x30, 0x46, 0x39,
	0x41, 0x35, 0xa7, 0xe7, 0xe7, 0x24, 0xe6, 0xa5, 0xeb, 0xe5, 0x17, 0xa5, 0xeb, 0xa7, 0xa7, 0xe6,
	0x81, 0xbd, 0x07, 0x0b, 0x90, 0x82, 0xcc, 0x62, 0x1c, 0xe1, 0x63, 0x0d, 0x26, 0x17, 0x31, 0x31,
	0xbb, 0x3b, 0x3a, 0xae, 0x62, 0x92, 0x75, 0x87, 0x18, 0xe5, 0x98, 0x52, 0xac, 0x07, 0x61, 0x82,
	0x58, 0x61, 0x06, 0x7a, 0x20, 0xff, 0x16, 0x9f, 0x82, 0xc9, 0xc7, 0x38, 0xa6, 0x14, 0xc7, 0xc0,
	0xe5, 0x63, 0xc2, 0x0c, 0x62, 0xc0, 0xf2, 0x49, 0x6c, 0x60, 0x4b, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xea, 0x1d, 0x66, 0x86, 0x93, 0x01, 0x00, 0x00,
}
