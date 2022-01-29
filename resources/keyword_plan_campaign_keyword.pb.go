// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: google/ads/googleads/v9/resources/keyword_plan_campaign_keyword.proto

package resources

import (
	enums "github.com/butlerhq/google-ads-go/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A Keyword Plan Campaign keyword.
// Only negative keywords are supported for Campaign Keyword.
type KeywordPlanCampaignKeyword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the Keyword Plan Campaign keyword.
	// KeywordPlanCampaignKeyword resource names have the form:
	//
	// `customers/{customer_id}/keywordPlanCampaignKeywords/{kp_campaign_keyword_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The Keyword Plan campaign to which this negative keyword belongs.
	KeywordPlanCampaign *string `protobuf:"bytes,8,opt,name=keyword_plan_campaign,json=keywordPlanCampaign,proto3,oneof" json:"keyword_plan_campaign,omitempty"`
	// Output only. The ID of the Keyword Plan negative keyword.
	Id *int64 `protobuf:"varint,9,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// The keyword text.
	Text *string `protobuf:"bytes,10,opt,name=text,proto3,oneof" json:"text,omitempty"`
	// The keyword match type.
	MatchType enums.KeywordMatchTypeEnum_KeywordMatchType `protobuf:"varint,5,opt,name=match_type,json=matchType,proto3,enum=google.ads.googleads.v9.enums.KeywordMatchTypeEnum_KeywordMatchType" json:"match_type,omitempty"`
	// Immutable. If true, the keyword is negative.
	// Must be set to true. Only negative campaign keywords are supported.
	Negative *bool `protobuf:"varint,11,opt,name=negative,proto3,oneof" json:"negative,omitempty"`
}

func (x *KeywordPlanCampaignKeyword) Reset() {
	*x = KeywordPlanCampaignKeyword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_keyword_plan_campaign_keyword_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeywordPlanCampaignKeyword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordPlanCampaignKeyword) ProtoMessage() {}

func (x *KeywordPlanCampaignKeyword) ProtoReflect() protoreflect.Message {
	mi := &file_resources_keyword_plan_campaign_keyword_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordPlanCampaignKeyword.ProtoReflect.Descriptor instead.
func (*KeywordPlanCampaignKeyword) Descriptor() ([]byte, []int) {
	return file_resources_keyword_plan_campaign_keyword_proto_rawDescGZIP(), []int{0}
}

func (x *KeywordPlanCampaignKeyword) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *KeywordPlanCampaignKeyword) GetKeywordPlanCampaign() string {
	if x != nil && x.KeywordPlanCampaign != nil {
		return *x.KeywordPlanCampaign
	}
	return ""
}

func (x *KeywordPlanCampaignKeyword) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *KeywordPlanCampaignKeyword) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

func (x *KeywordPlanCampaignKeyword) GetMatchType() enums.KeywordMatchTypeEnum_KeywordMatchType {
	if x != nil {
		return x.MatchType
	}
	return enums.KeywordMatchTypeEnum_KeywordMatchType(0)
}

func (x *KeywordPlanCampaignKeyword) GetNegative() bool {
	if x != nil && x.Negative != nil {
		return *x.Negative
	}
	return false
}

var File_resources_keyword_plan_campaign_keyword_proto protoreflect.FileDescriptor

var file_resources_keyword_plan_campaign_keyword_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e,
	0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x36, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2f, 0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x04, 0x0a,
	0x1a, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x61, 0x6d, 0x70,
	0x61, 0x69, 0x67, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x60, 0x0a, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x3b, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x35, 0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e,
	0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x6a, 0x0a,
	0x15, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x63, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x31, 0xfa, 0x41,
	0x2e, 0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x48,
	0x00, 0x52, 0x13, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x48, 0x01, 0x52, 0x02, 0x69, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x12, 0x63, 0x0a, 0x0a,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x44, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x24, 0x0a, 0x08, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x48, 0x03, 0x52, 0x08, 0x6e, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x3a, 0x91, 0x01, 0xea, 0x41, 0x8d, 0x01, 0x0a, 0x33,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x50, 0x6c, 0x61, 0x6e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x4b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x56, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e,
	0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x18, 0x0a, 0x16, 0x5f,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x63, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x74, 0x65, 0x78, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x42, 0x8c, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x1f, 0x4b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41,
	0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x39, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41,
	0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x39, 0x5c,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea, 0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x56, 0x39, 0x3a, 0x3a, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_keyword_plan_campaign_keyword_proto_rawDescOnce sync.Once
	file_resources_keyword_plan_campaign_keyword_proto_rawDescData = file_resources_keyword_plan_campaign_keyword_proto_rawDesc
)

func file_resources_keyword_plan_campaign_keyword_proto_rawDescGZIP() []byte {
	file_resources_keyword_plan_campaign_keyword_proto_rawDescOnce.Do(func() {
		file_resources_keyword_plan_campaign_keyword_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_keyword_plan_campaign_keyword_proto_rawDescData)
	})
	return file_resources_keyword_plan_campaign_keyword_proto_rawDescData
}

var file_resources_keyword_plan_campaign_keyword_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_resources_keyword_plan_campaign_keyword_proto_goTypes = []interface{}{
	(*KeywordPlanCampaignKeyword)(nil),               // 0: google.ads.googleads.v9.resources.KeywordPlanCampaignKeyword
	(enums.KeywordMatchTypeEnum_KeywordMatchType)(0), // 1: google.ads.googleads.v9.enums.KeywordMatchTypeEnum.KeywordMatchType
}
var file_resources_keyword_plan_campaign_keyword_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v9.resources.KeywordPlanCampaignKeyword.match_type:type_name -> google.ads.googleads.v9.enums.KeywordMatchTypeEnum.KeywordMatchType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_resources_keyword_plan_campaign_keyword_proto_init() }
func file_resources_keyword_plan_campaign_keyword_proto_init() {
	if File_resources_keyword_plan_campaign_keyword_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_keyword_plan_campaign_keyword_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeywordPlanCampaignKeyword); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_resources_keyword_plan_campaign_keyword_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_keyword_plan_campaign_keyword_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_keyword_plan_campaign_keyword_proto_goTypes,
		DependencyIndexes: file_resources_keyword_plan_campaign_keyword_proto_depIdxs,
		MessageInfos:      file_resources_keyword_plan_campaign_keyword_proto_msgTypes,
	}.Build()
	File_resources_keyword_plan_campaign_keyword_proto = out.File
	file_resources_keyword_plan_campaign_keyword_proto_rawDesc = nil
	file_resources_keyword_plan_campaign_keyword_proto_goTypes = nil
	file_resources_keyword_plan_campaign_keyword_proto_depIdxs = nil
}
