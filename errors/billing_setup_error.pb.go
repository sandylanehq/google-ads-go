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
// source: google/ads/googleads/v9/errors/billing_setup_error.proto

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

// Enum describing possible billing setup errors.
type BillingSetupErrorEnum_BillingSetupError int32

const (
	// Enum unspecified.
	BillingSetupErrorEnum_UNSPECIFIED BillingSetupErrorEnum_BillingSetupError = 0
	// The received error code is not known in this version.
	BillingSetupErrorEnum_UNKNOWN BillingSetupErrorEnum_BillingSetupError = 1
	// Cannot specify both an existing payments account and a new payments
	// account when setting up billing.
	BillingSetupErrorEnum_CANNOT_USE_EXISTING_AND_NEW_ACCOUNT BillingSetupErrorEnum_BillingSetupError = 2
	// Cannot cancel an approved billing setup whose start time has passed.
	BillingSetupErrorEnum_CANNOT_REMOVE_STARTED_BILLING_SETUP BillingSetupErrorEnum_BillingSetupError = 3
	// Cannot perform a Change of Bill-To (CBT) to the same payments account.
	BillingSetupErrorEnum_CANNOT_CHANGE_BILLING_TO_SAME_PAYMENTS_ACCOUNT BillingSetupErrorEnum_BillingSetupError = 4
	// Billing setups can only be used by customers with ENABLED or DRAFT
	// status.
	BillingSetupErrorEnum_BILLING_SETUP_NOT_PERMITTED_FOR_CUSTOMER_STATUS BillingSetupErrorEnum_BillingSetupError = 5
	// Billing setups must either include a correctly formatted existing
	// payments account id, or a non-empty new payments account name.
	BillingSetupErrorEnum_INVALID_PAYMENTS_ACCOUNT BillingSetupErrorEnum_BillingSetupError = 6
	// Only billable and third-party customers can create billing setups.
	BillingSetupErrorEnum_BILLING_SETUP_NOT_PERMITTED_FOR_CUSTOMER_CATEGORY BillingSetupErrorEnum_BillingSetupError = 7
	// Billing setup creations can only use NOW for start time type.
	BillingSetupErrorEnum_INVALID_START_TIME_TYPE BillingSetupErrorEnum_BillingSetupError = 8
	// Billing setups can only be created for a third-party customer if they do
	// not already have a setup.
	BillingSetupErrorEnum_THIRD_PARTY_ALREADY_HAS_BILLING BillingSetupErrorEnum_BillingSetupError = 9
	// Billing setups cannot be created if there is already a pending billing in
	// progress.
	BillingSetupErrorEnum_BILLING_SETUP_IN_PROGRESS BillingSetupErrorEnum_BillingSetupError = 10
	// Billing setups can only be created by customers who have permission to
	// setup billings. Users can contact a representative for help setting up
	// permissions.
	BillingSetupErrorEnum_NO_SIGNUP_PERMISSION BillingSetupErrorEnum_BillingSetupError = 11
	// Billing setups cannot be created if there is already a future-approved
	// billing.
	BillingSetupErrorEnum_CHANGE_OF_BILL_TO_IN_PROGRESS BillingSetupErrorEnum_BillingSetupError = 12
	// Requested payments profile not found.
	BillingSetupErrorEnum_PAYMENTS_PROFILE_NOT_FOUND BillingSetupErrorEnum_BillingSetupError = 13
	// Requested payments account not found.
	BillingSetupErrorEnum_PAYMENTS_ACCOUNT_NOT_FOUND BillingSetupErrorEnum_BillingSetupError = 14
	// Billing setup creation failed because the payments profile is ineligible.
	BillingSetupErrorEnum_PAYMENTS_PROFILE_INELIGIBLE BillingSetupErrorEnum_BillingSetupError = 15
	// Billing setup creation failed because the payments account is ineligible.
	BillingSetupErrorEnum_PAYMENTS_ACCOUNT_INELIGIBLE BillingSetupErrorEnum_BillingSetupError = 16
	// Billing setup creation failed because the payments profile needs internal
	// approval.
	BillingSetupErrorEnum_CUSTOMER_NEEDS_INTERNAL_APPROVAL BillingSetupErrorEnum_BillingSetupError = 17
	// Payments account has different currency code than the current customer
	// and hence cannot be used to setup billing.
	BillingSetupErrorEnum_PAYMENTS_ACCOUNT_INELIGIBLE_CURRENCY_CODE_MISMATCH BillingSetupErrorEnum_BillingSetupError = 19
	// A start time in the future cannot be used because there is currently no
	// active billing setup for this customer.
	BillingSetupErrorEnum_FUTURE_START_TIME_PROHIBITED BillingSetupErrorEnum_BillingSetupError = 20
)

