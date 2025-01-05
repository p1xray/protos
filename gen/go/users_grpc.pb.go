// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: users.proto

package usersv1

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
	Users_FollowedUsers_FullMethodName = "/users.Users/FollowedUsers"
	Users_FollowUser_FullMethodName    = "/users.Users/FollowUser"
	Users_UnfollowUser_FullMethodName  = "/users.Users/UnfollowUser"
)

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsersClient interface {
	FollowedUsers(ctx context.Context, in *FollowedUsersRequest, opts ...grpc.CallOption) (*FollowedUsersResponse, error)
	FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*FollowUserResponse, error)
	UnfollowUser(ctx context.Context, in *UnfollowUserRequest, opts ...grpc.CallOption) (*UnfollowUserResponse, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) FollowedUsers(ctx context.Context, in *FollowedUsersRequest, opts ...grpc.CallOption) (*FollowedUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FollowedUsersResponse)
	err := c.cc.Invoke(ctx, Users_FollowedUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*FollowUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FollowUserResponse)
	err := c.cc.Invoke(ctx, Users_FollowUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersClient) UnfollowUser(ctx context.Context, in *UnfollowUserRequest, opts ...grpc.CallOption) (*UnfollowUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnfollowUserResponse)
	err := c.cc.Invoke(ctx, Users_UnfollowUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
// All implementations must embed UnimplementedUsersServer
// for forward compatibility.
type UsersServer interface {
	FollowedUsers(context.Context, *FollowedUsersRequest) (*FollowedUsersResponse, error)
	FollowUser(context.Context, *FollowUserRequest) (*FollowUserResponse, error)
	UnfollowUser(context.Context, *UnfollowUserRequest) (*UnfollowUserResponse, error)
	mustEmbedUnimplementedUsersServer()
}

// UnimplementedUsersServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUsersServer struct{}

func (UnimplementedUsersServer) FollowedUsers(context.Context, *FollowedUsersRequest) (*FollowedUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowedUsers not implemented")
}
func (UnimplementedUsersServer) FollowUser(context.Context, *FollowUserRequest) (*FollowUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (UnimplementedUsersServer) UnfollowUser(context.Context, *UnfollowUserRequest) (*UnfollowUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfollowUser not implemented")
}
func (UnimplementedUsersServer) mustEmbedUnimplementedUsersServer() {}
func (UnimplementedUsersServer) testEmbeddedByValue()               {}

// UnsafeUsersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServer will
// result in compilation errors.
type UnsafeUsersServer interface {
	mustEmbedUnimplementedUsersServer()
}

func RegisterUsersServer(s grpc.ServiceRegistrar, srv UsersServer) {
	// If the following call pancis, it indicates UnimplementedUsersServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Users_ServiceDesc, srv)
}

func _Users_FollowedUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowedUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).FollowedUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_FollowedUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).FollowedUsers(ctx, req.(*FollowedUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_FollowUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).FollowUser(ctx, req.(*FollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Users_UnfollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnfollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).UnfollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Users_UnfollowUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).UnfollowUser(ctx, req.(*UnfollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Users_ServiceDesc is the grpc.ServiceDesc for Users service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Users_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FollowedUsers",
			Handler:    _Users_FollowedUsers_Handler,
		},
		{
			MethodName: "FollowUser",
			Handler:    _Users_FollowUser_Handler,
		},
		{
			MethodName: "UnfollowUser",
			Handler:    _Users_UnfollowUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
