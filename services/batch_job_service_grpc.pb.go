// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	resources "github.com/butlerhq/google-ads-go/resources"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BatchJobServiceClient is the client API for BatchJobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BatchJobServiceClient interface {
	// Mutates a batch job.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	MutateBatchJob(ctx context.Context, in *MutateBatchJobRequest, opts ...grpc.CallOption) (*MutateBatchJobResponse, error)
	// Returns the batch job.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetBatchJob(ctx context.Context, in *GetBatchJobRequest, opts ...grpc.CallOption) (*resources.BatchJob, error)
	// Returns the results of the batch job. The job must be done.
	// Supports standard list paging.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [BatchJobError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	ListBatchJobResults(ctx context.Context, in *ListBatchJobResultsRequest, opts ...grpc.CallOption) (*ListBatchJobResultsResponse, error)
	// Runs the batch job.
	//
	// The Operation.metadata field type is BatchJobMetadata. When finished, the
	// long running operation will not contain errors or a response. Instead, use
	// ListBatchJobResults to get the results of the job.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [BatchJobError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	RunBatchJob(ctx context.Context, in *RunBatchJobRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Add operations to the batch job.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [BatchJobError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	AddBatchJobOperations(ctx context.Context, in *AddBatchJobOperationsRequest, opts ...grpc.CallOption) (*AddBatchJobOperationsResponse, error)
}

type batchJobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBatchJobServiceClient(cc grpc.ClientConnInterface) BatchJobServiceClient {
	return &batchJobServiceClient{cc}
}

