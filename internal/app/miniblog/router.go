// // Copyright 2023 Innkeeper Belm(郁凯) <yukai98@foxmain.com>. All rights reserved.
// // Use of this source code is governed by a MIT style
// // license that can be found in the LICENSE file. The original repo for
// // this file is https://github.com/ProgramKai/miniblog

package miniblog

import (
	v1 "cn.xdmnb/study/miniblog/internal/app/miniblog/controller/v1/user"
	"cn.xdmnb/study/miniblog/internal/app/miniblog/store"
	"cn.xdmnb/study/miniblog/internal/pkg/authz"
	"cn.xdmnb/study/miniblog/internal/pkg/errno"
	"cn.xdmnb/study/miniblog/internal/pkg/log"
	"cn.xdmnb/study/miniblog/internal/pkg/middleware"
	"cn.xdmnb/study/miniblog/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

// installRouters 注册路由.
func installRouters(g *gin.Engine) error {
	authz, err := authz.NewAuthz(store.S.DB)
	if err != nil {
		return err
	}

	middlewares := []gin.HandlerFunc{
		middleware.RequestLog(),
		gin.Recovery(),
		middleware.NoCache,
		middleware.Secure,
		middleware.RequestID(),
		middleware.Authn(),
		middleware.Authz(authz),
	}
	g.Use(middlewares...)

	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		response.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		response.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})
	v1Router := g.Group("/api/v1")
	{
		userController := v1.NewUserController(store.S, authz)
		authRouter := v1Router.Group("/auth")
		{
			authRouter.POST("/register", userController.CreateUser)
			authRouter.POST("/login", userController.Login)
		}

		userRouter := v1Router.Group("/user")
		{
			userRouter.PUT(":name/change-pwd", userController.ChangePassword)
			userRouter.GET("/info/by/:name", userController.QueryByUserName)
		}

	}
	return nil
}
