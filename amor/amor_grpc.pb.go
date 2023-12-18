// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: amor/amor.proto

package amor

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

const (
	ProjectAmor_AddHome_FullMethodName    = "/ProjectAmor/AddHome"
	ProjectAmor_DeleteHome_FullMethodName = "/ProjectAmor/DeleteHome"
	ProjectAmor_UpdateHome_FullMethodName = "/ProjectAmor/UpdateHome"
	ProjectAmor_GetHome_FullMethodName    = "/ProjectAmor/GetHome"
	ProjectAmor_ListHome_FullMethodName   = "/ProjectAmor/ListHome"
)

// ProjectAmorClient is the client API for ProjectAmor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectAmorClient interface {
	AddHome(ctx context.Context, in *AddHomeRequest, opts ...grpc.CallOption) (*AddHomeResponse, error)
	DeleteHome(ctx context.Context, in *DeleteHomeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateHome(ctx context.Context, in *UpdateHomeRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetHome(ctx context.Context, in *GetHomeRequest, opts ...grpc.CallOption) (*GetHomeResponse, error)
	ListHome(ctx context.Context, in *ListHomeRequest, opts ...grpc.CallOption) (*ListHomeResponse, error)
}

type projectAmorClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectAmorClient(cc grpc.ClientConnInterface) ProjectAmorClient {
	return &projectAmorClient{cc}
}

func (c *projectAmorClient) AddHome(ctx context.Context, in *AddHomeRequest, opts ...grpc.CallOption) (*AddHomeResponse, error) {
	out := new(AddHomeResponse)
	err := c.cc.Invoke(ctx, ProjectAmor_AddHome_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) DeleteHome(ctx context.Context, in *DeleteHomeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ProjectAmor_DeleteHome_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) UpdateHome(ctx context.Context, in *UpdateHomeRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, ProjectAmor_UpdateHome_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) GetHome(ctx context.Context, in *GetHomeRequest, opts ...grpc.CallOption) (*GetHomeResponse, error) {
	out := new(GetHomeResponse)
	err := c.cc.Invoke(ctx, ProjectAmor_GetHome_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectAmorClient) ListHome(ctx context.Context, in *ListHomeRequest, opts ...grpc.CallOption) (*ListHomeResponse, error) {
	out := new(ListHomeResponse)
	err := c.cc.Invoke(ctx, ProjectAmor_ListHome_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectAmorServer is the server API for ProjectAmor service.
// All implementations must embed UnimplementedProjectAmorServer
// for forward compatibility
type ProjectAmorServer interface {
	AddHome(context.Context, *AddHomeRequest) (*AddHomeResponse, error)
	DeleteHome(context.Context, *DeleteHomeRequest) (*empty.Empty, error)
	UpdateHome(context.Context, *UpdateHomeRequest) (*empty.Empty, error)
	GetHome(context.Context, *GetHomeRequest) (*GetHomeResponse, error)
	ListHome(context.Context, *ListHomeRequest) (*ListHomeResponse, error)
	mustEmbedUnimplementedProjectAmorServer()
}

// UnimplementedProjectAmorServer must be embedded to have forward compatible implementations.
type UnimplementedProjectAmorServer struct {
}

func (UnimplementedProjectAmorServer) AddHome(context.Context, *AddHomeRequest) (*AddHomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHome not implemented")
}
func (UnimplementedProjectAmorServer) DeleteHome(context.Context, *DeleteHomeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHome not implemented")
}
func (UnimplementedProjectAmorServer) UpdateHome(context.Context, *UpdateHomeRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHome not implemented")
}
func (UnimplementedProjectAmorServer) GetHome(context.Context, *GetHomeRequest) (*GetHomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHome not implemented")
}
func (UnimplementedProjectAmorServer) ListHome(context.Context, *ListHomeRequest) (*ListHomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHome not implemented")
}
func (UnimplementedProjectAmorServer) mustEmbedUnimplementedProjectAmorServer() {}

// UnsafeProjectAmorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectAmorServer will
// result in compilation errors.
type UnsafeProjectAmorServer interface {
	mustEmbedUnimplementedProjectAmorServer()
}

func RegisterProjectAmorServer(s grpc.ServiceRegistrar, srv ProjectAmorServer) {
	s.RegisterService(&ProjectAmor_ServiceDesc, srv)
}

func _ProjectAmor_AddHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).AddHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_AddHome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).AddHome(ctx, req.(*AddHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_DeleteHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).DeleteHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_DeleteHome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).DeleteHome(ctx, req.(*DeleteHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_UpdateHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).UpdateHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_UpdateHome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).UpdateHome(ctx, req.(*UpdateHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_GetHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).GetHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_GetHome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).GetHome(ctx, req.(*GetHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectAmor_ListHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectAmorServer).ListHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectAmor_ListHome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectAmorServer).ListHome(ctx, req.(*ListHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectAmor_ServiceDesc is the grpc.ServiceDesc for ProjectAmor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectAmor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProjectAmor",
	HandlerType: (*ProjectAmorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddHome",
			Handler:    _ProjectAmor_AddHome_Handler,
		},
		{
			MethodName: "DeleteHome",
			Handler:    _ProjectAmor_DeleteHome_Handler,
		},
		{
			MethodName: "UpdateHome",
			Handler:    _ProjectAmor_UpdateHome_Handler,
		},
		{
			MethodName: "GetHome",
			Handler:    _ProjectAmor_GetHome_Handler,
		},
		{
			MethodName: "ListHome",
			Handler:    _ProjectAmor_ListHome_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "amor/amor.proto",
}
