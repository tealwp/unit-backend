// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: pkg/grpc/testing.proto

package pb

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

// InventoryClient is the client API for Inventory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryClient interface {
	GetItemList(ctx context.Context, in *GetItemListRequest, opts ...grpc.CallOption) (*GetItemListResponse, error)
}

type inventoryClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryClient(cc grpc.ClientConnInterface) InventoryClient {
	return &inventoryClient{cc}
}

func (c *inventoryClient) GetItemList(ctx context.Context, in *GetItemListRequest, opts ...grpc.CallOption) (*GetItemListResponse, error) {
	out := new(GetItemListResponse)
	err := c.cc.Invoke(ctx, "/Inventory/GetItemList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServer is the server API for Inventory service.
// All implementations must embed UnimplementedInventoryServer
// for forward compatibility
type InventoryServer interface {
	GetItemList(context.Context, *GetItemListRequest) (*GetItemListResponse, error)
	mustEmbedUnimplementedInventoryServer()
}

// UnimplementedInventoryServer must be embedded to have forward compatible implementations.
type UnimplementedInventoryServer struct {
}

func (UnimplementedInventoryServer) GetItemList(context.Context, *GetItemListRequest) (*GetItemListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemList not implemented")
}
func (UnimplementedInventoryServer) mustEmbedUnimplementedInventoryServer() {}

// UnsafeInventoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServer will
// result in compilation errors.
type UnsafeInventoryServer interface {
	mustEmbedUnimplementedInventoryServer()
}

func RegisterInventoryServer(s grpc.ServiceRegistrar, srv InventoryServer) {
	s.RegisterService(&Inventory_ServiceDesc, srv)
}

func _Inventory_GetItemList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServer).GetItemList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Inventory/GetItemList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServer).GetItemList(ctx, req.(*GetItemListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Inventory_ServiceDesc is the grpc.ServiceDesc for Inventory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Inventory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Inventory",
	HandlerType: (*InventoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetItemList",
			Handler:    _Inventory_GetItemList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc/testing.proto",
}