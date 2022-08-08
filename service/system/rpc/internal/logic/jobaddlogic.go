package logic

import (
	"context"

	"gogoshop/service/system/rpc/internal/svc"
	"gogoshop/service/system/rpc/systemclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type JobAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJobAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobAddLogic {
	return &JobAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JobAddLogic) JobAdd(in *systemclient.JobAddReq) (*systemclient.JobAddResp, error) {
	// todo: add your logic here and delete this line

	return &systemclient.JobAddResp{}, nil
}
