package user

import (
	"net/http"

	"gogoshop/common/result"
	"gogoshop/service/system/adminapi/internal/logic/system/user"
	"gogoshop/service/system/adminapi/internal/svc"
	"gogoshop/service/system/adminapi/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}
		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req, r.RemoteAddr)
		result.HttpResult(r, w, resp, err)
	}
}
