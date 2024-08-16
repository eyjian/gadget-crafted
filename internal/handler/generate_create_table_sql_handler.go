package handler

import (
	"net/http"
	"strings"

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

		userIp := GetUserIp(r)
		l := logic.NewGenerateCreateTableSqlLogic(r.Context(), svcCtx)
		resp, err := l.GenerateCreateTableSql(userIp, &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

func GetUserIp(r *http.Request) string {
	realIP := r.Header.Get("X-Real-IP")
	if realIP == "" {
		realIP = r.Header.Get("X-Forwarded-For")
		if realIP != "" {
			// 如果有多个代理，取第一个IP地址
			realIP = strings.Split(realIP, ",")[0]
		}
	}
	if realIP == "" {
		realIP = r.RemoteAddr
	}

	return realIP
}
