package logic

import (
	"context"

	"douyin-simple/app/interaction/rpc/internal/svc"
	"douyin-simple/app/interaction/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteListLogic {
	return &GetFavoriteListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFavoriteListLogic) GetFavoriteList(in *pb.GetFavoriteListReq) (*pb.GetFavoriteListReply, error) {
	// todo: add your logic here and delete this line

	return &pb.GetFavoriteListReply{}, nil
}
