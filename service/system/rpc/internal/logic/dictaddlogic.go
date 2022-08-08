package logic

import (
	"context"

	"gogoshop/service/system/rpc/internal/svc"
	"gogoshop/service/system/rpc/systemclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DictAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDictAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DictAddLogic {
	return &DictAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DictAddLogic) DictAdd(in *systemclient.DictAddReq) (*systemclient.DictAddResp, error) {
	// todo: add your logic here and delete this line

	return &systemclient.DictAddResp{}, nil
}
