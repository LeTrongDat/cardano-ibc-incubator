// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package types

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

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// IncentivizedPackets returns all incentivized packets and their associated fees
	IncentivizedPackets(ctx context.Context, in *QueryIncentivizedPacketsRequest, opts ...grpc.CallOption) (*QueryIncentivizedPacketsResponse, error)
	// IncentivizedPacket returns all packet fees for a packet given its identifier
	IncentivizedPacket(ctx context.Context, in *QueryIncentivizedPacketRequest, opts ...grpc.CallOption) (*QueryIncentivizedPacketResponse, error)
	// Gets all incentivized packets for a specific channel
	IncentivizedPacketsForChannel(ctx context.Context, in *QueryIncentivizedPacketsForChannelRequest, opts ...grpc.CallOption) (*QueryIncentivizedPacketsForChannelResponse, error)
	// TotalRecvFees returns the total receive fees for a packet given its identifier
	TotalRecvFees(ctx context.Context, in *QueryTotalRecvFeesRequest, opts ...grpc.CallOption) (*QueryTotalRecvFeesResponse, error)
	// TotalAckFees returns the total acknowledgement fees for a packet given its identifier
	TotalAckFees(ctx context.Context, in *QueryTotalAckFeesRequest, opts ...grpc.CallOption) (*QueryTotalAckFeesResponse, error)
	// TotalTimeoutFees returns the total timeout fees for a packet given its identifier
	TotalTimeoutFees(ctx context.Context, in *QueryTotalTimeoutFeesRequest, opts ...grpc.CallOption) (*QueryTotalTimeoutFeesResponse, error)
	// Payee returns the registered payee address for a specific channel given the relayer address
	Payee(ctx context.Context, in *QueryPayeeRequest, opts ...grpc.CallOption) (*QueryPayeeResponse, error)
	// CounterpartyPayee returns the registered counterparty payee for forward relaying
	CounterpartyPayee(ctx context.Context, in *QueryCounterpartyPayeeRequest, opts ...grpc.CallOption) (*QueryCounterpartyPayeeResponse, error)
	// FeeEnabledChannels returns a list of all fee enabled channels
	FeeEnabledChannels(ctx context.Context, in *QueryFeeEnabledChannelsRequest, opts ...grpc.CallOption) (*QueryFeeEnabledChannelsResponse, error)
	// FeeEnabledChannel returns true if the provided port and channel identifiers belong to a fee enabled channel
	FeeEnabledChannel(ctx context.Context, in *QueryFeeEnabledChannelRequest, opts ...grpc.CallOption) (*QueryFeeEnabledChannelResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) IncentivizedPackets(ctx context.Context, in *QueryIncentivizedPacketsRequest, opts ...grpc.CallOption) (*QueryIncentivizedPacketsResponse, error) {
	out := new(QueryIncentivizedPacketsResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.fee.v1.Query/IncentivizedPackets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IncentivizedPacket(ctx context.Context, in *QueryIncentivizedPacketRequest, opts ...grpc.CallOption) (*QueryIncentivizedPacketResponse, error) {
	out := new(QueryIncentivizedPacketResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.fee.v1.Query/IncentivizedPacket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IncentivizedPacketsForChannel(ctx context.Context, in *QueryIncentivizedPacketsForChannelRequest, opts ...grpc.CallOption) (*QueryIncentivizedPacketsForChannelResponse, error) {
	out := new(QueryIncentivizedPacketsForChannelResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.fee.v1.Query/IncentivizedPacketsForChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TotalRecvFees(ctx context.Context, in *QueryTotalRecvFeesRequest, opts ...grpc.CallOption) (*QueryTotalRecvFeesResponse, error) {
	out := new(QueryTotalRecvFeesResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.fee.v1.Query/TotalRecvFees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TotalAckFees(ctx context.Context, in *QueryTotalAckFeesRequest, opts ...grpc.CallOption) (*QueryTotalAckFeesResponse, error) {
	out := new(QueryTotalAckFeesResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.fee.v1.Query/TotalAckFees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TotalTimeoutFees(ctx context.Context, in *QueryTotalTimeoutFeesRequest, opts ...grpc.CallOption) (*QueryTotalTimeoutFeesResponse, error) {
	out := new(QueryTotalTimeoutFeesResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.fee.v1.Query/TotalTimeoutFees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Payee(ctx context.Context, in *QueryPayeeRequest, opts ...grpc.CallOption) (*QueryPayeeResponse, error) {
	out := new(QueryPayeeResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.fee.v1.Query/Payee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CounterpartyPayee(ctx context.Context, in *QueryCounterpartyPayeeRequest, opts ...grpc.CallOption) (*QueryCounterpartyPayeeResponse, error) {
	out := new(QueryCounterpartyPayeeResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.fee.v1.Query/CounterpartyPayee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FeeEnabledChannels(ctx context.Context, in *QueryFeeEnabledChannelsRequest, opts ...grpc.CallOption) (*QueryFeeEnabledChannelsResponse, error) {
	out := new(QueryFeeEnabledChannelsResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.fee.v1.Query/FeeEnabledChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FeeEnabledChannel(ctx context.Context, in *QueryFeeEnabledChannelRequest, opts ...grpc.CallOption) (*QueryFeeEnabledChannelResponse, error) {
	out := new(QueryFeeEnabledChannelResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.fee.v1.Query/FeeEnabledChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// IncentivizedPackets returns all incentivized packets and their associated fees
	IncentivizedPackets(context.Context, *QueryIncentivizedPacketsRequest) (*QueryIncentivizedPacketsResponse, error)
	// IncentivizedPacket returns all packet fees for a packet given its identifier
	IncentivizedPacket(context.Context, *QueryIncentivizedPacketRequest) (*QueryIncentivizedPacketResponse, error)
	// Gets all incentivized packets for a specific channel
	IncentivizedPacketsForChannel(context.Context, *QueryIncentivizedPacketsForChannelRequest) (*QueryIncentivizedPacketsForChannelResponse, error)
	// TotalRecvFees returns the total receive fees for a packet given its identifier
	TotalRecvFees(context.Context, *QueryTotalRecvFeesRequest) (*QueryTotalRecvFeesResponse, error)
	// TotalAckFees returns the total acknowledgement fees for a packet given its identifier
	TotalAckFees(context.Context, *QueryTotalAckFeesRequest) (*QueryTotalAckFeesResponse, error)
	// TotalTimeoutFees returns the total timeout fees for a packet given its identifier
	TotalTimeoutFees(context.Context, *QueryTotalTimeoutFeesRequest) (*QueryTotalTimeoutFeesResponse, error)
	// Payee returns the registered payee address for a specific channel given the relayer address
	Payee(context.Context, *QueryPayeeRequest) (*QueryPayeeResponse, error)
	// CounterpartyPayee returns the registered counterparty payee for forward relaying
	CounterpartyPayee(context.Context, *QueryCounterpartyPayeeRequest) (*QueryCounterpartyPayeeResponse, error)
	// FeeEnabledChannels returns a list of all fee enabled channels
	FeeEnabledChannels(context.Context, *QueryFeeEnabledChannelsRequest) (*QueryFeeEnabledChannelsResponse, error)
	// FeeEnabledChannel returns true if the provided port and channel identifiers belong to a fee enabled channel
	FeeEnabledChannel(context.Context, *QueryFeeEnabledChannelRequest) (*QueryFeeEnabledChannelResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) IncentivizedPackets(context.Context, *QueryIncentivizedPacketsRequest) (*QueryIncentivizedPacketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncentivizedPackets not implemented")
}
func (UnimplementedQueryServer) IncentivizedPacket(context.Context, *QueryIncentivizedPacketRequest) (*QueryIncentivizedPacketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncentivizedPacket not implemented")
}
func (UnimplementedQueryServer) IncentivizedPacketsForChannel(context.Context, *QueryIncentivizedPacketsForChannelRequest) (*QueryIncentivizedPacketsForChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncentivizedPacketsForChannel not implemented")
}
func (UnimplementedQueryServer) TotalRecvFees(context.Context, *QueryTotalRecvFeesRequest) (*QueryTotalRecvFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalRecvFees not implemented")
}
func (UnimplementedQueryServer) TotalAckFees(context.Context, *QueryTotalAckFeesRequest) (*QueryTotalAckFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalAckFees not implemented")
}
func (UnimplementedQueryServer) TotalTimeoutFees(context.Context, *QueryTotalTimeoutFeesRequest) (*QueryTotalTimeoutFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalTimeoutFees not implemented")
}
func (UnimplementedQueryServer) Payee(context.Context, *QueryPayeeRequest) (*QueryPayeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Payee not implemented")
}
func (UnimplementedQueryServer) CounterpartyPayee(context.Context, *QueryCounterpartyPayeeRequest) (*QueryCounterpartyPayeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CounterpartyPayee not implemented")
}
func (UnimplementedQueryServer) FeeEnabledChannels(context.Context, *QueryFeeEnabledChannelsRequest) (*QueryFeeEnabledChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeeEnabledChannels not implemented")
}
func (UnimplementedQueryServer) FeeEnabledChannel(context.Context, *QueryFeeEnabledChannelRequest) (*QueryFeeEnabledChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeeEnabledChannel not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_IncentivizedPackets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIncentivizedPacketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IncentivizedPackets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.fee.v1.Query/IncentivizedPackets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IncentivizedPackets(ctx, req.(*QueryIncentivizedPacketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IncentivizedPacket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIncentivizedPacketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IncentivizedPacket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.fee.v1.Query/IncentivizedPacket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IncentivizedPacket(ctx, req.(*QueryIncentivizedPacketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IncentivizedPacketsForChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIncentivizedPacketsForChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IncentivizedPacketsForChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.fee.v1.Query/IncentivizedPacketsForChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IncentivizedPacketsForChannel(ctx, req.(*QueryIncentivizedPacketsForChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TotalRecvFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalRecvFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalRecvFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.fee.v1.Query/TotalRecvFees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalRecvFees(ctx, req.(*QueryTotalRecvFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TotalAckFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalAckFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalAckFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.fee.v1.Query/TotalAckFees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalAckFees(ctx, req.(*QueryTotalAckFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TotalTimeoutFees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalTimeoutFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalTimeoutFees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.fee.v1.Query/TotalTimeoutFees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalTimeoutFees(ctx, req.(*QueryTotalTimeoutFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Payee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPayeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Payee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.fee.v1.Query/Payee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Payee(ctx, req.(*QueryPayeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CounterpartyPayee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCounterpartyPayeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CounterpartyPayee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.fee.v1.Query/CounterpartyPayee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CounterpartyPayee(ctx, req.(*QueryCounterpartyPayeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FeeEnabledChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFeeEnabledChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FeeEnabledChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.fee.v1.Query/FeeEnabledChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FeeEnabledChannels(ctx, req.(*QueryFeeEnabledChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FeeEnabledChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFeeEnabledChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FeeEnabledChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.fee.v1.Query/FeeEnabledChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FeeEnabledChannel(ctx, req.(*QueryFeeEnabledChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.applications.fee.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IncentivizedPackets",
			Handler:    _Query_IncentivizedPackets_Handler,
		},
		{
			MethodName: "IncentivizedPacket",
			Handler:    _Query_IncentivizedPacket_Handler,
		},
		{
			MethodName: "IncentivizedPacketsForChannel",
			Handler:    _Query_IncentivizedPacketsForChannel_Handler,
		},
		{
			MethodName: "TotalRecvFees",
			Handler:    _Query_TotalRecvFees_Handler,
		},
		{
			MethodName: "TotalAckFees",
			Handler:    _Query_TotalAckFees_Handler,
		},
		{
			MethodName: "TotalTimeoutFees",
			Handler:    _Query_TotalTimeoutFees_Handler,
		},
		{
			MethodName: "Payee",
			Handler:    _Query_Payee_Handler,
		},
		{
			MethodName: "CounterpartyPayee",
			Handler:    _Query_CounterpartyPayee_Handler,
		},
		{
			MethodName: "FeeEnabledChannels",
			Handler:    _Query_FeeEnabledChannels_Handler,
		},
		{
			MethodName: "FeeEnabledChannel",
			Handler:    _Query_FeeEnabledChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/applications/fee/v1/query.proto",
}
