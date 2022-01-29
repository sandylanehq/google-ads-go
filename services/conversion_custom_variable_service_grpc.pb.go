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

// ConversionCustomVariableServiceClient is the client API for ConversionCustomVariableService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConversionCustomVariableServiceClient interface {
	// Returns the requested conversion custom variable.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetConversionCustomVariable(ctx context.Context, in *GetConversionCustomVariableRequest, opts ...grpc.CallOption) (*resources.ConversionCustomVariable, error)
	// Creates or updates conversion custom variables. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ConversionCustomVariableError]()
	//   [DatabaseError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateConversionCustomVariables(ctx context.Context, in *MutateConversionCustomVariablesRequest, opts ...grpc.CallOption) (*MutateConversionCustomVariablesResponse, error)
}

type conversionCustomVariableServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConversionCustomVariableServiceClient(cc grpc.ClientConnInterface) ConversionCustomVariableServiceClient {
	return &conversionCustomVariableServiceClient{cc}
}

func (c *conversionCustomVariableServiceClient) GetConversionCustomVariable(ctx context.Context, in *GetConversionCustomVariableRequest, opts ...grpc.CallOption) (*resources.ConversionCustomVariable, error) {
	out := new(resources.ConversionCustomVariable)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.ConversionCustomVariableService/GetConversionCustomVariable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversionCustomVariableServiceClient) MutateConversionCustomVariables(ctx context.Context, in *MutateConversionCustomVariablesRequest, opts ...grpc.CallOption) (*MutateConversionCustomVariablesResponse, error) {
	out := new(MutateConversionCustomVariablesResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.ConversionCustomVariableService/MutateConversionCustomVariables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversionCustomVariableServiceServer is the server API for ConversionCustomVariableService service.
// All implementations must embed UnimplementedConversionCustomVariableServiceServer
// for forward compatibility
type ConversionCustomVariableServiceServer interface {
	// Returns the requested conversion custom variable.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetConversionCustomVariable(context.Context, *GetConversionCustomVariableRequest) (*resources.ConversionCustomVariable, error)
	// Creates or updates conversion custom variables. Operation statuses are
	// returned.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [ConversionCustomVariableError]()
	//   [DatabaseError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateConversionCustomVariables(context.Context, *MutateConversionCustomVariablesRequest) (*MutateConversionCustomVariablesResponse, error)
	mustEmbedUnimplementedConversionCustomVariableServiceServer()
}

// UnimplementedConversionCustomVariableServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConversionCustomVariableServiceServer struct {
}

func (UnimplementedConversionCustomVariableServiceServer) GetConversionCustomVariable(context.Context, *GetConversionCustomVariableRequest) (*resources.ConversionCustomVariable, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversionCustomVariable not implemented")
}
func (UnimplementedConversionCustomVariableServiceServer) MutateConversionCustomVariables(context.Context, *MutateConversionCustomVariablesRequest) (*MutateConversionCustomVariablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateConversionCustomVariables not implemented")
}
func (UnimplementedConversionCustomVariableServiceServer) mustEmbedUnimplementedConversionCustomVariableServiceServer() {
}

// UnsafeConversionCustomVariableServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConversionCustomVariableServiceServer will
// result in compilation errors.
type UnsafeConversionCustomVariableServiceServer interface {
	mustEmbedUnimplementedConversionCustomVariableServiceServer()
}

func RegisterConversionCustomVariableServiceServer(s grpc.ServiceRegistrar, srv ConversionCustomVariableServiceServer) {
	s.RegisterService(&ConversionCustomVariableService_ServiceDesc, srv)
}

func _ConversionCustomVariableService_GetConversionCustomVariable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversionCustomVariableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionCustomVariableServiceServer).GetConversionCustomVariable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.ConversionCustomVariableService/GetConversionCustomVariable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionCustomVariableServiceServer).GetConversionCustomVariable(ctx, req.(*GetConversionCustomVariableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConversionCustomVariableService_MutateConversionCustomVariables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateConversionCustomVariablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversionCustomVariableServiceServer).MutateConversionCustomVariables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.ConversionCustomVariableService/MutateConversionCustomVariables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversionCustomVariableServiceServer).MutateConversionCustomVariables(ctx, req.(*MutateConversionCustomVariablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConversionCustomVariableService_ServiceDesc is the grpc.ServiceDesc for ConversionCustomVariableService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConversionCustomVariableService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.ConversionCustomVariableService",
	HandlerType: (*ConversionCustomVariableServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConversionCustomVariable",
			Handler:    _ConversionCustomVariableService_GetConversionCustomVariable_Handler,
		},
		{
			MethodName: "MutateConversionCustomVariables",
			Handler:    _ConversionCustomVariableService_MutateConversionCustomVariables_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/conversion_custom_variable_service.proto",
}
