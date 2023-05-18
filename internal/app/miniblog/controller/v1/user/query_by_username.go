package user

import (
	"cn.xdmnb/study/miniblog/internal/pkg/errno"
	"cn.xdmnb/study/miniblog/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

func (u *UserController) QueryByUserName(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		response.WriteResponse(c, errno.ErrInvalidParameter, nil)
		return
	}
	userInfo, err := u.b.UserBiz().QueryByUsername(c, username)
	if err != nil {
		response.WriteResponse(c, err, nil)
		return
	}
	response.WriteResponse(c, nil, userInfo)
}
