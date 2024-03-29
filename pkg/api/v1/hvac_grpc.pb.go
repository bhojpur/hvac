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

// HvacServiceClient is the client API for HvacService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HvacServiceClient interface {
	// StartLocalCooler starts a Cooler on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the hvac/config.yaml
	//   3. all bytes constituting the Cooler YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalCooler(ctx context.Context, opts ...grpc.CallOption) (HvacService_StartLocalCoolerClient, error)
	// StartFromPreviousCooler starts a new Cooler based on a previous one.
	// If the previous Cooler does not have the can-replay condition set this call will result in an error.
	StartFromPreviousCooler(ctx context.Context, in *StartFromPreviousCoolerRequest, opts ...grpc.CallOption) (*StartCoolerResponse, error)
	// StartCoolerRequest starts a new Cooler based on its specification.
	StartCooler(ctx context.Context, in *StartCoolerRequest, opts ...grpc.CallOption) (*StartCoolerResponse, error)
	// Searches for Cooler(s) known to this instance
	ListCoolers(ctx context.Context, in *ListCoolersRequest, opts ...grpc.CallOption) (*ListCoolersResponse, error)
	// Subscribe listens to new Cooler(s) updates
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (HvacService_SubscribeClient, error)
	// GetCooler retrieves details of a single Cooler
	GetCooler(ctx context.Context, in *GetCoolerRequest, opts ...grpc.CallOption) (*GetCoolerResponse, error)
	// Listen listens to Cooler updates and log output of a running Cooler
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (HvacService_ListenClient, error)
	// StopCooler stops a currently running Cooler
	StopCooler(ctx context.Context, in *StopCoolerRequest, opts ...grpc.CallOption) (*StopCoolerResponse, error)
}

type hvacServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHvacServiceClient(cc grpc.ClientConnInterface) HvacServiceClient {
	return &hvacServiceClient{cc}
}

func (c *hvacServiceClient) StartLocalCooler(ctx context.Context, opts ...grpc.CallOption) (HvacService_StartLocalCoolerClient, error) {
	stream, err := c.cc.NewStream(ctx, &HvacService_ServiceDesc.Streams[0], "/v1.HvacService/StartLocalCooler", opts...)
	if err != nil {
		return nil, err
	}
	x := &hvacServiceStartLocalCoolerClient{stream}
	return x, nil
}

type HvacService_StartLocalCoolerClient interface {
	Send(*StartLocalCoolerRequest) error
	CloseAndRecv() (*StartCoolerResponse, error)
	grpc.ClientStream
}

type hvacServiceStartLocalCoolerClient struct {
	grpc.ClientStream
}

