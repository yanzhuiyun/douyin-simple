// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: pb/interaction.proto

package pb

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

// InteractionClient is the client API for Interaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InteractionClient interface {
	FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*Empty, error)
	GetFavoriteList(ctx context.Context, in *GetFavoriteListReq, opts ...grpc.CallOption) (*GetFavoriteListReply, error)
	CommentAction(ctx context.Context, in *CommentActionReq, opts ...grpc.CallOption) (*CommentActionReply, error)
	GetCommentList(ctx context.Context, in *GetCommentListReq, opts ...grpc.CallOption) (*GetCommentListReply, error)
}

type interactionClient struct {
	cc grpc.ClientConnInterface
}

func NewInteractionClient(cc grpc.ClientConnInterface) InteractionClient {
	return &interactionClient{cc}
}

func (c *interactionClient) FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.interaction/FavoriteAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionClient) GetFavoriteList(ctx context.Context, in *GetFavoriteListReq, opts ...grpc.CallOption) (*GetFavoriteListReply, error) {
	out := new(GetFavoriteListReply)
	err := c.cc.Invoke(ctx, "/pb.interaction/GetFavoriteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionClient) CommentAction(ctx context.Context, in *CommentActionReq, opts ...grpc.CallOption) (*CommentActionReply, error) {
	out := new(CommentActionReply)
	err := c.cc.Invoke(ctx, "/pb.interaction/CommentAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interactionClient) GetCommentList(ctx context.Context, in *GetCommentListReq, opts ...grpc.CallOption) (*GetCommentListReply, error) {
	out := new(GetCommentListReply)
	err := c.cc.Invoke(ctx, "/pb.interaction/GetCommentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InteractionServer is the server API for Interaction service.
// All implementations must embed UnimplementedInteractionServer
// for forward compatibility
type InteractionServer interface {
	FavoriteAction(context.Context, *FavoriteActionReq) (*Empty, error)
	GetFavoriteList(context.Context, *GetFavoriteListReq) (*GetFavoriteListReply, error)
	CommentAction(context.Context, *CommentActionReq) (*CommentActionReply, error)
	GetCommentList(context.Context, *GetCommentListReq) (*GetCommentListReply, error)
	mustEmbedUnimplementedInteractionServer()
}

// UnimplementedInteractionServer must be embedded to have forward compatible implementations.
type UnimplementedInteractionServer struct {
}

func (UnimplementedInteractionServer) FavoriteAction(context.Context, *FavoriteActionReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteAction not implemented")
}
func (UnimplementedInteractionServer) GetFavoriteList(context.Context, *GetFavoriteListReq) (*GetFavoriteListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavoriteList not implemented")
}
func (UnimplementedInteractionServer) CommentAction(context.Context, *CommentActionReq) (*CommentActionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
func (UnimplementedInteractionServer) GetCommentList(context.Context, *GetCommentListReq) (*GetCommentListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentList not implemented")
}
func (UnimplementedInteractionServer) mustEmbedUnimplementedInteractionServer() {}

// UnsafeInteractionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InteractionServer will
// result in compilation errors.
type UnsafeInteractionServer interface {
	mustEmbedUnimplementedInteractionServer()
}

func RegisterInteractionServer(s grpc.ServiceRegistrar, srv InteractionServer) {
	s.RegisterService(&Interaction_ServiceDesc, srv)
}

func _Interaction_FavoriteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServer).FavoriteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.interaction/FavoriteAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServer).FavoriteAction(ctx, req.(*FavoriteActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interaction_GetFavoriteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFavoriteListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServer).GetFavoriteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.interaction/GetFavoriteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServer).GetFavoriteList(ctx, req.(*GetFavoriteListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interaction_CommentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServer).CommentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.interaction/CommentAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServer).CommentAction(ctx, req.(*CommentActionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Interaction_GetCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InteractionServer).GetCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.interaction/GetCommentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InteractionServer).GetCommentList(ctx, req.(*GetCommentListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Interaction_ServiceDesc is the grpc.ServiceDesc for Interaction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Interaction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.interaction",
	HandlerType: (*InteractionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FavoriteAction",
			Handler:    _Interaction_FavoriteAction_Handler,
		},
		{
			MethodName: "GetFavoriteList",
			Handler:    _Interaction_GetFavoriteList_Handler,
		},
		{
			MethodName: "CommentAction",
			Handler:    _Interaction_CommentAction_Handler,
		},
		{
			MethodName: "GetCommentList",
			Handler:    _Interaction_GetCommentList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/interaction.proto",
}
