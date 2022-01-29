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

// CustomInterestServiceClient is the client API for CustomInterestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomInterestServiceClient interface {
	// Returns the requested custom interest in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCustomInterest(ctx context.Context, in *GetCustomInterestRequest, opts ...grpc.CallOption) (*resources.CustomInterest, error)
	// Creates or updates custom interests. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CriterionError]()
	//   [CustomInterestError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [PolicyViolationError]()
	//   [QuotaError]()
	//   [RequestError]()
	//   [StringLengthError]()
	MutateCustomInterests(ctx context.Context, in *MutateCustomInterestsRequest, opts ...grpc.CallOption) (*MutateCustomInterestsResponse, error)
}

type customInterestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomInterestServiceClient(cc grpc.ClientConnInterface) CustomInterestServiceClient {
	return &customInterestServiceClient{cc}
}

func (c *customInterestServiceClient) GetCustomInterest(ctx context.Context, in *GetCustomInterestRequest, opts ...grpc.CallOption) (*resources.CustomInterest, error) {
	out := new(resources.CustomInterest)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.CustomInterestService/GetCustomInterest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customInterestServiceClient) MutateCustomInterests(ctx context.Context, in *MutateCustomInterestsRequest, opts ...grpc.CallOption) (*MutateCustomInterestsResponse, error) {
	out := new(MutateCustomInterestsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.CustomInterestService/MutateCustomInterests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomInterestServiceServer is the server API for CustomInterestService service.
// All implementations must embed UnimplementedCustomInterestServiceServer
// for forward compatibility
type CustomInterestServiceServer interface {
	// Returns the requested custom interest in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCustomInterest(context.Context, *GetCustomInterestRequest) (*resources.CustomInterest, error)
	// Creates or updates custom interests. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CriterionError]()
	//   [CustomInterestError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [PolicyViolationError]()
	//   [QuotaError]()
	//   [RequestError]()
	//   [StringLengthError]()
	MutateCustomInterests(context.Context, *MutateCustomInterestsRequest) (*MutateCustomInterestsResponse, error)
	mustEmbedUnimplementedCustomInterestServiceServer()
}

// UnimplementedCustomInterestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomInterestServiceServer struct {
}

func (UnimplementedCustomInterestServiceServer) GetCustomInterest(context.Context, *GetCustomInterestRequest) (*resources.CustomInterest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomInterest not implemented")
}
func (UnimplementedCustomInterestServiceServer) MutateCustomInterests(context.Context, *MutateCustomInterestsRequest) (*MutateCustomInterestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomInterests not implemented")
}
func (UnimplementedCustomInterestServiceServer) mustEmbedUnimplementedCustomInterestServiceServer() {}

// UnsafeCustomInterestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomInterestServiceServer will
// result in compilation errors.
type UnsafeCustomInterestServiceServer interface {
	mustEmbedUnimplementedCustomInterestServiceServer()
}

func RegisterCustomInterestServiceServer(s grpc.ServiceRegistrar, srv CustomInterestServiceServer) {
	s.RegisterService(&CustomInterestService_ServiceDesc, srv)
}

func _CustomInterestService_GetCustomInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomInterestServiceServer).GetCustomInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.CustomInterestService/GetCustomInterest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomInterestServiceServer).GetCustomInterest(ctx, req.(*GetCustomInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomInterestService_MutateCustomInterests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomInterestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomInterestServiceServer).MutateCustomInterests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.CustomInterestService/MutateCustomInterests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomInterestServiceServer).MutateCustomInterests(ctx, req.(*MutateCustomInterestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomInterestService_ServiceDesc is the grpc.ServiceDesc for CustomInterestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomInterestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.CustomInterestService",
	HandlerType: (*CustomInterestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomInterest",
			Handler:    _CustomInterestService_GetCustomInterest_Handler,
		},
		{
			MethodName: "MutateCustomInterests",
			Handler:    _CustomInterestService_MutateCustomInterests_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/custom_interest_service.proto",
}
