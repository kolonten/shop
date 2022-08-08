package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"gogoshop/service/system/adminapi/internal/config"
	"gogoshop/service/system/adminapi/internal/middleware"
	"gogoshop/service/system/rpc/system"
)

type ServiceContext struct {
	Config    config.Config
	CheckAuth rest.Middleware
	System    system.System
	Redis     *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.New(c.Redis.Address, redisConfig(c))
	return &ServiceContext{
		Config:    c,
		CheckAuth: middleware.NewCheckAuthMiddleware().Handle,
		System:    system.NewSystem(zrpc.MustNewClient(c.SystemRpc)),
		Redis:     newRedis,
	}
}
func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
