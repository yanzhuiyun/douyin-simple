package svc

import (
	"example/api/internal/config"
	"example/model/mexample"
	"example/rpc/example"
)

type ServiceContext struct {
	Config     config.Config
	mexample   mexample.ExampleModel
	ExampleRpc example.Example
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
