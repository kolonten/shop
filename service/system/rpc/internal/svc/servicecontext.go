package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gogoshop/service/system/model"
	"gogoshop/service/system/rpc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.SystemUserModel
	LoginLogModel model.SystemLoginLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewSystemUserModel(sqlConn, c.CacheRedis),
		LoginLogModel: model.NewSystemLoginLogModel(sqlConn, c.CacheRedis),
	}
}
