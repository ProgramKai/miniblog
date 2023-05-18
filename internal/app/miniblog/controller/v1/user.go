package v1

import (
	"cn.xdmnb/study/miniblog/internal/app/miniblog/biz"
	"cn.xdmnb/study/miniblog/internal/app/miniblog/store"
	"cn.xdmnb/study/miniblog/internal/pkg/core"
	"cn.xdmnb/study/miniblog/internal/pkg/errno"
	"cn.xdmnb/study/miniblog/internal/pkg/log"
	v1 "cn.xdmnb/study/miniblog/internal/pkg/request_body/v1"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	b biz.IBiz
}

func NewUserController(ds store.IStore) *UserController {
	return &UserController{b: biz.NewBiz(ds)}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	log.C(c).Infow("Create user function called")
	var reqBody v1.CreateUserReq
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	if err := ctrl.b.UserBiz().Create(c, &reqBody); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}
	core.WriteResponse(c, nil, nil)
}
