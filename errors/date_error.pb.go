// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/date_error.proto

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

// Enum describing possible date errors.
type DateErrorEnum_DateError int32

const (
	// Enum unspecified.
	DateErrorEnum_UNSPECIFIED DateErrorEnum_DateError = 0
	// The received error code is not known in this version.
	DateErrorEnum_UNKNOWN DateErrorEnum_DateError = 1
	// Given field values do not correspond to a valid date.
	DateErrorEnum_INVALID_FIELD_VALUES_IN_DATE DateErrorEnum_DateError = 2
	// Given field values do not correspond to a valid date time.
	DateErrorEnum_INVALID_FIELD_VALUES_IN_DATE_TIME DateErrorEnum_DateError = 3
	// The string date's format should be yyyy-mm-dd.
	DateErrorEnum_INVALID_STRING_DATE DateErrorEnum_DateError = 4
	// The string date time's format should be yyyy-mm-dd hh:mm:ss.ssssss.
	DateErrorEnum_INVALID_STRING_DATE_TIME_MICROS DateErrorEnum_DateError = 6
	// The string date time's format should be yyyy-mm-dd hh:mm:ss.
	DateErrorEnum_INVALID_STRING_DATE_TIME_SECONDS DateErrorEnum_DateError = 11
	// Date is before allowed minimum.
	DateErrorEnum_EARLIER_THAN_MINIMUM_DATE DateErrorEnum_DateError = 7
	// Date is after allowed maximum.
	DateErrorEnum_LATER_THAN_MAXIMUM_DATE DateErrorEnum_DateError = 8
	// Date range bounds are not in order.
	DateErrorEnum_DATE_RANGE_MINIMUM_DATE_LATER_THAN_MAXIMUM_DATE DateErrorEnum_DateError = 9
	// Both dates in range are null.
	DateErrorEnum_DATE_RANGE_MINIMUM_AND_MAXIMUM_DATES_BOTH_NULL DateErrorEnum_DateError = 10
)

var DateErrorEnum_DateError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "INVALID_FIELD_VALUES_IN_DATE",
	3:  "INVALID_FIELD_VALUES_IN_DATE_TIME",
	4:  "INVALID_STRING_DATE",
	6:  "INVALID_STRING_DATE_TIME_MICROS",
	11: "INVALID_STRING_DATE_TIME_SECONDS",
	7:  "EARLIER_THAN_MINIMUM_DATE",
	8:  "LATER_THAN_MAXIMUM_DATE",
	9:  "DATE_RANGE_MINIMUM_DATE_LATER_THAN_MAXIMUM_DATE",
	10: "DATE_RANGE_MINIMUM_AND_MAXIMUM_DATES_BOTH_NULL",
}

var DateErrorEnum_DateError_value = map[string]int32{
	"UNSPECIFIED":                                     0,
	"UNKNOWN":                                         1,
	"INVALID_FIELD_VALUES_IN_DATE":                    2,
	"INVALID_FIELD_VALUES_IN_DATE_TIME":               3,
	"INVALID_STRING_DATE":                             4,
	"INVALID_STRING_DATE_TIME_MICROS":                 6,
	"INVALID_STRING_DATE_TIME_SECONDS":                11,
	"EARLIER_THAN_MINIMUM_DATE":                       7,
	"LATER_THAN_MAXIMUM_DATE":                         8,
	"DATE_RANGE_MINIMUM_DATE_LATER_THAN_MAXIMUM_DATE": 9,
	"DATE_RANGE_MINIMUM_AND_MAXIMUM_DATES_BOTH_NULL":  10,
}

func (x DateErrorEnum_DateError) String() string {
	return proto.EnumName(DateErrorEnum_DateError_name, int32(x))
}

func (DateErrorEnum_DateError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6b525c7f3dce743, []int{0, 0}
}

// Container for enum describing possible date errors.
type DateErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DateErrorEnum) Reset()         { *m = DateErrorEnum{} }
func (m *DateErrorEnum) String() string { return proto.CompactTextString(m) }
func (*DateErrorEnum) ProtoMessage()    {}
func (*DateErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6b525c7f3dce743, []int{0}
}

func (m *DateErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateErrorEnum.Unmarshal(m, b)
}
func (m *DateErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateErrorEnum.Marshal(b, m, deterministic)
}
func (m *DateErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateErrorEnum.Merge(m, src)
}
func (m *DateErrorEnum) XXX_Size() int {
	return xxx_messageInfo_DateErrorEnum.Size(m)
}
func (m *DateErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DateErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DateErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.errors.DateErrorEnum_DateError", DateErrorEnum_DateError_name, DateErrorEnum_DateError_value)
	proto.RegisterType((*DateErrorEnum)(nil), "google.ads.googleads.v0.errors.DateErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/date_error.proto", fileDescriptor_a6b525c7f3dce743)
}

