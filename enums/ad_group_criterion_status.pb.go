// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/ad_group_criterion_status.proto

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

// The possible statuses of an AdGroupCriterion.
type AdGroupCriterionStatusEnum_AdGroupCriterionStatus int32

const (
	// No value has been specified.
	AdGroupCriterionStatusEnum_UNSPECIFIED AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdGroupCriterionStatusEnum_UNKNOWN AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 1
	// The ad group criterion is enabled.
	AdGroupCriterionStatusEnum_ENABLED AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 2
	// The ad group criterion is paused.
	AdGroupCriterionStatusEnum_PAUSED AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 3
	// The ad group criterion is removed.
	AdGroupCriterionStatusEnum_REMOVED AdGroupCriterionStatusEnum_AdGroupCriterionStatus = 4
)

var AdGroupCriterionStatusEnum_AdGroupCriterionStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "PAUSED",
	4: "REMOVED",
}

var AdGroupCriterionStatusEnum_AdGroupCriterionStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"PAUSED":      3,
	"REMOVED":     4,
}

func (x AdGroupCriterionStatusEnum_AdGroupCriterionStatus) String() string {
	return proto.EnumName(AdGroupCriterionStatusEnum_AdGroupCriterionStatus_name, int32(x))
}

func (AdGroupCriterionStatusEnum_AdGroupCriterionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2ee8cdd92377e78, []int{0, 0}
}

// Message describing AdGroupCriterion statuses.
type AdGroupCriterionStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupCriterionStatusEnum) Reset()         { *m = AdGroupCriterionStatusEnum{} }
func (m *AdGroupCriterionStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterionStatusEnum) ProtoMessage()    {}
func (*AdGroupCriterionStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2ee8cdd92377e78, []int{0}
}

func (m *AdGroupCriterionStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterionStatusEnum.Unmarshal(m, b)
}
func (m *AdGroupCriterionStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterionStatusEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupCriterionStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterionStatusEnum.Merge(m, src)
}
func (m *AdGroupCriterionStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterionStatusEnum.Size(m)
}
func (m *AdGroupCriterionStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterionStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterionStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.AdGroupCriterionStatusEnum_AdGroupCriterionStatus", AdGroupCriterionStatusEnum_AdGroupCriterionStatus_name, AdGroupCriterionStatusEnum_AdGroupCriterionStatus_value)
	proto.RegisterType((*AdGroupCriterionStatusEnum)(nil), "google.ads.googleads.v0.enums.AdGroupCriterionStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/ad_group_criterion_status.proto", fileDescriptor_c2ee8cdd92377e78)
}

var fileDescriptor_c2ee8cdd92377e78 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x5d, 0x27, 0x13, 0xb2, 0x83, 0xa5, 0x07, 0x0f, 0xca, 0x0e, 0xee, 0x01, 0xd2, 0x82,
	0x47, 0xf1, 0x90, 0xae, 0xb1, 0x0c, 0xb5, 0x2b, 0x96, 0x56, 0x90, 0x40, 0x89, 0x4b, 0x09, 0x85,
	0x35, 0x29, 0x49, 0xbb, 0x83, 0x8f, 0xe3, 0xd1, 0x47, 0x11, 0x1f, 0x4a, 0x92, 0xda, 0x9d, 0xa6,
	0x97, 0xf2, 0x2f, 0x3f, 0x7e, 0x5f, 0xbe, 0xef, 0x0f, 0xee, 0xb8, 0x94, 0x7c, 0x57, 0xf9, 0x94,
	0x69, 0x7f, 0x88, 0x26, 0xed, 0x03, 0xbf, 0x12, 0x7d, 0xa3, 0x7d, 0xca, 0x4a, 0xae, 0x64, 0xdf,
	0x96, 0x5b, 0x55, 0x77, 0x95, 0xaa, 0xa5, 0x28, 0x75, 0x47, 0xbb, 0x5e, 0xc3, 0x56, 0xc9, 0x4e,
	0x7a, 0x8b, 0xc1, 0x81, 0x94, 0x69, 0x78, 0xd0, 0xe1, 0x3e, 0x80, 0x56, 0x5f, 0xbe, 0x83, 0x4b,
	0xc4, 0x62, 0x33, 0x60, 0x35, 0xfa, 0x99, 0xd5, 0xb1, 0xe8, 0x9b, 0x25, 0x01, 0x17, 0xc7, 0xa9,
	0x77, 0x0e, 0xe6, 0x79, 0x92, 0xa5, 0x78, 0xb5, 0xbe, 0x5f, 0xe3, 0xc8, 0x3d, 0xf1, 0xe6, 0xe0,
	0x2c, 0x4f, 0x1e, 0x92, 0xcd, 0x4b, 0xe2, 0x4e, 0xcc, 0x0f, 0x4e, 0x50, 0xf8, 0x88, 0x23, 0xd7,
	0xf1, 0x00, 0x98, 0xa5, 0x28, 0xcf, 0x70, 0xe4, 0x4e, 0x0d, 0x78, 0xc6, 0x4f, 0x9b, 0x02, 0x47,
	0xee, 0x69, 0xf8, 0x3d, 0x01, 0xd7, 0x5b, 0xd9, 0xc0, 0x7f, 0x37, 0x0c, 0xaf, 0x8e, 0x6f, 0x90,
	0x9a, 0xeb, 0xd2, 0xc9, 0x6b, 0xf8, 0x6b, 0x73, 0xb9, 0xa3, 0x82, 0x43, 0xa9, 0xb8, 0xcf, 0x2b,
	0x61, 0x6f, 0x1f, 0xeb, 0x6a, 0x6b, 0xfd, 0x47, 0x7b, 0xb7, 0xf6, 0xfb, 0xe1, 0x4c, 0x63, 0x84,
	0x3e, 0x9d, 0x45, 0x3c, 0x8c, 0x42, 0x4c, 0xc3, 0x21, 0x9a, 0x54, 0x04, 0xd0, 0x74, 0xa1, 0xbf,
	0x46, 0x4e, 0x10, 0xd3, 0xe4, 0xc0, 0x49, 0x11, 0x10, 0xcb, 0xdf, 0x66, 0xf6, 0xd1, 0x9b, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xd0, 0xd4, 0x02, 0xb1, 0x01, 0x00, 0x00,
}
