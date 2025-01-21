// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: api/service.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	StoreService_GetMenu_FullMethodName     = "/StoreService/getMenu"
	StoreService_PlaceOrder_FullMethodName  = "/StoreService/placeOrder"
	StoreService_CheckStatus_FullMethodName = "/StoreService/checkStatus"
	StoreService_CancelOrder_FullMethodName = "/StoreService/cancelOrder"
)

// StoreServiceClient is the client API for StoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreServiceClient interface {
	GetMenu(ctx context.Context, in *MenuRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OrderMenu], error)
	PlaceOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Receipt, error)
	CheckStatus(ctx context.Context, in *Receipt, opts ...grpc.CallOption) (*OrderStatus, error)
	CancelOrder(ctx context.Context, in *Receipt, opts ...grpc.CallOption) (*Receipt, error)
}

type storeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreServiceClient(cc grpc.ClientConnInterface) StoreServiceClient {
	return &storeServiceClient{cc}
}

func (c *storeServiceClient) GetMenu(ctx context.Context, in *MenuRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[OrderMenu], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StoreService_ServiceDesc.Streams[0], StoreService_GetMenu_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[MenuRequest, OrderMenu]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StoreService_GetMenuClient = grpc.ServerStreamingClient[OrderMenu]

func (c *storeServiceClient) PlaceOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Receipt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Receipt)
	err := c.cc.Invoke(ctx, StoreService_PlaceOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) CheckStatus(ctx context.Context, in *Receipt, opts ...grpc.CallOption) (*OrderStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrderStatus)
	err := c.cc.Invoke(ctx, StoreService_CheckStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeServiceClient) CancelOrder(ctx context.Context, in *Receipt, opts ...grpc.CallOption) (*Receipt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Receipt)
	err := c.cc.Invoke(ctx, StoreService_CancelOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServiceServer is the server API for StoreService service.
// All implementations must embed UnimplementedStoreServiceServer
// for forward compatibility.
type StoreServiceServer interface {
	GetMenu(*MenuRequest, grpc.ServerStreamingServer[OrderMenu]) error
	PlaceOrder(context.Context, *Order) (*Receipt, error)
	CheckStatus(context.Context, *Receipt) (*OrderStatus, error)
	CancelOrder(context.Context, *Receipt) (*Receipt, error)
	mustEmbedUnimplementedStoreServiceServer()
}

// UnimplementedStoreServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStoreServiceServer struct{}

func (UnimplementedStoreServiceServer) GetMenu(*MenuRequest, grpc.ServerStreamingServer[OrderMenu]) error {
	return status.Errorf(codes.Unimplemented, "method GetMenu not implemented")
}
func (UnimplementedStoreServiceServer) PlaceOrder(context.Context, *Order) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}
func (UnimplementedStoreServiceServer) CheckStatus(context.Context, *Receipt) (*OrderStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckStatus not implemented")
}
func (UnimplementedStoreServiceServer) CancelOrder(context.Context, *Receipt) (*Receipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedStoreServiceServer) mustEmbedUnimplementedStoreServiceServer() {}
func (UnimplementedStoreServiceServer) testEmbeddedByValue()                      {}

// UnsafeStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreServiceServer will
// result in compilation errors.
type UnsafeStoreServiceServer interface {
	mustEmbedUnimplementedStoreServiceServer()
}

func RegisterStoreServiceServer(s grpc.ServiceRegistrar, srv StoreServiceServer) {
	// If the following call pancis, it indicates UnimplementedStoreServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StoreService_ServiceDesc, srv)
}

func _StoreService_GetMenu_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MenuRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StoreServiceServer).GetMenu(m, &grpc.GenericServerStream[MenuRequest, OrderMenu]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StoreService_GetMenuServer = grpc.ServerStreamingServer[OrderMenu]

func _StoreService_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreService_PlaceOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).PlaceOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_CheckStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Receipt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).CheckStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreService_CheckStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).CheckStatus(ctx, req.(*Receipt))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Receipt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreService_CancelOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServiceServer).CancelOrder(ctx, req.(*Receipt))
	}
	return interceptor(ctx, in, info, handler)
}

// StoreService_ServiceDesc is the grpc.ServiceDesc for StoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StoreService",
	HandlerType: (*StoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "placeOrder",
			Handler:    _StoreService_PlaceOrder_Handler,
		},
		{
			MethodName: "checkStatus",
			Handler:    _StoreService_CheckStatus_Handler,
		},
		{
			MethodName: "cancelOrder",
			Handler:    _StoreService_CancelOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getMenu",
			Handler:       _StoreService_GetMenu_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/service.proto",
}
