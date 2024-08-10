package handler

import (
	"net/http"

	"gadget-crafted/internal/logic"
	"gadget-crafted/internal/svc"
	"gadget-crafted/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GenerateCreateTableSqlHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GenerateCreateTableSqlReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGenerateCreateTableSqlLogic(r.Context(), svcCtx)
		resp, err := l.GenerateCreateTableSql(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
