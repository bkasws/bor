// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// BorClient is the client API for Bor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BorClient interface {
	Pprof(ctx context.Context, in *PprofRequest, opts ...grpc.CallOption) (Bor_PprofClient, error)
	PeersAdd(ctx context.Context, in *PeersAddRequest, opts ...grpc.CallOption) (*PeersAddResponse, error)
	PeersRemove(ctx context.Context, in *PeersRemoveRequest, opts ...grpc.CallOption) (*PeersRemoveResponse, error)
	PeersList(ctx context.Context, in *PeersListRequest, opts ...grpc.CallOption) (*PeersListResponse, error)
	PeersStatus(ctx context.Context, in *PeersStatusRequest, opts ...grpc.CallOption) (*PeersStatusResponse, error)
	ChainSetHead(ctx context.Context, in *ChainSetHeadRequest, opts ...grpc.CallOption) (*ChainSetHeadResponse, error)
	Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusResponse, error)
}

type borClient struct {
	cc grpc.ClientConnInterface
}

func NewBorClient(cc grpc.ClientConnInterface) BorClient {
	return &borClient{cc}
}

func (c *borClient) Pprof(ctx context.Context, in *PprofRequest, opts ...grpc.CallOption) (Bor_PprofClient, error) {
	stream, err := c.cc.NewStream(ctx, &Bor_ServiceDesc.Streams[0], "/proto.Bor/Pprof", opts...)
	if err != nil {
		return nil, err
	}
	x := &borPprofClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Bor_PprofClient interface {
	Recv() (*PprofResponse, error)
	grpc.ClientStream
}

type borPprofClient struct {
	grpc.ClientStream
}

func (x *borPprofClient) Recv() (*PprofResponse, error) {
	m := new(PprofResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *borClient) PeersAdd(ctx context.Context, in *PeersAddRequest, opts ...grpc.CallOption) (*PeersAddResponse, error) {
	out := new(PeersAddResponse)
	err := c.cc.Invoke(ctx, "/proto.Bor/PeersAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borClient) PeersRemove(ctx context.Context, in *PeersRemoveRequest, opts ...grpc.CallOption) (*PeersRemoveResponse, error) {
	out := new(PeersRemoveResponse)
	err := c.cc.Invoke(ctx, "/proto.Bor/PeersRemove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borClient) PeersList(ctx context.Context, in *PeersListRequest, opts ...grpc.CallOption) (*PeersListResponse, error) {
	out := new(PeersListResponse)
	err := c.cc.Invoke(ctx, "/proto.Bor/PeersList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borClient) PeersStatus(ctx context.Context, in *PeersStatusRequest, opts ...grpc.CallOption) (*PeersStatusResponse, error) {
	out := new(PeersStatusResponse)
	err := c.cc.Invoke(ctx, "/proto.Bor/PeersStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borClient) ChainSetHead(ctx context.Context, in *ChainSetHeadRequest, opts ...grpc.CallOption) (*ChainSetHeadResponse, error) {
	out := new(ChainSetHeadResponse)
	err := c.cc.Invoke(ctx, "/proto.Bor/ChainSetHead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borClient) Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/proto.Bor/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BorServer is the server API for Bor service.
// All implementations must embed UnimplementedBorServer
// for forward compatibility
type BorServer interface {
	Pprof(*PprofRequest, Bor_PprofServer) error
	PeersAdd(context.Context, *PeersAddRequest) (*PeersAddResponse, error)
	PeersRemove(context.Context, *PeersRemoveRequest) (*PeersRemoveResponse, error)
	PeersList(context.Context, *PeersListRequest) (*PeersListResponse, error)
	PeersStatus(context.Context, *PeersStatusRequest) (*PeersStatusResponse, error)
	ChainSetHead(context.Context, *ChainSetHeadRequest) (*ChainSetHeadResponse, error)
	Status(context.Context, *empty.Empty) (*StatusResponse, error)
	mustEmbedUnimplementedBorServer()
}

// UnimplementedBorServer must be embedded to have forward compatible implementations.
type UnimplementedBorServer struct {
}

func (UnimplementedBorServer) Pprof(*PprofRequest, Bor_PprofServer) error {
	return status.Errorf(codes.Unimplemented, "method Pprof not implemented")
}
func (UnimplementedBorServer) PeersAdd(context.Context, *PeersAddRequest) (*PeersAddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeersAdd not implemented")
}
func (UnimplementedBorServer) PeersRemove(context.Context, *PeersRemoveRequest) (*PeersRemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeersRemove not implemented")
}
func (UnimplementedBorServer) PeersList(context.Context, *PeersListRequest) (*PeersListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeersList not implemented")
}
func (UnimplementedBorServer) PeersStatus(context.Context, *PeersStatusRequest) (*PeersStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeersStatus not implemented")
}
func (UnimplementedBorServer) ChainSetHead(context.Context, *ChainSetHeadRequest) (*ChainSetHeadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainSetHead not implemented")
}
func (UnimplementedBorServer) Status(context.Context, *empty.Empty) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedBorServer) mustEmbedUnimplementedBorServer() {}

// UnsafeBorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BorServer will
// result in compilation errors.
type UnsafeBorServer interface {
	mustEmbedUnimplementedBorServer()
}

func RegisterBorServer(s grpc.ServiceRegistrar, srv BorServer) {
	s.RegisterService(&Bor_ServiceDesc, srv)
}

func _Bor_Pprof_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PprofRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BorServer).Pprof(m, &borPprofServer{stream})
}

type Bor_PprofServer interface {
	Send(*PprofResponse) error
	grpc.ServerStream
}

type borPprofServer struct {
	grpc.ServerStream
}

func (x *borPprofServer) Send(m *PprofResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Bor_PeersAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeersAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorServer).PeersAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Bor/PeersAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorServer).PeersAdd(ctx, req.(*PeersAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bor_PeersRemove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeersRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorServer).PeersRemove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Bor/PeersRemove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorServer).PeersRemove(ctx, req.(*PeersRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bor_PeersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeersListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorServer).PeersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Bor/PeersList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorServer).PeersList(ctx, req.(*PeersListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bor_PeersStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeersStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorServer).PeersStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Bor/PeersStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorServer).PeersStatus(ctx, req.(*PeersStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bor_ChainSetHead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainSetHeadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorServer).ChainSetHead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Bor/ChainSetHead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorServer).ChainSetHead(ctx, req.(*ChainSetHeadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bor_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Bor/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorServer).Status(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Bor_ServiceDesc is the grpc.ServiceDesc for Bor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Bor",
	HandlerType: (*BorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PeersAdd",
			Handler:    _Bor_PeersAdd_Handler,
		},
		{
			MethodName: "PeersRemove",
			Handler:    _Bor_PeersRemove_Handler,
		},
		{
			MethodName: "PeersList",
			Handler:    _Bor_PeersList_Handler,
		},
		{
			MethodName: "PeersStatus",
			Handler:    _Bor_PeersStatus_Handler,
		},
		{
			MethodName: "ChainSetHead",
			Handler:    _Bor_ChainSetHead_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Bor_Status_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Pprof",
			Handler:       _Bor_Pprof_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "command/server/proto/server.proto",
}
