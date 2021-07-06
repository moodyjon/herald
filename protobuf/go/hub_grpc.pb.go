// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HubClient is the client API for Hub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HubClient interface {
	SubscribeHeaders(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (Hub_SubscribeHeadersClient, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*Outputs, error)
	GetBlock(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockOutput, error)
	GetBlockHeader(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockHeaderOutput, error)
	GetServerHeight(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*wrapperspb.UInt64Value, error)
	GetHeaders(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (Hub_GetHeadersClient, error)
	Ping(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	Version(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	Features(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*wrapperspb.StringValue, error)
	Broadcast(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*wrapperspb.UInt64Value, error)
}

type hubClient struct {
	cc grpc.ClientConnInterface
}

func NewHubClient(cc grpc.ClientConnInterface) HubClient {
	return &hubClient{cc}
}

func (c *hubClient) SubscribeHeaders(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (Hub_SubscribeHeadersClient, error) {
	stream, err := c.cc.NewStream(ctx, &Hub_ServiceDesc.Streams[0], "/pb.Hub/SubscribeHeaders", opts...)
	if err != nil {
		return nil, err
	}
	x := &hubSubscribeHeadersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Hub_SubscribeHeadersClient interface {
	Recv() (*BlockHeaderOutput, error)
	grpc.ClientStream
}

type hubSubscribeHeadersClient struct {
	grpc.ClientStream
}

func (x *hubSubscribeHeadersClient) Recv() (*BlockHeaderOutput, error) {
	m := new(BlockHeaderOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hubClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*Outputs, error) {
	out := new(Outputs)
	err := c.cc.Invoke(ctx, "/pb.Hub/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) GetBlock(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockOutput, error) {
	out := new(BlockOutput)
	err := c.cc.Invoke(ctx, "/pb.Hub/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) GetBlockHeader(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (*BlockHeaderOutput, error) {
	out := new(BlockHeaderOutput)
	err := c.cc.Invoke(ctx, "/pb.Hub/GetBlockHeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) GetServerHeight(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*wrapperspb.UInt64Value, error) {
	out := new(wrapperspb.UInt64Value)
	err := c.cc.Invoke(ctx, "/pb.Hub/GetServerHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) GetHeaders(ctx context.Context, in *BlockRequest, opts ...grpc.CallOption) (Hub_GetHeadersClient, error) {
	stream, err := c.cc.NewStream(ctx, &Hub_ServiceDesc.Streams[1], "/pb.Hub/GetHeaders", opts...)
	if err != nil {
		return nil, err
	}
	x := &hubGetHeadersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Hub_GetHeadersClient interface {
	Recv() (*BlockHeaderOutput, error)
	grpc.ClientStream
}

type hubGetHeadersClient struct {
	grpc.ClientStream
}

func (x *hubGetHeadersClient) Recv() (*BlockHeaderOutput, error) {
	m := new(BlockHeaderOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hubClient) Ping(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/pb.Hub/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) Version(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/pb.Hub/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) Features(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*wrapperspb.StringValue, error) {
	out := new(wrapperspb.StringValue)
	err := c.cc.Invoke(ctx, "/pb.Hub/Features", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) Broadcast(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*wrapperspb.UInt64Value, error) {
	out := new(wrapperspb.UInt64Value)
	err := c.cc.Invoke(ctx, "/pb.Hub/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HubServer is the server API for Hub service.
// All implementations must embed UnimplementedHubServer
// for forward compatibility
type HubServer interface {
	SubscribeHeaders(*BlockRequest, Hub_SubscribeHeadersServer) error
	Search(context.Context, *SearchRequest) (*Outputs, error)
	GetBlock(context.Context, *BlockRequest) (*BlockOutput, error)
	GetBlockHeader(context.Context, *BlockRequest) (*BlockHeaderOutput, error)
	GetServerHeight(context.Context, *EmptyMessage) (*wrapperspb.UInt64Value, error)
	GetHeaders(*BlockRequest, Hub_GetHeadersServer) error
	Ping(context.Context, *EmptyMessage) (*wrapperspb.StringValue, error)
	Version(context.Context, *EmptyMessage) (*wrapperspb.StringValue, error)
	Features(context.Context, *EmptyMessage) (*wrapperspb.StringValue, error)
	Broadcast(context.Context, *EmptyMessage) (*wrapperspb.UInt64Value, error)
	mustEmbedUnimplementedHubServer()
}

// UnimplementedHubServer must be embedded to have forward compatible implementations.
type UnimplementedHubServer struct {
}

func (UnimplementedHubServer) SubscribeHeaders(*BlockRequest, Hub_SubscribeHeadersServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeHeaders not implemented")
}
func (UnimplementedHubServer) Search(context.Context, *SearchRequest) (*Outputs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedHubServer) GetBlock(context.Context, *BlockRequest) (*BlockOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (UnimplementedHubServer) GetBlockHeader(context.Context, *BlockRequest) (*BlockHeaderOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeader not implemented")
}
func (UnimplementedHubServer) GetServerHeight(context.Context, *EmptyMessage) (*wrapperspb.UInt64Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerHeight not implemented")
}
func (UnimplementedHubServer) GetHeaders(*BlockRequest, Hub_GetHeadersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetHeaders not implemented")
}
func (UnimplementedHubServer) Ping(context.Context, *EmptyMessage) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedHubServer) Version(context.Context, *EmptyMessage) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedHubServer) Features(context.Context, *EmptyMessage) (*wrapperspb.StringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Features not implemented")
}
func (UnimplementedHubServer) Broadcast(context.Context, *EmptyMessage) (*wrapperspb.UInt64Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedHubServer) mustEmbedUnimplementedHubServer() {}

// UnsafeHubServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HubServer will
// result in compilation errors.
type UnsafeHubServer interface {
	mustEmbedUnimplementedHubServer()
}

func RegisterHubServer(s grpc.ServiceRegistrar, srv HubServer) {
	s.RegisterService(&Hub_ServiceDesc, srv)
}

func _Hub_SubscribeHeaders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlockRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HubServer).SubscribeHeaders(m, &hubSubscribeHeadersServer{stream})
}

type Hub_SubscribeHeadersServer interface {
	Send(*BlockHeaderOutput) error
	grpc.ServerStream
}

type hubSubscribeHeadersServer struct {
	grpc.ServerStream
}

func (x *hubSubscribeHeadersServer) Send(m *BlockHeaderOutput) error {
	return x.ServerStream.SendMsg(m)
}

func _Hub_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Hub/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Hub/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).GetBlock(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_GetBlockHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).GetBlockHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Hub/GetBlockHeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).GetBlockHeader(ctx, req.(*BlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_GetServerHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).GetServerHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Hub/GetServerHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).GetServerHeight(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_GetHeaders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlockRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HubServer).GetHeaders(m, &hubGetHeadersServer{stream})
}

type Hub_GetHeadersServer interface {
	Send(*BlockHeaderOutput) error
	grpc.ServerStream
}

type hubGetHeadersServer struct {
	grpc.ServerStream
}

func (x *hubGetHeadersServer) Send(m *BlockHeaderOutput) error {
	return x.ServerStream.SendMsg(m)
}

func _Hub_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Hub/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Ping(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Hub/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Version(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_Features_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Features(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Hub/Features",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Features(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Hub/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Broadcast(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Hub_ServiceDesc is the grpc.ServiceDesc for Hub service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hub_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Hub",
	HandlerType: (*HubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Hub_Search_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _Hub_GetBlock_Handler,
		},
		{
			MethodName: "GetBlockHeader",
			Handler:    _Hub_GetBlockHeader_Handler,
		},
		{
			MethodName: "GetServerHeight",
			Handler:    _Hub_GetServerHeight_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Hub_Ping_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _Hub_Version_Handler,
		},
		{
			MethodName: "Features",
			Handler:    _Hub_Features_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _Hub_Broadcast_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeHeaders",
			Handler:       _Hub_SubscribeHeaders_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetHeaders",
			Handler:       _Hub_GetHeaders_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hub.proto",
}
