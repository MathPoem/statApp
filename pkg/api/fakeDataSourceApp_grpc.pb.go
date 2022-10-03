// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/fakeDataSourceApp.proto

package api

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

// FakeDataSourceAppServiceClient is the client API for FakeDataSourceAppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FakeDataSourceAppServiceClient interface {
	GetData(ctx context.Context, in *RequestCurrencyInfo, opts ...grpc.CallOption) (*ResponseCurrencyInfo, error)
}

type fakeDataSourceAppServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFakeDataSourceAppServiceClient(cc grpc.ClientConnInterface) FakeDataSourceAppServiceClient {
	return &fakeDataSourceAppServiceClient{cc}
}

func (c *fakeDataSourceAppServiceClient) GetData(ctx context.Context, in *RequestCurrencyInfo, opts ...grpc.CallOption) (*ResponseCurrencyInfo, error) {
	out := new(ResponseCurrencyInfo)
	err := c.cc.Invoke(ctx, "/fakeDataSourceApp.fakeDataSourceAppService/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FakeDataSourceAppServiceServer is the server API for FakeDataSourceAppService service.
// All implementations must embed UnimplementedFakeDataSourceAppServiceServer
// for forward compatibility
type FakeDataSourceAppServiceServer interface {
	GetData(context.Context, *RequestCurrencyInfo) (*ResponseCurrencyInfo, error)
	mustEmbedUnimplementedFakeDataSourceAppServiceServer()
}

// UnimplementedFakeDataSourceAppServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFakeDataSourceAppServiceServer struct {
}

func (UnimplementedFakeDataSourceAppServiceServer) GetData(context.Context, *RequestCurrencyInfo) (*ResponseCurrencyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedFakeDataSourceAppServiceServer) mustEmbedUnimplementedFakeDataSourceAppServiceServer() {
}

// UnsafeFakeDataSourceAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FakeDataSourceAppServiceServer will
// result in compilation errors.
type UnsafeFakeDataSourceAppServiceServer interface {
	mustEmbedUnimplementedFakeDataSourceAppServiceServer()
}

func RegisterFakeDataSourceAppServiceServer(s grpc.ServiceRegistrar, srv FakeDataSourceAppServiceServer) {
	s.RegisterService(&FakeDataSourceAppService_ServiceDesc, srv)
}

func _FakeDataSourceAppService_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCurrencyInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FakeDataSourceAppServiceServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fakeDataSourceApp.fakeDataSourceAppService/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FakeDataSourceAppServiceServer).GetData(ctx, req.(*RequestCurrencyInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// FakeDataSourceAppService_ServiceDesc is the grpc.ServiceDesc for FakeDataSourceAppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FakeDataSourceAppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fakeDataSourceApp.fakeDataSourceAppService",
	HandlerType: (*FakeDataSourceAppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _FakeDataSourceAppService_GetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/fakeDataSourceApp.proto",
}
