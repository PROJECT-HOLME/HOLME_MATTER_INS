// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: soundbar.proto

package InstanceSoundbar

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
	Soundbar_HandleFrame_FullMethodName = "/HOLME_instance.Soundbar/HandleFrame"
)

// SoundbarClient is the client API for Soundbar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SoundbarClient interface {
	HandleFrame(ctx context.Context, in *SoundbarFrame, opts ...grpc.CallOption) (*SoundbarRes, error)
}

type soundbarClient struct {
	cc grpc.ClientConnInterface
}

func NewSoundbarClient(cc grpc.ClientConnInterface) SoundbarClient {
	return &soundbarClient{cc}
}

func (c *soundbarClient) HandleFrame(ctx context.Context, in *SoundbarFrame, opts ...grpc.CallOption) (*SoundbarRes, error) {
	out := new(SoundbarRes)
	err := c.cc.Invoke(ctx, Soundbar_HandleFrame_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SoundbarServer is the server API for Soundbar service.
// All implementations must embed UnimplementedSoundbarServer
// for forward compatibility
type SoundbarServer interface {
	HandleFrame(context.Context, *SoundbarFrame) (*SoundbarRes, error)
	mustEmbedUnimplementedSoundbarServer()
}

// UnimplementedSoundbarServer must be embedded to have forward compatible implementations.
type UnimplementedSoundbarServer struct {
}

func (UnimplementedSoundbarServer) HandleFrame(context.Context, *SoundbarFrame) (*SoundbarRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleFrame not implemented")
}
func (UnimplementedSoundbarServer) mustEmbedUnimplementedSoundbarServer() {}

// UnsafeSoundbarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SoundbarServer will
// result in compilation errors.
type UnsafeSoundbarServer interface {
	mustEmbedUnimplementedSoundbarServer()
}

func RegisterSoundbarServer(s grpc.ServiceRegistrar, srv SoundbarServer) {
	s.RegisterService(&Soundbar_ServiceDesc, srv)
}

func _Soundbar_HandleFrame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SoundbarFrame)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SoundbarServer).HandleFrame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Soundbar_HandleFrame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SoundbarServer).HandleFrame(ctx, req.(*SoundbarFrame))
	}
	return interceptor(ctx, in, info, handler)
}

// Soundbar_ServiceDesc is the grpc.ServiceDesc for Soundbar service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Soundbar_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HOLME_instance.Soundbar",
	HandlerType: (*SoundbarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleFrame",
			Handler:    _Soundbar_HandleFrame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "soundbar.proto",
}
