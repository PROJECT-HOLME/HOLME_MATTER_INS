// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: config/config.proto

package config

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

// ConfigReqRespClient is the client API for ConfigReqResp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigReqRespClient interface {
	TestResponse(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error)
}

type configReqRespClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigReqRespClient(cc grpc.ClientConnInterface) ConfigReqRespClient {
	return &configReqRespClient{cc}
}

func (c *configReqRespClient) TestResponse(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := c.cc.Invoke(ctx, "/config.ConfigReqResp/TestResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigReqRespServer is the server API for ConfigReqResp service.
// All implementations must embed UnimplementedConfigReqRespServer
// for forward compatibility
type ConfigReqRespServer interface {
	TestResponse(context.Context, *ConfigRequest) (*ConfigResponse, error)
	mustEmbedUnimplementedConfigReqRespServer()
}

// UnimplementedConfigReqRespServer must be embedded to have forward compatible implementations.
type UnimplementedConfigReqRespServer struct {
}

func (UnimplementedConfigReqRespServer) TestResponse(context.Context, *ConfigRequest) (*ConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestResponse not implemented")
}
func (UnimplementedConfigReqRespServer) mustEmbedUnimplementedConfigReqRespServer() {}

// UnsafeConfigReqRespServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigReqRespServer will
// result in compilation errors.
type UnsafeConfigReqRespServer interface {
	mustEmbedUnimplementedConfigReqRespServer()
}

func RegisterConfigReqRespServer(s grpc.ServiceRegistrar, srv ConfigReqRespServer) {
	s.RegisterService(&ConfigReqResp_ServiceDesc, srv)
}

func _ConfigReqResp_TestResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigReqRespServer).TestResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigReqResp/TestResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigReqRespServer).TestResponse(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigReqResp_ServiceDesc is the grpc.ServiceDesc for ConfigReqResp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigReqResp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "config.ConfigReqResp",
	HandlerType: (*ConfigReqRespServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestResponse",
			Handler:    _ConfigReqResp_TestResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config/config.proto",
}
