// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: chiave.proto

package chiave

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

// ChiaveClient is the client API for Chiave service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChiaveClient interface {
	Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*GetResponse, error)
	Increment(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Decrement(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	ProcessEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type chiaveClient struct {
	cc grpc.ClientConnInterface
}

func NewChiaveClient(cc grpc.ClientConnInterface) ChiaveClient {
	return &chiaveClient{cc}
}

func (c *chiaveClient) Get(ctx context.Context, in *Request, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/chiave.Chiave/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chiaveClient) Increment(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/chiave.Chiave/Increment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chiaveClient) Decrement(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/chiave.Chiave/Decrement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chiaveClient) Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/chiave.Chiave/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chiaveClient) ProcessEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chiave.Chiave/ProcessEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChiaveServer is the server API for Chiave service.
// All implementations must embed UnimplementedChiaveServer
// for forward compatibility
type ChiaveServer interface {
	Get(context.Context, *Request) (*GetResponse, error)
	Increment(context.Context, *Request) (*Response, error)
	Decrement(context.Context, *Request) (*Response, error)
	Add(context.Context, *Request) (*Response, error)
	ProcessEvent(context.Context, *Event) (*emptypb.Empty, error)
	mustEmbedUnimplementedChiaveServer()
}

// UnimplementedChiaveServer must be embedded to have forward compatible implementations.
type UnimplementedChiaveServer struct {
}

func (UnimplementedChiaveServer) Get(context.Context, *Request) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedChiaveServer) Increment(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Increment not implemented")
}
func (UnimplementedChiaveServer) Decrement(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrement not implemented")
}
func (UnimplementedChiaveServer) Add(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedChiaveServer) ProcessEvent(context.Context, *Event) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessEvent not implemented")
}
func (UnimplementedChiaveServer) mustEmbedUnimplementedChiaveServer() {}

// UnsafeChiaveServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChiaveServer will
// result in compilation errors.
type UnsafeChiaveServer interface {
	mustEmbedUnimplementedChiaveServer()
}

func RegisterChiaveServer(s grpc.ServiceRegistrar, srv ChiaveServer) {
	s.RegisterService(&Chiave_ServiceDesc, srv)
}

func _Chiave_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChiaveServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chiave.Chiave/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChiaveServer).Get(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chiave_Increment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChiaveServer).Increment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chiave.Chiave/Increment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChiaveServer).Increment(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chiave_Decrement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChiaveServer).Decrement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chiave.Chiave/Decrement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChiaveServer).Decrement(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chiave_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChiaveServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chiave.Chiave/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChiaveServer).Add(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chiave_ProcessEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChiaveServer).ProcessEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chiave.Chiave/ProcessEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChiaveServer).ProcessEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

// Chiave_ServiceDesc is the grpc.ServiceDesc for Chiave service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chiave_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chiave.Chiave",
	HandlerType: (*ChiaveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Chiave_Get_Handler,
		},
		{
			MethodName: "Increment",
			Handler:    _Chiave_Increment_Handler,
		},
		{
			MethodName: "Decrement",
			Handler:    _Chiave_Decrement_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Chiave_Add_Handler,
		},
		{
			MethodName: "ProcessEvent",
			Handler:    _Chiave_ProcessEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chiave.proto",
}
