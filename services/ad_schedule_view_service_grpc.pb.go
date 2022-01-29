// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	resources "github.com/butlerhq/google-ads-go/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AdScheduleViewServiceClient is the client API for AdScheduleViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdScheduleViewServiceClient interface {
	// Returns the requested ad schedule view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetAdScheduleView(ctx context.Context, in *GetAdScheduleViewRequest, opts ...grpc.CallOption) (*resources.AdScheduleView, error)
}

type adScheduleViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdScheduleViewServiceClient(cc grpc.ClientConnInterface) AdScheduleViewServiceClient {
	return &adScheduleViewServiceClient{cc}
}

func (c *adScheduleViewServiceClient) GetAdScheduleView(ctx context.Context, in *GetAdScheduleViewRequest, opts ...grpc.CallOption) (*resources.AdScheduleView, error) {
	out := new(resources.AdScheduleView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.AdScheduleViewService/GetAdScheduleView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdScheduleViewServiceServer is the server API for AdScheduleViewService service.
// All implementations must embed UnimplementedAdScheduleViewServiceServer
// for forward compatibility
type AdScheduleViewServiceServer interface {
	// Returns the requested ad schedule view in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetAdScheduleView(context.Context, *GetAdScheduleViewRequest) (*resources.AdScheduleView, error)
	mustEmbedUnimplementedAdScheduleViewServiceServer()
}

// UnimplementedAdScheduleViewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdScheduleViewServiceServer struct {
}

func (UnimplementedAdScheduleViewServiceServer) GetAdScheduleView(context.Context, *GetAdScheduleViewRequest) (*resources.AdScheduleView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAdScheduleView not implemented")
}
func (UnimplementedAdScheduleViewServiceServer) mustEmbedUnimplementedAdScheduleViewServiceServer() {}

// UnsafeAdScheduleViewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdScheduleViewServiceServer will
// result in compilation errors.
type UnsafeAdScheduleViewServiceServer interface {
	mustEmbedUnimplementedAdScheduleViewServiceServer()
}

func RegisterAdScheduleViewServiceServer(s grpc.ServiceRegistrar, srv AdScheduleViewServiceServer) {
	s.RegisterService(&AdScheduleViewService_ServiceDesc, srv)
}

func _AdScheduleViewService_GetAdScheduleView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdScheduleViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdScheduleViewServiceServer).GetAdScheduleView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.AdScheduleViewService/GetAdScheduleView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdScheduleViewServiceServer).GetAdScheduleView(ctx, req.(*GetAdScheduleViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdScheduleViewService_ServiceDesc is the grpc.ServiceDesc for AdScheduleViewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdScheduleViewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.AdScheduleViewService",
	HandlerType: (*AdScheduleViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdScheduleView",
			Handler:    _AdScheduleViewService_GetAdScheduleView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/ad_schedule_view_service.proto",
}
