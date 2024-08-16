package handler

import (
	"net/http"

	"gadget-crafted/internal/logic"
	"gadget-crafted/internal/svc"
	"gadget-crafted/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Sql2StructHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Sql2StructReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		userIp := GetUserIp(r)
		l := logic.NewSql2StructLogic(r.Context(), svcCtx)
		resp, err := l.Sql2Struct(userIp, &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
