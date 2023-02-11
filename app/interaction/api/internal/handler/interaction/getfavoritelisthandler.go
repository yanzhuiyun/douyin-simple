package interaction

import (
	"douyin-simple/common/result"
	"net/http"

	"douyin-simple/app/interaction/api/internal/logic/interaction"
	"douyin-simple/app/interaction/api/internal/svc"
	"douyin-simple/app/interaction/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetfavoriteListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFavoriteListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ResponseBadRequest(w, err)
			return
		}

		l := interaction.NewGetfavoriteListLogic(r.Context(), svcCtx)
		resp, err := l.GetfavoriteList(&req)
		if err != nil {
			result.ResponseBadRequest(w, err)
		} else {
			tagName := result.GetTagName(resp)
			result.ResponseSuccess(w, resp.FavoriteList, tagName)
		}
	}
}
