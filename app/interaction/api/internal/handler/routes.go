// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"douyin-simple/app/interaction/api/internal/handler/interaction"
	"douyin-simple/app/interaction/api/internal/svc"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/favorite/action",
				Handler: interaction.FavoriteActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/favorite/list",
				Handler: interaction.GetfavoriteListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/comment/action",
				Handler: interaction.CommentActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/comment/list",
				Handler: interaction.GetcommentListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin"),
	)
}
