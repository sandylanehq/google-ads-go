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

// SharedSetServiceClient is the client API for SharedSetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SharedSetServiceClient interface {
	// Returns the requested shared set in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetSharedSet(ctx context.Context, in *GetSharedSetRequest, opts ...grpc.CallOption) (*resources.SharedSet, error)
	// Creates, updates, or removes shared sets. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [SharedSetError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateSharedSets(ctx context.Context, in *MutateSharedSetsRequest, opts ...grpc.CallOption) (*MutateSharedSetsResponse, error)
}

type sharedSetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSharedSetServiceClient(cc grpc.ClientConnInterface) SharedSetServiceClient {
	return &sharedSetServiceClient{cc}
}

func (c *sharedSetServiceClient) GetSharedSet(ctx context.Context, in *GetSharedSetRequest, opts ...grpc.CallOption) (*resources.SharedSet, error) {
	out := new(resources.SharedSet)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.SharedSetService/GetSharedSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedSetServiceClient) MutateSharedSets(ctx context.Context, in *MutateSharedSetsRequest, opts ...grpc.CallOption) (*MutateSharedSetsResponse, error) {
	out := new(MutateSharedSetsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.SharedSetService/MutateSharedSets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharedSetServiceServer is the server API for SharedSetService service.
// All implementations must embed UnimplementedSharedSetServiceServer
// for forward compatibility
type SharedSetServiceServer interface {
	// Returns the requested shared set in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetSharedSet(context.Context, *GetSharedSetRequest) (*resources.SharedSet, error)
	// Creates, updates, or removes shared sets. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [DatabaseError]()
	//   [DateError]()
	//   [DistinctError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [IdError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [NewResourceCreationError]()
	//   [NotEmptyError]()
	//   [NullError]()
	//   [OperatorError]()
	//   [QuotaError]()
	//   [RangeError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	//   [SharedSetError]()
	//   [SizeLimitError]()
	//   [StringFormatError]()
	//   [StringLengthError]()
	MutateSharedSets(context.Context, *MutateSharedSetsRequest) (*MutateSharedSetsResponse, error)
	mustEmbedUnimplementedSharedSetServiceServer()
}

// UnimplementedSharedSetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSharedSetServiceServer struct {
}

func (UnimplementedSharedSetServiceServer) GetSharedSet(context.Context, *GetSharedSetRequest) (*resources.SharedSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSharedSet not implemented")
}
func (UnimplementedSharedSetServiceServer) MutateSharedSets(context.Context, *MutateSharedSetsRequest) (*MutateSharedSetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateSharedSets not implemented")
}
func (UnimplementedSharedSetServiceServer) mustEmbedUnimplementedSharedSetServiceServer() {}

// UnsafeSharedSetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SharedSetServiceServer will
// result in compilation errors.
type UnsafeSharedSetServiceServer interface {
	mustEmbedUnimplementedSharedSetServiceServer()
}

func RegisterSharedSetServiceServer(s grpc.ServiceRegistrar, srv SharedSetServiceServer) {
	s.RegisterService(&SharedSetService_ServiceDesc, srv)
}

func _SharedSetService_GetSharedSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSharedSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedSetServiceServer).GetSharedSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.SharedSetService/GetSharedSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedSetServiceServer).GetSharedSet(ctx, req.(*GetSharedSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SharedSetService_MutateSharedSets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateSharedSetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedSetServiceServer).MutateSharedSets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.SharedSetService/MutateSharedSets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedSetServiceServer).MutateSharedSets(ctx, req.(*MutateSharedSetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SharedSetService_ServiceDesc is the grpc.ServiceDesc for SharedSetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SharedSetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.SharedSetService",
	HandlerType: (*SharedSetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSharedSet",
			Handler:    _SharedSetService_GetSharedSet_Handler,
		},
		{
			MethodName: "MutateSharedSets",
			Handler:    _SharedSetService_MutateSharedSets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/shared_set_service.proto",
}
