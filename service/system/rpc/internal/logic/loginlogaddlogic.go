package logic

import (
	"context"
	"gogoshop/service/system/model"
	"time"

	"gogoshop/service/system/rpc/internal/svc"
	"gogoshop/service/system/rpc/systemclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogAddLogic {
	return &LoginLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogAddLogic) LoginLogAdd(in *systemclient.LoginLogAddReq) (*systemclient.LoginLogAddResp, error) {
	_, err := l.svcCtx.LoginLogModel.Insert(l.ctx, &model.SystemLoginLog{
		UserName:   in.UserName,
		Ip:         in.Ip,
		Status:     in.Status,
		CreateBy:   in.CreateBy,
		UpdateBy:   in.CreateBy,
		CreateTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &systemclient.LoginLogAddResp{}, nil
}
