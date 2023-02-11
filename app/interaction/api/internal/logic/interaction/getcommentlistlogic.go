package interaction

import (
	"context"

	"douyin-simple/app/interaction/api/internal/svc"
	"douyin-simple/app/interaction/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetcommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetcommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetcommentListLogic {
	return &GetcommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetcommentListLogic) GetcommentList(req *types.GetCommentListReq) (resp *types.CommentListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
