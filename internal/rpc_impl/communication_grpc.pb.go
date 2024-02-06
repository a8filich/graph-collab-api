// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: communication.proto

package rpc_impl

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
	QuantifierCommunication_SendBig5Msg_FullMethodName            = "/rpc_impl.QuantifierCommunication/SendBig5Msg"
	QuantifierCommunication_SendTeamRoleMsg_FullMethodName        = "/rpc_impl.QuantifierCommunication/SendTeamRoleMsg"
	QuantifierCommunication_SendExpertiseMsg_FullMethodName       = "/rpc_impl.QuantifierCommunication/SendExpertiseMsg"
	QuantifierCommunication_SendManagerialMsg_FullMethodName      = "/rpc_impl.QuantifierCommunication/SendManagerialMsg"
	QuantifierCommunication_GetGraphComputationRes_FullMethodName = "/rpc_impl.QuantifierCommunication/GetGraphComputationRes"
)

// QuantifierCommunicationClient is the client API for QuantifierCommunication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuantifierCommunicationClient interface {
	SendBig5Msg(ctx context.Context, in *Big5Msg, opts ...grpc.CallOption) (*Empty, error)
	SendTeamRoleMsg(ctx context.Context, in *TeamRoleMsg, opts ...grpc.CallOption) (*Empty, error)
	SendExpertiseMsg(ctx context.Context, in *ExpertiseMsg, opts ...grpc.CallOption) (*Empty, error)
	SendManagerialMsg(ctx context.Context, in *ManagerialMsg, opts ...grpc.CallOption) (*Empty, error)
	GetGraphComputationRes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GraphComputationResMsg, error)
}

type quantifierCommunicationClient struct {
	cc grpc.ClientConnInterface
}

func NewQuantifierCommunicationClient(cc grpc.ClientConnInterface) QuantifierCommunicationClient {
	return &quantifierCommunicationClient{cc}
}

func (c *quantifierCommunicationClient) SendBig5Msg(ctx context.Context, in *Big5Msg, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, QuantifierCommunication_SendBig5Msg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quantifierCommunicationClient) SendTeamRoleMsg(ctx context.Context, in *TeamRoleMsg, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, QuantifierCommunication_SendTeamRoleMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quantifierCommunicationClient) SendExpertiseMsg(ctx context.Context, in *ExpertiseMsg, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, QuantifierCommunication_SendExpertiseMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quantifierCommunicationClient) SendManagerialMsg(ctx context.Context, in *ManagerialMsg, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, QuantifierCommunication_SendManagerialMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quantifierCommunicationClient) GetGraphComputationRes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GraphComputationResMsg, error) {
	out := new(GraphComputationResMsg)
	err := c.cc.Invoke(ctx, QuantifierCommunication_GetGraphComputationRes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuantifierCommunicationServer is the server API for QuantifierCommunication service.
// All implementations must embed UnimplementedQuantifierCommunicationServer
// for forward compatibility
type QuantifierCommunicationServer interface {
	SendBig5Msg(context.Context, *Big5Msg) (*Empty, error)
	SendTeamRoleMsg(context.Context, *TeamRoleMsg) (*Empty, error)
	SendExpertiseMsg(context.Context, *ExpertiseMsg) (*Empty, error)
	SendManagerialMsg(context.Context, *ManagerialMsg) (*Empty, error)
	GetGraphComputationRes(context.Context, *Empty) (*GraphComputationResMsg, error)
	mustEmbedUnimplementedQuantifierCommunicationServer()
}

// UnimplementedQuantifierCommunicationServer must be embedded to have forward compatible implementations.
type UnimplementedQuantifierCommunicationServer struct {
}

func (UnimplementedQuantifierCommunicationServer) SendBig5Msg(context.Context, *Big5Msg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBig5Msg not implemented")
}
func (UnimplementedQuantifierCommunicationServer) SendTeamRoleMsg(context.Context, *TeamRoleMsg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTeamRoleMsg not implemented")
}
func (UnimplementedQuantifierCommunicationServer) SendExpertiseMsg(context.Context, *ExpertiseMsg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendExpertiseMsg not implemented")
}
func (UnimplementedQuantifierCommunicationServer) SendManagerialMsg(context.Context, *ManagerialMsg) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendManagerialMsg not implemented")
}
func (UnimplementedQuantifierCommunicationServer) GetGraphComputationRes(context.Context, *Empty) (*GraphComputationResMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGraphComputationRes not implemented")
}
func (UnimplementedQuantifierCommunicationServer) mustEmbedUnimplementedQuantifierCommunicationServer() {
}

// UnsafeQuantifierCommunicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuantifierCommunicationServer will
// result in compilation errors.
type UnsafeQuantifierCommunicationServer interface {
	mustEmbedUnimplementedQuantifierCommunicationServer()
}

func RegisterQuantifierCommunicationServer(s grpc.ServiceRegistrar, srv QuantifierCommunicationServer) {
	s.RegisterService(&QuantifierCommunication_ServiceDesc, srv)
}

func _QuantifierCommunication_SendBig5Msg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Big5Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuantifierCommunicationServer).SendBig5Msg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuantifierCommunication_SendBig5Msg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuantifierCommunicationServer).SendBig5Msg(ctx, req.(*Big5Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuantifierCommunication_SendTeamRoleMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamRoleMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuantifierCommunicationServer).SendTeamRoleMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuantifierCommunication_SendTeamRoleMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuantifierCommunicationServer).SendTeamRoleMsg(ctx, req.(*TeamRoleMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuantifierCommunication_SendExpertiseMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpertiseMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuantifierCommunicationServer).SendExpertiseMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuantifierCommunication_SendExpertiseMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuantifierCommunicationServer).SendExpertiseMsg(ctx, req.(*ExpertiseMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuantifierCommunication_SendManagerialMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManagerialMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuantifierCommunicationServer).SendManagerialMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuantifierCommunication_SendManagerialMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuantifierCommunicationServer).SendManagerialMsg(ctx, req.(*ManagerialMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuantifierCommunication_GetGraphComputationRes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuantifierCommunicationServer).GetGraphComputationRes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuantifierCommunication_GetGraphComputationRes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuantifierCommunicationServer).GetGraphComputationRes(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// QuantifierCommunication_ServiceDesc is the grpc.ServiceDesc for QuantifierCommunication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuantifierCommunication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc_impl.QuantifierCommunication",
	HandlerType: (*QuantifierCommunicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendBig5Msg",
			Handler:    _QuantifierCommunication_SendBig5Msg_Handler,
		},
		{
			MethodName: "SendTeamRoleMsg",
			Handler:    _QuantifierCommunication_SendTeamRoleMsg_Handler,
		},
		{
			MethodName: "SendExpertiseMsg",
			Handler:    _QuantifierCommunication_SendExpertiseMsg_Handler,
		},
		{
			MethodName: "SendManagerialMsg",
			Handler:    _QuantifierCommunication_SendManagerialMsg_Handler,
		},
		{
			MethodName: "GetGraphComputationRes",
			Handler:    _QuantifierCommunication_GetGraphComputationRes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "communication.proto",
}
