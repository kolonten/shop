package logic

import (
	"context"

	"gogoshop/service/system/rpc/internal/svc"
	"gogoshop/service/system/rpc/systemclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConfigAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConfigAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConfigAddLogic {
	return &ConfigAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConfigAddLogic) ConfigAdd(in *systemclient.ConfigAddReq) (*systemclient.ConfigAddResp, error) {
	// todo: add your logic here and delete this line

	return &systemclient.ConfigAddResp{}, nil
}
