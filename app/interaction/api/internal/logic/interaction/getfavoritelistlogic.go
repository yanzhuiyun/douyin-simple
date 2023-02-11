package interaction

import (
	"context"

	"douyin-simple/app/interaction/api/internal/svc"
	"douyin-simple/app/interaction/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetfavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetfavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetfavoriteListLogic {
	return &GetfavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetfavoriteListLogic) GetfavoriteList(req *types.GetFavoriteListReq) (resp *types.FavoriteListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
