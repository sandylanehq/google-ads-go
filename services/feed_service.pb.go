// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/feed_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	resources "github.com/kritzware/google-ads-go/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
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

// Request message for [FeedService.GetFeed][google.ads.googleads.v0.services.FeedService.GetFeed].
type GetFeedRequest struct {
	// The resource name of the feed to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedRequest) Reset()         { *m = GetFeedRequest{} }
func (m *GetFeedRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedRequest) ProtoMessage()    {}
func (*GetFeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a46368c45e6d06e, []int{0}
}

func (m *GetFeedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedRequest.Unmarshal(m, b)
}
func (m *GetFeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedRequest.Merge(m, src)
}
func (m *GetFeedRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedRequest.Size(m)
}
func (m *GetFeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedRequest proto.InternalMessageInfo

func (m *GetFeedRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [FeedService.MutateFeeds][google.ads.googleads.v0.services.FeedService.MutateFeeds].
type MutateFeedsRequest struct {
	// The ID of the customer whose feeds are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual feeds.
	Operations           []*FeedOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MutateFeedsRequest) Reset()         { *m = MutateFeedsRequest{} }
func (m *MutateFeedsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateFeedsRequest) ProtoMessage()    {}
func (*MutateFeedsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a46368c45e6d06e, []int{1}
}

func (m *MutateFeedsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedsRequest.Unmarshal(m, b)
}
func (m *MutateFeedsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedsRequest.Marshal(b, m, deterministic)
}
func (m *MutateFeedsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedsRequest.Merge(m, src)
}
func (m *MutateFeedsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateFeedsRequest.Size(m)
}
func (m *MutateFeedsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedsRequest proto.InternalMessageInfo

func (m *MutateFeedsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateFeedsRequest) GetOperations() []*FeedOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// A single operation (create, update, remove) on an feed.
type FeedOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*FeedOperation_Create
	//	*FeedOperation_Update
	//	*FeedOperation_Remove
	Operation            isFeedOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *FeedOperation) Reset()         { *m = FeedOperation{} }
func (m *FeedOperation) String() string { return proto.CompactTextString(m) }
func (*FeedOperation) ProtoMessage()    {}
func (*FeedOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a46368c45e6d06e, []int{2}
}

func (m *FeedOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedOperation.Unmarshal(m, b)
}
func (m *FeedOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedOperation.Marshal(b, m, deterministic)
}
func (m *FeedOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedOperation.Merge(m, src)
}
func (m *FeedOperation) XXX_Size() int {
	return xxx_messageInfo_FeedOperation.Size(m)
}
func (m *FeedOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedOperation.DiscardUnknown(m)
}

var xxx_messageInfo_FeedOperation proto.InternalMessageInfo

func (m *FeedOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isFeedOperation_Operation interface {
	isFeedOperation_Operation()
}

type FeedOperation_Create struct {
	Create *resources.Feed `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type FeedOperation_Update struct {
	Update *resources.Feed `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type FeedOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*FeedOperation_Create) isFeedOperation_Operation() {}

func (*FeedOperation_Update) isFeedOperation_Operation() {}

func (*FeedOperation_Remove) isFeedOperation_Operation() {}

func (m *FeedOperation) GetOperation() isFeedOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *FeedOperation) GetCreate() *resources.Feed {
	if x, ok := m.GetOperation().(*FeedOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *FeedOperation) GetUpdate() *resources.Feed {
	if x, ok := m.GetOperation().(*FeedOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *FeedOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*FeedOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FeedOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FeedOperation_Create)(nil),
		(*FeedOperation_Update)(nil),
		(*FeedOperation_Remove)(nil),
	}
}

// Response message for an feed mutate.
type MutateFeedsResponse struct {
	// All results for the mutate.
	Results              []*MutateFeedResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MutateFeedsResponse) Reset()         { *m = MutateFeedsResponse{} }
func (m *MutateFeedsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateFeedsResponse) ProtoMessage()    {}
func (*MutateFeedsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a46368c45e6d06e, []int{3}
}

func (m *MutateFeedsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedsResponse.Unmarshal(m, b)
}
func (m *MutateFeedsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedsResponse.Marshal(b, m, deterministic)
}
func (m *MutateFeedsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedsResponse.Merge(m, src)
}
func (m *MutateFeedsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateFeedsResponse.Size(m)
}
func (m *MutateFeedsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedsResponse proto.InternalMessageInfo

func (m *MutateFeedsResponse) GetResults() []*MutateFeedResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the feed mutate.
type MutateFeedResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedResult) Reset()         { *m = MutateFeedResult{} }
func (m *MutateFeedResult) String() string { return proto.CompactTextString(m) }
func (*MutateFeedResult) ProtoMessage()    {}
func (*MutateFeedResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a46368c45e6d06e, []int{4}
}

func (m *MutateFeedResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedResult.Unmarshal(m, b)
}
func (m *MutateFeedResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedResult.Marshal(b, m, deterministic)
}
func (m *MutateFeedResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedResult.Merge(m, src)
}
func (m *MutateFeedResult) XXX_Size() int {
	return xxx_messageInfo_MutateFeedResult.Size(m)
}
func (m *MutateFeedResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedResult proto.InternalMessageInfo

func (m *MutateFeedResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedRequest)(nil), "google.ads.googleads.v0.services.GetFeedRequest")
	proto.RegisterType((*MutateFeedsRequest)(nil), "google.ads.googleads.v0.services.MutateFeedsRequest")
	proto.RegisterType((*FeedOperation)(nil), "google.ads.googleads.v0.services.FeedOperation")
	proto.RegisterType((*MutateFeedsResponse)(nil), "google.ads.googleads.v0.services.MutateFeedsResponse")
	proto.RegisterType((*MutateFeedResult)(nil), "google.ads.googleads.v0.services.MutateFeedResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/feed_service.proto", fileDescriptor_0a46368c45e6d06e)
}

var fileDescriptor_0a46368c45e6d06e = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0x76, 0x37, 0x92, 0xd2, 0x59, 0x2b, 0x65, 0xbc, 0x2c, 0x41, 0x30, 0xac, 0x42, 0x63, 0x90,
	0x99, 0x90, 0x1a, 0x84, 0x96, 0x1e, 0x92, 0x43, 0x53, 0xc1, 0xda, 0xb2, 0x42, 0x0f, 0x12, 0x08,
	0xd3, 0xdd, 0x97, 0x65, 0x69, 0x76, 0x67, 0xdd, 0x99, 0xcd, 0xa5, 0xf4, 0xe2, 0xc1, 0x3f, 0x20,
	0xfe, 0x01, 0x8f, 0xde, 0xfc, 0x1b, 0x9e, 0x04, 0x4f, 0xde, 0xfd, 0x13, 0xde, 0x64, 0x76, 0x76,
	0xd2, 0x44, 0x08, 0x69, 0x6e, 0x6f, 0x76, 0xde, 0xf7, 0xbd, 0x6f, 0xbe, 0xf7, 0xde, 0xa2, 0xfd,
	0x88, 0xf3, 0x68, 0x0a, 0x94, 0x85, 0x82, 0xea, 0x50, 0x45, 0xb3, 0x0e, 0x15, 0x90, 0xcf, 0xe2,
	0x00, 0x04, 0x9d, 0x00, 0x84, 0xe3, 0xea, 0x44, 0xb2, 0x9c, 0x4b, 0x8e, 0x9b, 0x3a, 0x93, 0xb0,
	0x50, 0x90, 0x39, 0x88, 0xcc, 0x3a, 0xc4, 0x80, 0x1a, 0x2f, 0x56, 0xd1, 0xe6, 0x20, 0x78, 0x91,
	0x1b, 0x5e, 0xcd, 0xd7, 0x78, 0x6c, 0xb2, 0xb3, 0x98, 0xb2, 0x34, 0xe5, 0x92, 0xc9, 0x98, 0xa7,
	0xa2, 0xba, 0xad, 0xaa, 0xd1, 0xf2, 0x74, 0x59, 0x4c, 0xe8, 0x24, 0x86, 0x69, 0x38, 0x4e, 0x98,
	0xb8, 0xd2, 0x19, 0x5e, 0x0f, 0x3d, 0x1c, 0x82, 0x3c, 0x06, 0x08, 0x7d, 0xf8, 0x50, 0x80, 0x90,
	0xf8, 0x29, 0xda, 0x31, 0x95, 0xc6, 0x29, 0x4b, 0xc0, 0xb5, 0x9a, 0x56, 0x6b, 0xdb, 0x7f, 0x60,
	0x3e, 0xbe, 0x65, 0x09, 0x78, 0x9f, 0x2c, 0x84, 0x4f, 0x0b, 0xc9, 0x24, 0x28, 0xa8, 0x30, 0xd8,
	0x27, 0xc8, 0x09, 0x0a, 0x21, 0x79, 0x02, 0xf9, 0x38, 0x0e, 0x2b, 0x24, 0x32, 0x9f, 0x5e, 0x87,
	0xf8, 0x0c, 0x21, 0x9e, 0x41, 0xae, 0x45, 0xba, 0x76, 0xb3, 0xd6, 0x72, 0xba, 0x94, 0xac, 0xf3,
	0x84, 0xa8, 0x22, 0x67, 0x06, 0xe7, 0x2f, 0x50, 0x78, 0x7f, 0x2d, 0xb4, 0xb3, 0x74, 0x8b, 0x0f,
	0x91, 0x53, 0x64, 0x21, 0x93, 0x50, 0x3e, 0xd3, 0xbd, 0xdf, 0xb4, 0x5a, 0x4e, 0xb7, 0x61, 0x6a,
	0x18, 0x27, 0xc8, 0xb1, 0x72, 0xe2, 0x94, 0x89, 0x2b, 0x1f, 0xe9, 0x74, 0x15, 0xe3, 0x3e, 0xaa,
	0x07, 0x39, 0x30, 0xa9, 0x5f, 0xed, 0x74, 0xf7, 0x56, 0x6a, 0x9b, 0x77, 0xa3, 0x14, 0x77, 0x72,
	0xcf, 0xaf, 0x80, 0x8a, 0x42, 0x13, 0xba, 0xf6, 0xc6, 0x14, 0x1a, 0x88, 0x5d, 0x54, 0xcf, 0x21,
	0xe1, 0x33, 0x70, 0x6b, 0xca, 0x41, 0x75, 0xa3, 0xcf, 0x03, 0x07, 0x6d, 0xcf, 0x1f, 0xef, 0x05,
	0xe8, 0xd1, 0x52, 0x0f, 0x44, 0xc6, 0x53, 0x01, 0xf8, 0x0d, 0xda, 0xca, 0x41, 0x14, 0x53, 0x69,
	0x0c, 0xee, 0xae, 0x37, 0xf8, 0x96, 0xc7, 0x2f, 0xa1, 0xbe, 0xa1, 0xf0, 0x5e, 0xa1, 0xdd, 0xff,
	0x2f, 0xef, 0x34, 0x22, 0xdd, 0xdf, 0x36, 0x72, 0x14, 0xe6, 0x9d, 0xae, 0x81, 0xbf, 0x58, 0x68,
	0xab, 0x1a, 0x35, 0xdc, 0x59, 0xaf, 0x68, 0x79, 0x2a, 0x1b, 0x77, 0x75, 0xd1, 0xa3, 0x1f, 0x7f,
	0xfd, 0xf9, 0x6c, 0x3f, 0xc7, 0x7b, 0x6a, 0x65, 0xae, 0x97, 0x64, 0x1e, 0x99, 0x41, 0x14, 0xb4,
	0x5d, 0xee, 0x90, 0xa0, 0xed, 0x1b, 0xfc, 0xdd, 0x42, 0xce, 0x82, 0x8d, 0xf8, 0xe5, 0x26, 0x6e,
	0x99, 0xc9, 0x6f, 0xf4, 0x36, 0x44, 0xe9, 0x5e, 0x79, 0xbd, 0x52, 0x2d, 0xf5, 0xda, 0x4a, 0xed,
	0xad, 0xbc, 0xeb, 0x85, 0x2d, 0x3a, 0x6a, 0xdf, 0x68, 0xb1, 0x07, 0x49, 0x49, 0x70, 0x60, 0xb5,
	0x07, 0x3f, 0x2d, 0xf4, 0x2c, 0xe0, 0xc9, 0xda, 0x9a, 0x83, 0xdd, 0x85, 0x0e, 0x9c, 0xab, 0xd1,
	0x3f, 0xb7, 0xde, 0x9f, 0x54, 0xa8, 0x88, 0x4f, 0x59, 0x1a, 0x11, 0x9e, 0x47, 0x34, 0x82, 0xb4,
	0x5c, 0x0c, 0xf3, 0xc3, 0xc9, 0x62, 0xb1, 0xfa, 0xb7, 0x76, 0x68, 0x82, 0xaf, 0x76, 0x6d, 0xd8,
	0xef, 0x7f, 0xb3, 0x9b, 0x43, 0x4d, 0xd8, 0x0f, 0x05, 0xd1, 0xa1, 0x8a, 0x2e, 0x3a, 0xa4, 0x2a,
	0x2c, 0x7e, 0x98, 0x94, 0x51, 0x3f, 0x14, 0xa3, 0x79, 0xca, 0xe8, 0xa2, 0x33, 0x32, 0x29, 0x97,
	0xf5, 0x52, 0xc0, 0xfe, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x35, 0xe0, 0x55, 0x56, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedServiceClient is the client API for FeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedServiceClient interface {
	// Returns the requested feed in full detail.
	GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*resources.Feed, error)
	// Creates, updates, or removes feeds. Operation statuses are
	// returned.
	MutateFeeds(ctx context.Context, in *MutateFeedsRequest, opts ...grpc.CallOption) (*MutateFeedsResponse, error)
}

type feedServiceClient struct {
	cc *grpc.ClientConn
}

func NewFeedServiceClient(cc *grpc.ClientConn) FeedServiceClient {
	return &feedServiceClient{cc}
}

func (c *feedServiceClient) GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*resources.Feed, error) {
	out := new(resources.Feed)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.FeedService/GetFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) MutateFeeds(ctx context.Context, in *MutateFeedsRequest, opts ...grpc.CallOption) (*MutateFeedsResponse, error) {
	out := new(MutateFeedsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.FeedService/MutateFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServiceServer is the server API for FeedService service.
type FeedServiceServer interface {
	// Returns the requested feed in full detail.
	GetFeed(context.Context, *GetFeedRequest) (*resources.Feed, error)
	// Creates, updates, or removes feeds. Operation statuses are
	// returned.
	MutateFeeds(context.Context, *MutateFeedsRequest) (*MutateFeedsResponse, error)
}

func RegisterFeedServiceServer(s *grpc.Server, srv FeedServiceServer) {
	s.RegisterService(&_FeedService_serviceDesc, srv)
}

func _FeedService_GetFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).GetFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.FeedService/GetFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).GetFeed(ctx, req.(*GetFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_MutateFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).MutateFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.FeedService/MutateFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).MutateFeeds(ctx, req.(*MutateFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.FeedService",
	HandlerType: (*FeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeed",
			Handler:    _FeedService_GetFeed_Handler,
		},
		{
			MethodName: "MutateFeeds",
			Handler:    _FeedService_MutateFeeds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/feed_service.proto",
}
