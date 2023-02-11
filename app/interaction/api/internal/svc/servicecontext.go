package svc

import (
	"douyin-simple/app/interaction/api/internal/config"
	"douyin-simple/app/interaction/rpc/interaction"
)

type ServiceContext struct {
	Config         config.Config
	InteractionRpc interaction.Interaction
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//InteractionRpc: interaction.NewInteraction(zrpc.MustNewClient(c.InterActionRpc)),
	}
}