// Enum value maps for BillingSetupErrorEnum_BillingSetupError.
var (
	BillingSetupErrorEnum_BillingSetupError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "CANNOT_USE_EXISTING_AND_NEW_ACCOUNT",
		3:  "CANNOT_REMOVE_STARTED_BILLING_SETUP",
		4:  "CANNOT_CHANGE_BILLING_TO_SAME_PAYMENTS_ACCOUNT",
		5:  "BILLING_SETUP_NOT_PERMITTED_FOR_CUSTOMER_STATUS",
		6:  "INVALID_PAYMENTS_ACCOUNT",
		7:  "BILLING_SETUP_NOT_PERMITTED_FOR_CUSTOMER_CATEGORY",
		8:  "INVALID_START_TIME_TYPE",
		9:  "THIRD_PARTY_ALREADY_HAS_BILLING",
		10: "BILLING_SETUP_IN_PROGRESS",
		11: "NO_SIGNUP_PERMISSION",
		12: "CHANGE_OF_BILL_TO_IN_PROGRESS",
		13: "PAYMENTS_PROFILE_NOT_FOUND",
		14: "PAYMENTS_ACCOUNT_NOT_FOUND",
		15: "PAYMENTS_PROFILE_INELIGIBLE",
		16: "PAYMENTS_ACCOUNT_INELIGIBLE",
		17: "CUSTOMER_NEEDS_INTERNAL_APPROVAL",
		19: "PAYMENTS_ACCOUNT_INELIGIBLE_CURRENCY_CODE_MISMATCH",
		20: "FUTURE_START_TIME_PROHIBITED",
	}
	BillingSetupErrorEnum_BillingSetupError_value = map[string]int32{
		"UNSPECIFIED":                                        0,
		"UNKNOWN":                                            1,
		"CANNOT_USE_EXISTING_AND_NEW_ACCOUNT":                2,
		"CANNOT_REMOVE_STARTED_BILLING_SETUP":                3,
		"CANNOT_CHANGE_BILLING_TO_SAME_PAYMENTS_ACCOUNT":     4,
		"BILLING_SETUP_NOT_PERMITTED_FOR_CUSTOMER_STATUS":    5,
		"INVALID_PAYMENTS_ACCOUNT":                           6,
		"BILLING_SETUP_NOT_PERMITTED_FOR_CUSTOMER_CATEGORY":  7,
		"INVALID_START_TIME_TYPE":                            8,
		"THIRD_PARTY_ALREADY_HAS_BILLING":                    9,
		"BILLING_SETUP_IN_PROGRESS":                          10,
		"NO_SIGNUP_PERMISSION":                               11,
		"CHANGE_OF_BILL_TO_IN_PROGRESS":                      12,
		"PAYMENTS_PROFILE_NOT_FOUND":                         13,
		"PAYMENTS_ACCOUNT_NOT_FOUND":                         14,
		"PAYMENTS_PROFILE_INELIGIBLE":                        15,
		"PAYMENTS_ACCOUNT_INELIGIBLE":                        16,
		"CUSTOMER_NEEDS_INTERNAL_APPROVAL":                   17,
		"PAYMENTS_ACCOUNT_INELIGIBLE_CURRENCY_CODE_MISMATCH": 19,
		"FUTURE_START_TIME_PROHIBITED":                       20,
	}
)

func (x BillingSetupErrorEnum_BillingSetupError) Enum() *BillingSetupErrorEnum_BillingSetupError {
	p := new(BillingSetupErrorEnum_BillingSetupError)
	*p = x
	return p
}

func (x BillingSetupErrorEnum_BillingSetupError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BillingSetupErrorEnum_BillingSetupError) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_billing_setup_error_proto_enumTypes[0].Descriptor()
}

func (BillingSetupErrorEnum_BillingSetupError) Type() protoreflect.EnumType {
	return &file_errors_billing_setup_error_proto_enumTypes[0]
}

