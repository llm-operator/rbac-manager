// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// RbacInternalServiceClient is the client API for RbacInternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RbacInternalServiceClient interface {
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
	// AuthorizeWorker authorizes requests from worker clusters.
	AuthorizeWorker(ctx context.Context, in *AuthorizeWorkerRequest, opts ...grpc.CallOption) (*AuthorizeWorkerResponse, error)
}

type rbacInternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRbacInternalServiceClient(cc grpc.ClientConnInterface) RbacInternalServiceClient {
	return &rbacInternalServiceClient{cc}
}

func (c *rbacInternalServiceClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	out := new(AuthorizeResponse)
	err := c.cc.Invoke(ctx, "/llmoperator.rbac.server.v1.RbacInternalService/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rbacInternalServiceClient) AuthorizeWorker(ctx context.Context, in *AuthorizeWorkerRequest, opts ...grpc.CallOption) (*AuthorizeWorkerResponse, error) {
	out := new(AuthorizeWorkerResponse)
	err := c.cc.Invoke(ctx, "/llmoperator.rbac.server.v1.RbacInternalService/AuthorizeWorker", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RbacInternalServiceServer is the server API for RbacInternalService service.
// All implementations must embed UnimplementedRbacInternalServiceServer
// for forward compatibility
type RbacInternalServiceServer interface {
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error)
	// AuthorizeWorker authorizes requests from worker clusters.
	AuthorizeWorker(context.Context, *AuthorizeWorkerRequest) (*AuthorizeWorkerResponse, error)
	mustEmbedUnimplementedRbacInternalServiceServer()
}

// UnimplementedRbacInternalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRbacInternalServiceServer struct {
}

func (UnimplementedRbacInternalServiceServer) Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedRbacInternalServiceServer) AuthorizeWorker(context.Context, *AuthorizeWorkerRequest) (*AuthorizeWorkerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeWorker not implemented")
}
func (UnimplementedRbacInternalServiceServer) mustEmbedUnimplementedRbacInternalServiceServer() {}

// UnsafeRbacInternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RbacInternalServiceServer will
// result in compilation errors.
type UnsafeRbacInternalServiceServer interface {
	mustEmbedUnimplementedRbacInternalServiceServer()
}

func RegisterRbacInternalServiceServer(s grpc.ServiceRegistrar, srv RbacInternalServiceServer) {
	s.RegisterService(&RbacInternalService_ServiceDesc, srv)
}

func _RbacInternalService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacInternalServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/llmoperator.rbac.server.v1.RbacInternalService/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacInternalServiceServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RbacInternalService_AuthorizeWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeWorkerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RbacInternalServiceServer).AuthorizeWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/llmoperator.rbac.server.v1.RbacInternalService/AuthorizeWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RbacInternalServiceServer).AuthorizeWorker(ctx, req.(*AuthorizeWorkerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RbacInternalService_ServiceDesc is the grpc.ServiceDesc for RbacInternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RbacInternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "llmoperator.rbac.server.v1.RbacInternalService",
	HandlerType: (*RbacInternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _RbacInternalService_Authorize_Handler,
		},
		{
			MethodName: "AuthorizeWorker",
			Handler:    _RbacInternalService_AuthorizeWorker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/rbac_manager_service.proto",
}
