// // Copyright 2023 Innkeeper Belm(郁凯) <yukai98@foxmain.com>. All rights reserved.
// // Use of this source code is governed by a MIT style
// // license that can be found in the LICENSE file. The original repo for
// // this file is https://github.com/ProgramKai/miniblog

package user

import (
	"cn.xdmnb/study/miniblog/internal/pkg/errno"
	"cn.xdmnb/study/miniblog/internal/pkg/log"
	v1 "cn.xdmnb/study/miniblog/internal/pkg/request_body/v1"
	"cn.xdmnb/study/miniblog/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) Login(c *gin.Context) {
	log.C(c).Infow("Login function called")
	var reqBody v1.LoginRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	resp, err := ctrl.b.UserBiz().Login(c, &reqBody)
	if err != nil {
		response.WriteResponse(c, err, nil)
		return
	}
	response.WriteResponse(c, nil, resp)
}
