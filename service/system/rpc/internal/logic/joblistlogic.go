package logic

import (
	"context"

	"gogoshop/service/system/rpc/internal/svc"
	"gogoshop/service/system/rpc/systemclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type JobListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJobListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobListLogic {
	return &JobListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JobListLogic) JobList(in *systemclient.JobListReq) (*systemclient.JobListResp, error) {
	// todo: add your logic here and delete this line

	return &systemclient.JobListResp{}, nil
}
