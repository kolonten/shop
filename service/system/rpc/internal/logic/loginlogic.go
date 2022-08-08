package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"gogoshop/common/errorx"
	"gogoshop/common/jwtx"
	"gogoshop/service/system/rpc/system"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gogoshop/service/system/rpc/internal/svc"
	"gogoshop/service/system/rpc/systemclient"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *systemclient.LoginReq) (*systemclient.LoginResp, error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByName(l.ctx, in.UserName)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		return nil, errors.Wrapf(errorx.NewErrMsg("用户不存在"), "用户不存在,参数:%s,异常:%s", in, err)
	default:
		return nil, errors.Wrapf(errorx.NewErrMsg("登录失败"), "用户登录失败,参数:%s,异常:%s", in, err)
	}

	if userInfo.Password != in.Password {
		return nil, errors.Wrapf(errorx.NewErrCode(errorx.SERVER_COMMON_ERROR), "用户密码不正确,参数:%s , err: %v", in, err)
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	jwtToken, err := jwtx.GetToken(l.svcCtx.Config.JWT.AccessSecret, now, l.svcCtx.Config.JWT.AccessExpire, userInfo.Id)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("生成token失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	resp := &system.LoginResp{
		Id:           userInfo.Id,
		SiteId:       userInfo.SiteId,
		UserName:     userInfo.Name,
		Authority:    "admin",
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(resp)
	logx.WithContext(l.ctx).Infof("登录成功,参数:%s,响应:%s", reqStr, listStr)
	return resp, nil
}
