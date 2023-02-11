package interaction

import (
	"douyin-simple/common/result"
	"net/http"

	"douyin-simple/app/interaction/api/internal/logic/interaction"
	"douyin-simple/app/interaction/api/internal/svc"
	"douyin-simple/app/interaction/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavoriteActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteActionReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ResponseBadRequest(w, err)
			return
		}

		l := interaction.NewFavoriteActionLogic(r.Context(), svcCtx)
		err := l.FavoriteAction(&req)
		if err != nil {
			result.ResponseBadRequest(w, err)
		} else {
			result.ResponseSuccess(w, nil, "")
		}
	}
}
