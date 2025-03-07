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

// CustomerUserAccessInvitationServiceClient is the client API for CustomerUserAccessInvitationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerUserAccessInvitationServiceClient interface {
	// Returns the requested access invitation in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCustomerUserAccessInvitation(ctx context.Context, in *GetCustomerUserAccessInvitationRequest, opts ...grpc.CallOption) (*resources.CustomerUserAccessInvitation, error)
	// Creates or removes an access invitation.
	//
	// List of thrown errors:
	//   [AccessInvitationError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCustomerUserAccessInvitation(ctx context.Context, in *MutateCustomerUserAccessInvitationRequest, opts ...grpc.CallOption) (*MutateCustomerUserAccessInvitationResponse, error)
}

type customerUserAccessInvitationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerUserAccessInvitationServiceClient(cc grpc.ClientConnInterface) CustomerUserAccessInvitationServiceClient {
	return &customerUserAccessInvitationServiceClient{cc}
}

func (c *customerUserAccessInvitationServiceClient) GetCustomerUserAccessInvitation(ctx context.Context, in *GetCustomerUserAccessInvitationRequest, opts ...grpc.CallOption) (*resources.CustomerUserAccessInvitation, error) {
	out := new(resources.CustomerUserAccessInvitation)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.CustomerUserAccessInvitationService/GetCustomerUserAccessInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerUserAccessInvitationServiceClient) MutateCustomerUserAccessInvitation(ctx context.Context, in *MutateCustomerUserAccessInvitationRequest, opts ...grpc.CallOption) (*MutateCustomerUserAccessInvitationResponse, error) {
	out := new(MutateCustomerUserAccessInvitationResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.CustomerUserAccessInvitationService/MutateCustomerUserAccessInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerUserAccessInvitationServiceServer is the server API for CustomerUserAccessInvitationService service.
// All implementations must embed UnimplementedCustomerUserAccessInvitationServiceServer
// for forward compatibility
type CustomerUserAccessInvitationServiceServer interface {
	// Returns the requested access invitation in full detail.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetCustomerUserAccessInvitation(context.Context, *GetCustomerUserAccessInvitationRequest) (*resources.CustomerUserAccessInvitation, error)
	// Creates or removes an access invitation.
	//
	// List of thrown errors:
	//   [AccessInvitationError]()
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	MutateCustomerUserAccessInvitation(context.Context, *MutateCustomerUserAccessInvitationRequest) (*MutateCustomerUserAccessInvitationResponse, error)
	mustEmbedUnimplementedCustomerUserAccessInvitationServiceServer()
}

// UnimplementedCustomerUserAccessInvitationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerUserAccessInvitationServiceServer struct {
}

func (UnimplementedCustomerUserAccessInvitationServiceServer) GetCustomerUserAccessInvitation(context.Context, *GetCustomerUserAccessInvitationRequest) (*resources.CustomerUserAccessInvitation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerUserAccessInvitation not implemented")
}
func (UnimplementedCustomerUserAccessInvitationServiceServer) MutateCustomerUserAccessInvitation(context.Context, *MutateCustomerUserAccessInvitationRequest) (*MutateCustomerUserAccessInvitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateCustomerUserAccessInvitation not implemented")
}
func (UnimplementedCustomerUserAccessInvitationServiceServer) mustEmbedUnimplementedCustomerUserAccessInvitationServiceServer() {
}

// UnsafeCustomerUserAccessInvitationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerUserAccessInvitationServiceServer will
// result in compilation errors.
type UnsafeCustomerUserAccessInvitationServiceServer interface {
	mustEmbedUnimplementedCustomerUserAccessInvitationServiceServer()
}

func RegisterCustomerUserAccessInvitationServiceServer(s grpc.ServiceRegistrar, srv CustomerUserAccessInvitationServiceServer) {
	s.RegisterService(&CustomerUserAccessInvitationService_ServiceDesc, srv)
}

func _CustomerUserAccessInvitationService_GetCustomerUserAccessInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomerUserAccessInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerUserAccessInvitationServiceServer).GetCustomerUserAccessInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.CustomerUserAccessInvitationService/GetCustomerUserAccessInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerUserAccessInvitationServiceServer).GetCustomerUserAccessInvitation(ctx, req.(*GetCustomerUserAccessInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerUserAccessInvitationService_MutateCustomerUserAccessInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateCustomerUserAccessInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerUserAccessInvitationServiceServer).MutateCustomerUserAccessInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.CustomerUserAccessInvitationService/MutateCustomerUserAccessInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerUserAccessInvitationServiceServer).MutateCustomerUserAccessInvitation(ctx, req.(*MutateCustomerUserAccessInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerUserAccessInvitationService_ServiceDesc is the grpc.ServiceDesc for CustomerUserAccessInvitationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerUserAccessInvitationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.CustomerUserAccessInvitationService",
	HandlerType: (*CustomerUserAccessInvitationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCustomerUserAccessInvitation",
			Handler:    _CustomerUserAccessInvitationService_GetCustomerUserAccessInvitation_Handler,
		},
		{
			MethodName: "MutateCustomerUserAccessInvitation",
			Handler:    _CustomerUserAccessInvitationService_MutateCustomerUserAccessInvitation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/customer_user_access_invitation_service.proto",
}
