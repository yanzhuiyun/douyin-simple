package logic

import (
	"context"

	"example/rpc/internal/svc"
	"example/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetContentByUsernameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetContentByUsernameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetContentByUsernameLogic {
	return &GetContentByUsernameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetContentByUsernameLogic) GetContentByUsername(in *pb.ExampleReq) (*pb.ExampleReply, error) {
	// todo: add your logic here and delete this line

	return &pb.ExampleReply{}, nil
}