func (x *hvacServiceStartLocalCoolerClient) Send(m *StartLocalCoolerRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hvacServiceStartLocalCoolerClient) CloseAndRecv() (*StartCoolerResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StartCoolerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hvacServiceClient) StartFromPreviousCooler(ctx context.Context, in *StartFromPreviousCoolerRequest, opts ...grpc.CallOption) (*StartCoolerResponse, error) {
	out := new(StartCoolerResponse)
	err := c.cc.Invoke(ctx, "/v1.HvacService/StartFromPreviousCooler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hvacServiceClient) StartCooler(ctx context.Context, in *StartCoolerRequest, opts ...grpc.CallOption) (*StartCoolerResponse, error) {
	out := new(StartCoolerResponse)
	err := c.cc.Invoke(ctx, "/v1.HvacService/StartCooler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hvacServiceClient) ListCoolers(ctx context.Context, in *ListCoolersRequest, opts ...grpc.CallOption) (*ListCoolersResponse, error) {
	out := new(ListCoolersResponse)
	err := c.cc.Invoke(ctx, "/v1.HvacService/ListCoolers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hvacServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (HvacService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &HvacService_ServiceDesc.Streams[1], "/v1.HvacService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &hvacServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HvacService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type hvacServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *hvacServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hvacServiceClient) GetCooler(ctx context.Context, in *GetCoolerRequest, opts ...grpc.CallOption) (*GetCoolerResponse, error) {
	out := new(GetCoolerResponse)
	err := c.cc.Invoke(ctx, "/v1.HvacService/GetCooler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hvacServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (HvacService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &HvacService_ServiceDesc.Streams[2], "/v1.HvacService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &hvacServiceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HvacService_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type hvacServiceListenClient struct {
	grpc.ClientStream
}

func (x *hvacServiceListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hvacServiceClient) StopCooler(ctx context.Context, in *StopCoolerRequest, opts ...grpc.CallOption) (*StopCoolerResponse, error) {
	out := new(StopCoolerResponse)
	err := c.cc.Invoke(ctx, "/v1.HvacService/StopCooler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HvacServiceServer is the server API for HvacService service.
// All implementations must embed UnimplementedHvacServiceServer
// for forward compatibility
type HvacServiceServer interface {
	// StartLocalCooler starts a Cooler on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the hvac/config.yaml
	//   3. all bytes constituting the Cooler YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalCooler(HvacService_StartLocalCoolerServer) error
	// StartFromPreviousCooler starts a new Cooler based on a previous one.
	// If the previous Cooler does not have the can-replay condition set this call will result in an error.
	StartFromPreviousCooler(context.Context, *StartFromPreviousCoolerRequest) (*StartCoolerResponse, error)
	// StartCoolerRequest starts a new Cooler based on its specification.
	StartCooler(context.Context, *StartCoolerRequest) (*StartCoolerResponse, error)
	// Searches for Cooler(s) known to this instance
	ListCoolers(context.Context, *ListCoolersRequest) (*ListCoolersResponse, error)
	// Subscribe listens to new Cooler(s) updates
	Subscribe(*SubscribeRequest, HvacService_SubscribeServer) error
	// GetCooler retrieves details of a single Cooler
	GetCooler(context.Context, *GetCoolerRequest) (*GetCoolerResponse, error)
	// Listen listens to Cooler updates and log output of a running Cooler
	Listen(*ListenRequest, HvacService_ListenServer) error
	// StopCooler stops a currently running Cooler
	StopCooler(context.Context, *StopCoolerRequest) (*StopCoolerResponse, error)
	mustEmbedUnimplementedHvacServiceServer()
}

// UnimplementedHvacServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHvacServiceServer struct {
}

func (UnimplementedHvacServiceServer) StartLocalCooler(HvacService_StartLocalCoolerServer) error {
	return status.Errorf(codes.Unimplemented, "method StartLocalCooler not implemented")
}
func (UnimplementedHvacServiceServer) StartFromPreviousCooler(context.Context, *StartFromPreviousCoolerRequest) (*StartCoolerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFromPreviousCooler not implemented")
}
func (UnimplementedHvacServiceServer) StartCooler(context.Context, *StartCoolerRequest) (*StartCoolerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartCooler not implemented")
}
func (UnimplementedHvacServiceServer) ListCoolers(context.Context, *ListCoolersRequest) (*ListCoolersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCoolers not implemented")
}
func (UnimplementedHvacServiceServer) Subscribe(*SubscribeRequest, HvacService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedHvacServiceServer) GetCooler(context.Context, *GetCoolerRequest) (*GetCoolerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCooler not implemented")
}
func (UnimplementedHvacServiceServer) Listen(*ListenRequest, HvacService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedHvacServiceServer) StopCooler(context.Context, *StopCoolerRequest) (*StopCoolerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopCooler not implemented")
}
func (UnimplementedHvacServiceServer) mustEmbedUnimplementedHvacServiceServer() {}

// UnsafeHvacServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HvacServiceServer will
// result in compilation errors.
type UnsafeHvacServiceServer interface {
	mustEmbedUnimplementedHvacServiceServer()
}

func RegisterHvacServiceServer(s grpc.ServiceRegistrar, srv HvacServiceServer) {
	s.RegisterService(&HvacService_ServiceDesc, srv)
}

func _HvacService_StartLocalCooler_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HvacServiceServer).StartLocalCooler(&hvacServiceStartLocalCoolerServer{stream})
}

type HvacService_StartLocalCoolerServer interface {
	SendAndClose(*StartCoolerResponse) error
	Recv() (*StartLocalCoolerRequest, error)
	grpc.ServerStream
}

type hvacServiceStartLocalCoolerServer struct {
	grpc.ServerStream
}

func (x *hvacServiceStartLocalCoolerServer) SendAndClose(m *StartCoolerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hvacServiceStartLocalCoolerServer) Recv() (*StartLocalCoolerRequest, error) {
	m := new(StartLocalCoolerRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HvacService_StartFromPreviousCooler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFromPreviousCoolerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HvacServiceServer).StartFromPreviousCooler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.HvacService/StartFromPreviousCooler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HvacServiceServer).StartFromPreviousCooler(ctx, req.(*StartFromPreviousCoolerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HvacService_StartCooler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartCoolerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HvacServiceServer).StartCooler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.HvacService/StartCooler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HvacServiceServer).StartCooler(ctx, req.(*StartCoolerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HvacService_ListCoolers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCoolersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HvacServiceServer).ListCoolers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.HvacService/ListCoolers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HvacServiceServer).ListCoolers(ctx, req.(*ListCoolersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HvacService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HvacServiceServer).Subscribe(m, &hvacServiceSubscribeServer{stream})
}

type HvacService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type hvacServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *hvacServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HvacService_GetCooler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoolerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HvacServiceServer).GetCooler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.HvacService/GetCooler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HvacServiceServer).GetCooler(ctx, req.(*GetCoolerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HvacService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HvacServiceServer).Listen(m, &hvacServiceListenServer{stream})
}

type HvacService_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type hvacServiceListenServer struct {
	grpc.ServerStream
}

func (x *hvacServiceListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HvacService_StopCooler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopCoolerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HvacServiceServer).StopCooler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.HvacService/StopCooler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HvacServiceServer).StopCooler(ctx, req.(*StopCoolerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HvacService_ServiceDesc is the grpc.ServiceDesc for HvacService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HvacService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.HvacService",
	HandlerType: (*HvacServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartFromPreviousCooler",
			Handler:    _HvacService_StartFromPreviousCooler_Handler,
		},
		{
			MethodName: "StartCooler",
			Handler:    _HvacService_StartCooler_Handler,
		},
		{
			MethodName: "ListCoolers",
			Handler:    _HvacService_ListCoolers_Handler,
		},
		{
			MethodName: "GetCooler",
			Handler:    _HvacService_GetCooler_Handler,
		},
		{
			MethodName: "StopCooler",
			Handler:    _HvacService_StopCooler_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartLocalCooler",
			Handler:       _HvacService_StartLocalCooler_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _HvacService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Listen",
			Handler:       _HvacService_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hvac.proto",
}
