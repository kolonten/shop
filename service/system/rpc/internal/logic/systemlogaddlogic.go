package logic

import (
	"context"

	"gogoshop/service/system/rpc/internal/svc"
	"gogoshop/service/system/rpc/systemclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SystemLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSystemLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemLogAddLogic {
	return &SystemLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SystemLogAddLogic) SystemLogAdd(in *systemclient.SystemLogAddReq) (*systemclient.SystemLogAddResp, error) {
	// todo: add your logic here and delete this line

	return &systemclient.SystemLogAddResp{}, nil
}
