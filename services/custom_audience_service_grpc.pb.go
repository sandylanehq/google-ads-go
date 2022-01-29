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

// CustomAudienceServiceClient is the client API for CustomAudienceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomAudienceServiceClient interface {
	// Returns the requested custom audience in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCustomAudience(ctx context.Context, in *GetCustomAudienceRequest, opts ...grpc.CallOption) (*resources.CustomAudience, error)
	// Creates or updates custom audiences. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CustomAudienceError]()
	//   [CustomInterestError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [OperationAccessDeniedError]()
	//   [PolicyViolationError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCustomAudiences(ctx context.Context, in *MutateCustomAudiencesRequest, opts ...grpc.CallOption) (*MutateCustomAudiencesResponse, error)
}

type customAudienceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomAudienceServiceClient(cc grpc.ClientConnInterface) CustomAudienceServiceClient {
	return &customAudienceServiceClient{cc}
}

func (c *customAudienceServiceClient) GetCustomAudience(ctx context.Context, in *GetCustomAudienceRequest, opts ...grpc.CallOption) (*resources.CustomAudience, error) {
	out := new(resources.CustomAudience)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.CustomAudienceService/GetCustomAudience", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customAudienceServiceClient) MutateCustomAudiences(ctx context.Context, in *MutateCustomAudiencesRequest, opts ...grpc.CallOption) (*MutateCustomAudiencesResponse, error) {
	out := new(MutateCustomAudiencesResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.CustomAudienceService/MutateCustomAudiences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomAudienceServiceServer is the server API for CustomAudienceService service.
// All implementations must embed UnimplementedCustomAudienceServiceServer
// for forward compatibility
type CustomAudienceServiceServer interface {
	// Returns the requested custom audience in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCustomAudience(context.Context, *GetCustomAudienceRequest) (*resources.CustomAudience, error)
	// Creates or updates custom audiences. Operation statuses are returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [CustomAudienceError]()
	//   [CustomInterestError]()
	//   [FieldError]()
	//   [FieldMaskError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [MutateError]()
	//   [OperationAccessDeniedError]()
	//   [PolicyViolationError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCustomAudiences(context.Context, *MutateCustomAudiencesRequest) (*MutateCustomAudiencesResponse, error)
	mustEmbedUnimplementedCustomAudienceServiceServer()
}

// UnimplementedCustomAudienceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomAudienceServiceServer struct {
}

func (UnimplementedCustomAudienceServiceServer) GetCustomAudience(context.Context, *GetCustomAudienceRequest) (*resources.CustomAudience, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomAudience not implemented")
}
func (UnimplementedCustomAudienceServiceServer) MutateCustomAudiences(context.Context, *MutateCustomAudiencesRequest) (*MutateCustomAudiencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomAudiences not implemented")
}
func (UnimplementedCustomAudienceServiceServer) mustEmbedUnimplementedCustomAudienceServiceServer() {}

// UnsafeCustomAudienceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomAudienceServiceServer will
// result in compilation errors.
type UnsafeCustomAudienceServiceServer interface {
	mustEmbedUnimplementedCustomAudienceServiceServer()
}

func RegisterCustomAudienceServiceServer(s grpc.ServiceRegistrar, srv CustomAudienceServiceServer) {
	s.RegisterService(&CustomAudienceService_ServiceDesc, srv)
}

func _CustomAudienceService_GetCustomAudience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomAudienceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAudienceServiceServer).GetCustomAudience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.CustomAudienceService/GetCustomAudience",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAudienceServiceServer).GetCustomAudience(ctx, req.(*GetCustomAudienceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomAudienceService_MutateCustomAudiences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomAudiencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomAudienceServiceServer).MutateCustomAudiences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.CustomAudienceService/MutateCustomAudiences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomAudienceServiceServer).MutateCustomAudiences(ctx, req.(*MutateCustomAudiencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomAudienceService_ServiceDesc is the grpc.ServiceDesc for CustomAudienceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomAudienceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.CustomAudienceService",
	HandlerType: (*CustomAudienceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomAudience",
			Handler:    _CustomAudienceService_GetCustomAudience_Handler,
		},
		{
			MethodName: "MutateCustomAudiences",
			Handler:    _CustomAudienceService_MutateCustomAudiences_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/custom_audience_service.proto",
}
