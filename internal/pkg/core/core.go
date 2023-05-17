// // Copyright 2023 Innkeeper Belm(郁凯) <yukai98@foxmain.com>. All rights reserved.
// // Use of this source code is governed by a MIT style
// // license that can be found in the LICENSE file. The original repo for
// // this file is https://github.com/ProgramKai/miniblog

package core

import (
	"cn.xdmnb/study/miniblog/internal/pkg/errno"
	"github.com/gin-gonic/gin"
)

type ErrResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		code, codeMsg, msg := errno.Decode(err)
		c.JSON(code, ErrResponse{
			Code:    codeMsg,
			Message: msg,
		})
		return
	}
	c.JSON(200, data)
}
