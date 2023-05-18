// // Copyright 2023 Innkeeper Belm(郁凯) <yukai98@foxmain.com>. All rights reserved.
// // Use of this source code is governed by a MIT style
// // license that can be found in the LICENSE file. The original repo for
// // this file is https://github.com/ProgramKai/miniblog

package user

import (
	v1 "cn.xdmnb/study/miniblog/internal/pkg/domain/v1"
	"cn.xdmnb/study/miniblog/internal/pkg/errno"
	"cn.xdmnb/study/miniblog/internal/pkg/log"
	"cn.xdmnb/study/miniblog/internal/pkg/response"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

const defaultMethods = "(GET)|(POST)|(PUT)|(DELETE)"

func (ctrl *UserController) CreateUser(c *gin.Context) {
	log.C(c).Infow("Create user function called")
	var reqBody v1.CreateUserReq
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	if _, err := govalidator.ValidateStruct(reqBody); err != nil {
		response.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	if err := ctrl.b.UserBiz().Create(c, &reqBody); err != nil {
		response.WriteResponse(c, err, nil)
		return
	}
	if _, err := ctrl.authz.AddNamedPolicy("p", reqBody.Username, "/api/v1/user/by/"+reqBody.Username, defaultMethods); err != nil {
		response.WriteResponse(c, err, nil)
		return
	}
	response.WriteResponse(c, nil, nil)
}
