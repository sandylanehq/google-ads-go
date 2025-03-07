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

// FeedItemSetServiceClient is the client API for FeedItemSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedItemSetServiceClient interface {
	// Returns the requested feed item set in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetFeedItemSet(ctx context.Context, in *GetFeedItemSetRequest, opts ...grpc.CallOption) (*resources.FeedItemSet, error)
	// Creates, updates or removes feed item sets. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateFeedItemSets(ctx context.Context, in *MutateFeedItemSetsRequest, opts ...grpc.CallOption) (*MutateFeedItemSetsResponse, error)
}

type feedItemSetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedItemSetServiceClient(cc grpc.ClientConnInterface) FeedItemSetServiceClient {
	return &feedItemSetServiceClient{cc}
}

func (c *feedItemSetServiceClient) GetFeedItemSet(ctx context.Context, in *GetFeedItemSetRequest, opts ...grpc.CallOption) (*resources.FeedItemSet, error) {
	out := new(resources.FeedItemSet)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.FeedItemSetService/GetFeedItemSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedItemSetServiceClient) MutateFeedItemSets(ctx context.Context, in *MutateFeedItemSetsRequest, opts ...grpc.CallOption) (*MutateFeedItemSetsResponse, error) {
	out := new(MutateFeedItemSetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.FeedItemSetService/MutateFeedItemSets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedItemSetServiceServer is the server API for FeedItemSetService service.
// All implementations must embed UnimplementedFeedItemSetServiceServer
// for forward compatibility
type FeedItemSetServiceServer interface {
	// Returns the requested feed item set in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetFeedItemSet(context.Context, *GetFeedItemSetRequest) (*resources.FeedItemSet, error)
	// Creates, updates or removes feed item sets. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateFeedItemSets(context.Context, *MutateFeedItemSetsRequest) (*MutateFeedItemSetsResponse, error)
	mustEmbedUnimplementedFeedItemSetServiceServer()
}

// UnimplementedFeedItemSetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFeedItemSetServiceServer struct {
}

func (UnimplementedFeedItemSetServiceServer) GetFeedItemSet(context.Context, *GetFeedItemSetRequest) (*resources.FeedItemSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedItemSet not implemented")
}
func (UnimplementedFeedItemSetServiceServer) MutateFeedItemSets(context.Context, *MutateFeedItemSetsRequest) (*MutateFeedItemSetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateFeedItemSets not implemented")
}
func (UnimplementedFeedItemSetServiceServer) mustEmbedUnimplementedFeedItemSetServiceServer() {}

// UnsafeFeedItemSetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedItemSetServiceServer will
// result in compilation errors.
type UnsafeFeedItemSetServiceServer interface {
	mustEmbedUnimplementedFeedItemSetServiceServer()
}

func RegisterFeedItemSetServiceServer(s grpc.ServiceRegistrar, srv FeedItemSetServiceServer) {
	s.RegisterService(&FeedItemSetService_ServiceDesc, srv)
}

func _FeedItemSetService_GetFeedItemSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedItemSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedItemSetServiceServer).GetFeedItemSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.FeedItemSetService/GetFeedItemSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedItemSetServiceServer).GetFeedItemSet(ctx, req.(*GetFeedItemSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedItemSetService_MutateFeedItemSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedItemSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedItemSetServiceServer).MutateFeedItemSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.FeedItemSetService/MutateFeedItemSets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedItemSetServiceServer).MutateFeedItemSets(ctx, req.(*MutateFeedItemSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedItemSetService_ServiceDesc is the grpc.ServiceDesc for FeedItemSetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedItemSetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.FeedItemSetService",
	HandlerType: (*FeedItemSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeedItemSet",
			Handler:    _FeedItemSetService_GetFeedItemSet_Handler,
		},
		{
			MethodName: "MutateFeedItemSets",
			Handler:    _FeedItemSetService_MutateFeedItemSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/feed_item_set_service.proto",
}
