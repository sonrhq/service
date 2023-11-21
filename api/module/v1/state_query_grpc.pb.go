// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package modulev1

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

// StateQueryServiceClient is the client API for StateQueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StateQueryServiceClient interface {
	// Get queries the Balance table by its primary key.
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
	// ListBalance queries the Balance table using prefix and range queries against defined indexes.
	ListBalance(ctx context.Context, in *ListBalanceRequest, opts ...grpc.CallOption) (*ListBalanceResponse, error)
	// Get queries the Supply table by its primary key.
	GetSupply(ctx context.Context, in *GetSupplyRequest, opts ...grpc.CallOption) (*GetSupplyResponse, error)
	// ListSupply queries the Supply table using prefix and range queries against defined indexes.
	ListSupply(ctx context.Context, in *ListSupplyRequest, opts ...grpc.CallOption) (*ListSupplyResponse, error)
}

type stateQueryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStateQueryServiceClient(cc grpc.ClientConnInterface) StateQueryServiceClient {
	return &stateQueryServiceClient{cc}
}

func (c *stateQueryServiceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := c.cc.Invoke(ctx, "/sonrhq.service.module.v1.StateQueryService/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateQueryServiceClient) ListBalance(ctx context.Context, in *ListBalanceRequest, opts ...grpc.CallOption) (*ListBalanceResponse, error) {
	out := new(ListBalanceResponse)
	err := c.cc.Invoke(ctx, "/sonrhq.service.module.v1.StateQueryService/ListBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateQueryServiceClient) GetSupply(ctx context.Context, in *GetSupplyRequest, opts ...grpc.CallOption) (*GetSupplyResponse, error) {
	out := new(GetSupplyResponse)
	err := c.cc.Invoke(ctx, "/sonrhq.service.module.v1.StateQueryService/GetSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateQueryServiceClient) ListSupply(ctx context.Context, in *ListSupplyRequest, opts ...grpc.CallOption) (*ListSupplyResponse, error) {
	out := new(ListSupplyResponse)
	err := c.cc.Invoke(ctx, "/sonrhq.service.module.v1.StateQueryService/ListSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StateQueryServiceServer is the server API for StateQueryService service.
// All implementations must embed UnimplementedStateQueryServiceServer
// for forward compatibility
type StateQueryServiceServer interface {
	// Get queries the Balance table by its primary key.
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
	// ListBalance queries the Balance table using prefix and range queries against defined indexes.
	ListBalance(context.Context, *ListBalanceRequest) (*ListBalanceResponse, error)
	// Get queries the Supply table by its primary key.
	GetSupply(context.Context, *GetSupplyRequest) (*GetSupplyResponse, error)
	// ListSupply queries the Supply table using prefix and range queries against defined indexes.
	ListSupply(context.Context, *ListSupplyRequest) (*ListSupplyResponse, error)
	mustEmbedUnimplementedStateQueryServiceServer()
}

// UnimplementedStateQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStateQueryServiceServer struct {
}

func (UnimplementedStateQueryServiceServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedStateQueryServiceServer) ListBalance(context.Context, *ListBalanceRequest) (*ListBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBalance not implemented")
}
func (UnimplementedStateQueryServiceServer) GetSupply(context.Context, *GetSupplyRequest) (*GetSupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupply not implemented")
}
func (UnimplementedStateQueryServiceServer) ListSupply(context.Context, *ListSupplyRequest) (*ListSupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSupply not implemented")
}
func (UnimplementedStateQueryServiceServer) mustEmbedUnimplementedStateQueryServiceServer() {}

// UnsafeStateQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StateQueryServiceServer will
// result in compilation errors.
type UnsafeStateQueryServiceServer interface {
	mustEmbedUnimplementedStateQueryServiceServer()
}

func RegisterStateQueryServiceServer(s grpc.ServiceRegistrar, srv StateQueryServiceServer) {
	s.RegisterService(&StateQueryService_ServiceDesc, srv)
}

func _StateQueryService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateQueryServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrhq.service.module.v1.StateQueryService/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateQueryServiceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StateQueryService_ListBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateQueryServiceServer).ListBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrhq.service.module.v1.StateQueryService/ListBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateQueryServiceServer).ListBalance(ctx, req.(*ListBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StateQueryService_GetSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateQueryServiceServer).GetSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrhq.service.module.v1.StateQueryService/GetSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateQueryServiceServer).GetSupply(ctx, req.(*GetSupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StateQueryService_ListSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateQueryServiceServer).ListSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonrhq.service.module.v1.StateQueryService/ListSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateQueryServiceServer).ListSupply(ctx, req.(*ListSupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StateQueryService_ServiceDesc is the grpc.ServiceDesc for StateQueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StateQueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sonrhq.service.module.v1.StateQueryService",
	HandlerType: (*StateQueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalance",
			Handler:    _StateQueryService_GetBalance_Handler,
		},
		{
			MethodName: "ListBalance",
			Handler:    _StateQueryService_ListBalance_Handler,
		},
		{
			MethodName: "GetSupply",
			Handler:    _StateQueryService_GetSupply_Handler,
		},
		{
			MethodName: "ListSupply",
			Handler:    _StateQueryService_ListSupply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sonrhq/service/module/v1/state_query.proto",
}
