package logic

import (
	"context"

	"douyin-simple/app/interaction/rpc/internal/svc"
	"douyin-simple/app/interaction/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentActionLogic) CommentAction(in *pb.CommentActionReq) (*pb.CommentActionReply, error) {
	// todo: add your logic here and delete this line

	return &pb.CommentActionReply{}, nil
}
