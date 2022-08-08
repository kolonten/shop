package logic

import (
	"context"

	"gogoshop/service/system/rpc/internal/svc"
	"gogoshop/service/system/rpc/systemclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SystemLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSystemLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemLogDeleteLogic {
	return &SystemLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SystemLogDeleteLogic) SystemLogDelete(in *systemclient.SystemLogDeleteReq) (*systemclient.SystemLogDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &systemclient.SystemLogDeleteResp{}, nil
}
