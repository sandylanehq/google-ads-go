// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/feed_item.proto

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

// A feed item.
type FeedItem struct {
	// The resource name of the feed item.
	// Feed item resource names have the form:
	//
	// `customers/{customer_id}/feedItems/{feed_id}_{feed_item_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The feed to which this feed item belongs.
	Feed *wrappers.StringValue `protobuf:"bytes,2,opt,name=feed,proto3" json:"feed,omitempty"`
	// The ID of this feed item.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Start time in which this feed item is effective and can begin serving.
	// The format is "YYYY-MM-DD HH:MM:SS".
	// Examples: "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	StartDateTime *wrappers.StringValue `protobuf:"bytes,4,opt,name=start_date_time,json=startDateTime,proto3" json:"start_date_time,omitempty"`
	// End time in which this feed item is no longer effective and will stop
	// The format is "YYYY-MM-DD HH:MM:SS".
	// Examples: "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	EndDateTime *wrappers.StringValue `protobuf:"bytes,5,opt,name=end_date_time,json=endDateTime,proto3" json:"end_date_time,omitempty"`
	// The feed item's attribute values.
	AttributeValues []*FeedItemAttributeValue `protobuf:"bytes,6,rep,name=attribute_values,json=attributeValues,proto3" json:"attribute_values,omitempty"`
	// Geo targeting restriction specifies the type of location that can be used
	// for targeting.
	GeoTargetingRestriction enums.GeoTargetingRestrictionEnum_GeoTargetingRestriction `protobuf:"varint,7,opt,name=geo_targeting_restriction,json=geoTargetingRestriction,proto3,enum=google.ads.googleads.v0.enums.GeoTargetingRestrictionEnum_GeoTargetingRestriction" json:"geo_targeting_restriction,omitempty"`
	// The list of mappings used to substitute custom parameter tags in a
	// `tracking_url_template`, `final_urls`, or `mobile_final_urls`.
	UrlCustomParameters []*common.CustomParameter `protobuf:"bytes,8,rep,name=url_custom_parameters,json=urlCustomParameters,proto3" json:"url_custom_parameters,omitempty"`
	// Status of the feed item.
	// This field is read-only.
	Status               enums.FeedItemStatusEnum_FeedItemStatus `protobuf:"varint,9,opt,name=status,proto3,enum=google.ads.googleads.v0.enums.FeedItemStatusEnum_FeedItemStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *FeedItem) Reset()         { *m = FeedItem{} }
func (m *FeedItem) String() string { return proto.CompactTextString(m) }
func (*FeedItem) ProtoMessage()    {}
func (*FeedItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_208b54b3cb63ad39, []int{0}
}

func (m *FeedItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItem.Unmarshal(m, b)
}
func (m *FeedItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItem.Marshal(b, m, deterministic)
}
func (m *FeedItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItem.Merge(m, src)
}
func (m *FeedItem) XXX_Size() int {
	return xxx_messageInfo_FeedItem.Size(m)
}
func (m *FeedItem) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItem.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItem proto.InternalMessageInfo

func (m *FeedItem) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *FeedItem) GetFeed() *wrappers.StringValue {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *FeedItem) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *FeedItem) GetStartDateTime() *wrappers.StringValue {
	if m != nil {
		return m.StartDateTime
	}
	return nil
}

func (m *FeedItem) GetEndDateTime() *wrappers.StringValue {
	if m != nil {
		return m.EndDateTime
	}
	return nil
}

func (m *FeedItem) GetAttributeValues() []*FeedItemAttributeValue {
	if m != nil {
		return m.AttributeValues
	}
	return nil
}

func (m *FeedItem) GetGeoTargetingRestriction() enums.GeoTargetingRestrictionEnum_GeoTargetingRestriction {
	if m != nil {
		return m.GeoTargetingRestriction
	}
	return enums.GeoTargetingRestrictionEnum_UNSPECIFIED
}

func (m *FeedItem) GetUrlCustomParameters() []*common.CustomParameter {
	if m != nil {
		return m.UrlCustomParameters
	}
	return nil
}

func (m *FeedItem) GetStatus() enums.FeedItemStatusEnum_FeedItemStatus {
	if m != nil {
		return m.Status
	}
	return enums.FeedItemStatusEnum_UNSPECIFIED
}

// A feed item attribute value.
type FeedItemAttributeValue struct {
	// Id of the feed attribute for which the value is associated with.
	FeedAttributeId *wrappers.Int64Value `protobuf:"bytes,1,opt,name=feed_attribute_id,json=feedAttributeId,proto3" json:"feed_attribute_id,omitempty"`
	// Int64 value. Should be set if feed_attribute_id refers to a feed attribute
	// of type INT64.
	IntegerValue *wrappers.Int64Value `protobuf:"bytes,2,opt,name=integer_value,json=integerValue,proto3" json:"integer_value,omitempty"`
	// Bool value. Should be set if feed_attribute_id refers to a feed attribute
	// of type BOOLEAN.
	BooleanValue *wrappers.BoolValue `protobuf:"bytes,3,opt,name=boolean_value,json=booleanValue,proto3" json:"boolean_value,omitempty"`
	// String value. Should be set if feed_attribute_id refers to a feed attribute
	// of type STRING, URL or DATE_TIME.
	// For STRING the maximum length is 1500 characters. For URL the maximum
	// length is 2076 characters. For DATE_TIME the format of the string must
	// be the same as start and end time for the feed item.
	StringValue *wrappers.StringValue `protobuf:"bytes,4,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	// Double value. Should be set if feed_attribute_id refers to a feed attribute
	// of type DOUBLE.
	DoubleValue *wrappers.DoubleValue `protobuf:"bytes,5,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	// Price value. Should be set if feed_attribute_id refers to a feed attribute
	// of type PRICE.
	PriceValue *common.Price `protobuf:"bytes,6,opt,name=price_value,json=priceValue,proto3" json:"price_value,omitempty"`
	// Repeated int64 value. Should be set if feed_attribute_id refers to a feed
	// attribute of type INT64_LIST.
	IntegerValues []*wrappers.Int64Value `protobuf:"bytes,7,rep,name=integer_values,json=integerValues,proto3" json:"integer_values,omitempty"`
	// Repeated bool value. Should be set if feed_attribute_id refers to a feed
	// attribute of type BOOLEAN_LIST.
	BooleanValues []*wrappers.BoolValue `protobuf:"bytes,8,rep,name=boolean_values,json=booleanValues,proto3" json:"boolean_values,omitempty"`
	// Repeated string value. Should be set if feed_attribute_id refers to a feed
	// attribute of type STRING_LIST, URL_LIST or DATE_TIME_LIST.
	// For STRING_LIST and URL_LIST the total size of the list in bytes may not
	// exceed 3000. For DATE_TIME_LIST the number of elements may not exceed 200.
	//
	// For STRING_LIST the maximum length of each string element is 1500
	// characters. For URL_LIST the maximum length is 2076 characters. For
	// DATE_TIME the format of the string must be the same as start and end time
	// for the feed item.
	StringValues []*wrappers.StringValue `protobuf:"bytes,9,rep,name=string_values,json=stringValues,proto3" json:"string_values,omitempty"`
	// Repeated double value. Should be set if feed_attribute_id refers to a feed
	// attribute of type DOUBLE_LIST.
	DoubleValues         []*wrappers.DoubleValue `protobuf:"bytes,10,rep,name=double_values,json=doubleValues,proto3" json:"double_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FeedItemAttributeValue) Reset()         { *m = FeedItemAttributeValue{} }
func (m *FeedItemAttributeValue) String() string { return proto.CompactTextString(m) }
func (*FeedItemAttributeValue) ProtoMessage()    {}
func (*FeedItemAttributeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_208b54b3cb63ad39, []int{1}
}

func (m *FeedItemAttributeValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemAttributeValue.Unmarshal(m, b)
}
func (m *FeedItemAttributeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemAttributeValue.Marshal(b, m, deterministic)
}
func (m *FeedItemAttributeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemAttributeValue.Merge(m, src)
}
func (m *FeedItemAttributeValue) XXX_Size() int {
	return xxx_messageInfo_FeedItemAttributeValue.Size(m)
}
func (m *FeedItemAttributeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemAttributeValue.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemAttributeValue proto.InternalMessageInfo

func (m *FeedItemAttributeValue) GetFeedAttributeId() *wrappers.Int64Value {
	if m != nil {
		return m.FeedAttributeId
	}
	return nil
}

func (m *FeedItemAttributeValue) GetIntegerValue() *wrappers.Int64Value {
	if m != nil {
		return m.IntegerValue
	}
	return nil
}

func (m *FeedItemAttributeValue) GetBooleanValue() *wrappers.BoolValue {
	if m != nil {
		return m.BooleanValue
	}
	return nil
}

func (m *FeedItemAttributeValue) GetStringValue() *wrappers.StringValue {
	if m != nil {
		return m.StringValue
	}
	return nil
}

func (m *FeedItemAttributeValue) GetDoubleValue() *wrappers.DoubleValue {
	if m != nil {
		return m.DoubleValue
	}
	return nil
}

func (m *FeedItemAttributeValue) GetPriceValue() *common.Price {
	if m != nil {
		return m.PriceValue
	}
	return nil
}

func (m *FeedItemAttributeValue) GetIntegerValues() []*wrappers.Int64Value {
	if m != nil {
		return m.IntegerValues
	}
	return nil
}

func (m *FeedItemAttributeValue) GetBooleanValues() []*wrappers.BoolValue {
	if m != nil {
		return m.BooleanValues
	}
	return nil
}

func (m *FeedItemAttributeValue) GetStringValues() []*wrappers.StringValue {
	if m != nil {
		return m.StringValues
	}
	return nil
}

func (m *FeedItemAttributeValue) GetDoubleValues() []*wrappers.DoubleValue {
	if m != nil {
		return m.DoubleValues
	}
	return nil
}

func init() {
	proto.RegisterType((*FeedItem)(nil), "google.ads.googleads.v0.resources.FeedItem")
	proto.RegisterType((*FeedItemAttributeValue)(nil), "google.ads.googleads.v0.resources.FeedItemAttributeValue")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/feed_item.proto", fileDescriptor_208b54b3cb63ad39)
}

var fileDescriptor_208b54b3cb63ad39 = []byte{
	// 722 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xd1, 0x6a, 0xdb, 0x3a,
	0x18, 0xc7, 0x49, 0xd2, 0xa6, 0xad, 0x12, 0xb7, 0xe7, 0xf8, 0x70, 0xce, 0xf1, 0xba, 0x31, 0xd2,
	0x8e, 0x42, 0x60, 0x60, 0x67, 0x5d, 0x37, 0x18, 0x63, 0xac, 0xc9, 0xba, 0x86, 0xee, 0x62, 0x04,
	0xb7, 0x94, 0x31, 0x02, 0x46, 0x89, 0xbe, 0x1a, 0x81, 0x6d, 0x05, 0x49, 0xee, 0xde, 0x62, 0x0f,
	0xb1, 0xcb, 0xbd, 0xc2, 0xde, 0x60, 0xec, 0x6a, 0x4f, 0x34, 0x2c, 0x4b, 0x8e, 0xd3, 0xd6, 0x49,
	0xef, 0x24, 0xfb, 0xfb, 0xfd, 0xbf, 0x4f, 0xff, 0xfc, 0x15, 0xa3, 0x67, 0x21, 0x63, 0x61, 0x04,
	0x1e, 0x26, 0xc2, 0xcb, 0x97, 0xd9, 0xea, 0xba, 0xe7, 0x71, 0x10, 0x2c, 0xe5, 0x53, 0x10, 0xde,
	0x15, 0x00, 0x09, 0xa8, 0x84, 0xd8, 0x9d, 0x71, 0x26, 0x99, 0xbd, 0x97, 0xd7, 0xb9, 0x98, 0x08,
	0xb7, 0x40, 0xdc, 0xeb, 0x9e, 0x5b, 0x20, 0xbb, 0x2f, 0xaa, 0x54, 0xa7, 0x2c, 0x8e, 0x59, 0xe2,
	0x4d, 0x53, 0x21, 0x59, 0x1c, 0xcc, 0x30, 0xc7, 0x31, 0x48, 0xe0, 0xb9, 0xf2, 0x6e, 0x6f, 0x05,
	0xa6, 0x26, 0xc9, 0xd7, 0x9a, 0x38, 0xaa, 0x22, 0x20, 0x49, 0xe3, 0xd2, 0xe8, 0x81, 0x90, 0x58,
	0xa6, 0x42, 0x53, 0x6f, 0x96, 0x53, 0x21, 0xb0, 0x40, 0x62, 0x1e, 0x82, 0xa4, 0x49, 0x18, 0x70,
	0x10, 0x92, 0xd3, 0xa9, 0xa4, 0x45, 0xd3, 0xc7, 0x1a, 0x57, 0xbb, 0x49, 0x7a, 0xe5, 0x7d, 0xe1,
	0x78, 0x36, 0x03, 0xae, 0xe5, 0xf7, 0x7f, 0xad, 0xa3, 0xcd, 0x53, 0x00, 0x72, 0x26, 0x21, 0xb6,
	0x9f, 0x20, 0xcb, 0xf8, 0x12, 0x24, 0x38, 0x06, 0xa7, 0xd6, 0xa9, 0x75, 0xb7, 0xfc, 0xb6, 0x79,
	0xf8, 0x11, 0xc7, 0x60, 0xf7, 0xd0, 0x5a, 0x36, 0xaa, 0x53, 0xef, 0xd4, 0xba, 0xad, 0xc3, 0x47,
	0xda, 0x56, 0xd7, 0x34, 0x70, 0xcf, 0x25, 0xa7, 0x49, 0x78, 0x89, 0xa3, 0x14, 0x7c, 0x55, 0x69,
	0x3f, 0x45, 0x75, 0x4a, 0x9c, 0x86, 0xaa, 0x7f, 0x78, 0xab, 0xfe, 0x2c, 0x91, 0x2f, 0x8f, 0xf2,
	0xf2, 0x3a, 0x25, 0xf6, 0x09, 0xda, 0x11, 0x12, 0x73, 0x19, 0x10, 0x2c, 0x21, 0x90, 0x34, 0x06,
	0x67, 0xed, 0x1e, 0x9d, 0x2c, 0x05, 0x9d, 0x60, 0x09, 0x17, 0x34, 0x06, 0xfb, 0x18, 0x59, 0x90,
	0x90, 0x92, 0xc6, 0xfa, 0x3d, 0x34, 0x5a, 0x90, 0x90, 0x42, 0x81, 0xa0, 0xbf, 0xb0, 0x94, 0x9c,
	0x4e, 0x52, 0x09, 0xc1, 0x75, 0xf6, 0x5e, 0x38, 0xcd, 0x4e, 0xa3, 0xdb, 0x3a, 0x7c, 0xe5, 0xae,
	0x0c, 0x95, 0x6b, 0x2c, 0xed, 0x1b, 0x89, 0xbc, 0xc3, 0x0e, 0x5e, 0xd8, 0x0b, 0xfb, 0x6b, 0x0d,
	0x3d, 0xa8, 0xfc, 0x09, 0x9d, 0x8d, 0x4e, 0xad, 0xbb, 0x7d, 0xe8, 0x57, 0xf6, 0x53, 0x11, 0x70,
	0x87, 0xc0, 0x2e, 0x0c, 0xee, 0xcf, 0xe9, 0xf7, 0x49, 0x1a, 0x57, 0xbd, 0xf3, 0xff, 0x0f, 0xef,
	0x7e, 0x61, 0x4f, 0xd1, 0xbf, 0x29, 0x8f, 0x82, 0x9b, 0xa1, 0x17, 0xce, 0xa6, 0x3a, 0xbb, 0x57,
	0x39, 0x8b, 0x8e, 0xfa, 0x3b, 0x05, 0x8e, 0x0c, 0xe7, 0xff, 0x93, 0xf2, 0xe8, 0xc6, 0x33, 0x61,
	0x7f, 0x42, 0xcd, 0x3c, 0xe3, 0xce, 0x96, 0x3a, 0xe1, 0xf1, 0x8a, 0x13, 0x1a, 0x37, 0xcf, 0x15,
	0xa4, 0x0e, 0xb6, 0xf8, 0xc8, 0xd7, 0x7a, 0xfb, 0x3f, 0xd6, 0xd1, 0x7f, 0x77, 0x7b, 0x6f, 0x0f,
	0xd1, 0xdf, 0xea, 0x8a, 0xcd, 0x7f, 0x55, 0x4a, 0x54, 0xc0, 0x57, 0x84, 0x72, 0x27, 0xa3, 0x0a,
	0xad, 0x33, 0x92, 0x65, 0x8b, 0x26, 0x12, 0x42, 0xe0, 0x79, 0x2e, 0xf4, 0x4d, 0x58, 0x2a, 0xd2,
	0xd6, 0x44, 0x3e, 0xca, 0x5b, 0x64, 0x4d, 0x18, 0x8b, 0x00, 0x27, 0x5a, 0x21, 0xbf, 0x1b, 0xbb,
	0xb7, 0x14, 0x06, 0x8c, 0x45, 0x5a, 0x40, 0x03, 0x46, 0xa0, 0x2d, 0x54, 0x70, 0x35, 0x7f, 0x9f,
	0x1b, 0xd2, 0x12, 0xf3, 0x4d, 0x26, 0x40, 0x58, 0x3a, 0x89, 0x74, 0xb4, 0x2b, 0xaf, 0xc7, 0x89,
	0x2a, 0xd2, 0x02, 0x64, 0xbe, 0xb1, 0x4f, 0x51, 0x6b, 0xc6, 0xe9, 0xd4, 0xf0, 0x4d, 0xc5, 0x1f,
	0xac, 0x4a, 0xc7, 0x28, 0x43, 0x7c, 0xa4, 0xc8, 0x5c, 0x67, 0x80, 0xb6, 0x17, 0xcc, 0x14, 0xce,
	0x86, 0x0a, 0xda, 0x52, 0x37, 0xad, 0xb2, 0x9b, 0xc2, 0xee, 0xa3, 0xed, 0x05, 0x3b, 0x4d, 0x58,
	0x97, 0xf9, 0x69, 0x95, 0xfd, 0xcc, 0x24, 0xac, 0xb2, 0xa1, 0x59, 0x30, 0x1b, 0x2b, 0x1d, 0x6d,
	0x97, 0x1c, 0x55, 0x12, 0x65, 0x4b, 0x85, 0x83, 0x2a, 0x24, 0xca, 0x9e, 0xb6, 0x4b, 0x9e, 0x8a,
	0xc1, 0xef, 0x1a, 0x3a, 0x98, 0xb2, 0x78, 0xf5, 0xff, 0xcb, 0xc0, 0x32, 0x21, 0x1f, 0x65, 0xaa,
	0xa3, 0xda, 0xe7, 0x0f, 0x9a, 0x09, 0x59, 0x84, 0x93, 0xd0, 0x65, 0x3c, 0xf4, 0x42, 0x48, 0x54,
	0x4f, 0xf3, 0xd9, 0x98, 0x51, 0xb1, 0xe4, 0xd3, 0xf9, 0xba, 0x58, 0x7d, 0xab, 0x37, 0x86, 0xfd,
	0xfe, 0xf7, 0xfa, 0xde, 0x30, 0x97, 0xec, 0x13, 0xe1, 0xe6, 0xcb, 0x6c, 0x75, 0xd9, 0x73, 0x7d,
	0x53, 0xf9, 0xd3, 0xd4, 0x8c, 0xfb, 0x44, 0x8c, 0x8b, 0x9a, 0xf1, 0x65, 0x6f, 0x5c, 0xd4, 0x4c,
	0x9a, 0x6a, 0x88, 0xe7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x75, 0x42, 0x94, 0x5b, 0xbe, 0x07,
	0x00, 0x00,
}
