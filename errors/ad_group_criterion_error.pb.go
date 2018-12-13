// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/ad_group_criterion_error.proto

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

// Enum describing possible ad group criterion errors.
type AdGroupCriterionErrorEnum_AdGroupCriterionError int32

const (
	// Enum unspecified.
	AdGroupCriterionErrorEnum_UNSPECIFIED AdGroupCriterionErrorEnum_AdGroupCriterionError = 0
	// The received error code is not known in this version.
	AdGroupCriterionErrorEnum_UNKNOWN AdGroupCriterionErrorEnum_AdGroupCriterionError = 1
	// No link found between the AdGroupCriterion and the label.
	AdGroupCriterionErrorEnum_AD_GROUP_CRITERION_LABEL_DOES_NOT_EXIST AdGroupCriterionErrorEnum_AdGroupCriterionError = 2
	// The label has already been attached to the AdGroupCriterion.
	AdGroupCriterionErrorEnum_AD_GROUP_CRITERION_LABEL_ALREADY_EXISTS AdGroupCriterionErrorEnum_AdGroupCriterionError = 3
	// Negative AdGroupCriterion cannot have labels.
	AdGroupCriterionErrorEnum_CANNOT_ADD_LABEL_TO_NEGATIVE_CRITERION AdGroupCriterionErrorEnum_AdGroupCriterionError = 4
	// Too many operations for a single call.
	AdGroupCriterionErrorEnum_TOO_MANY_OPERATIONS AdGroupCriterionErrorEnum_AdGroupCriterionError = 5
	// Negative ad group criteria are not updateable.
	AdGroupCriterionErrorEnum_CANT_UPDATE_NEGATIVE AdGroupCriterionErrorEnum_AdGroupCriterionError = 6
	// Concrete type of criterion (keyword v.s. placement) is required for ADD
	// and SET operations.
	AdGroupCriterionErrorEnum_CONCRETE_TYPE_REQUIRED AdGroupCriterionErrorEnum_AdGroupCriterionError = 7
	// Bid is incompatible with ad group's bidding settings.
	AdGroupCriterionErrorEnum_BID_INCOMPATIBLE_WITH_ADGROUP AdGroupCriterionErrorEnum_AdGroupCriterionError = 8
	// Cannot target and exclude the same criterion at once.
	AdGroupCriterionErrorEnum_CANNOT_TARGET_AND_EXCLUDE AdGroupCriterionErrorEnum_AdGroupCriterionError = 9
	// The URL of a placement is invalid.
	AdGroupCriterionErrorEnum_ILLEGAL_URL AdGroupCriterionErrorEnum_AdGroupCriterionError = 10
	// Keyword text was invalid.
	AdGroupCriterionErrorEnum_INVALID_KEYWORD_TEXT AdGroupCriterionErrorEnum_AdGroupCriterionError = 11
	// Destination URL was invalid.
	AdGroupCriterionErrorEnum_INVALID_DESTINATION_URL AdGroupCriterionErrorEnum_AdGroupCriterionError = 12
	// The destination url must contain at least one tag (e.g. {lpurl})
	AdGroupCriterionErrorEnum_MISSING_DESTINATION_URL_TAG AdGroupCriterionErrorEnum_AdGroupCriterionError = 13
	// Keyword-level cpm bid is not supported
	AdGroupCriterionErrorEnum_KEYWORD_LEVEL_BID_NOT_SUPPORTED_FOR_MANUALCPM AdGroupCriterionErrorEnum_AdGroupCriterionError = 14
	// For example, cannot add a biddable ad group criterion that had been
	// removed.
	AdGroupCriterionErrorEnum_INVALID_USER_STATUS AdGroupCriterionErrorEnum_AdGroupCriterionError = 15
	// Criteria type cannot be targeted for the ad group. Either the account is
	// restricted to keywords only, the criteria type is incompatible with the
	// campaign's bidding strategy, or the criteria type can only be applied to
	// campaigns.
	AdGroupCriterionErrorEnum_CANNOT_ADD_CRITERIA_TYPE AdGroupCriterionErrorEnum_AdGroupCriterionError = 16
	// Criteria type cannot be excluded for the ad group. Refer to the
	// documentation for a specific criterion to check if it is excludable.
	AdGroupCriterionErrorEnum_CANNOT_EXCLUDE_CRITERIA_TYPE AdGroupCriterionErrorEnum_AdGroupCriterionError = 17
	// Partial failure is not supported for shopping campaign mutate operations.
	AdGroupCriterionErrorEnum_CAMPAIGN_TYPE_NOT_COMPATIBLE_WITH_PARTIAL_FAILURE AdGroupCriterionErrorEnum_AdGroupCriterionError = 27
	// Operations in the mutate request changes too many shopping ad groups.
	// Please split requests for multiple shopping ad groups across multiple
	// requests.
	AdGroupCriterionErrorEnum_OPERATIONS_FOR_TOO_MANY_SHOPPING_ADGROUPS AdGroupCriterionErrorEnum_AdGroupCriterionError = 28
	// Not allowed to modify url fields of an ad group criterion if there are
	// duplicate elements for that ad group criterion in the request.
	AdGroupCriterionErrorEnum_CANNOT_MODIFY_URL_FIELDS_WITH_DUPLICATE_ELEMENTS AdGroupCriterionErrorEnum_AdGroupCriterionError = 29
	// Cannot set url fields without also setting final urls.
	AdGroupCriterionErrorEnum_CANNOT_SET_WITHOUT_FINAL_URLS AdGroupCriterionErrorEnum_AdGroupCriterionError = 30
	// Cannot clear final urls if final mobile urls exist.
	AdGroupCriterionErrorEnum_CANNOT_CLEAR_FINAL_URLS_IF_FINAL_MOBILE_URLS_EXIST AdGroupCriterionErrorEnum_AdGroupCriterionError = 31
	// Cannot clear final urls if final app urls exist.
	AdGroupCriterionErrorEnum_CANNOT_CLEAR_FINAL_URLS_IF_FINAL_APP_URLS_EXIST AdGroupCriterionErrorEnum_AdGroupCriterionError = 32
	// Cannot clear final urls if tracking url template exists.
	AdGroupCriterionErrorEnum_CANNOT_CLEAR_FINAL_URLS_IF_TRACKING_URL_TEMPLATE_EXISTS AdGroupCriterionErrorEnum_AdGroupCriterionError = 33
	// Cannot clear final urls if url custom parameters exist.
	AdGroupCriterionErrorEnum_CANNOT_CLEAR_FINAL_URLS_IF_URL_CUSTOM_PARAMETERS_EXIST AdGroupCriterionErrorEnum_AdGroupCriterionError = 34
	// Cannot set both destination url and final urls.
	AdGroupCriterionErrorEnum_CANNOT_SET_BOTH_DESTINATION_URL_AND_FINAL_URLS AdGroupCriterionErrorEnum_AdGroupCriterionError = 35
	// Cannot set both destination url and tracking url template.
	AdGroupCriterionErrorEnum_CANNOT_SET_BOTH_DESTINATION_URL_AND_TRACKING_URL_TEMPLATE AdGroupCriterionErrorEnum_AdGroupCriterionError = 36
	// Final urls are not supported for this criterion type.
	AdGroupCriterionErrorEnum_FINAL_URLS_NOT_SUPPORTED_FOR_CRITERION_TYPE AdGroupCriterionErrorEnum_AdGroupCriterionError = 37
	// Final mobile urls are not supported for this criterion type.
	AdGroupCriterionErrorEnum_FINAL_MOBILE_URLS_NOT_SUPPORTED_FOR_CRITERION_TYPE AdGroupCriterionErrorEnum_AdGroupCriterionError = 38
	// Ad group is invalid due to the listing groups it contains.
	AdGroupCriterionErrorEnum_INVALID_LISTING_GROUP_HIERARCHY AdGroupCriterionErrorEnum_AdGroupCriterionError = 39
	// Listing group unit cannot have children.
	AdGroupCriterionErrorEnum_LISTING_GROUP_UNIT_CANNOT_HAVE_CHILDREN AdGroupCriterionErrorEnum_AdGroupCriterionError = 40
	// Subdivided listing groups must have an "others" case.
	AdGroupCriterionErrorEnum_LISTING_GROUP_SUBDIVISION_REQUIRES_OTHERS_CASE AdGroupCriterionErrorEnum_AdGroupCriterionError = 41
	// Dimension type of listing group must be the same as that of its siblings.
	AdGroupCriterionErrorEnum_LISTING_GROUP_REQUIRES_SAME_DIMENSION_TYPE_AS_SIBLINGS AdGroupCriterionErrorEnum_AdGroupCriterionError = 42
	// Listing group cannot be added to the ad group because it already exists.
	AdGroupCriterionErrorEnum_LISTING_GROUP_ALREADY_EXISTS AdGroupCriterionErrorEnum_AdGroupCriterionError = 43
	// Listing group referenced in the operation was not found in the ad group.
	AdGroupCriterionErrorEnum_LISTING_GROUP_DOES_NOT_EXIST AdGroupCriterionErrorEnum_AdGroupCriterionError = 44
	// Recursive removal failed because listing group subdivision is being
	// created or modified in this request.
	AdGroupCriterionErrorEnum_LISTING_GROUP_CANNOT_BE_REMOVED AdGroupCriterionErrorEnum_AdGroupCriterionError = 45
	// Listing group type is not allowed for specified ad group criterion type.
	AdGroupCriterionErrorEnum_INVALID_LISTING_GROUP_TYPE AdGroupCriterionErrorEnum_AdGroupCriterionError = 46
	// Listing group in an ADD operation specifies a non temporary criterion id.
	AdGroupCriterionErrorEnum_LISTING_GROUP_ADD_MAY_ONLY_USE_TEMP_ID AdGroupCriterionErrorEnum_AdGroupCriterionError = 47
)

