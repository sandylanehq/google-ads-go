// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/customer_client_link_service.proto

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

// Request message for [CustomerClientLinkService.GetCustomerClientLink][google.ads.googleads.v0.services.CustomerClientLinkService.GetCustomerClientLink].
type GetCustomerClientLinkRequest struct {
	// The resource name of the customer client link to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCustomerClientLinkRequest) Reset()         { *m = GetCustomerClientLinkRequest{} }
func (m *GetCustomerClientLinkRequest) String() string { return proto.CompactTextString(m) }
func (*GetCustomerClientLinkRequest) ProtoMessage()    {}
func (*GetCustomerClientLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea9531ab047027e2, []int{0}
}

func (m *GetCustomerClientLinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCustomerClientLinkRequest.Unmarshal(m, b)
}
func (m *GetCustomerClientLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCustomerClientLinkRequest.Marshal(b, m, deterministic)
}
func (m *GetCustomerClientLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCustomerClientLinkRequest.Merge(m, src)
}
func (m *GetCustomerClientLinkRequest) XXX_Size() int {
	return xxx_messageInfo_GetCustomerClientLinkRequest.Size(m)
}
func (m *GetCustomerClientLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCustomerClientLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCustomerClientLinkRequest proto.InternalMessageInfo

func (m *GetCustomerClientLinkRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCustomerClientLinkRequest)(nil), "google.ads.googleads.v0.services.GetCustomerClientLinkRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/customer_client_link_service.proto", fileDescriptor_ea9531ab047027e2)
}

var fileDescriptor_ea9531ab047027e2 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xbf, 0x4a, 0xc3, 0x40,
	0x18, 0x27, 0x15, 0x04, 0x83, 0x2e, 0x01, 0x41, 0x4b, 0x91, 0x52, 0x1d, 0xa4, 0xc3, 0x5d, 0x50,
	0x8a, 0x83, 0x56, 0x48, 0x33, 0xd4, 0x41, 0xa4, 0x54, 0xe8, 0x20, 0x81, 0x70, 0x26, 0x47, 0x38,
	0x9a, 0xdc, 0xd5, 0xfb, 0xae, 0x5d, 0xc4, 0xc5, 0x57, 0xf0, 0x0d, 0x1c, 0x7d, 0x14, 0xc1, 0xc9,
	0xdd, 0xc9, 0xcd, 0x97, 0x90, 0xe4, 0x72, 0x11, 0xa9, 0xb1, 0xdb, 0x8f, 0xdc, 0xf7, 0xfb, 0xf3,
	0xfd, 0xbe, 0xd8, 0x7e, 0x22, 0x44, 0x92, 0x52, 0x4c, 0x62, 0xc0, 0x1a, 0xe6, 0x68, 0xe1, 0x62,
	0xa0, 0x72, 0xc1, 0x22, 0x0a, 0x38, 0x9a, 0x83, 0x12, 0x19, 0x95, 0x61, 0x94, 0x32, 0xca, 0x55,
	0x98, 0x32, 0x3e, 0x0d, 0xcb, 0x57, 0x34, 0x93, 0x42, 0x09, 0xa7, 0xad, 0x99, 0x88, 0xc4, 0x80,
	0x2a, 0x11, 0xb4, 0x70, 0x91, 0x11, 0x69, 0x9e, 0xd5, 0xd9, 0x48, 0x0a, 0x62, 0x2e, 0xeb, 0x7c,
	0xb4, 0x7e, 0xb3, 0x65, 0xd8, 0x33, 0x86, 0x09, 0xe7, 0x42, 0x11, 0xc5, 0x04, 0x07, 0xfd, 0xda,
	0xf1, 0xed, 0xd6, 0x90, 0x2a, 0xbf, 0xa4, 0xfb, 0x05, 0xfb, 0x92, 0xf1, 0xe9, 0x98, 0xde, 0xcd,
	0x29, 0x28, 0x67, 0xdf, 0xde, 0x32, 0x2e, 0x21, 0x27, 0x19, 0xdd, 0xb1, 0xda, 0xd6, 0xe1, 0xc6,
	0x78, 0xd3, 0x7c, 0xbc, 0x22, 0x19, 0x3d, 0xfa, 0xb2, 0xec, 0xdd, 0x65, 0x89, 0x6b, 0x9d, 0xdf,
	0x79, 0xb3, 0xec, 0xed, 0x3f, 0x3d, 0x9c, 0x73, 0xb4, 0x6a, 0x77, 0xf4, 0x5f, 0xb8, 0x66, 0xaf,
	0x96, 0x5f, 0x35, 0x83, 0x96, 0xd9, 0x9d, 0xfe, 0xe3, 0xfb, 0xe7, 0x53, 0xe3, 0xc4, 0xe9, 0xe5,
	0x1d, 0xde, 0xff, 0x5a, 0xaf, 0x6f, 0x8a, 0x04, 0xdc, 0xad, 0x4a, 0xfd, 0xa1, 0x02, 0xee, 0x3e,
	0x0c, 0x3e, 0x2c, 0xfb, 0x20, 0x12, 0xd9, 0xca, 0xec, 0x83, 0xbd, 0xda, 0x4e, 0x46, 0x79, 0xf7,
	0x23, 0xeb, 0xe6, 0xa2, 0xd4, 0x48, 0x44, 0x4a, 0x78, 0x82, 0x84, 0x4c, 0x70, 0x42, 0x79, 0x71,
	0x19, 0x73, 0xe9, 0x19, 0x83, 0xfa, 0xff, 0xeb, 0xd4, 0x80, 0xe7, 0xc6, 0xda, 0xd0, 0xf3, 0x5e,
	0x1a, 0xed, 0xa1, 0x16, 0xf4, 0x62, 0x40, 0x1a, 0xe6, 0x68, 0xe2, 0xa2, 0xd2, 0x18, 0x5e, 0xcd,
	0x48, 0xe0, 0xc5, 0x10, 0x54, 0x23, 0xc1, 0xc4, 0x0d, 0xcc, 0xc8, 0xed, 0x7a, 0x11, 0xe0, 0xf8,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x53, 0x7b, 0xfc, 0x03, 0xdf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerClientLinkServiceClient is the client API for CustomerClientLinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerClientLinkServiceClient interface {
	// Returns the requested CustomerClientLink in full detail.
	GetCustomerClientLink(ctx context.Context, in *GetCustomerClientLinkRequest, opts ...grpc.CallOption) (*resources.CustomerClientLink, error)
}

