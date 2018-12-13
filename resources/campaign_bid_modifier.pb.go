// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/campaign_bid_modifier.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/kritzware/google-ads-go/common"
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

// Represents a bid-modifiable only criterion at the campaign level.
type CampaignBidModifier struct {
	// The resource name of the campaign bid modifier.
	// Campaign bid modifier resource names have the form:
	//
	// `customers/{customer_id}/campaignBidModifiers/{campaign_id}_{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The campaign to which this criterion belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,2,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The ID of the criterion to bid modify.
	//
	// This field is ignored for mutates.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The modifier for the bid when the criterion matches.
	BidModifier *wrappers.DoubleValue `protobuf:"bytes,4,opt,name=bid_modifier,json=bidModifier,proto3" json:"bid_modifier,omitempty"`
	// The criterion of this campaign bid modifier.
	//
	// Types that are valid to be assigned to Criterion:
	//	*CampaignBidModifier_InteractionType
	Criterion            isCampaignBidModifier_Criterion `protobuf_oneof:"criterion"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *CampaignBidModifier) Reset()         { *m = CampaignBidModifier{} }
func (m *CampaignBidModifier) String() string { return proto.CompactTextString(m) }
func (*CampaignBidModifier) ProtoMessage()    {}
func (*CampaignBidModifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7b70b68df78a4e1, []int{0}
}

func (m *CampaignBidModifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignBidModifier.Unmarshal(m, b)
}
func (m *CampaignBidModifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignBidModifier.Marshal(b, m, deterministic)
}
func (m *CampaignBidModifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignBidModifier.Merge(m, src)
}
func (m *CampaignBidModifier) XXX_Size() int {
	return xxx_messageInfo_CampaignBidModifier.Size(m)
}
func (m *CampaignBidModifier) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignBidModifier.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignBidModifier proto.InternalMessageInfo

func (m *CampaignBidModifier) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignBidModifier) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignBidModifier) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *CampaignBidModifier) GetBidModifier() *wrappers.DoubleValue {
	if m != nil {
		return m.BidModifier
	}
	return nil
}

type isCampaignBidModifier_Criterion interface {
	isCampaignBidModifier_Criterion()
}

type CampaignBidModifier_InteractionType struct {
	InteractionType *common.InteractionTypeInfo `protobuf:"bytes,5,opt,name=interaction_type,json=interactionType,proto3,oneof"`
}

func (*CampaignBidModifier_InteractionType) isCampaignBidModifier_Criterion() {}

func (m *CampaignBidModifier) GetCriterion() isCampaignBidModifier_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (m *CampaignBidModifier) GetInteractionType() *common.InteractionTypeInfo {
	if x, ok := m.GetCriterion().(*CampaignBidModifier_InteractionType); ok {
		return x.InteractionType
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignBidModifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignBidModifier_InteractionType)(nil),
	}
}

func init() {
	proto.RegisterType((*CampaignBidModifier)(nil), "google.ads.googleads.v0.resources.CampaignBidModifier")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/campaign_bid_modifier.proto", fileDescriptor_f7b70b68df78a4e1)
}

var fileDescriptor_f7b70b68df78a4e1 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xd1, 0xea, 0xd3, 0x30,
	0x14, 0xc6, 0x6d, 0xa7, 0xe2, 0xb2, 0x89, 0x52, 0x6f, 0xca, 0x14, 0xd9, 0x14, 0x61, 0x37, 0xa6,
	0xc5, 0x89, 0x08, 0xa2, 0xd2, 0x2a, 0xcc, 0x09, 0xca, 0xa8, 0xb2, 0x0b, 0x29, 0xd4, 0xb4, 0xc9,
	0x42, 0xa0, 0x4d, 0x4a, 0xd2, 0x4e, 0xf6, 0x3a, 0x5e, 0xfa, 0x28, 0xbe, 0x83, 0x77, 0x3e, 0x88,
	0xb4, 0x69, 0xc2, 0x44, 0xf7, 0xff, 0xdf, 0x7d, 0x4d, 0xbf, 0xdf, 0x39, 0xdf, 0xc9, 0x09, 0x78,
	0x49, 0x85, 0xa0, 0x25, 0x09, 0x10, 0x56, 0x81, 0x96, 0x9d, 0x3a, 0x84, 0x81, 0x24, 0x4a, 0xb4,
	0xb2, 0x20, 0x2a, 0x28, 0x50, 0x55, 0x23, 0x46, 0x79, 0x96, 0x33, 0x9c, 0x55, 0x02, 0xb3, 0x3d,
	0x23, 0x12, 0xd6, 0x52, 0x34, 0xc2, 0x5b, 0x68, 0x06, 0x22, 0xac, 0xa0, 0xc5, 0xe1, 0x21, 0x84,
	0x16, 0x9f, 0x3d, 0x3e, 0xd7, 0xa1, 0x10, 0x55, 0x25, 0x78, 0x50, 0x48, 0xd6, 0x10, 0xc9, 0x90,
	0xae, 0x38, 0xbb, 0x3f, 0xd8, 0xfb, 0xaf, 0xbc, 0xdd, 0x07, 0xdf, 0x24, 0xaa, 0x6b, 0x22, 0x95,
	0xfe, 0xff, 0xe0, 0xb7, 0x0b, 0xee, 0xbc, 0x19, 0x12, 0xc5, 0x0c, 0x7f, 0x18, 0xf2, 0x78, 0x0f,
	0xc1, 0x4d, 0xd3, 0x33, 0xe3, 0xa8, 0x22, 0xbe, 0x33, 0x77, 0x96, 0xe3, 0x64, 0x6a, 0x0e, 0x3f,
	0xa2, 0x8a, 0x78, 0xcf, 0xc1, 0x0d, 0x33, 0x8d, 0xef, 0xce, 0x9d, 0xe5, 0xe4, 0xc9, 0xbd, 0x21,
	0x36, 0x34, 0xfd, 0xe0, 0xa7, 0x46, 0x32, 0x4e, 0x77, 0xa8, 0x6c, 0x49, 0x62, 0xdd, 0xde, 0x2b,
	0x30, 0x1d, 0x82, 0x0a, 0x9e, 0x31, 0xec, 0x8f, 0x7a, 0xfa, 0xee, 0x3f, 0xf4, 0x86, 0x37, 0xcf,
	0x9e, 0x6a, 0x78, 0x62, 0x81, 0x0d, 0xf6, 0x5e, 0x83, 0xe9, 0xe9, 0xf5, 0xf9, 0x57, 0xcf, 0x74,
	0x7f, 0x2b, 0xda, 0xbc, 0x24, 0x43, 0x81, 0xfc, 0x64, 0xbe, 0xaf, 0xe0, 0x36, 0xe3, 0x0d, 0x91,
	0xa8, 0x68, 0xba, 0x08, 0xcd, 0xb1, 0x26, 0xfe, 0xb5, 0xbe, 0xc8, 0x0a, 0x9e, 0x5b, 0x82, 0xbe,
	0xe1, 0x2e, 0x93, 0xe1, 0x3e, 0x1f, 0x6b, 0xb2, 0xe1, 0x7b, 0xf1, 0xee, 0x4a, 0x72, 0x8b, 0xfd,
	0x7d, 0x1c, 0x4f, 0xc0, 0xd8, 0x26, 0x8e, 0x7f, 0x39, 0xe0, 0x51, 0x21, 0x2a, 0x78, 0xe9, 0x7e,
	0x63, 0xff, 0x3f, 0xdb, 0xd8, 0x76, 0xe3, 0x6c, 0x9d, 0x2f, 0xef, 0x07, 0x9c, 0x8a, 0x12, 0x71,
	0x0a, 0x85, 0xa4, 0x01, 0x25, 0xbc, 0x1f, 0xd6, 0xbc, 0x85, 0x9a, 0xa9, 0x0b, 0x1e, 0xdf, 0x0b,
	0xab, 0xbe, 0xbb, 0xa3, 0x75, 0x14, 0xfd, 0x70, 0x17, 0x6b, 0x5d, 0x32, 0xc2, 0x0a, 0x6a, 0xd9,
	0xa9, 0x5d, 0x08, 0x13, 0xe3, 0xfc, 0x69, 0x3c, 0x69, 0x84, 0x55, 0x6a, 0x3d, 0xe9, 0x2e, 0x4c,
	0xad, 0x27, 0xbf, 0xde, 0x87, 0x58, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xd9, 0xa7, 0x4a,
	0x00, 0x03, 0x00, 0x00,
}
