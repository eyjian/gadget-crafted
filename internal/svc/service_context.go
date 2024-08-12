package svc

import (
	"gadget-crafted/internal/config"
	"gadget-crafted/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config         config.Config
	CorsMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CorsMiddleware: middleware.NewCorsMiddleware().Handle,
	}
}
