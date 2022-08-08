package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gogoshop/service/system/adminapi/internal/logic/system/user"
	"gogoshop/service/system/adminapi/internal/svc"
)

func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
