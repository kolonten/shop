package user

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"gogoshop/common/errorx"
	"gogoshop/service/system/adminapi/internal/svc"
	"gogoshop/service/system/adminapi/internal/types"
	"gogoshop/service/system/rpc/systemclient"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq, ip string) (*types.LoginResp, error) {
	if len(strings.TrimSpace(req.UserName)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		reqStr, _ := json.Marshal(req)
		return nil, errors.Wrapf(errorx.NewErrCode(errorx.REQUEST_ERROR), "用户名或密码不能为空 参数: %s", reqStr)
	}
	resp, err := l.svcCtx.System.Login(l.ctx, &systemclient.LoginReq{
		SiteId:   req.SiteId,
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	//保存登录日志
	_, _ = l.svcCtx.System.LoginLogAdd(l.ctx, &systemclient.LoginLogAddReq{
		UserName: resp.UserName,
		Status:   1,
		Ip:       ip,
		CreateBy: resp.UserName,
	})
	return &types.LoginResp{
		SiteId:       resp.SiteId,
		Authority:    resp.Authority,
		Id:           resp.Id,
		UserName:     resp.UserName,
		AccessToken:  resp.AccessToken,
		AccessExpire: resp.AccessExpire,
		RefreshAfter: resp.RefreshAfter,
	}, nil
}
