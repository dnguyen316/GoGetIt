// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: api/go_get_it.proto

package go_get_it

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GoGetItService_CreateAccount_FullMethodName       = "/go_get_it.GoGetItService/CreateAccount"
	GoGetItService_CreateSession_FullMethodName       = "/go_get_it.GoGetItService/CreateSession"
	GoGetItService_CreateDownloadTask_FullMethodName  = "/go_get_it.GoGetItService/CreateDownloadTask"
	GoGetItService_GetDownloadTaskList_FullMethodName = "/go_get_it.GoGetItService/GetDownloadTaskList"
	GoGetItService_UpdateDownloadTask_FullMethodName  = "/go_get_it.GoGetItService/UpdateDownloadTask"
	GoGetItService_DeleteDownloadTask_FullMethodName  = "/go_get_it.GoGetItService/DeleteDownloadTask"
	GoGetItService_GetDownloadTaskFile_FullMethodName = "/go_get_it.GoGetItService/GetDownloadTaskFile"
)

// GoGetItServiceClient is the client API for GoGetItService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service definition with the 7 API endpoints.
type GoGetItServiceClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error)
	CreateDownloadTask(ctx context.Context, in *CreateDownloadTaskRequest, opts ...grpc.CallOption) (*CreateDownloadTaskResponse, error)
	GetDownloadTaskList(ctx context.Context, in *GetDownloadTaskListRequest, opts ...grpc.CallOption) (*GetDownloadTaskListResponse, error)
	UpdateDownloadTask(ctx context.Context, in *UpdateDownloadTaskRequest, opts ...grpc.CallOption) (*UpdateDownloadTaskResponse, error)
	DeleteDownloadTask(ctx context.Context, in *DeleteDownloadTaskRequest, opts ...grpc.CallOption) (*DeleteDownloadTaskResponse, error)
	GetDownloadTaskFile(ctx context.Context, in *GetDownloadTaskFileRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetDownloadTaskFileResponse], error)
}

type goGetItServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoGetItServiceClient(cc grpc.ClientConnInterface) GoGetItServiceClient {
	return &goGetItServiceClient{cc}
}

func (c *goGetItServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, GoGetItService_CreateAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goGetItServiceClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSessionResponse)
	err := c.cc.Invoke(ctx, GoGetItService_CreateSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goGetItServiceClient) CreateDownloadTask(ctx context.Context, in *CreateDownloadTaskRequest, opts ...grpc.CallOption) (*CreateDownloadTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDownloadTaskResponse)
	err := c.cc.Invoke(ctx, GoGetItService_CreateDownloadTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goGetItServiceClient) GetDownloadTaskList(ctx context.Context, in *GetDownloadTaskListRequest, opts ...grpc.CallOption) (*GetDownloadTaskListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDownloadTaskListResponse)
	err := c.cc.Invoke(ctx, GoGetItService_GetDownloadTaskList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goGetItServiceClient) UpdateDownloadTask(ctx context.Context, in *UpdateDownloadTaskRequest, opts ...grpc.CallOption) (*UpdateDownloadTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDownloadTaskResponse)
	err := c.cc.Invoke(ctx, GoGetItService_UpdateDownloadTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goGetItServiceClient) DeleteDownloadTask(ctx context.Context, in *DeleteDownloadTaskRequest, opts ...grpc.CallOption) (*DeleteDownloadTaskResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDownloadTaskResponse)
	err := c.cc.Invoke(ctx, GoGetItService_DeleteDownloadTask_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goGetItServiceClient) GetDownloadTaskFile(ctx context.Context, in *GetDownloadTaskFileRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetDownloadTaskFileResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &GoGetItService_ServiceDesc.Streams[0], GoGetItService_GetDownloadTaskFile_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetDownloadTaskFileRequest, GetDownloadTaskFileResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GoGetItService_GetDownloadTaskFileClient = grpc.ServerStreamingClient[GetDownloadTaskFileResponse]

// GoGetItServiceServer is the server API for GoGetItService service.
// All implementations must embed UnimplementedGoGetItServiceServer
// for forward compatibility.
//
// Service definition with the 7 API endpoints.
type GoGetItServiceServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error)
	CreateDownloadTask(context.Context, *CreateDownloadTaskRequest) (*CreateDownloadTaskResponse, error)
	GetDownloadTaskList(context.Context, *GetDownloadTaskListRequest) (*GetDownloadTaskListResponse, error)
	UpdateDownloadTask(context.Context, *UpdateDownloadTaskRequest) (*UpdateDownloadTaskResponse, error)
	DeleteDownloadTask(context.Context, *DeleteDownloadTaskRequest) (*DeleteDownloadTaskResponse, error)
	GetDownloadTaskFile(*GetDownloadTaskFileRequest, grpc.ServerStreamingServer[GetDownloadTaskFileResponse]) error
	mustEmbedUnimplementedGoGetItServiceServer()
}

// UnimplementedGoGetItServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGoGetItServiceServer struct{}

func (UnimplementedGoGetItServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedGoGetItServiceServer) CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedGoGetItServiceServer) CreateDownloadTask(context.Context, *CreateDownloadTaskRequest) (*CreateDownloadTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDownloadTask not implemented")
}
func (UnimplementedGoGetItServiceServer) GetDownloadTaskList(context.Context, *GetDownloadTaskListRequest) (*GetDownloadTaskListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadTaskList not implemented")
}
func (UnimplementedGoGetItServiceServer) UpdateDownloadTask(context.Context, *UpdateDownloadTaskRequest) (*UpdateDownloadTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDownloadTask not implemented")
}
func (UnimplementedGoGetItServiceServer) DeleteDownloadTask(context.Context, *DeleteDownloadTaskRequest) (*DeleteDownloadTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDownloadTask not implemented")
}
func (UnimplementedGoGetItServiceServer) GetDownloadTaskFile(*GetDownloadTaskFileRequest, grpc.ServerStreamingServer[GetDownloadTaskFileResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetDownloadTaskFile not implemented")
}
func (UnimplementedGoGetItServiceServer) mustEmbedUnimplementedGoGetItServiceServer() {}
func (UnimplementedGoGetItServiceServer) testEmbeddedByValue()                        {}

// UnsafeGoGetItServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoGetItServiceServer will
// result in compilation errors.
type UnsafeGoGetItServiceServer interface {
	mustEmbedUnimplementedGoGetItServiceServer()
}

func RegisterGoGetItServiceServer(s grpc.ServiceRegistrar, srv GoGetItServiceServer) {
	// If the following call pancis, it indicates UnimplementedGoGetItServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GoGetItService_ServiceDesc, srv)
}

func _GoGetItService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoGetItServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoGetItService_CreateAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoGetItServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoGetItService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoGetItServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoGetItService_CreateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoGetItServiceServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoGetItService_CreateDownloadTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDownloadTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoGetItServiceServer).CreateDownloadTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoGetItService_CreateDownloadTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoGetItServiceServer).CreateDownloadTask(ctx, req.(*CreateDownloadTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoGetItService_GetDownloadTaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDownloadTaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoGetItServiceServer).GetDownloadTaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoGetItService_GetDownloadTaskList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoGetItServiceServer).GetDownloadTaskList(ctx, req.(*GetDownloadTaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoGetItService_UpdateDownloadTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDownloadTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoGetItServiceServer).UpdateDownloadTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoGetItService_UpdateDownloadTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoGetItServiceServer).UpdateDownloadTask(ctx, req.(*UpdateDownloadTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoGetItService_DeleteDownloadTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDownloadTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoGetItServiceServer).DeleteDownloadTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GoGetItService_DeleteDownloadTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoGetItServiceServer).DeleteDownloadTask(ctx, req.(*DeleteDownloadTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoGetItService_GetDownloadTaskFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDownloadTaskFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoGetItServiceServer).GetDownloadTaskFile(m, &grpc.GenericServerStream[GetDownloadTaskFileRequest, GetDownloadTaskFileResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type GoGetItService_GetDownloadTaskFileServer = grpc.ServerStreamingServer[GetDownloadTaskFileResponse]

// GoGetItService_ServiceDesc is the grpc.ServiceDesc for GoGetItService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoGetItService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_get_it.GoGetItService",
	HandlerType: (*GoGetItServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _GoGetItService_CreateAccount_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _GoGetItService_CreateSession_Handler,
		},
		{
			MethodName: "CreateDownloadTask",
			Handler:    _GoGetItService_CreateDownloadTask_Handler,
		},
		{
			MethodName: "GetDownloadTaskList",
			Handler:    _GoGetItService_GetDownloadTaskList_Handler,
		},
		{
			MethodName: "UpdateDownloadTask",
			Handler:    _GoGetItService_UpdateDownloadTask_Handler,
		},
		{
			MethodName: "DeleteDownloadTask",
			Handler:    _GoGetItService_DeleteDownloadTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDownloadTaskFile",
			Handler:       _GoGetItService_GetDownloadTaskFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/go_get_it.proto",
}
