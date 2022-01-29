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
// source: google/ads/googleads/v9/services/bidding_data_exclusion_service.proto

package services

import (
	enums "github.com/butlerhq/google-ads-go/enums"
	resources "github.com/butlerhq/google-ads-go/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for
// [BiddingDataExclusionService.GetBiddingDataExclusion][google.ads.googleads.v9.services.BiddingDataExclusionService.GetBiddingDataExclusion].
type GetBiddingDataExclusionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The resource name of the data exclusion to fetch.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
}

func (x *GetBiddingDataExclusionRequest) Reset() {
	*x = GetBiddingDataExclusionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_bidding_data_exclusion_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBiddingDataExclusionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBiddingDataExclusionRequest) ProtoMessage() {}

func (x *GetBiddingDataExclusionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_bidding_data_exclusion_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBiddingDataExclusionRequest.ProtoReflect.Descriptor instead.
func (*GetBiddingDataExclusionRequest) Descriptor() ([]byte, []int) {
	return file_services_bidding_data_exclusion_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetBiddingDataExclusionRequest) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

// Request message for
// [BiddingDataExclusionService.MutateBiddingDataExclusions][google.ads.googleads.v9.services.BiddingDataExclusionService.MutateBiddingDataExclusions].
type MutateBiddingDataExclusionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the customer whose data exclusions are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual data exclusions.
	Operations []*BiddingDataExclusionOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly bool `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// The response content type setting. Determines whether the mutable resource
	// or just the resource name should be returned post mutation.
	ResponseContentType enums.ResponseContentTypeEnum_ResponseContentType `protobuf:"varint,5,opt,name=response_content_type,json=responseContentType,proto3,enum=google.ads.googleads.v9.enums.ResponseContentTypeEnum_ResponseContentType" json:"response_content_type,omitempty"`
}

func (x *MutateBiddingDataExclusionsRequest) Reset() {
	*x = MutateBiddingDataExclusionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_bidding_data_exclusion_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateBiddingDataExclusionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateBiddingDataExclusionsRequest) ProtoMessage() {}

func (x *MutateBiddingDataExclusionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_bidding_data_exclusion_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateBiddingDataExclusionsRequest.ProtoReflect.Descriptor instead.
func (*MutateBiddingDataExclusionsRequest) Descriptor() ([]byte, []int) {
	return file_services_bidding_data_exclusion_service_proto_rawDescGZIP(), []int{1}
}

func (x *MutateBiddingDataExclusionsRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *MutateBiddingDataExclusionsRequest) GetOperations() []*BiddingDataExclusionOperation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *MutateBiddingDataExclusionsRequest) GetPartialFailure() bool {
	if x != nil {
		return x.PartialFailure
	}
	return false
}

func (x *MutateBiddingDataExclusionsRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *MutateBiddingDataExclusionsRequest) GetResponseContentType() enums.ResponseContentTypeEnum_ResponseContentType {
	if x != nil {
		return x.ResponseContentType
	}
	return enums.ResponseContentTypeEnum_ResponseContentType(0)
}

// A single operation (create, remove, update) on a data exclusion.
type BiddingDataExclusionOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are assignable to Operation:
	//	*BiddingDataExclusionOperation_Create
	//	*BiddingDataExclusionOperation_Update
	//	*BiddingDataExclusionOperation_Remove
	Operation isBiddingDataExclusionOperation_Operation `protobuf_oneof:"operation"`
}

func (x *BiddingDataExclusionOperation) Reset() {
	*x = BiddingDataExclusionOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_bidding_data_exclusion_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BiddingDataExclusionOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BiddingDataExclusionOperation) ProtoMessage() {}

func (x *BiddingDataExclusionOperation) ProtoReflect() protoreflect.Message {
	mi := &file_services_bidding_data_exclusion_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BiddingDataExclusionOperation.ProtoReflect.Descriptor instead.
func (*BiddingDataExclusionOperation) Descriptor() ([]byte, []int) {
	return file_services_bidding_data_exclusion_service_proto_rawDescGZIP(), []int{2}
}

func (x *BiddingDataExclusionOperation) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (m *BiddingDataExclusionOperation) GetOperation() isBiddingDataExclusionOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (x *BiddingDataExclusionOperation) GetCreate() *resources.BiddingDataExclusion {
	if x, ok := x.GetOperation().(*BiddingDataExclusionOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (x *BiddingDataExclusionOperation) GetUpdate() *resources.BiddingDataExclusion {
	if x, ok := x.GetOperation().(*BiddingDataExclusionOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (x *BiddingDataExclusionOperation) GetRemove() string {
	if x, ok := x.GetOperation().(*BiddingDataExclusionOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

type isBiddingDataExclusionOperation_Operation interface {
	isBiddingDataExclusionOperation_Operation()
}

type BiddingDataExclusionOperation_Create struct {
	// Create operation: No resource name is expected for the new data
	// exclusion.
	Create *resources.BiddingDataExclusion `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type BiddingDataExclusionOperation_Update struct {
	// Update operation: The data exclusion is expected to have a valid
	// resource name.
	Update *resources.BiddingDataExclusion `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type BiddingDataExclusionOperation_Remove struct {
	// Remove operation: A resource name for the removed data exclusion
	// is expected, in this format:
	//
	// `customers/{customer_id}/biddingDataExclusions/{data_exclusion_id}`
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*BiddingDataExclusionOperation_Create) isBiddingDataExclusionOperation_Operation() {}

func (*BiddingDataExclusionOperation_Update) isBiddingDataExclusionOperation_Operation() {}

func (*BiddingDataExclusionOperation_Remove) isBiddingDataExclusionOperation_Operation() {}

// Response message for data exlusions mutate.
type MutateBiddingDataExclusionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results []*MutateBiddingDataExclusionsResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *MutateBiddingDataExclusionsResponse) Reset() {
	*x = MutateBiddingDataExclusionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_bidding_data_exclusion_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateBiddingDataExclusionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateBiddingDataExclusionsResponse) ProtoMessage() {}

func (x *MutateBiddingDataExclusionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_bidding_data_exclusion_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateBiddingDataExclusionsResponse.ProtoReflect.Descriptor instead.
func (*MutateBiddingDataExclusionsResponse) Descriptor() ([]byte, []int) {
	return file_services_bidding_data_exclusion_service_proto_rawDescGZIP(), []int{3}
}

func (x *MutateBiddingDataExclusionsResponse) GetPartialFailureError() *status.Status {
	if x != nil {
		return x.PartialFailureError
	}
	return nil
}

func (x *MutateBiddingDataExclusionsResponse) GetResults() []*MutateBiddingDataExclusionsResult {
	if x != nil {
		return x.Results
	}
	return nil
}

// The result for the data exclusion mutate.
type MutateBiddingDataExclusionsResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned for successful operations.
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The mutated bidding data exclusion with only mutable fields after mutate.
	// The field will only be returned when response_content_type is set
	// to "MUTABLE_RESOURCE".
	BiddingDataExclusion *resources.BiddingDataExclusion `protobuf:"bytes,2,opt,name=bidding_data_exclusion,json=biddingDataExclusion,proto3" json:"bidding_data_exclusion,omitempty"`
}

func (x *MutateBiddingDataExclusionsResult) Reset() {
	*x = MutateBiddingDataExclusionsResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_bidding_data_exclusion_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutateBiddingDataExclusionsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutateBiddingDataExclusionsResult) ProtoMessage() {}

func (x *MutateBiddingDataExclusionsResult) ProtoReflect() protoreflect.Message {
	mi := &file_services_bidding_data_exclusion_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutateBiddingDataExclusionsResult.ProtoReflect.Descriptor instead.
func (*MutateBiddingDataExclusionsResult) Descriptor() ([]byte, []int) {
	return file_services_bidding_data_exclusion_service_proto_rawDescGZIP(), []int{4}
}

func (x *MutateBiddingDataExclusionsResult) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *MutateBiddingDataExclusionsResult) GetBiddingDataExclusion() *resources.BiddingDataExclusion {
	if x != nil {
		return x.BiddingDataExclusion
	}
	return nil
}

var File_services_bidding_data_exclusion_service_proto protoreflect.FileDescriptor

var file_services_bidding_data_exclusion_service_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f,
	0x76, 0x39, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x5a, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0xe0, 0x41, 0x02,
	0xfa, 0x41, 0x2f, 0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x69,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xfe, 0x02, 0x0a, 0x22, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x42, 0x69, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x64, 0x0a,
	0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c,
	0x79, 0x12, 0x7e, 0x0a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x4a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x13, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x22, 0xa9, 0x02, 0x0a, 0x1d, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x51, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x51, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xcc, 0x01,
	0x0a, 0x23, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c,
	0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x13, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x5d, 0x0a,
	0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0xb7, 0x01, 0x0a,
	0x21, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x6d, 0x0a, 0x16, 0x62, 0x69, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76,
	0x39, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x64, 0x64,
	0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x14, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63,
	0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0xd9, 0x04, 0x0a, 0x1b, 0x42, 0x69, 0x64, 0x64, 0x69,
	0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xe5, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x42, 0x69,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x39, 0x12, 0x37, 0x2f, 0x76, 0x39, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x2f, 0x2a, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x8a,
	0x02, 0x0a, 0x1b, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x44,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x42, 0x69,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x3f, 0x22, 0x3a, 0x2f, 0x76, 0x39, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3d,
	0x2a, 0x7d, 0x2f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78,
	0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x6d, 0x75, 0x74, 0x61, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0xda, 0x41, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x2c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x45, 0xca, 0x41, 0x18,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x64, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x42, 0x87, 0x02, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x39, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x20, 0x42, 0x69, 0x64,
	0x64, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa,
	0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x39, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x39, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0xea, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x39, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_bidding_data_exclusion_service_proto_rawDescOnce sync.Once
	file_services_bidding_data_exclusion_service_proto_rawDescData = file_services_bidding_data_exclusion_service_proto_rawDesc
)

func file_services_bidding_data_exclusion_service_proto_rawDescGZIP() []byte {
	file_services_bidding_data_exclusion_service_proto_rawDescOnce.Do(func() {
		file_services_bidding_data_exclusion_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_bidding_data_exclusion_service_proto_rawDescData)
	})
	return file_services_bidding_data_exclusion_service_proto_rawDescData
}

var file_services_bidding_data_exclusion_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_services_bidding_data_exclusion_service_proto_goTypes = []interface{}{
	(*GetBiddingDataExclusionRequest)(nil),                 // 0: google.ads.googleads.v9.services.GetBiddingDataExclusionRequest
	(*MutateBiddingDataExclusionsRequest)(nil),             // 1: google.ads.googleads.v9.services.MutateBiddingDataExclusionsRequest
	(*BiddingDataExclusionOperation)(nil),                  // 2: google.ads.googleads.v9.services.BiddingDataExclusionOperation
	(*MutateBiddingDataExclusionsResponse)(nil),            // 3: google.ads.googleads.v9.services.MutateBiddingDataExclusionsResponse
	(*MutateBiddingDataExclusionsResult)(nil),              // 4: google.ads.googleads.v9.services.MutateBiddingDataExclusionsResult
	(enums.ResponseContentTypeEnum_ResponseContentType)(0), // 5: google.ads.googleads.v9.enums.ResponseContentTypeEnum.ResponseContentType
	(*fieldmaskpb.FieldMask)(nil),                          // 6: google.protobuf.FieldMask
	(*resources.BiddingDataExclusion)(nil),                 // 7: google.ads.googleads.v9.resources.BiddingDataExclusion
	(*status.Status)(nil),                                  // 8: google.rpc.Status
}
var file_services_bidding_data_exclusion_service_proto_depIdxs = []int32{
	2,  // 0: google.ads.googleads.v9.services.MutateBiddingDataExclusionsRequest.operations:type_name -> google.ads.googleads.v9.services.BiddingDataExclusionOperation
	5,  // 1: google.ads.googleads.v9.services.MutateBiddingDataExclusionsRequest.response_content_type:type_name -> google.ads.googleads.v9.enums.ResponseContentTypeEnum.ResponseContentType
	6,  // 2: google.ads.googleads.v9.services.BiddingDataExclusionOperation.update_mask:type_name -> google.protobuf.FieldMask
	7,  // 3: google.ads.googleads.v9.services.BiddingDataExclusionOperation.create:type_name -> google.ads.googleads.v9.resources.BiddingDataExclusion
	7,  // 4: google.ads.googleads.v9.services.BiddingDataExclusionOperation.update:type_name -> google.ads.googleads.v9.resources.BiddingDataExclusion
	8,  // 5: google.ads.googleads.v9.services.MutateBiddingDataExclusionsResponse.partial_failure_error:type_name -> google.rpc.Status
	4,  // 6: google.ads.googleads.v9.services.MutateBiddingDataExclusionsResponse.results:type_name -> google.ads.googleads.v9.services.MutateBiddingDataExclusionsResult
	7,  // 7: google.ads.googleads.v9.services.MutateBiddingDataExclusionsResult.bidding_data_exclusion:type_name -> google.ads.googleads.v9.resources.BiddingDataExclusion
	0,  // 8: google.ads.googleads.v9.services.BiddingDataExclusionService.GetBiddingDataExclusion:input_type -> google.ads.googleads.v9.services.GetBiddingDataExclusionRequest
	1,  // 9: google.ads.googleads.v9.services.BiddingDataExclusionService.MutateBiddingDataExclusions:input_type -> google.ads.googleads.v9.services.MutateBiddingDataExclusionsRequest
	7,  // 10: google.ads.googleads.v9.services.BiddingDataExclusionService.GetBiddingDataExclusion:output_type -> google.ads.googleads.v9.resources.BiddingDataExclusion
	3,  // 11: google.ads.googleads.v9.services.BiddingDataExclusionService.MutateBiddingDataExclusions:output_type -> google.ads.googleads.v9.services.MutateBiddingDataExclusionsResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_services_bidding_data_exclusion_service_proto_init() }
func file_services_bidding_data_exclusion_service_proto_init() {
	if File_services_bidding_data_exclusion_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_bidding_data_exclusion_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBiddingDataExclusionRequest); i {
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
		file_services_bidding_data_exclusion_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateBiddingDataExclusionsRequest); i {
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
		file_services_bidding_data_exclusion_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BiddingDataExclusionOperation); i {
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
		file_services_bidding_data_exclusion_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateBiddingDataExclusionsResponse); i {
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
		file_services_bidding_data_exclusion_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutateBiddingDataExclusionsResult); i {
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
	file_services_bidding_data_exclusion_service_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*BiddingDataExclusionOperation_Create)(nil),
		(*BiddingDataExclusionOperation_Update)(nil),
		(*BiddingDataExclusionOperation_Remove)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_bidding_data_exclusion_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_bidding_data_exclusion_service_proto_goTypes,
		DependencyIndexes: file_services_bidding_data_exclusion_service_proto_depIdxs,
		MessageInfos:      file_services_bidding_data_exclusion_service_proto_msgTypes,
	}.Build()
	File_services_bidding_data_exclusion_service_proto = out.File
	file_services_bidding_data_exclusion_service_proto_rawDesc = nil
	file_services_bidding_data_exclusion_service_proto_goTypes = nil
	file_services_bidding_data_exclusion_service_proto_depIdxs = nil
}
