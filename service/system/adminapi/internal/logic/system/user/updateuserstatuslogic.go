package user

import (
	"context"

	"gogoshop/service/system/adminapi/internal/svc"
	"gogoshop/service/system/adminapi/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserStatusLogic {
	return &UpdateUserStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserStatusLogic) UpdateUserStatus(req *types.UserStatusReq) (resp *types.UserStatusResp, err error) {
	// todo: add your logic here and delete this line

	return
}
