package logic

import (
	"context"

	"gogoshop/service/system/rpc/internal/svc"
	"gogoshop/service/system/rpc/systemclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuByRoleIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMenuByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuByRoleIdLogic {
	return &QueryMenuByRoleIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryMenuByRoleIdLogic) QueryMenuByRoleId(in *systemclient.QueryMenuByRoleIdReq) (*systemclient.QueryMenuByRoleIdResp, error) {
	// todo: add your logic here and delete this line

	return &systemclient.QueryMenuByRoleIdResp{}, nil
}
