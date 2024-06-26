// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: proto/api.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	AggregateScores_SendAggregateScores_FullMethodName = "/proto.AggregateScores/SendAggregateScores"
)

// AggregateScoresClient is the client API for AggregateScores service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AggregateScoresClient interface {
	SendAggregateScores(ctx context.Context, in *AggregateScoresRequest, opts ...grpc.CallOption) (*AggregateScoresReply, error)
}

type aggregateScoresClient struct {
	cc grpc.ClientConnInterface
}

func NewAggregateScoresClient(cc grpc.ClientConnInterface) AggregateScoresClient {
	return &aggregateScoresClient{cc}
}

func (c *aggregateScoresClient) SendAggregateScores(ctx context.Context, in *AggregateScoresRequest, opts ...grpc.CallOption) (*AggregateScoresReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AggregateScoresReply)
	err := c.cc.Invoke(ctx, AggregateScores_SendAggregateScores_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AggregateScoresServer is the server API for AggregateScores service.
// All implementations must embed UnimplementedAggregateScoresServer
// for forward compatibility
type AggregateScoresServer interface {
	SendAggregateScores(context.Context, *AggregateScoresRequest) (*AggregateScoresReply, error)
	mustEmbedUnimplementedAggregateScoresServer()
}

// UnimplementedAggregateScoresServer must be embedded to have forward compatible implementations.
type UnimplementedAggregateScoresServer struct {
}

func (UnimplementedAggregateScoresServer) SendAggregateScores(context.Context, *AggregateScoresRequest) (*AggregateScoresReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAggregateScores not implemented")
}
func (UnimplementedAggregateScoresServer) mustEmbedUnimplementedAggregateScoresServer() {}

// UnsafeAggregateScoresServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AggregateScoresServer will
// result in compilation errors.
type UnsafeAggregateScoresServer interface {
	mustEmbedUnimplementedAggregateScoresServer()
}

func RegisterAggregateScoresServer(s grpc.ServiceRegistrar, srv AggregateScoresServer) {
	s.RegisterService(&AggregateScores_ServiceDesc, srv)
}

func _AggregateScores_SendAggregateScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AggregateScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregateScoresServer).SendAggregateScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AggregateScores_SendAggregateScores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregateScoresServer).SendAggregateScores(ctx, req.(*AggregateScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AggregateScores_ServiceDesc is the grpc.ServiceDesc for AggregateScores service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AggregateScores_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AggregateScores",
	HandlerType: (*AggregateScoresServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendAggregateScores",
			Handler:    _AggregateScores_SendAggregateScores_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}

const (
	TicketScores_SendTicketScores_FullMethodName = "/proto.TicketScores/SendTicketScores"
)

// TicketScoresClient is the client API for TicketScores service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketScoresClient interface {
	SendTicketScores(ctx context.Context, in *TicketScoresRequest, opts ...grpc.CallOption) (*TicketScoresReply, error)
}

type ticketScoresClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketScoresClient(cc grpc.ClientConnInterface) TicketScoresClient {
	return &ticketScoresClient{cc}
}

func (c *ticketScoresClient) SendTicketScores(ctx context.Context, in *TicketScoresRequest, opts ...grpc.CallOption) (*TicketScoresReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TicketScoresReply)
	err := c.cc.Invoke(ctx, TicketScores_SendTicketScores_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketScoresServer is the server API for TicketScores service.
// All implementations must embed UnimplementedTicketScoresServer
// for forward compatibility
type TicketScoresServer interface {
	SendTicketScores(context.Context, *TicketScoresRequest) (*TicketScoresReply, error)
	mustEmbedUnimplementedTicketScoresServer()
}

// UnimplementedTicketScoresServer must be embedded to have forward compatible implementations.
type UnimplementedTicketScoresServer struct {
}

func (UnimplementedTicketScoresServer) SendTicketScores(context.Context, *TicketScoresRequest) (*TicketScoresReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTicketScores not implemented")
}
func (UnimplementedTicketScoresServer) mustEmbedUnimplementedTicketScoresServer() {}

// UnsafeTicketScoresServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketScoresServer will
// result in compilation errors.
type UnsafeTicketScoresServer interface {
	mustEmbedUnimplementedTicketScoresServer()
}

func RegisterTicketScoresServer(s grpc.ServiceRegistrar, srv TicketScoresServer) {
	s.RegisterService(&TicketScores_ServiceDesc, srv)
}

func _TicketScores_SendTicketScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketScoresServer).SendTicketScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketScores_SendTicketScores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketScoresServer).SendTicketScores(ctx, req.(*TicketScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketScores_ServiceDesc is the grpc.ServiceDesc for TicketScores service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketScores_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TicketScores",
	HandlerType: (*TicketScoresServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTicketScores",
			Handler:    _TicketScores_SendTicketScores_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}

const (
	OverallScore_SendOverallScore_FullMethodName = "/proto.OverallScore/SendOverallScore"
)

// OverallScoreClient is the client API for OverallScore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OverallScoreClient interface {
	SendOverallScore(ctx context.Context, in *OverallScoreRequest, opts ...grpc.CallOption) (*OverallScoreReply, error)
}

type overallScoreClient struct {
	cc grpc.ClientConnInterface
}

func NewOverallScoreClient(cc grpc.ClientConnInterface) OverallScoreClient {
	return &overallScoreClient{cc}
}

func (c *overallScoreClient) SendOverallScore(ctx context.Context, in *OverallScoreRequest, opts ...grpc.CallOption) (*OverallScoreReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OverallScoreReply)
	err := c.cc.Invoke(ctx, OverallScore_SendOverallScore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OverallScoreServer is the server API for OverallScore service.
// All implementations must embed UnimplementedOverallScoreServer
// for forward compatibility
type OverallScoreServer interface {
	SendOverallScore(context.Context, *OverallScoreRequest) (*OverallScoreReply, error)
	mustEmbedUnimplementedOverallScoreServer()
}

// UnimplementedOverallScoreServer must be embedded to have forward compatible implementations.
type UnimplementedOverallScoreServer struct {
}

func (UnimplementedOverallScoreServer) SendOverallScore(context.Context, *OverallScoreRequest) (*OverallScoreReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOverallScore not implemented")
}
func (UnimplementedOverallScoreServer) mustEmbedUnimplementedOverallScoreServer() {}

// UnsafeOverallScoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OverallScoreServer will
// result in compilation errors.
type UnsafeOverallScoreServer interface {
	mustEmbedUnimplementedOverallScoreServer()
}

func RegisterOverallScoreServer(s grpc.ServiceRegistrar, srv OverallScoreServer) {
	s.RegisterService(&OverallScore_ServiceDesc, srv)
}

func _OverallScore_SendOverallScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OverallScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OverallScoreServer).SendOverallScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OverallScore_SendOverallScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OverallScoreServer).SendOverallScore(ctx, req.(*OverallScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OverallScore_ServiceDesc is the grpc.ServiceDesc for OverallScore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OverallScore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.OverallScore",
	HandlerType: (*OverallScoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOverallScore",
			Handler:    _OverallScore_SendOverallScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}

const (
	ChangeInScore_SendChangeInScore_FullMethodName = "/proto.ChangeInScore/SendChangeInScore"
)

// ChangeInScoreClient is the client API for ChangeInScore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChangeInScoreClient interface {
	SendChangeInScore(ctx context.Context, in *ChangeInScoreRequest, opts ...grpc.CallOption) (*ChangeInScoreReply, error)
}

type changeInScoreClient struct {
	cc grpc.ClientConnInterface
}

func NewChangeInScoreClient(cc grpc.ClientConnInterface) ChangeInScoreClient {
	return &changeInScoreClient{cc}
}

func (c *changeInScoreClient) SendChangeInScore(ctx context.Context, in *ChangeInScoreRequest, opts ...grpc.CallOption) (*ChangeInScoreReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangeInScoreReply)
	err := c.cc.Invoke(ctx, ChangeInScore_SendChangeInScore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChangeInScoreServer is the server API for ChangeInScore service.
// All implementations must embed UnimplementedChangeInScoreServer
// for forward compatibility
type ChangeInScoreServer interface {
	SendChangeInScore(context.Context, *ChangeInScoreRequest) (*ChangeInScoreReply, error)
	mustEmbedUnimplementedChangeInScoreServer()
}

// UnimplementedChangeInScoreServer must be embedded to have forward compatible implementations.
type UnimplementedChangeInScoreServer struct {
}

func (UnimplementedChangeInScoreServer) SendChangeInScore(context.Context, *ChangeInScoreRequest) (*ChangeInScoreReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendChangeInScore not implemented")
}
func (UnimplementedChangeInScoreServer) mustEmbedUnimplementedChangeInScoreServer() {}

// UnsafeChangeInScoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChangeInScoreServer will
// result in compilation errors.
type UnsafeChangeInScoreServer interface {
	mustEmbedUnimplementedChangeInScoreServer()
}

func RegisterChangeInScoreServer(s grpc.ServiceRegistrar, srv ChangeInScoreServer) {
	s.RegisterService(&ChangeInScore_ServiceDesc, srv)
}

func _ChangeInScore_SendChangeInScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeInScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChangeInScoreServer).SendChangeInScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChangeInScore_SendChangeInScore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChangeInScoreServer).SendChangeInScore(ctx, req.(*ChangeInScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChangeInScore_ServiceDesc is the grpc.ServiceDesc for ChangeInScore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChangeInScore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ChangeInScore",
	HandlerType: (*ChangeInScoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendChangeInScore",
			Handler:    _ChangeInScore_SendChangeInScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}
