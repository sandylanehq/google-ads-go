// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/account_budget_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	resources "github.com/kritzware/google-ads-go/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for
// [AccountBudgetService.GetAccountBudget][google.ads.googleads.v0.services.AccountBudgetService.GetAccountBudget].
type GetAccountBudgetRequest struct {
	// The resource name of the account-level budget to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAccountBudgetRequest) Reset()         { *m = GetAccountBudgetRequest{} }
func (m *GetAccountBudgetRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountBudgetRequest) ProtoMessage()    {}
func (*GetAccountBudgetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a71cca60e5ec334e, []int{0}
}

func (m *GetAccountBudgetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountBudgetRequest.Unmarshal(m, b)
}
func (m *GetAccountBudgetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountBudgetRequest.Marshal(b, m, deterministic)
}
func (m *GetAccountBudgetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountBudgetRequest.Merge(m, src)
}
func (m *GetAccountBudgetRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountBudgetRequest.Size(m)
}
func (m *GetAccountBudgetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountBudgetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountBudgetRequest proto.InternalMessageInfo

func (m *GetAccountBudgetRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAccountBudgetRequest)(nil), "google.ads.googleads.v0.services.GetAccountBudgetRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/account_budget_service.proto", fileDescriptor_a71cca60e5ec334e)
}

var fileDescriptor_a71cca60e5ec334e = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4d, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xe2, 0xd4,
	0xa2, 0xb2, 0xcc, 0xe4, 0xd4, 0x62, 0xfd, 0xc4, 0xe4, 0xe4, 0xfc, 0xd2, 0xbc, 0x92, 0xf8, 0xa4,
	0xd2, 0x94, 0xf4, 0xd4, 0x92, 0x78, 0xa8, 0xb8, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x02,
	0x44, 0x8f, 0x5e, 0x62, 0x4a, 0xb1, 0x1e, 0x5c, 0xbb, 0x5e, 0x99, 0x81, 0x1e, 0x4c, 0xbb, 0x94,
	0x19, 0x2e, 0x0b, 0x8a, 0x52, 0x8b, 0xf3, 0x4b, 0x8b, 0x30, 0x6d, 0x80, 0x98, 0x2c, 0x25, 0x03,
	0xd3, 0x57, 0x90, 0xa9, 0x9f, 0x98, 0x97, 0x97, 0x5f, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x57, 0x0c,
	0x91, 0x55, 0xb2, 0xe3, 0x12, 0x77, 0x4f, 0x2d, 0x71, 0x84, 0x68, 0x74, 0x02, 0xeb, 0x0b, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x52, 0xe6, 0xe2, 0x85, 0x19, 0x1d, 0x9f, 0x97, 0x98, 0x9b,
	0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x03, 0x13, 0xf4, 0x4b, 0xcc, 0x4d, 0x35, 0xba,
	0xc4, 0xc8, 0x25, 0x82, 0xa2, 0x3b, 0x18, 0xe2, 0x5e, 0xa1, 0x9d, 0x8c, 0x5c, 0x02, 0xe8, 0x26,
	0x0b, 0x59, 0xea, 0x11, 0xf2, 0xa6, 0x1e, 0x0e, 0xd7, 0x48, 0x19, 0xe0, 0xd4, 0x0a, 0xf7, 0xbf,
	0x1e, 0x8a, 0x46, 0x25, 0x8b, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x19, 0x09, 0x19, 0x80, 0x02, 0xa9,
	0x1a, 0xc5, 0x2b, 0xb6, 0xc9, 0xa5, 0xc5, 0x25, 0xf9, 0xb9, 0xa9, 0x45, 0xc5, 0xfa, 0x5a, 0xb0,
	0x50, 0x83, 0xe8, 0x2a, 0xd6, 0xd7, 0xaa, 0x75, 0xba, 0xc9, 0xc8, 0xa5, 0x92, 0x9c, 0x9f, 0x4b,
	0xd0, 0xb1, 0x4e, 0x92, 0xd8, 0xbc, 0x1e, 0x00, 0x0a, 0xd8, 0x00, 0xc6, 0x28, 0x0f, 0xa8, 0xf6,
	0xf4, 0xfc, 0x9c, 0xc4, 0xbc, 0x74, 0xbd, 0xfc, 0xa2, 0x74, 0xfd, 0xf4, 0xd4, 0x3c, 0x70, 0xb0,
	0xc3, 0x22, 0xb0, 0x20, 0xb3, 0x18, 0x77, 0x82, 0xb1, 0x86, 0x31, 0x16, 0x31, 0x31, 0xbb, 0x3b,
	0x3a, 0xae, 0x62, 0x52, 0x70, 0x87, 0x18, 0xe8, 0x98, 0x52, 0xac, 0x07, 0x61, 0x82, 0x58, 0x61,
	0x06, 0x7a, 0x50, 0x8b, 0x8b, 0x4f, 0xc1, 0x94, 0xc4, 0x38, 0xa6, 0x14, 0xc7, 0xc0, 0x95, 0xc4,
	0x84, 0x19, 0xc4, 0xc0, 0x94, 0x24, 0xb1, 0x81, 0x1d, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xcd, 0xfb, 0x18, 0x58, 0xb0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountBudgetServiceClient is the client API for AccountBudgetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountBudgetServiceClient interface {
	// Returns an account-level budget in full detail.
	GetAccountBudget(ctx context.Context, in *GetAccountBudgetRequest, opts ...grpc.CallOption) (*resources.AccountBudget, error)
}

type accountBudgetServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountBudgetServiceClient(cc *grpc.ClientConn) AccountBudgetServiceClient {
	return &accountBudgetServiceClient{cc}
}

func (c *accountBudgetServiceClient) GetAccountBudget(ctx context.Context, in *GetAccountBudgetRequest, opts ...grpc.CallOption) (*resources.AccountBudget, error) {
	out := new(resources.AccountBudget)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.AccountBudgetService/GetAccountBudget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountBudgetServiceServer is the server API for AccountBudgetService service.
type AccountBudgetServiceServer interface {
	// Returns an account-level budget in full detail.
	GetAccountBudget(context.Context, *GetAccountBudgetRequest) (*resources.AccountBudget, error)
}

func RegisterAccountBudgetServiceServer(s *grpc.Server, srv AccountBudgetServiceServer) {
	s.RegisterService(&_AccountBudgetService_serviceDesc, srv)
}

func _AccountBudgetService_GetAccountBudget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountBudgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountBudgetServiceServer).GetAccountBudget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.AccountBudgetService/GetAccountBudget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountBudgetServiceServer).GetAccountBudget(ctx, req.(*GetAccountBudgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountBudgetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.AccountBudgetService",
	HandlerType: (*AccountBudgetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountBudget",
			Handler:    _AccountBudgetService_GetAccountBudget_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/account_budget_service.proto",
}
