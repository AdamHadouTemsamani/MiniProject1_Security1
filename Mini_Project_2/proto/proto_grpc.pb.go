// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// PingClient is the client API for Ping service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type pingClient struct {
	cc grpc.ClientConnInterface
}

func NewPingClient(cc grpc.ClientConnInterface) PingClient {
	return &pingClient{cc}
}

func (c *pingClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/Mini_Project_2.Ping/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingServer is the server API for Ping service.
// All implementations must embed UnimplementedPingServer
// for forward compatibility
type PingServer interface {
	Ping(context.Context, *Request) (*Reply, error)
	mustEmbedUnimplementedPingServer()
}

// UnimplementedPingServer must be embedded to have forward compatible implementations.
type UnimplementedPingServer struct {
}

func (UnimplementedPingServer) Ping(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedPingServer) mustEmbedUnimplementedPingServer() {}

// UnsafePingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PingServer will
// result in compilation errors.
type UnsafePingServer interface {
	mustEmbedUnimplementedPingServer()
}

func RegisterPingServer(s grpc.ServiceRegistrar, srv PingServer) {
	s.RegisterService(&Ping_ServiceDesc, srv)
}

func _Ping_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mini_Project_2.Ping/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Ping_ServiceDesc is the grpc.ServiceDesc for Ping service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ping_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Mini_Project_2.Ping",
	HandlerType: (*PingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Ping_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/proto.proto",
}

// SendSharesClient is the client API for SendShares service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendSharesClient interface {
	SendShares(ctx context.Context, in *Share, opts ...grpc.CallOption) (*Acknoledgement, error)
}

type sendSharesClient struct {
	cc grpc.ClientConnInterface
}

func NewSendSharesClient(cc grpc.ClientConnInterface) SendSharesClient {
	return &sendSharesClient{cc}
}

func (c *sendSharesClient) SendShares(ctx context.Context, in *Share, opts ...grpc.CallOption) (*Acknoledgement, error) {
	out := new(Acknoledgement)
	err := c.cc.Invoke(ctx, "/Mini_Project_2.SendShares/SendShares", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendSharesServer is the server API for SendShares service.
// All implementations must embed UnimplementedSendSharesServer
// for forward compatibility
type SendSharesServer interface {
	SendShares(context.Context, *Share) (*Acknoledgement, error)
	mustEmbedUnimplementedSendSharesServer()
}

// UnimplementedSendSharesServer must be embedded to have forward compatible implementations.
type UnimplementedSendSharesServer struct {
}

func (UnimplementedSendSharesServer) SendShares(context.Context, *Share) (*Acknoledgement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendShares not implemented")
}
func (UnimplementedSendSharesServer) mustEmbedUnimplementedSendSharesServer() {}

// UnsafeSendSharesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendSharesServer will
// result in compilation errors.
type UnsafeSendSharesServer interface {
	mustEmbedUnimplementedSendSharesServer()
}

func RegisterSendSharesServer(s grpc.ServiceRegistrar, srv SendSharesServer) {
	s.RegisterService(&SendShares_ServiceDesc, srv)
}

func _SendShares_SendShares_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Share)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendSharesServer).SendShares(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Mini_Project_2.SendShares/SendShares",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendSharesServer).SendShares(ctx, req.(*Share))
	}
	return interceptor(ctx, in, info, handler)
}

// SendShares_ServiceDesc is the grpc.ServiceDesc for SendShares service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendShares_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Mini_Project_2.SendShares",
	HandlerType: (*SendSharesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendShares",
			Handler:    _SendShares_SendShares_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/proto.proto",
}