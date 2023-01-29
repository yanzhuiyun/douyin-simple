package example

import (
	"net/http"

	"example/api/internal/logic/example"
	"example/api/internal/svc"
	"example/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ExampleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExampleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := example.NewExampleLogic(r.Context(), svcCtx)
		resp, err := l.Example(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