var fileDescriptor_a6b525c7f3dce743 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0xae, 0xd2, 0x40,
	0x14, 0x87, 0xa5, 0x98, 0x7b, 0xbd, 0x87, 0xa8, 0xcd, 0xb8, 0xb8, 0x31, 0xea, 0x15, 0x51, 0xb7,
	0xd3, 0x46, 0x96, 0xae, 0x06, 0x66, 0x28, 0x13, 0xdb, 0x29, 0xe9, 0x3f, 0x8d, 0x69, 0x32, 0xa9,
	0xb6, 0x69, 0x4c, 0x80, 0x21, 0x2d, 0xf2, 0x16, 0xbe, 0x84, 0x4b, 0x9f, 0xc3, 0x95, 0xaf, 0xe4,
	0xc6, 0xb4, 0x03, 0x35, 0x24, 0xc0, 0xaa, 0xbf, 0xce, 0xf9, 0xbe, 0x73, 0x16, 0xe7, 0x80, 0x55,
	0x2a, 0x55, 0x2e, 0x0b, 0x2b, 0xcb, 0xeb, 0x7d, 0x6c, 0xd2, 0xce, 0xb6, 0x8a, 0xaa, 0x52, 0x55,
	0x6d, 0xe5, 0xd9, 0xb6, 0x90, 0x6d, 0xc6, 0x9b, 0x4a, 0x6d, 0x15, 0xba, 0xd3, 0x14, 0xce, 0xf2,
	0x1a, 0x77, 0x02, 0xde, 0xd9, 0x58, 0x0b, 0xa3, 0x1f, 0x7d, 0x78, 0x48, 0xb3, 0x6d, 0xc1, 0x9a,
	0x5f, 0xb6, 0xfe, 0xbe, 0x1a, 0xfd, 0x35, 0xe0, 0xa6, 0x7b, 0x41, 0x8f, 0x61, 0x10, 0x8b, 0x70,
	0xc1, 0xa6, 0x7c, 0xc6, 0x19, 0x35, 0xef, 0xa1, 0x01, 0x5c, 0xc7, 0xe2, 0x83, 0xf0, 0x3f, 0x0a,
	0xb3, 0x87, 0x86, 0xf0, 0x9c, 0x8b, 0x84, 0xb8, 0x9c, 0xca, 0x19, 0x67, 0x2e, 0x95, 0x09, 0x71,
	0x63, 0x16, 0x4a, 0x2e, 0x24, 0x25, 0x11, 0x33, 0x0d, 0xf4, 0x16, 0x5e, 0x5d, 0x22, 0x64, 0xc4,
	0x3d, 0x66, 0xf6, 0xd1, 0x2d, 0x3c, 0x39, 0x60, 0x61, 0x14, 0x70, 0xe1, 0x68, 0xff, 0x3e, 0x7a,
	0x0d, 0x2f, 0x4f, 0x14, 0x5a, 0x4d, 0x7a, 0x7c, 0x1a, 0xf8, 0xa1, 0x79, 0x85, 0xde, 0xc0, 0xf0,
	0x2c, 0x14, 0xb2, 0xa9, 0x2f, 0x68, 0x68, 0x0e, 0xd0, 0x0b, 0x78, 0xca, 0x48, 0xe0, 0x72, 0x16,
	0xc8, 0x68, 0x4e, 0x84, 0xf4, 0xb8, 0xe0, 0x5e, 0xec, 0xe9, 0x49, 0xd7, 0xe8, 0x19, 0xdc, 0xba,
	0x24, 0xea, 0x8a, 0xe4, 0xd3, 0xff, 0xe2, 0x03, 0x34, 0x06, 0xab, 0x6d, 0x19, 0x10, 0xe1, 0xb0,
	0x23, 0x53, 0x9e, 0x93, 0x6e, 0xd0, 0x3b, 0xc0, 0x27, 0x24, 0x22, 0xe8, 0x11, 0x18, 0xca, 0x89,
	0x1f, 0xcd, 0xa5, 0x88, 0x5d, 0xd7, 0x84, 0xc9, 0xef, 0x1e, 0x8c, 0xbe, 0xaa, 0x15, 0xbe, 0xbc,
	0xb6, 0xc9, 0xa3, 0x6e, 0x43, 0x8b, 0x66, 0xcd, 0x8b, 0xde, 0x67, 0xba, 0x37, 0x4a, 0xb5, 0xcc,
	0xd6, 0x25, 0x56, 0x55, 0x69, 0x95, 0xc5, 0xba, 0x3d, 0x82, 0xc3, 0xa5, 0x6c, 0xbe, 0xd5, 0xe7,
	0x0e, 0xe7, 0xbd, 0xfe, 0xfc, 0x34, 0xfa, 0x0e, 0x21, 0xbf, 0x8c, 0x3b, 0x47, 0x37, 0x23, 0x79,
	0x8d, 0x75, 0x6c, 0x52, 0x62, 0xe3, 0x76, 0x64, 0xfd, 0xe7, 0x00, 0xa4, 0x24, 0xaf, 0xd3, 0x0e,
	0x48, 0x13, 0x3b, 0xd5, 0xc0, 0x97, 0xab, 0x76, 0xf0, 0xf8, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc4, 0x62, 0x40, 0x2a, 0xb0, 0x02, 0x00, 0x00,
}
