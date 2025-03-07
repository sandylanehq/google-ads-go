// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AssetSetServiceClient is the client API for AssetSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetSetServiceClient interface {
	// Creates, updates or removes asset sets. Operation statuses are
	// returned.
	MutateAssetSets(ctx context.Context, in *MutateAssetSetsRequest, opts ...grpc.CallOption) (*MutateAssetSetsResponse, error)
}

type assetSetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetSetServiceClient(cc grpc.ClientConnInterface) AssetSetServiceClient {
	return &assetSetServiceClient{cc}
}

func (c *assetSetServiceClient) MutateAssetSets(ctx context.Context, in *MutateAssetSetsRequest, opts ...grpc.CallOption) (*MutateAssetSetsResponse, error) {
	out := new(MutateAssetSetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.AssetSetService/MutateAssetSets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetSetServiceServer is the server API for AssetSetService service.
// All implementations must embed UnimplementedAssetSetServiceServer
// for forward compatibility
type AssetSetServiceServer interface {
	// Creates, updates or removes asset sets. Operation statuses are
	// returned.
	MutateAssetSets(context.Context, *MutateAssetSetsRequest) (*MutateAssetSetsResponse, error)
	mustEmbedUnimplementedAssetSetServiceServer()
}

// UnimplementedAssetSetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAssetSetServiceServer struct {
}

func (UnimplementedAssetSetServiceServer) MutateAssetSets(context.Context, *MutateAssetSetsRequest) (*MutateAssetSetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAssetSets not implemented")
}
func (UnimplementedAssetSetServiceServer) mustEmbedUnimplementedAssetSetServiceServer() {}

// UnsafeAssetSetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetSetServiceServer will
// result in compilation errors.
type UnsafeAssetSetServiceServer interface {
	mustEmbedUnimplementedAssetSetServiceServer()
}

func RegisterAssetSetServiceServer(s grpc.ServiceRegistrar, srv AssetSetServiceServer) {
	s.RegisterService(&AssetSetService_ServiceDesc, srv)
}

func _AssetSetService_MutateAssetSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAssetSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetSetServiceServer).MutateAssetSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.AssetSetService/MutateAssetSets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetSetServiceServer).MutateAssetSets(ctx, req.(*MutateAssetSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetSetService_ServiceDesc is the grpc.ServiceDesc for AssetSetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetSetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.AssetSetService",
	HandlerType: (*AssetSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateAssetSets",
			Handler:    _AssetSetService_MutateAssetSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/asset_set_service.proto",
}
