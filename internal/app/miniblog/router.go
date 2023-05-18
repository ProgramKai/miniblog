package miniblog

import (
	v1 "cn.xdmnb/study/miniblog/internal/app/miniblog/controller/v1"
	"cn.xdmnb/study/miniblog/internal/app/miniblog/store"
	"cn.xdmnb/study/miniblog/internal/pkg/core"
	"cn.xdmnb/study/miniblog/internal/pkg/errno"
	"cn.xdmnb/study/miniblog/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

// installRouters 注册路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})
	userController := v1.NewUserController(store.S)

	v1 := g.Group("/v1")
	{
		userRouter := v1.Group("/user")
		{
			userRouter.POST("/", userController.CreateUser)
		}

	}
	return nil
}
