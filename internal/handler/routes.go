// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"gadget-crafted/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CorsMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/generate_create_table_sql",
					Handler: GenerateCreateTableSqlHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/sql2struct",
					Handler: Sql2StructHandler(serverCtx),
				},
			}...,
		),
	)
}