var AdGroupCriterionErrorEnum_AdGroupCriterionError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "AD_GROUP_CRITERION_LABEL_DOES_NOT_EXIST",
	3:  "AD_GROUP_CRITERION_LABEL_ALREADY_EXISTS",
	4:  "CANNOT_ADD_LABEL_TO_NEGATIVE_CRITERION",
	5:  "TOO_MANY_OPERATIONS",
	6:  "CANT_UPDATE_NEGATIVE",
	7:  "CONCRETE_TYPE_REQUIRED",
	8:  "BID_INCOMPATIBLE_WITH_ADGROUP",
	9:  "CANNOT_TARGET_AND_EXCLUDE",
	10: "ILLEGAL_URL",
	11: "INVALID_KEYWORD_TEXT",
	12: "INVALID_DESTINATION_URL",
	13: "MISSING_DESTINATION_URL_TAG",
	14: "KEYWORD_LEVEL_BID_NOT_SUPPORTED_FOR_MANUALCPM",
	15: "INVALID_USER_STATUS",
	16: "CANNOT_ADD_CRITERIA_TYPE",
	17: "CANNOT_EXCLUDE_CRITERIA_TYPE",
	27: "CAMPAIGN_TYPE_NOT_COMPATIBLE_WITH_PARTIAL_FAILURE",
	28: "OPERATIONS_FOR_TOO_MANY_SHOPPING_ADGROUPS",
	29: "CANNOT_MODIFY_URL_FIELDS_WITH_DUPLICATE_ELEMENTS",
	30: "CANNOT_SET_WITHOUT_FINAL_URLS",
	31: "CANNOT_CLEAR_FINAL_URLS_IF_FINAL_MOBILE_URLS_EXIST",
	32: "CANNOT_CLEAR_FINAL_URLS_IF_FINAL_APP_URLS_EXIST",
	33: "CANNOT_CLEAR_FINAL_URLS_IF_TRACKING_URL_TEMPLATE_EXISTS",
	34: "CANNOT_CLEAR_FINAL_URLS_IF_URL_CUSTOM_PARAMETERS_EXIST",
	35: "CANNOT_SET_BOTH_DESTINATION_URL_AND_FINAL_URLS",
	36: "CANNOT_SET_BOTH_DESTINATION_URL_AND_TRACKING_URL_TEMPLATE",
	37: "FINAL_URLS_NOT_SUPPORTED_FOR_CRITERION_TYPE",
	38: "FINAL_MOBILE_URLS_NOT_SUPPORTED_FOR_CRITERION_TYPE",
	39: "INVALID_LISTING_GROUP_HIERARCHY",
	40: "LISTING_GROUP_UNIT_CANNOT_HAVE_CHILDREN",
	41: "LISTING_GROUP_SUBDIVISION_REQUIRES_OTHERS_CASE",
	42: "LISTING_GROUP_REQUIRES_SAME_DIMENSION_TYPE_AS_SIBLINGS",
	43: "LISTING_GROUP_ALREADY_EXISTS",
	44: "LISTING_GROUP_DOES_NOT_EXIST",
	45: "LISTING_GROUP_CANNOT_BE_REMOVED",
	46: "INVALID_LISTING_GROUP_TYPE",
	47: "LISTING_GROUP_ADD_MAY_ONLY_USE_TEMP_ID",
}

