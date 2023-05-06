// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: internal/domain/proto/auto_market.proto

package grpc

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

const (
	AutoMarket_CreateAccount_FullMethodName = "/AutoMarket/CreateAccount"
)

// AutoMarketClient is the client API for AutoMarket service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AutoMarketClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
}

type autoMarketClient struct {
	cc grpc.ClientConnInterface
}

func NewAutoMarketClient(cc grpc.ClientConnInterface) AutoMarketClient {
	return &autoMarketClient{cc}
}

func (c *autoMarketClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, AutoMarket_CreateAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutoMarketServer is the server API for AutoMarket service.
// All implementations must embed UnimplementedAutoMarketServer
// for forward compatibility
type AutoMarketServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	mustEmbedUnimplementedAutoMarketServer()
}

// UnimplementedAutoMarketServer must be embedded to have forward compatible implementations.
type UnimplementedAutoMarketServer struct {
}

func (UnimplementedAutoMarketServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAutoMarketServer) mustEmbedUnimplementedAutoMarketServer() {}

// UnsafeAutoMarketServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AutoMarketServer will
// result in compilation errors.
type UnsafeAutoMarketServer interface {
	mustEmbedUnimplementedAutoMarketServer()
}

func RegisterAutoMarketServer(s grpc.ServiceRegistrar, srv AutoMarketServer) {
	s.RegisterService(&AutoMarket_ServiceDesc, srv)
}

func _AutoMarket_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutoMarketServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AutoMarket_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutoMarketServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AutoMarket_ServiceDesc is the grpc.ServiceDesc for AutoMarket service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AutoMarket_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AutoMarket",
	HandlerType: (*AutoMarketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AutoMarket_CreateAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/domain/proto/auto_market.proto",
}