type customerClientLinkServiceClient struct {
	cc *grpc.ClientConn
}

func NewCustomerClientLinkServiceClient(cc *grpc.ClientConn) CustomerClientLinkServiceClient {
	return &customerClientLinkServiceClient{cc}
}

func (c *customerClientLinkServiceClient) GetCustomerClientLink(ctx context.Context, in *GetCustomerClientLinkRequest, opts ...grpc.CallOption) (*resources.CustomerClientLink, error) {
	out := new(resources.CustomerClientLink)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.CustomerClientLinkService/GetCustomerClientLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerClientLinkServiceServer is the server API for CustomerClientLinkService service.
type CustomerClientLinkServiceServer interface {
	// Returns the requested CustomerClientLink in full detail.
	GetCustomerClientLink(context.Context, *GetCustomerClientLinkRequest) (*resources.CustomerClientLink, error)
}

func RegisterCustomerClientLinkServiceServer(s *grpc.Server, srv CustomerClientLinkServiceServer) {
	s.RegisterService(&_CustomerClientLinkService_serviceDesc, srv)
}

func _CustomerClientLinkService_GetCustomerClientLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerClientLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerClientLinkServiceServer).GetCustomerClientLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.CustomerClientLinkService/GetCustomerClientLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerClientLinkServiceServer).GetCustomerClientLink(ctx, req.(*GetCustomerClientLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerClientLinkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.CustomerClientLinkService",
	HandlerType: (*CustomerClientLinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerClientLink",
			Handler:    _CustomerClientLinkService_GetCustomerClientLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/customer_client_link_service.proto",
}