var AdGroupCriterionErrorEnum_AdGroupCriterionError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"AD_GROUP_CRITERION_LABEL_DOES_NOT_EXIST":                   2,
	"AD_GROUP_CRITERION_LABEL_ALREADY_EXISTS":                   3,
	"CANNOT_ADD_LABEL_TO_NEGATIVE_CRITERION":                    4,
	"TOO_MANY_OPERATIONS":                                       5,
	"CANT_UPDATE_NEGATIVE":                                      6,
	"CONCRETE_TYPE_REQUIRED":                                    7,
	"BID_INCOMPATIBLE_WITH_ADGROUP":                             8,
	"CANNOT_TARGET_AND_EXCLUDE":                                 9,
	"ILLEGAL_URL":                                               10,
	"INVALID_KEYWORD_TEXT":                                      11,
	"INVALID_DESTINATION_URL":                                   12,
	"MISSING_DESTINATION_URL_TAG":                               13,
	"KEYWORD_LEVEL_BID_NOT_SUPPORTED_FOR_MANUALCPM":             14,
	"INVALID_USER_STATUS":                                       15,
	"CANNOT_ADD_CRITERIA_TYPE":                                  16,
	"CANNOT_EXCLUDE_CRITERIA_TYPE":                              17,
	"CAMPAIGN_TYPE_NOT_COMPATIBLE_WITH_PARTIAL_FAILURE":         27,
	"OPERATIONS_FOR_TOO_MANY_SHOPPING_ADGROUPS":                 28,
	"CANNOT_MODIFY_URL_FIELDS_WITH_DUPLICATE_ELEMENTS":          29,
	"CANNOT_SET_WITHOUT_FINAL_URLS":                             30,
	"CANNOT_CLEAR_FINAL_URLS_IF_FINAL_MOBILE_URLS_EXIST":        31,
	"CANNOT_CLEAR_FINAL_URLS_IF_FINAL_APP_URLS_EXIST":           32,
	"CANNOT_CLEAR_FINAL_URLS_IF_TRACKING_URL_TEMPLATE_EXISTS":   33,
	"CANNOT_CLEAR_FINAL_URLS_IF_URL_CUSTOM_PARAMETERS_EXIST":    34,
	"CANNOT_SET_BOTH_DESTINATION_URL_AND_FINAL_URLS":            35,
	"CANNOT_SET_BOTH_DESTINATION_URL_AND_TRACKING_URL_TEMPLATE": 36,
	"FINAL_URLS_NOT_SUPPORTED_FOR_CRITERION_TYPE":               37,
	"FINAL_MOBILE_URLS_NOT_SUPPORTED_FOR_CRITERION_TYPE":        38,
	"INVALID_LISTING_GROUP_HIERARCHY":                           39,
	"LISTING_GROUP_UNIT_CANNOT_HAVE_CHILDREN":                   40,
	"LISTING_GROUP_SUBDIVISION_REQUIRES_OTHERS_CASE":            41,
	"LISTING_GROUP_REQUIRES_SAME_DIMENSION_TYPE_AS_SIBLINGS":    42,
	"LISTING_GROUP_ALREADY_EXISTS":                              43,
	"LISTING_GROUP_DOES_NOT_EXIST":                              44,
	"LISTING_GROUP_CANNOT_BE_REMOVED":                           45,
	"INVALID_LISTING_GROUP_TYPE":                                46,
	"LISTING_GROUP_ADD_MAY_ONLY_USE_TEMP_ID":                    47,
}

