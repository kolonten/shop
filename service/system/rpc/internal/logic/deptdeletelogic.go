package logic

import (
	"context"

	"gogoshop/service/system/rpc/internal/svc"
	"gogoshop/service/system/rpc/systemclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptDeleteLogic {
	return &DeptDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptDeleteLogic) DeptDelete(in *systemclient.DeptDeleteReq) (*systemclient.DeptDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &systemclient.DeptDeleteResp{}, nil
}