func (c *batchJobServiceClient) MutateBatchJob(ctx context.Context, in *MutateBatchJobRequest, opts ...grpc.CallOption) (*MutateBatchJobResponse, error) {
	out := new(MutateBatchJobResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.BatchJobService/MutateBatchJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *batchJobServiceClient) GetBatchJob(ctx context.Context, in *GetBatchJobRequest, opts ...grpc.CallOption) (*resources.BatchJob, error) {
	out := new(resources.BatchJob)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.BatchJobService/GetBatchJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *batchJobServiceClient) ListBatchJobResults(ctx context.Context, in *ListBatchJobResultsRequest, opts ...grpc.CallOption) (*ListBatchJobResultsResponse, error) {
	out := new(ListBatchJobResultsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.BatchJobService/ListBatchJobResults", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *batchJobServiceClient) RunBatchJob(ctx context.Context, in *RunBatchJobRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.BatchJobService/RunBatchJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *batchJobServiceClient) AddBatchJobOperations(ctx context.Context, in *AddBatchJobOperationsRequest, opts ...grpc.CallOption) (*AddBatchJobOperationsResponse, error) {
	out := new(AddBatchJobOperationsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v9.services.BatchJobService/AddBatchJobOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BatchJobServiceServer is the server API for BatchJobService service.
// All implementations must embed UnimplementedBatchJobServiceServer
// for forward compatibility
type BatchJobServiceServer interface {
	// Mutates a batch job.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	MutateBatchJob(context.Context, *MutateBatchJobRequest) (*MutateBatchJobResponse, error)
	// Returns the batch job.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	GetBatchJob(context.Context, *GetBatchJobRequest) (*resources.BatchJob, error)
	// Returns the results of the batch job. The job must be done.
	// Supports standard list paging.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [BatchJobError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	ListBatchJobResults(context.Context, *ListBatchJobResultsRequest) (*ListBatchJobResultsResponse, error)
	// Runs the batch job.
	//
	// The Operation.metadata field type is BatchJobMetadata. When finished, the
	// long running operation will not contain errors or a response. Instead, use
	// ListBatchJobResults to get the results of the job.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [BatchJobError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	RunBatchJob(context.Context, *RunBatchJobRequest) (*longrunning.Operation, error)
	// Add operations to the batch job.
	//
	// List of thrown errors:
	//   [AuthenticationError]()
	//   [AuthorizationError]()
	//   [BatchJobError]()
	//   [HeaderError]()
	//   [InternalError]()
	//   [QuotaError]()
	//   [RequestError]()
	//   [ResourceCountLimitExceededError]()
	AddBatchJobOperations(context.Context, *AddBatchJobOperationsRequest) (*AddBatchJobOperationsResponse, error)
	mustEmbedUnimplementedBatchJobServiceServer()
}

// UnimplementedBatchJobServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBatchJobServiceServer struct {
}

func (UnimplementedBatchJobServiceServer) MutateBatchJob(context.Context, *MutateBatchJobRequest) (*MutateBatchJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateBatchJob not implemented")
}
func (UnimplementedBatchJobServiceServer) GetBatchJob(context.Context, *GetBatchJobRequest) (*resources.BatchJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBatchJob not implemented")
}
func (UnimplementedBatchJobServiceServer) ListBatchJobResults(context.Context, *ListBatchJobResultsRequest) (*ListBatchJobResultsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBatchJobResults not implemented")
}
func (UnimplementedBatchJobServiceServer) RunBatchJob(context.Context, *RunBatchJobRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunBatchJob not implemented")
}
func (UnimplementedBatchJobServiceServer) AddBatchJobOperations(context.Context, *AddBatchJobOperationsRequest) (*AddBatchJobOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBatchJobOperations not implemented")
}
func (UnimplementedBatchJobServiceServer) mustEmbedUnimplementedBatchJobServiceServer() {}

// UnsafeBatchJobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BatchJobServiceServer will
// result in compilation errors.
type UnsafeBatchJobServiceServer interface {
	mustEmbedUnimplementedBatchJobServiceServer()
}

func RegisterBatchJobServiceServer(s grpc.ServiceRegistrar, srv BatchJobServiceServer) {
	s.RegisterService(&BatchJobService_ServiceDesc, srv)
}

func _BatchJobService_MutateBatchJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateBatchJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchJobServiceServer).MutateBatchJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.BatchJobService/MutateBatchJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchJobServiceServer).MutateBatchJob(ctx, req.(*MutateBatchJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BatchJobService_GetBatchJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBatchJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchJobServiceServer).GetBatchJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.BatchJobService/GetBatchJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchJobServiceServer).GetBatchJob(ctx, req.(*GetBatchJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BatchJobService_ListBatchJobResults_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBatchJobResultsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchJobServiceServer).ListBatchJobResults(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.BatchJobService/ListBatchJobResults",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchJobServiceServer).ListBatchJobResults(ctx, req.(*ListBatchJobResultsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BatchJobService_RunBatchJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunBatchJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchJobServiceServer).RunBatchJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.BatchJobService/RunBatchJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchJobServiceServer).RunBatchJob(ctx, req.(*RunBatchJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BatchJobService_AddBatchJobOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBatchJobOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BatchJobServiceServer).AddBatchJobOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v9.services.BatchJobService/AddBatchJobOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BatchJobServiceServer).AddBatchJobOperations(ctx, req.(*AddBatchJobOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BatchJobService_ServiceDesc is the grpc.ServiceDesc for BatchJobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BatchJobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v9.services.BatchJobService",
	HandlerType: (*BatchJobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateBatchJob",
			Handler:    _BatchJobService_MutateBatchJob_Handler,
		},
		{
			MethodName: "GetBatchJob",
			Handler:    _BatchJobService_GetBatchJob_Handler,
		},
		{
			MethodName: "ListBatchJobResults",
			Handler:    _BatchJobService_ListBatchJobResults_Handler,
		},
		{
			MethodName: "RunBatchJob",
			Handler:    _BatchJobService_RunBatchJob_Handler,
		},
		{
			MethodName: "AddBatchJobOperations",
			Handler:    _BatchJobService_AddBatchJobOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v9/services/batch_job_service.proto",
}