func (x AdGroupCriterionErrorEnum_AdGroupCriterionError) String() string {
	return proto.EnumName(AdGroupCriterionErrorEnum_AdGroupCriterionError_name, int32(x))
}

func (AdGroupCriterionErrorEnum_AdGroupCriterionError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ab297178b75dc9da, []int{0, 0}
}

// Container for enum describing possible ad group criterion errors.
type AdGroupCriterionErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupCriterionErrorEnum) Reset()         { *m = AdGroupCriterionErrorEnum{} }
func (m *AdGroupCriterionErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterionErrorEnum) ProtoMessage()    {}
func (*AdGroupCriterionErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab297178b75dc9da, []int{0}
}

func (m *AdGroupCriterionErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterionErrorEnum.Unmarshal(m, b)
}
func (m *AdGroupCriterionErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterionErrorEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupCriterionErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterionErrorEnum.Merge(m, src)
}
func (m *AdGroupCriterionErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterionErrorEnum.Size(m)
}
func (m *AdGroupCriterionErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterionErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterionErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.errors.AdGroupCriterionErrorEnum_AdGroupCriterionError", AdGroupCriterionErrorEnum_AdGroupCriterionError_name, AdGroupCriterionErrorEnum_AdGroupCriterionError_value)
	proto.RegisterType((*AdGroupCriterionErrorEnum)(nil), "google.ads.googleads.v0.errors.AdGroupCriterionErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/ad_group_criterion_error.proto", fileDescriptor_ab297178b75dc9da)
}

