// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	systemuser "gogoshop/service/system/adminapi/internal/handler/system/user"
	"gogoshop/service/system/adminapi/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.CheckAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/currentUser",
					Handler: systemuser.InfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/add",
					Handler: systemuser.AddHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/list",
					Handler: systemuser.ListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: systemuser.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: systemuser.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/setPassword",
					Handler: systemuser.SetPasswordHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/UpdateUserStatus",
					Handler: systemuser.UpdateUserStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/getAllData",
					Handler: systemuser.GetAllDataHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/adminapi/system/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/adminapi/system/user/login",
				Handler: systemuser.LoginHandler(serverCtx),
			},
		},
	)
}
