// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: sso.proto

package Web_With_Articles_ssov1

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
	WebService_Register_FullMethodName      = "/auth.WebService/Register"
	WebService_Login_FullMethodName         = "/auth.WebService/Login"
	WebService_CreateArticle_FullMethodName = "/auth.WebService/CreateArticle"
	WebService_GetArticles_FullMethodName   = "/auth.WebService/GetArticles"
	WebService_AddComment_FullMethodName    = "/auth.WebService/AddComment"
	WebService_DeleteComment_FullMethodName = "/auth.WebService/DeleteComment"
)

// WebServiceClient is the client API for WebService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebServiceClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	CreateArticle(ctx context.Context, in *Article, opts ...grpc.CallOption) (*CreateArticleResponse, error)
	GetArticles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Article], error)
	AddComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*AddCommentResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommRequest, opts ...grpc.CallOption) (*Empty, error)
}

type webServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebServiceClient(cc grpc.ClientConnInterface) WebServiceClient {
	return &webServiceClient{cc}
}

func (c *webServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, WebService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, WebService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webServiceClient) CreateArticle(ctx context.Context, in *Article, opts ...grpc.CallOption) (*CreateArticleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateArticleResponse)
	err := c.cc.Invoke(ctx, WebService_CreateArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webServiceClient) GetArticles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Article], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &WebService_ServiceDesc.Streams[0], WebService_GetArticles_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Empty, Article]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WebService_GetArticlesClient = grpc.ServerStreamingClient[Article]

func (c *webServiceClient) AddComment(ctx context.Context, in *Comment, opts ...grpc.CallOption) (*AddCommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCommentResponse)
	err := c.cc.Invoke(ctx, WebService_AddComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webServiceClient) DeleteComment(ctx context.Context, in *DeleteCommRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, WebService_DeleteComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebServiceServer is the server API for WebService service.
// All implementations must embed UnimplementedWebServiceServer
// for forward compatibility.
type WebServiceServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	CreateArticle(context.Context, *Article) (*CreateArticleResponse, error)
	GetArticles(*Empty, grpc.ServerStreamingServer[Article]) error
	AddComment(context.Context, *Comment) (*AddCommentResponse, error)
	DeleteComment(context.Context, *DeleteCommRequest) (*Empty, error)
	mustEmbedUnimplementedWebServiceServer()
}

// UnimplementedWebServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWebServiceServer struct{}

func (UnimplementedWebServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedWebServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedWebServiceServer) CreateArticle(context.Context, *Article) (*CreateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedWebServiceServer) GetArticles(*Empty, grpc.ServerStreamingServer[Article]) error {
	return status.Errorf(codes.Unimplemented, "method GetArticles not implemented")
}
func (UnimplementedWebServiceServer) AddComment(context.Context, *Comment) (*AddCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedWebServiceServer) DeleteComment(context.Context, *DeleteCommRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedWebServiceServer) mustEmbedUnimplementedWebServiceServer() {}
func (UnimplementedWebServiceServer) testEmbeddedByValue()                    {}

// UnsafeWebServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebServiceServer will
// result in compilation errors.
type UnsafeWebServiceServer interface {
	mustEmbedUnimplementedWebServiceServer()
}

func RegisterWebServiceServer(s grpc.ServiceRegistrar, srv WebServiceServer) {
	// If the following call pancis, it indicates UnimplementedWebServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WebService_ServiceDesc, srv)
}

func _WebService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebService_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Article)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServiceServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebService_CreateArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServiceServer).CreateArticle(ctx, req.(*Article))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebService_GetArticles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WebServiceServer).GetArticles(m, &grpc.GenericServerStream[Empty, Article]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type WebService_GetArticlesServer = grpc.ServerStreamingServer[Article]

func _WebService_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Comment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServiceServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebService_AddComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServiceServer).AddComment(ctx, req.(*Comment))
	}
	return interceptor(ctx, in, info, handler)
}

func _WebService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebService_DeleteComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebServiceServer).DeleteComment(ctx, req.(*DeleteCommRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WebService_ServiceDesc is the grpc.ServiceDesc for WebService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.WebService",
	HandlerType: (*WebServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _WebService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _WebService_Login_Handler,
		},
		{
			MethodName: "CreateArticle",
			Handler:    _WebService_CreateArticle_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _WebService_AddComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _WebService_DeleteComment_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetArticles",
			Handler:       _WebService_GetArticles_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sso.proto",
}
