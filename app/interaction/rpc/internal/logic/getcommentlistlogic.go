package logic

import (
	"context"

	"douyin-simple/app/interaction/rpc/internal/svc"
	"douyin-simple/app/interaction/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentListLogic) GetCommentList(in *pb.GetCommentListReq) (*pb.GetCommentListReply, error) {
	// todo: add your logic here and delete this line

	return &pb.GetCommentListReply{}, nil
}
