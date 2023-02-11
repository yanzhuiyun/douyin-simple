package interaction

import (
	"douyin-simple/app/interaction/api/internal/logic/interaction"
	"douyin-simple/app/interaction/api/internal/svc"
	"douyin-simple/app/interaction/api/internal/types"
	"douyin-simple/common/result"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func GetcommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCommentListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ResponseBadRequest(w, err)
			return
		}

		l := interaction.NewGetcommentListLogic(r.Context(), svcCtx)
		resp, err := l.GetcommentList(&req)
		if err != nil {
			result.ResponseBadRequest(w, err)
		} else {
			tagName := result.GetTagName(resp)
			result.ResponseSuccess(w, resp.CommentList, tagName)
		}
	}
}
