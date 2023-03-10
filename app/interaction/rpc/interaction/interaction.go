// Code generated by goctl. DO NOT EDIT!
// Source: interaction.proto

package interaction

import (
	"context"

	"douyin-simple/app/interaction/rpc/pb/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Comment              = pb.Comment
	CommentActionReply   = pb.CommentActionReply
	CommentActionReq     = pb.CommentActionReq
	Empty                = pb.Empty
	FavoriteActionReq    = pb.FavoriteActionReq
	GetCommentListReply  = pb.GetCommentListReply
	GetCommentListReq    = pb.GetCommentListReq
	GetFavoriteListReply = pb.GetFavoriteListReply
	GetFavoriteListReq   = pb.GetFavoriteListReq
	User                 = pb.User
	Video                = pb.Video

	Interaction interface {
		FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*Empty, error)
		GetFavoriteList(ctx context.Context, in *GetFavoriteListReq, opts ...grpc.CallOption) (*GetFavoriteListReply, error)
		CommentAction(ctx context.Context, in *CommentActionReq, opts ...grpc.CallOption) (*CommentActionReply, error)
		GetCommentList(ctx context.Context, in *GetCommentListReq, opts ...grpc.CallOption) (*GetCommentListReply, error)
	}

	defaultInteraction struct {
		cli zrpc.Client
	}
)

func NewInteraction(cli zrpc.Client) Interaction {
	return &defaultInteraction{
		cli: cli,
	}
}

func (m *defaultInteraction) FavoriteAction(ctx context.Context, in *FavoriteActionReq, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewInteractionClient(m.cli.Conn())
	return client.FavoriteAction(ctx, in, opts...)
}

func (m *defaultInteraction) GetFavoriteList(ctx context.Context, in *GetFavoriteListReq, opts ...grpc.CallOption) (*GetFavoriteListReply, error) {
	client := pb.NewInteractionClient(m.cli.Conn())
	return client.GetFavoriteList(ctx, in, opts...)
}

func (m *defaultInteraction) CommentAction(ctx context.Context, in *CommentActionReq, opts ...grpc.CallOption) (*CommentActionReply, error) {
	client := pb.NewInteractionClient(m.cli.Conn())
	return client.CommentAction(ctx, in, opts...)
}

func (m *defaultInteraction) GetCommentList(ctx context.Context, in *GetCommentListReq, opts ...grpc.CallOption) (*GetCommentListReply, error) {
	client := pb.NewInteractionClient(m.cli.Conn())
	return client.GetCommentList(ctx, in, opts...)
}
