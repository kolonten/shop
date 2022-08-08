package user

import (
	"net/http"

	"gogoshop/service/system/adminapi/internal/logic/system/user"
	"gogoshop/service/system/adminapi/internal/svc"
	"gogoshop/service/system/adminapi/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAllDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetDataReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewGetAllDataLogic(r.Context(), svcCtx)
		resp, err := l.GetAllData(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
