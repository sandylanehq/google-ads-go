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
// source: google/ads/googleads/v9/errors/keyword_plan_campaign_error.proto

package errors

import (
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

// Enum describing possible errors from applying a keyword plan campaign.
type KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError int32

const (
	// Enum unspecified.
	KeywordPlanCampaignErrorEnum_UNSPECIFIED KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 0
	// The received error code is not known in this version.
	KeywordPlanCampaignErrorEnum_UNKNOWN KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 1
	// A keyword plan campaign name is missing, empty, longer than allowed limit
	// or contains invalid chars.
	KeywordPlanCampaignErrorEnum_INVALID_NAME KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 2
	// A keyword plan campaign contains one or more untargetable languages.
	KeywordPlanCampaignErrorEnum_INVALID_LANGUAGES KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 3
	// A keyword plan campaign contains one or more invalid geo targets.
	KeywordPlanCampaignErrorEnum_INVALID_GEOS KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 4
	// The keyword plan campaign name is duplicate to an existing keyword plan
	// campaign name or other keyword plan campaign name in the request.
	KeywordPlanCampaignErrorEnum_DUPLICATE_NAME KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 5
	// The number of geo targets in the keyword plan campaign exceeds limits.
	KeywordPlanCampaignErrorEnum_MAX_GEOS_EXCEEDED KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 6
	// The number of languages in the keyword plan campaign exceeds limits.
	KeywordPlanCampaignErrorEnum_MAX_LANGUAGES_EXCEEDED KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 7
)

// Enum value maps for KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError.
var (
	KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "INVALID_NAME",
		3: "INVALID_LANGUAGES",
		4: "INVALID_GEOS",
		5: "DUPLICATE_NAME",
		6: "MAX_GEOS_EXCEEDED",
		7: "MAX_LANGUAGES_EXCEEDED",
	}
	KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_value = map[string]int32{
		"UNSPECIFIED":            0,
		"UNKNOWN":                1,
		"INVALID_NAME":           2,
		"INVALID_LANGUAGES":      3,
		"INVALID_GEOS":           4,
		"DUPLICATE_NAME":         5,
		"MAX_GEOS_EXCEEDED":      6,
		"MAX_LANGUAGES_EXCEEDED": 7,
	}
)

func (x KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError) Enum() *KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError {
	p := new(KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError)
	*p = x
	return p
}

func (x KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_keyword_plan_campaign_error_proto_enumTypes[0].Descriptor()
}

func (KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError) Type() protoreflect.EnumType {
	return &file_errors_keyword_plan_campaign_error_proto_enumTypes[0]
}

func (x KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError.Descriptor instead.
func (KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError) EnumDescriptor() ([]byte, []int) {
	return file_errors_keyword_plan_campaign_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible errors from applying a keyword plan
// campaign.
type KeywordPlanCampaignErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *KeywordPlanCampaignErrorEnum) Reset() {
	*x = KeywordPlanCampaignErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_keyword_plan_campaign_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeywordPlanCampaignErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordPlanCampaignErrorEnum) ProtoMessage() {}

func (x *KeywordPlanCampaignErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_keyword_plan_campaign_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordPlanCampaignErrorEnum.ProtoReflect.Descriptor instead.
func (*KeywordPlanCampaignErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_keyword_plan_campaign_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_keyword_plan_campaign_error_proto protoreflect.FileDescriptor

var file_errors_keyword_plan_campaign_error_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x63, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdb, 0x01, 0x0a, 0x1c, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61, 0x6e,
	0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75,
	0x6d, 0x22, 0xba, 0x01, 0x0a, 0x18, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c, 0x61,
	0x6e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x02, 0x12, 0x15,
	0x0a, 0x11, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41,
	0x47, 0x45, 0x53, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x47, 0x45, 0x4f, 0x53, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x55, 0x50, 0x4c, 0x49,
	0x43, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x4d,
	0x41, 0x58, 0x5f, 0x47, 0x45, 0x4f, 0x53, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44,
	0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x41, 0x58, 0x5f, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41,
	0x47, 0x45, 0x53, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x07, 0x42, 0xf8,
	0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x1d, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x6c,
	0x61, 0x6e, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47,
	0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x39, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73,
	0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x39, 0x5c, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41,
	0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x39, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_errors_keyword_plan_campaign_error_proto_rawDescOnce sync.Once
	file_errors_keyword_plan_campaign_error_proto_rawDescData = file_errors_keyword_plan_campaign_error_proto_rawDesc
)

func file_errors_keyword_plan_campaign_error_proto_rawDescGZIP() []byte {
	file_errors_keyword_plan_campaign_error_proto_rawDescOnce.Do(func() {
		file_errors_keyword_plan_campaign_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_keyword_plan_campaign_error_proto_rawDescData)
	})
	return file_errors_keyword_plan_campaign_error_proto_rawDescData
}

var file_errors_keyword_plan_campaign_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_keyword_plan_campaign_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_keyword_plan_campaign_error_proto_goTypes = []interface{}{
	(KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError)(0), // 0: google.ads.googleads.v9.errors.KeywordPlanCampaignErrorEnum.KeywordPlanCampaignError
	(*KeywordPlanCampaignErrorEnum)(nil),                       // 1: google.ads.googleads.v9.errors.KeywordPlanCampaignErrorEnum
}
var file_errors_keyword_plan_campaign_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_keyword_plan_campaign_error_proto_init() }
func file_errors_keyword_plan_campaign_error_proto_init() {
	if File_errors_keyword_plan_campaign_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_keyword_plan_campaign_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeywordPlanCampaignErrorEnum); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errors_keyword_plan_campaign_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_keyword_plan_campaign_error_proto_goTypes,
		DependencyIndexes: file_errors_keyword_plan_campaign_error_proto_depIdxs,
		EnumInfos:         file_errors_keyword_plan_campaign_error_proto_enumTypes,
		MessageInfos:      file_errors_keyword_plan_campaign_error_proto_msgTypes,
	}.Build()
	File_errors_keyword_plan_campaign_error_proto = out.File
	file_errors_keyword_plan_campaign_error_proto_rawDesc = nil
	file_errors_keyword_plan_campaign_error_proto_goTypes = nil
	file_errors_keyword_plan_campaign_error_proto_depIdxs = nil
}
