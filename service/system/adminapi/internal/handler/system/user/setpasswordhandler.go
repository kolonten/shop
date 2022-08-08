package user

import (
	"net/http"

	"gogoshop/service/system/adminapi/internal/logic/system/user"
	"gogoshop/service/system/adminapi/internal/svc"
	"gogoshop/service/system/adminapi/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SetPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetPasswordReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewSetPasswordLogic(r.Context(), svcCtx)
		resp, err := l.SetPassword(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
