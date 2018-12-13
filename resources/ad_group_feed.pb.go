// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/ad_group_feed.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/kritzware/google-ads-go/common"
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

// An ad group feed.
type AdGroupFeed struct {
	// The resource name of the ad group feed.
	// Ad group feed resource names have the form:
	//
	// `customers/{customer_id}/adGroupFeeds/{ad_group_id}_{feed_id}
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The feed being linked to the ad group.
	Feed *wrappers.StringValue `protobuf:"bytes,2,opt,name=feed,proto3" json:"feed,omitempty"`
	// The ad group being linked to the feed.
	AdGroup *wrappers.StringValue `protobuf:"bytes,3,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Indicates which placeholder types the feed may populate under the connected
	// ad group. Required.
	PlaceholderTypes []enums.PlaceholderTypeEnum_PlaceholderType `protobuf:"varint,4,rep,packed,name=placeholder_types,json=placeholderTypes,proto3,enum=google.ads.googleads.v0.enums.PlaceholderTypeEnum_PlaceholderType" json:"placeholder_types,omitempty"`
	// Matching function associated with the AdGroupFeed.
	// The matching function is used to filter the set of feed items selected.
	// Required.
	MatchingFunction *common.MatchingFunction `protobuf:"bytes,5,opt,name=matching_function,json=matchingFunction,proto3" json:"matching_function,omitempty"`
	// Status of the ad group feed.
	// This field is read-only.
	Status               enums.FeedLinkStatusEnum_FeedLinkStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v0.enums.FeedLinkStatusEnum_FeedLinkStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *AdGroupFeed) Reset()         { *m = AdGroupFeed{} }
func (m *AdGroupFeed) String() string { return proto.CompactTextString(m) }
func (*AdGroupFeed) ProtoMessage()    {}
func (*AdGroupFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f048fb4293eaa8a, []int{0}
}

func (m *AdGroupFeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupFeed.Unmarshal(m, b)
}
func (m *AdGroupFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupFeed.Marshal(b, m, deterministic)
}
func (m *AdGroupFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupFeed.Merge(m, src)
}
func (m *AdGroupFeed) XXX_Size() int {
	return xxx_messageInfo_AdGroupFeed.Size(m)
}
func (m *AdGroupFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupFeed.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupFeed proto.InternalMessageInfo

func (m *AdGroupFeed) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupFeed) GetFeed() *wrappers.StringValue {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *AdGroupFeed) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupFeed) GetPlaceholderTypes() []enums.PlaceholderTypeEnum_PlaceholderType {
	if m != nil {
		return m.PlaceholderTypes
	}
	return nil
}

func (m *AdGroupFeed) GetMatchingFunction() *common.MatchingFunction {
	if m != nil {
		return m.MatchingFunction
	}
	return nil
}

func (m *AdGroupFeed) GetStatus() enums.FeedLinkStatusEnum_FeedLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.FeedLinkStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*AdGroupFeed)(nil), "google.ads.googleads.v0.resources.AdGroupFeed")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/ad_group_feed.proto", fileDescriptor_1f048fb4293eaa8a)
}

var fileDescriptor_1f048fb4293eaa8a = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6b, 0xd4, 0x40,
	0x10, 0xc7, 0xc9, 0x5d, 0x3d, 0x75, 0xab, 0xe5, 0x9a, 0xa7, 0x50, 0x44, 0xae, 0x8a, 0x70, 0x4f,
	0x9b, 0x70, 0xfe, 0x7a, 0xf0, 0xc5, 0x1c, 0xd8, 0x03, 0x51, 0x39, 0x52, 0x39, 0x44, 0x4e, 0xc2,
	0x36, 0x3b, 0xb7, 0x0d, 0xcd, 0xfe, 0x60, 0x37, 0xa9, 0xf4, 0xdf, 0xf1, 0xd1, 0x7f, 0x44, 0xf0,
	0xdd, 0xff, 0x47, 0x92, 0xdd, 0x44, 0x7b, 0x12, 0x7b, 0x6f, 0x93, 0xe1, 0xfb, 0x99, 0x9d, 0x99,
	0xef, 0x04, 0x3d, 0x67, 0x52, 0xb2, 0x02, 0x42, 0x42, 0x4d, 0x68, 0xc3, 0x3a, 0xba, 0x8c, 0x42,
	0x0d, 0x46, 0x56, 0x3a, 0x03, 0x13, 0x12, 0x9a, 0x32, 0x2d, 0x2b, 0x95, 0x6e, 0x00, 0x28, 0x56,
	0x5a, 0x96, 0xd2, 0x3f, 0xb6, 0x5a, 0x4c, 0xa8, 0xc1, 0x1d, 0x86, 0x2f, 0x23, 0xdc, 0x61, 0x47,
	0x2f, 0xfa, 0x2a, 0x67, 0x92, 0x73, 0x29, 0x42, 0x4e, 0xca, 0xec, 0x3c, 0x17, 0x2c, 0xdd, 0x54,
	0x22, 0x2b, 0x73, 0x29, 0x6c, 0xe9, 0xa3, 0x67, 0x7d, 0x1c, 0x88, 0x8a, 0x9b, 0xb0, 0x6e, 0x22,
	0x2d, 0x72, 0x71, 0x91, 0x9a, 0x92, 0x94, 0x95, 0xd9, 0x8d, 0x52, 0x05, 0xc9, 0xe0, 0x5c, 0x16,
	0x14, 0x74, 0x5a, 0x5e, 0x29, 0x70, 0xd4, 0x43, 0x47, 0x35, 0x5f, 0x67, 0xd5, 0x26, 0xfc, 0xaa,
	0x89, 0x52, 0xa0, 0x5d, 0xd5, 0x47, 0x3f, 0x86, 0x68, 0x3f, 0xa6, 0x8b, 0x7a, 0xfa, 0x13, 0x00,
	0xea, 0x3f, 0x46, 0xf7, 0xdb, 0x01, 0x53, 0x41, 0x38, 0x04, 0xde, 0xc4, 0x9b, 0xde, 0x4d, 0xee,
	0xb5, 0xc9, 0x0f, 0x84, 0x83, 0x1f, 0xa1, 0xbd, 0xba, 0xc9, 0x60, 0x30, 0xf1, 0xa6, 0xfb, 0xb3,
	0x07, 0x6e, 0x3f, 0xb8, 0x7d, 0x03, 0x9f, 0x96, 0x3a, 0x17, 0x6c, 0x45, 0x8a, 0x0a, 0x92, 0x46,
	0xe9, 0xbf, 0x44, 0x77, 0xda, 0x25, 0x07, 0xc3, 0x1d, 0xa8, 0xdb, 0xc4, 0xf6, 0xe4, 0x4b, 0x74,
	0xb8, 0x3d, 0x99, 0x09, 0xf6, 0x26, 0xc3, 0xe9, 0xc1, 0x6c, 0x8e, 0xfb, 0x2c, 0x6a, 0x36, 0x82,
	0x97, 0x7f, 0xb8, 0x8f, 0x57, 0x0a, 0xde, 0x88, 0x8a, 0x6f, 0xe7, 0x92, 0xb1, 0xba, 0x9e, 0x30,
	0xfe, 0x17, 0x74, 0xf8, 0x8f, 0x6f, 0xc1, 0xad, 0xa6, 0xe5, 0xa8, 0xf7, 0x41, 0x6b, 0x38, 0x7e,
	0xef, 0xc0, 0x13, 0xc7, 0x25, 0x63, 0xbe, 0x95, 0xf1, 0x3f, 0xa1, 0x91, 0x75, 0x35, 0x18, 0x4d,
	0xbc, 0xe9, 0xc1, 0xec, 0xf5, 0x0d, 0x43, 0xd4, 0xa6, 0xbc, 0xcb, 0xc5, 0xc5, 0x69, 0x03, 0x35,
	0x33, 0x5c, 0x4f, 0x25, 0xae, 0xde, 0xfc, 0x97, 0x87, 0x9e, 0x64, 0x92, 0xe3, 0x1b, 0xef, 0x76,
	0x3e, 0xfe, 0xcb, 0xf0, 0x65, 0xbd, 0xfd, 0xa5, 0xf7, 0xf9, 0xad, 0xc3, 0x98, 0x2c, 0x88, 0x60,
	0x58, 0x6a, 0x16, 0x32, 0x10, 0x8d, 0x37, 0xed, 0xb5, 0xa9, 0xdc, 0xfc, 0xe7, 0x27, 0x7a, 0xd5,
	0x45, 0xdf, 0x06, 0xc3, 0x45, 0x1c, 0x7f, 0x1f, 0x1c, 0x2f, 0x6c, 0xc9, 0x98, 0x1a, 0x6c, 0xc3,
	0x3a, 0x5a, 0x45, 0x38, 0x69, 0x95, 0x3f, 0x5b, 0xcd, 0x3a, 0xa6, 0x66, 0xdd, 0x69, 0xd6, 0xab,
	0x68, 0xdd, 0x69, 0xce, 0x46, 0x4d, 0x13, 0x4f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x06,
	0x3f, 0xc0, 0xc8, 0x03, 0x00, 0x00,
}
