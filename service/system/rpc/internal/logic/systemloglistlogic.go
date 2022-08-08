package logic

import (
	"context"

	"gogoshop/service/system/rpc/internal/svc"
	"gogoshop/service/system/rpc/systemclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SystemLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSystemLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemLogListLogic {
	return &SystemLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SystemLogListLogic) SystemLogList(in *systemclient.SystemLogListReq) (*systemclient.SystemLogListResp, error) {
	// todo: add your logic here and delete this line

	return &systemclient.SystemLogListResp{}, nil
}
