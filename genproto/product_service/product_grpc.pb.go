// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: product.proto

package product_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	Create(ctx context.Context, in *CreateProduct, opts ...grpc.CallOption) (*GetProduct, error)
	GetByID(ctx context.Context, in *ProductPrimaryKey, opts ...grpc.CallOption) (*GetProduct, error)
	GetList(ctx context.Context, in *GetListProductRequest, opts ...grpc.CallOption) (*GetListProductResponse, error)
	Update(ctx context.Context, in *UpdateProduct, opts ...grpc.CallOption) (*GetProduct, error)
	Delete(ctx context.Context, in *ProductPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) Create(ctx context.Context, in *CreateProduct, opts ...grpc.CallOption) (*GetProduct, error) {
	out := new(GetProduct)
	err := c.cc.Invoke(ctx, "/product_service_go.ProductService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetByID(ctx context.Context, in *ProductPrimaryKey, opts ...grpc.CallOption) (*GetProduct, error) {
	out := new(GetProduct)
	err := c.cc.Invoke(ctx, "/product_service_go.ProductService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetList(ctx context.Context, in *GetListProductRequest, opts ...grpc.CallOption) (*GetListProductResponse, error) {
	out := new(GetListProductResponse)
	err := c.cc.Invoke(ctx, "/product_service_go.ProductService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) Update(ctx context.Context, in *UpdateProduct, opts ...grpc.CallOption) (*GetProduct, error) {
	out := new(GetProduct)
	err := c.cc.Invoke(ctx, "/product_service_go.ProductService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) Delete(ctx context.Context, in *ProductPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/product_service_go.ProductService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations should embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	Create(context.Context, *CreateProduct) (*GetProduct, error)
	GetByID(context.Context, *ProductPrimaryKey) (*GetProduct, error)
	GetList(context.Context, *GetListProductRequest) (*GetListProductResponse, error)
	Update(context.Context, *UpdateProduct) (*GetProduct, error)
	Delete(context.Context, *ProductPrimaryKey) (*empty.Empty, error)
}

// UnimplementedProductServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) Create(context.Context, *CreateProduct) (*GetProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProductServiceServer) GetByID(context.Context, *ProductPrimaryKey) (*GetProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedProductServiceServer) GetList(context.Context, *GetListProductRequest) (*GetListProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedProductServiceServer) Update(context.Context, *UpdateProduct) (*GetProduct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProductServiceServer) Delete(context.Context, *ProductPrimaryKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProduct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service_go.ProductService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Create(ctx, req.(*CreateProduct))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service_go.ProductService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetByID(ctx, req.(*ProductPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service_go.ProductService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetList(ctx, req.(*GetListProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProduct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service_go.ProductService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Update(ctx, req.(*UpdateProduct))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/product_service_go.ProductService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).Delete(ctx, req.(*ProductPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product_service_go.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProductService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _ProductService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _ProductService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProductService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProductService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}