var fileDescriptor_ab297178b75dc9da = []byte{
	// 936 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x72, 0x1b, 0x35,
	0x14, 0x26, 0x29, 0xb4, 0xa0, 0x04, 0x2a, 0xc4, 0x4f, 0xdb, 0xfc, 0xb6, 0x29, 0xb4, 0xb4, 0x21,
	0xeb, 0xb4, 0x85, 0x32, 0xd0, 0xe9, 0xc5, 0xd9, 0xd5, 0xf1, 0x5a, 0x13, 0xad, 0xb4, 0x48, 0x5a,
	0x27, 0x66, 0x32, 0xa3, 0x09, 0x75, 0xc6, 0x93, 0x99, 0x36, 0x9b, 0xb1, 0xdb, 0x3e, 0x10, 0x97,
	0xbc, 0x01, 0xaf, 0xc0, 0x15, 0x77, 0xbc, 0x0e, 0xa3, 0xdd, 0x75, 0x62, 0x3b, 0x21, 0xed, 0x55,
	0x14, 0x9f, 0xef, 0x93, 0xbe, 0xf3, 0x9d, 0xb3, 0xe7, 0x90, 0xe7, 0x83, 0xb2, 0x1c, 0xbc, 0x3c,
	0x6c, 0x1d, 0xf4, 0x47, 0xad, 0xfa, 0x18, 0x4e, 0x6f, 0xb7, 0x5b, 0x87, 0xc3, 0x61, 0x39, 0x1c,
	0xb5, 0x0e, 0xfa, 0x7e, 0x30, 0x2c, 0xdf, 0x9c, 0xf8, 0x17, 0xc3, 0xa3, 0xd7, 0x87, 0xc3, 0xa3,
	0xf2, 0xd8, 0x57, 0x91, 0xe8, 0x64, 0x58, 0xbe, 0x2e, 0xd9, 0x5a, 0xcd, 0x89, 0x0e, 0xfa, 0xa3,
	0xe8, 0x94, 0x1e, 0xbd, 0xdd, 0x8e, 0x6a, 0xfa, 0xc6, 0xbf, 0x8b, 0xe4, 0x16, 0xf4, 0xd3, 0x70,
	0x43, 0x32, 0xbe, 0x00, 0x43, 0x08, 0x8f, 0xdf, 0xbc, 0xda, 0xf8, 0x6b, 0x91, 0x7c, 0x75, 0x61,
	0x94, 0x5d, 0x27, 0x0b, 0x85, 0xb2, 0x39, 0x26, 0xa2, 0x2d, 0x90, 0xd3, 0x0f, 0xd8, 0x02, 0xb9,
	0x56, 0xa8, 0x1d, 0xa5, 0x77, 0x15, 0x9d, 0x63, 0x9b, 0xe4, 0x3e, 0x70, 0x9f, 0x1a, 0x5d, 0xe4,
	0x3e, 0x31, 0xc2, 0xa1, 0x11, 0x5a, 0x79, 0x09, 0x31, 0x4a, 0xcf, 0x35, 0x5a, 0xaf, 0xb4, 0xf3,
	0xb8, 0x27, 0xac, 0xa3, 0xf3, 0x97, 0x82, 0x41, 0x1a, 0x04, 0xde, 0xab, 0xb1, 0x96, 0x5e, 0x61,
	0x0f, 0xc9, 0xbd, 0x04, 0x54, 0xa0, 0x03, 0xe7, 0x0d, 0xc8, 0x69, 0xaf, 0x30, 0x05, 0x27, 0xba,
	0x78, 0x76, 0x01, 0xfd, 0x90, 0xdd, 0x20, 0x5f, 0x38, 0xad, 0x7d, 0x06, 0xaa, 0xe7, 0x75, 0x8e,
	0x06, 0x9c, 0xd0, 0xca, 0xd2, 0x8f, 0xd8, 0x4d, 0xf2, 0x65, 0x02, 0xca, 0xf9, 0x22, 0xe7, 0xe0,
	0xf0, 0x94, 0x4c, 0xaf, 0xb2, 0x25, 0xf2, 0x75, 0xa2, 0x55, 0x62, 0xd0, 0xa1, 0x77, 0xbd, 0x1c,
	0xbd, 0xc1, 0x5f, 0x0b, 0x61, 0x90, 0xd3, 0x6b, 0xec, 0x0e, 0x59, 0x8d, 0x05, 0xf7, 0x42, 0x25,
	0x3a, 0xcb, 0xc1, 0x89, 0x58, 0xa2, 0xdf, 0x15, 0xae, 0xe3, 0x81, 0x57, 0xe2, 0xe9, 0xc7, 0x6c,
	0x95, 0xdc, 0x6a, 0xd4, 0x39, 0x30, 0x29, 0x3a, 0x0f, 0x8a, 0x7b, 0xdc, 0x4b, 0x64, 0xc1, 0x91,
	0x7e, 0x12, 0x4c, 0x13, 0x52, 0x62, 0x0a, 0xd2, 0x17, 0x46, 0x52, 0x12, 0x84, 0x08, 0xd5, 0x05,
	0x29, 0xb8, 0xdf, 0xc1, 0xde, 0xae, 0x36, 0xdc, 0x3b, 0xdc, 0x73, 0x74, 0x81, 0x2d, 0x93, 0x1b,
	0xe3, 0x08, 0x47, 0xeb, 0x84, 0xaa, 0xc4, 0x57, 0xb4, 0x45, 0xb6, 0x4e, 0x96, 0x33, 0x61, 0xad,
	0x50, 0xe9, 0x6c, 0xd0, 0x3b, 0x48, 0xe9, 0xa7, 0xec, 0x11, 0xd9, 0x1a, 0xdf, 0x27, 0xb1, 0x8b,
	0xd2, 0x07, 0xe1, 0x41, 0x96, 0x2d, 0xf2, 0x5c, 0x1b, 0x87, 0xdc, 0xb7, 0xb5, 0x09, 0xce, 0x14,
	0x20, 0x93, 0x3c, 0xa3, 0x9f, 0x05, 0xb3, 0xc6, 0x0f, 0x16, 0x16, 0x8d, 0xb7, 0x0e, 0x5c, 0x61,
	0xe9, 0x75, 0xb6, 0x42, 0x6e, 0x4e, 0x38, 0xde, 0xf8, 0x0b, 0x95, 0x3b, 0x94, 0xb2, 0xdb, 0x64,
	0xa5, 0x89, 0x36, 0x69, 0xce, 0x20, 0x3e, 0x67, 0x3f, 0x92, 0x47, 0x09, 0x64, 0x39, 0x88, 0x54,
	0xd5, 0x96, 0x06, 0xf0, 0xac, 0x85, 0x39, 0x18, 0x27, 0x40, 0xfa, 0x36, 0x08, 0x59, 0x18, 0xa4,
	0xcb, 0x6c, 0x8b, 0x3c, 0x38, 0xab, 0x59, 0xa5, 0xf6, 0xb4, 0x96, 0xb6, 0xa3, 0xf3, 0x3c, 0x24,
	0xdf, 0x18, 0x6f, 0xe9, 0x0a, 0xfb, 0x81, 0x6c, 0x37, 0x3a, 0x32, 0xcd, 0x45, 0xbb, 0x57, 0x99,
	0xd1, 0x16, 0x28, 0xb9, 0xad, 0x5f, 0xe0, 0x45, 0x2e, 0x45, 0x12, 0xca, 0x8d, 0x12, 0x33, 0x54,
	0xce, 0xd2, 0xd5, 0x50, 0xd2, 0x86, 0x65, 0xd1, 0x55, 0x38, 0x5d, 0x38, 0xdf, 0x16, 0xaa, 0xae,
	0x90, 0xa5, 0x6b, 0xec, 0x29, 0x79, 0xdc, 0x40, 0x12, 0x89, 0x60, 0x26, 0x82, 0x5e, 0xb4, 0x9b,
	0xff, 0x32, 0x1d, 0x0b, 0x89, 0xf5, 0x8f, 0x75, 0x57, 0xaf, 0xb3, 0x27, 0xa4, 0xf5, 0x4e, 0x1e,
	0xe4, 0xf9, 0x24, 0xe9, 0x36, 0x7b, 0x46, 0x7e, 0xba, 0x84, 0xe4, 0x0c, 0x24, 0x3b, 0x21, 0xef,
	0xaa, 0xd0, 0x98, 0xe5, 0xb2, 0x4a, 0xa7, 0xfe, 0x34, 0xee, 0xb0, 0x5f, 0xc8, 0xd3, 0x4b, 0xc8,
	0x81, 0x93, 0x14, 0xd6, 0xe9, 0x2c, 0x98, 0x0d, 0x19, 0x3a, 0x34, 0xe3, 0x87, 0x37, 0xd8, 0x63,
	0x12, 0x4d, 0x18, 0x11, 0xeb, 0x60, 0xd8, 0x4c, 0x67, 0x85, 0x56, 0x9e, 0x70, 0xe6, 0x2e, 0x7b,
	0x4e, 0x7e, 0x7e, 0x1f, 0xce, 0x85, 0xaa, 0xe9, 0x37, 0xac, 0x45, 0x36, 0x27, 0x14, 0x9e, 0x6f,
	0xce, 0xb3, 0x51, 0x50, 0x35, 0xd2, 0xb7, 0xa1, 0x12, 0xe7, 0xed, 0x7e, 0x27, 0xef, 0x1e, 0xbb,
	0x4b, 0xd6, 0xc7, 0x9d, 0x2d, 0x45, 0xd0, 0x97, 0x36, 0xc3, 0xa6, 0x23, 0xd0, 0x80, 0x49, 0x3a,
	0x3d, 0x7a, 0x3f, 0x0c, 0xa1, 0xe9, 0x60, 0xa1, 0x84, 0xf3, 0x4d, 0x7e, 0x1d, 0x08, 0x93, 0xa5,
	0x23, 0x24, 0x37, 0xa8, 0xe8, 0x77, 0xc1, 0xad, 0x69, 0xb0, 0x2d, 0x62, 0x2e, 0xba, 0xc2, 0x86,
	0x57, 0x9b, 0x89, 0x61, 0xbd, 0x76, 0x9d, 0x60, 0x71, 0x02, 0x16, 0xe9, 0x83, 0x50, 0x9d, 0x69,
	0xce, 0x29, 0xce, 0x42, 0x86, 0x9e, 0x8b, 0x0c, 0x95, 0x1d, 0xab, 0xf6, 0x60, 0xbd, 0x15, 0xb1,
	0x14, 0x2a, 0xb5, 0xf4, 0x61, 0xf8, 0xc8, 0xa6, 0xb9, 0x33, 0x63, 0x71, 0xf3, 0x3c, 0x62, 0x66,
	0xca, 0x7e, 0x1f, 0x5c, 0x98, 0x46, 0x34, 0xb9, 0xc5, 0x61, 0xc6, 0x65, 0xba, 0x8b, 0x9c, 0x6e,
	0xb1, 0x35, 0xb2, 0x74, 0xb1, 0x55, 0x95, 0x95, 0x51, 0x98, 0xbe, 0x33, 0x42, 0x38, 0xf7, 0x19,
	0xf4, 0xbc, 0x56, 0xb2, 0x17, 0xe6, 0x46, 0x55, 0x60, 0x2f, 0x38, 0x6d, 0xc5, 0xff, 0xcc, 0x91,
	0x8d, 0x17, 0xe5, 0xab, 0xe8, 0xf2, 0x05, 0x14, 0x2f, 0x5d, 0xb8, 0x5f, 0xf2, 0xb0, 0xbc, 0xf2,
	0xb9, 0xdf, 0x78, 0xc3, 0x1e, 0x94, 0x2f, 0x0f, 0x8e, 0x07, 0x51, 0x39, 0x1c, 0xb4, 0x06, 0x87,
	0xc7, 0xd5, 0x6a, 0x1b, 0x6f, 0xc3, 0x93, 0xa3, 0xd1, 0xff, 0x2d, 0xc7, 0x67, 0xf5, 0x9f, 0x3f,
	0xe6, 0xaf, 0xa4, 0x00, 0x7f, 0xce, 0xaf, 0xa5, 0xf5, 0x65, 0xd0, 0x1f, 0x45, 0xf5, 0x31, 0x9c,
	0xba, 0xdb, 0x51, 0xf5, 0xe4, 0xe8, 0xef, 0x31, 0x60, 0x1f, 0xfa, 0xa3, 0xfd, 0x53, 0xc0, 0x7e,
	0x77, 0x7b, 0xbf, 0x06, 0xfc, 0x7e, 0xb5, 0x7a, 0xf8, 0xc9, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x3f, 0x5b, 0xb6, 0x2f, 0x94, 0x07, 0x00, 0x00,
}
