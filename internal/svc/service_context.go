package svc

import (
	"gadget-crafted/internal/config"
	"gadget-crafted/internal/middleware"
	"gadget-crafted/internal/model"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config         config.Config
	CorsMiddleware rest.Middleware
	Db             *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		CorsMiddleware: middleware.NewCorsMiddleware().Handle,
	}
}

func (s *ServiceContext) Init() bool {
	dbName := "gadget_crafted.db"
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		logx.Errorf("open sqlite://%s error: %s", dbName, err.Error())
		return false
	}
	db.AutoMigrate(&model.FeatureUsage{}) // 添加 FeatureUsage 结构体
	s.Db = db
	return true
}
