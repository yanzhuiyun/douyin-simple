package svc

import (
	"example/model/mexample"
	"example/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	Mexample mexample.ExampleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Mexample: mexample.NewExampleModel(sqlx.NewMysql(c.Mysql.DataSource), nil),
	}
}
