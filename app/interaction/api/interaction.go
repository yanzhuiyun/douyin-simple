package main

import (
	"douyin-simple/common/middleware"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"

	"douyin-simple/app/interaction/api/internal/config"
	"douyin-simple/app/interaction/api/internal/handler"
	"douyin-simple/app/interaction/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/interaction-api.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	server := rest.MustNewServer(c.RestConf)
	server.Use(middleware.RecoverFromPanic)
	defer server.Stop()
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	logx.DisableStat()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
