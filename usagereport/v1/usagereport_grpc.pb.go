// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: usagereport/v1/usagereport.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UsageReportServiceClient is the client API for UsageReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsageReportServiceClient interface {
	Report(ctx context.Context, in *UsageReportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type usageReportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsageReportServiceClient(cc grpc.ClientConnInterface) UsageReportServiceClient {
	return &usageReportServiceClient{cc}
}

func (c *usageReportServiceClient) Report(ctx context.Context, in *UsageReportRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/usagereport.v1.UsageReportService/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsageReportServiceServer is the server API for UsageReportService service.
// All implementations must embed UnimplementedUsageReportServiceServer
// for forward compatibility
type UsageReportServiceServer interface {
	Report(context.Context, *UsageReportRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedUsageReportServiceServer()
}

// UnimplementedUsageReportServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUsageReportServiceServer struct {
}

func (UnimplementedUsageReportServiceServer) Report(context.Context, *UsageReportRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}
func (UnimplementedUsageReportServiceServer) mustEmbedUnimplementedUsageReportServiceServer() {}

// UnsafeUsageReportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsageReportServiceServer will
// result in compilation errors.
type UnsafeUsageReportServiceServer interface {
	mustEmbedUnimplementedUsageReportServiceServer()
}

func RegisterUsageReportServiceServer(s grpc.ServiceRegistrar, srv UsageReportServiceServer) {
	s.RegisterService(&UsageReportService_ServiceDesc, srv)
}

func _UsageReportService_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsageReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsageReportServiceServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usagereport.v1.UsageReportService/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsageReportServiceServer).Report(ctx, req.(*UsageReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsageReportService_ServiceDesc is the grpc.ServiceDesc for UsageReportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsageReportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usagereport.v1.UsageReportService",
	HandlerType: (*UsageReportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _UsageReportService_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usagereport/v1/usagereport.proto",
}
