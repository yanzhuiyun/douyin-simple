// Code generated by goctl. DO NOT EDIT!
// Source: interaction.proto

package server

import (
	"context"

	"douyin-simple/app/interaction/rpc/internal/logic"
	"douyin-simple/app/interaction/rpc/internal/svc"
	"douyin-simple/app/interaction/rpc/pb/pb"
)

type InteractionServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedInteractionServer
}

func NewInteractionServer(svcCtx *svc.ServiceContext) *InteractionServer {
	return &InteractionServer{
		svcCtx: svcCtx,
	}
}

func (s *InteractionServer) FavoriteAction(ctx context.Context, in *pb.FavoriteActionReq) (*pb.Empty, error) {
	l := logic.NewFavoriteActionLogic(ctx, s.svcCtx)
	return l.FavoriteAction(in)
}

func (s *InteractionServer) GetFavoriteList(ctx context.Context, in *pb.GetFavoriteListReq) (*pb.GetFavoriteListReply, error) {
	l := logic.NewGetFavoriteListLogic(ctx, s.svcCtx)
	return l.GetFavoriteList(in)
}

func (s *InteractionServer) CommentAction(ctx context.Context, in *pb.CommentActionReq) (*pb.CommentActionReply, error) {
	l := logic.NewCommentActionLogic(ctx, s.svcCtx)
	return l.CommentAction(in)
}

func (s *InteractionServer) GetCommentList(ctx context.Context, in *pb.GetCommentListReq) (*pb.GetCommentListReply, error) {
	l := logic.NewGetCommentListLogic(ctx, s.svcCtx)
	return l.GetCommentList(in)
}
