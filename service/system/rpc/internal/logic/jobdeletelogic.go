package logic

import (
	"context"

	"gogoshop/service/system/rpc/internal/svc"
	"gogoshop/service/system/rpc/systemclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type JobDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJobDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobDeleteLogic {
	return &JobDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JobDeleteLogic) JobDelete(in *systemclient.JobDeleteReq) (*systemclient.JobDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &systemclient.JobDeleteResp{}, nil
}
