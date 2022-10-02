// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/statApp.proto

package pkg

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

// StatAppServiceClient is the client API for StatAppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatAppServiceClient interface {
	DayPriceAvg(ctx context.Context, in *RequestDayPriceAvg, opts ...grpc.CallOption) (*ResponseDayPriceAvg, error)
}

type statAppServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatAppServiceClient(cc grpc.ClientConnInterface) StatAppServiceClient {
	return &statAppServiceClient{cc}
}

func (c *statAppServiceClient) DayPriceAvg(ctx context.Context, in *RequestDayPriceAvg, opts ...grpc.CallOption) (*ResponseDayPriceAvg, error) {
	out := new(ResponseDayPriceAvg)
	err := c.cc.Invoke(ctx, "/statApp.StatAppService/DayPriceAvg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatAppServiceServer is the server API for StatAppService service.
// All implementations must embed UnimplementedStatAppServiceServer
// for forward compatibility
type StatAppServiceServer interface {
	DayPriceAvg(context.Context, *RequestDayPriceAvg) (*ResponseDayPriceAvg, error)
	mustEmbedUnimplementedStatAppServiceServer()
}

// UnimplementedStatAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStatAppServiceServer struct {
}

func (UnimplementedStatAppServiceServer) DayPriceAvg(context.Context, *RequestDayPriceAvg) (*ResponseDayPriceAvg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DayPriceAvg not implemented")
}
func (UnimplementedStatAppServiceServer) mustEmbedUnimplementedStatAppServiceServer() {}

// UnsafeStatAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatAppServiceServer will
// result in compilation errors.
type UnsafeStatAppServiceServer interface {
	mustEmbedUnimplementedStatAppServiceServer()
}

func RegisterStatAppServiceServer(s grpc.ServiceRegistrar, srv StatAppServiceServer) {
	s.RegisterService(&StatAppService_ServiceDesc, srv)
}

func _StatAppService_DayPriceAvg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDayPriceAvg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatAppServiceServer).DayPriceAvg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statApp.StatAppService/DayPriceAvg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatAppServiceServer).DayPriceAvg(ctx, req.(*RequestDayPriceAvg))
	}
	return interceptor(ctx, in, info, handler)
}

// StatAppService_ServiceDesc is the grpc.ServiceDesc for StatAppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatAppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "statApp.StatAppService",
	HandlerType: (*StatAppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DayPriceAvg",
			Handler:    _StatAppService_DayPriceAvg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/statApp.proto",
}