func (x BillingSetupErrorEnum_BillingSetupError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BillingSetupErrorEnum_BillingSetupError.Descriptor instead.
func (BillingSetupErrorEnum_BillingSetupError) EnumDescriptor() ([]byte, []int) {
	return file_errors_billing_setup_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible billing setup errors.
type BillingSetupErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BillingSetupErrorEnum) Reset() {
	*x = BillingSetupErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_billing_setup_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingSetupErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingSetupErrorEnum) ProtoMessage() {}

func (x *BillingSetupErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_errors_billing_setup_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingSetupErrorEnum.ProtoReflect.Descriptor instead.
func (*BillingSetupErrorEnum) Descriptor() ([]byte, []int) {
	return file_errors_billing_setup_error_proto_rawDescGZIP(), []int{0}
}

var File_errors_billing_setup_error_proto protoreflect.FileDescriptor

var file_errors_billing_setup_error_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x39, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x05, 0x0a, 0x15, 0x42, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e,
	0x75, 0x6d, 0x22, 0xe1, 0x05, 0x0a, 0x11, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x74, 0x75, 0x70, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54,
	0x5f, 0x55, 0x53, 0x45, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x4e,
	0x44, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x02, 0x12,
	0x27, 0x0a, 0x23, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45,
	0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x42, 0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47,
	0x5f, 0x53, 0x45, 0x54, 0x55, 0x50, 0x10, 0x03, 0x12, 0x32, 0x0a, 0x2e, 0x43, 0x41, 0x4e, 0x4e,
	0x4f, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x42, 0x49, 0x4c, 0x4c, 0x49, 0x4e,
	0x47, 0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x41, 0x4d, 0x45, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e,
	0x54, 0x53, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x33, 0x0a, 0x2f,
	0x42, 0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x54, 0x55, 0x50, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x5f,
	0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10,
	0x05, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x59,
	0x4d, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x06, 0x12,
	0x35, 0x0a, 0x31, 0x42, 0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x54, 0x55, 0x50,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x54, 0x54, 0x45, 0x44, 0x5f, 0x46,
	0x4f, 0x52, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x5f, 0x43, 0x41, 0x54, 0x45,
	0x47, 0x4f, 0x52, 0x59, 0x10, 0x07, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x10, 0x08, 0x12, 0x23, 0x0a, 0x1f, 0x54, 0x48, 0x49, 0x52, 0x44, 0x5f, 0x50, 0x41, 0x52,
	0x54, 0x59, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x42,
	0x49, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x09, 0x12, 0x1d, 0x0a, 0x19, 0x42, 0x49, 0x4c, 0x4c,
	0x49, 0x4e, 0x47, 0x5f, 0x53, 0x45, 0x54, 0x55, 0x50, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f,
	0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x0a, 0x12, 0x18, 0x0a, 0x14, 0x4e, 0x4f, 0x5f, 0x53, 0x49,
	0x47, 0x4e, 0x55, 0x50, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x10,
	0x0b, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x4f, 0x46, 0x5f, 0x42,
	0x49, 0x4c, 0x4c, 0x5f, 0x54, 0x4f, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45,
	0x53, 0x53, 0x10, 0x0c, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x53,
	0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x0d, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x53,
	0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x0e, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x53,
	0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x49, 0x4e, 0x45, 0x4c, 0x49, 0x47, 0x49,
	0x42, 0x4c, 0x45, 0x10, 0x0f, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54,
	0x53, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x49, 0x4e, 0x45, 0x4c, 0x49, 0x47,
	0x49, 0x42, 0x4c, 0x45, 0x10, 0x10, 0x12, 0x24, 0x0a, 0x20, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d,
	0x45, 0x52, 0x5f, 0x4e, 0x45, 0x45, 0x44, 0x53, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x41, 0x4c, 0x10, 0x11, 0x12, 0x36, 0x0a, 0x32,
	0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x53, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54,
	0x5f, 0x49, 0x4e, 0x45, 0x4c, 0x49, 0x47, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x43, 0x55, 0x52, 0x52,
	0x45, 0x4e, 0x43, 0x59, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x4d, 0x49, 0x53, 0x4d, 0x41, 0x54,
	0x43, 0x48, 0x10, 0x13, 0x12, 0x20, 0x0a, 0x1c, 0x46, 0x55, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x53,
	0x54, 0x41, 0x52, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x48, 0x49, 0x42,
	0x49, 0x54, 0x45, 0x44, 0x10, 0x14, 0x42, 0xf1, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x76, 0x39, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x16, 0x42,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x74, 0x75, 0x70, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61,
	0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x39, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03,
	0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73,
	0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56, 0x39, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64,
	0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x5c, 0x56, 0x39, 0x5c, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a,
	0x56, 0x39, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_errors_billing_setup_error_proto_rawDescOnce sync.Once
	file_errors_billing_setup_error_proto_rawDescData = file_errors_billing_setup_error_proto_rawDesc
)

func file_errors_billing_setup_error_proto_rawDescGZIP() []byte {
	file_errors_billing_setup_error_proto_rawDescOnce.Do(func() {
		file_errors_billing_setup_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_billing_setup_error_proto_rawDescData)
	})
	return file_errors_billing_setup_error_proto_rawDescData
}

var file_errors_billing_setup_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_billing_setup_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_billing_setup_error_proto_goTypes = []interface{}{
	(BillingSetupErrorEnum_BillingSetupError)(0), // 0: google.ads.googleads.v9.errors.BillingSetupErrorEnum.BillingSetupError
	(*BillingSetupErrorEnum)(nil),                // 1: google.ads.googleads.v9.errors.BillingSetupErrorEnum
}
var file_errors_billing_setup_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_billing_setup_error_proto_init() }
func file_errors_billing_setup_error_proto_init() {
	if File_errors_billing_setup_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_billing_setup_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingSetupErrorEnum); i {
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
			RawDescriptor: file_errors_billing_setup_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_billing_setup_error_proto_goTypes,
		DependencyIndexes: file_errors_billing_setup_error_proto_depIdxs,
		EnumInfos:         file_errors_billing_setup_error_proto_enumTypes,
		MessageInfos:      file_errors_billing_setup_error_proto_msgTypes,
	}.Build()
	File_errors_billing_setup_error_proto = out.File
	file_errors_billing_setup_error_proto_rawDesc = nil
	file_errors_billing_setup_error_proto_goTypes = nil
	file_errors_billing_setup_error_proto_depIdxs = nil
}
