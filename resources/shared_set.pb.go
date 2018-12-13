// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/shared_set.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "github.com/kritzware/google-ads-go/enums"
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

// SharedSets are used for sharing criterion exclusions across multiple
// campaigns.
type SharedSet struct {
	// The resource name of the shared set.
	// Shared set resource names have the form:
	//
	// `customers/{customer_id}/sharedSets/{shared_set_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of this shared set. Read only.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The type of this shared set: each shared set holds only a single kind
	// of entity. Required. Immutable.
	Type enums.SharedSetTypeEnum_SharedSetType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v0.enums.SharedSetTypeEnum_SharedSetType" json:"type,omitempty"`
	// The name of this shared set. Required.
	// Shared Sets must have names that are unique among active shared sets of
	// the same type.
	// The length of this string should be between 1 and 255 UTF-8 bytes,
	// inclusive.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The status of this shared set. Read only.
	Status enums.SharedSetStatusEnum_SharedSetStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v0.enums.SharedSetStatusEnum_SharedSetStatus" json:"status,omitempty"`
	// The number of shared criteria within this shared set. Read only.
	MemberCount *wrappers.Int64Value `protobuf:"bytes,6,opt,name=member_count,json=memberCount,proto3" json:"member_count,omitempty"`
	// The number of campaigns associated with this shared set. Read only.
	ReferenceCount       *wrappers.Int64Value `protobuf:"bytes,7,opt,name=reference_count,json=referenceCount,proto3" json:"reference_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SharedSet) Reset()         { *m = SharedSet{} }
func (m *SharedSet) String() string { return proto.CompactTextString(m) }
func (*SharedSet) ProtoMessage()    {}
func (*SharedSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_a08bbd9f39f73150, []int{0}
}

func (m *SharedSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedSet.Unmarshal(m, b)
}
func (m *SharedSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedSet.Marshal(b, m, deterministic)
}
func (m *SharedSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedSet.Merge(m, src)
}
func (m *SharedSet) XXX_Size() int {
	return xxx_messageInfo_SharedSet.Size(m)
}
func (m *SharedSet) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedSet.DiscardUnknown(m)
}

var xxx_messageInfo_SharedSet proto.InternalMessageInfo

func (m *SharedSet) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *SharedSet) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SharedSet) GetType() enums.SharedSetTypeEnum_SharedSetType {
	if m != nil {
		return m.Type
	}
	return enums.SharedSetTypeEnum_UNSPECIFIED
}

func (m *SharedSet) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SharedSet) GetStatus() enums.SharedSetStatusEnum_SharedSetStatus {
	if m != nil {
		return m.Status
	}
	return enums.SharedSetStatusEnum_UNSPECIFIED
}

func (m *SharedSet) GetMemberCount() *wrappers.Int64Value {
	if m != nil {
		return m.MemberCount
	}
	return nil
}

func (m *SharedSet) GetReferenceCount() *wrappers.Int64Value {
	if m != nil {
		return m.ReferenceCount
	}
	return nil
}

func init() {
	proto.RegisterType((*SharedSet)(nil), "google.ads.googleads.v0.resources.SharedSet")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/shared_set.proto", fileDescriptor_a08bbd9f39f73150)
}

var fileDescriptor_a08bbd9f39f73150 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0x86, 0x49, 0x5a, 0x2b, 0x3b, 0xbb, 0x56, 0x98, 0x53, 0x58, 0x45, 0xba, 0x8a, 0x50, 0x10,
	0x26, 0xa1, 0xab, 0x5e, 0x84, 0x85, 0x54, 0x65, 0xd1, 0x83, 0x2c, 0x89, 0xf4, 0x50, 0x02, 0x61,
	0x9a, 0x7c, 0x8d, 0x81, 0x66, 0x26, 0xcc, 0x4c, 0x2a, 0xfd, 0x3b, 0x1e, 0xfd, 0x1f, 0x5e, 0xbc,
	0xfa, 0x87, 0x24, 0x33, 0xc9, 0x60, 0x91, 0x6e, 0x7b, 0xfb, 0x3a, 0xbc, 0xef, 0xd3, 0x67, 0x3e,
	0x26, 0x68, 0x56, 0x70, 0x5e, 0x6c, 0xc0, 0xa7, 0xb9, 0xf4, 0xcd, 0xd8, 0x4e, 0xdb, 0xc0, 0x17,
	0x20, 0x79, 0x23, 0x32, 0x90, 0xbe, 0xfc, 0x46, 0x05, 0xe4, 0xa9, 0x04, 0x45, 0x6a, 0xc1, 0x15,
	0xc7, 0x57, 0x26, 0x48, 0x68, 0x2e, 0x89, 0xed, 0x90, 0x6d, 0x40, 0x6c, 0xe7, 0xf2, 0xcd, 0x21,
	0x2c, 0xb0, 0xa6, 0xfa, 0x17, 0x99, 0x4a, 0x45, 0x55, 0x23, 0x0d, 0xf9, 0xf2, 0xfa, 0xe4, 0x9a,
	0xda, 0xd5, 0xd0, 0x95, 0x9e, 0x75, 0x25, 0xfd, 0x6b, 0xd5, 0xac, 0xfd, 0xef, 0x82, 0xd6, 0x35,
	0x88, 0x0e, 0xfa, 0xfc, 0xd7, 0x00, 0x9d, 0xc5, 0xba, 0x19, 0x83, 0xc2, 0x2f, 0xd0, 0xa3, 0x5e,
	0x33, 0x65, 0xb4, 0x02, 0xcf, 0x99, 0x38, 0xd3, 0xb3, 0xe8, 0xa2, 0x3f, 0xfc, 0x42, 0x2b, 0xc0,
	0xaf, 0x90, 0x5b, 0xe6, 0x9e, 0x3b, 0x71, 0xa6, 0xe7, 0xb3, 0x27, 0xdd, 0x1d, 0x49, 0xcf, 0x27,
	0x9f, 0x98, 0x7a, 0xfb, 0x7a, 0x41, 0x37, 0x0d, 0x44, 0x6e, 0x99, 0xe3, 0x08, 0x0d, 0x5b, 0x1b,
	0x6f, 0x30, 0x71, 0xa6, 0xe3, 0xd9, 0x0d, 0x39, 0xb4, 0x1d, 0x7d, 0x07, 0x62, 0x4d, 0xbe, 0xee,
	0x6a, 0xf8, 0xc8, 0x9a, 0x6a, 0xff, 0x24, 0xd2, 0x2c, 0x1c, 0xa0, 0xa1, 0x96, 0x1b, 0x6a, 0x85,
	0xa7, 0xff, 0x29, 0xc4, 0x4a, 0x94, 0xac, 0x30, 0x0e, 0x3a, 0x89, 0x97, 0x68, 0x64, 0x56, 0xe9,
	0x3d, 0xd0, 0x1e, 0xf3, 0x53, 0x3d, 0x62, 0xdd, 0xda, 0x37, 0x31, 0x67, 0x51, 0x47, 0xc4, 0x37,
	0xe8, 0xa2, 0x82, 0x6a, 0x05, 0x22, 0xcd, 0x78, 0xc3, 0x94, 0x37, 0x3a, 0xbe, 0x98, 0x73, 0x53,
	0x78, 0xdf, 0xe6, 0xf1, 0x07, 0xf4, 0x58, 0xc0, 0x1a, 0x04, 0xb0, 0x0c, 0x3a, 0xc4, 0xc3, 0xe3,
	0x88, 0xb1, 0xed, 0x68, 0xca, 0xfc, 0x8f, 0x83, 0x5e, 0x66, 0xbc, 0x22, 0x47, 0x5f, 0xdf, 0x7c,
	0x6c, 0x2f, 0x72, 0xd7, 0x72, 0xef, 0x9c, 0xe5, 0xe7, 0xae, 0x54, 0xf0, 0x0d, 0x65, 0x05, 0xe1,
	0xa2, 0xf0, 0x0b, 0x60, 0xfa, 0x5f, 0xfb, 0x87, 0x56, 0x97, 0xf2, 0x9e, 0xaf, 0xe0, 0x9d, 0x9d,
	0x7e, 0xb8, 0x83, 0xdb, 0x30, 0xfc, 0xe9, 0x5e, 0xdd, 0x1a, 0x64, 0x98, 0x4b, 0x62, 0xc6, 0x76,
	0x5a, 0x04, 0x24, 0xea, 0x93, 0xbf, 0xfb, 0x4c, 0x12, 0xe6, 0x32, 0xb1, 0x99, 0x64, 0x11, 0x24,
	0x36, 0xb3, 0x1a, 0x69, 0x89, 0xeb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x69, 0x38, 0xf9,
	0x89, 0x03, 0x00, 0x00,
}
