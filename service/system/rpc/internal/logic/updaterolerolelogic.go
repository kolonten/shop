package logic

import (
	"context"

	"gogoshop/service/system/rpc/internal/svc"
	"gogoshop/service/system/rpc/systemclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleRoleLogic {
	return &UpdateRoleRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleRoleLogic) UpdateRoleRole(in *systemclient.UpdateRoleRoleReq) (*systemclient.UpdateRoleRoleResp, error) {
	// todo: add your logic here and delete this line

	return &systemclient.UpdateRoleRoleResp{}, nil
}
