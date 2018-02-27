// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
*/
package grpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import xk "github.com/xephonhq/xephon-k/xk/xkpb"

import context "golang.org/x/net/context"
import grpc1 "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for Xephonk service

type XephonkClient interface {
	Ping(ctx context.Context, in *xk.Ping, opts ...grpc1.CallOption) (*xk.Pong, error)
}

type xephonkClient struct {
	cc *grpc1.ClientConn
}

func NewXephonkClient(cc *grpc1.ClientConn) XephonkClient {
	return &xephonkClient{cc}
}

func (c *xephonkClient) Ping(ctx context.Context, in *xk.Ping, opts ...grpc1.CallOption) (*xk.Pong, error) {
	out := new(xk.Pong)
	err := grpc1.Invoke(ctx, "/xkrpc.Xephonk/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Xephonk service

type XephonkServer interface {
	Ping(context.Context, *xk.Ping) (*xk.Pong, error)
}

func RegisterXephonkServer(s *grpc1.Server, srv XephonkServer) {
	s.RegisterService(&_Xephonk_serviceDesc, srv)
}

func _Xephonk_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(xk.Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(XephonkServer).Ping(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xkrpc.Xephonk/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(XephonkServer).Ping(ctx, req.(*xk.Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _Xephonk_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "xkrpc.Xephonk",
	HandlerType: (*XephonkServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Xephonk_Ping_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptorRpc) }

var fileDescriptorRpc = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2a, 0x48, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xad, 0xc8, 0x2e, 0x2a, 0x48, 0x96, 0xd2, 0x4d, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xcb,
	0x26, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0x25, 0x65, 0x80, 0xa4, 0xbc, 0x22,
	0xb5, 0x20, 0x23, 0x3f, 0x2f, 0xa3, 0x10, 0xca, 0xd0, 0xcd, 0xd6, 0xaf, 0x00, 0xa1, 0x82, 0x24,
	0xfd, 0x92, 0xca, 0x82, 0xd4, 0x62, 0x88, 0x0e, 0x23, 0x75, 0x2e, 0xf6, 0x08, 0xb0, 0x7c, 0xb6,
	0x90, 0x0c, 0x17, 0x4b, 0x40, 0x66, 0x5e, 0xba, 0x10, 0x87, 0x5e, 0x45, 0xb6, 0x1e, 0x88, 0x25,
	0x05, 0x61, 0xe5, 0xe7, 0xa5, 0x2b, 0x31, 0x38, 0x89, 0x9d, 0x78, 0x28, 0xc7, 0x70, 0xe2, 0x91,
	0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x46, 0xb1, 0xa4, 0x17, 0x15, 0x24,
	0x27, 0xb1, 0x81, 0xcd, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x24, 0xf2, 0x5f, 0xbc,
	0x00, 0x00, 0x00,
}