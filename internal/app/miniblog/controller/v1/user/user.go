// // Copyright 2023 Innkeeper Belm(郁凯) <yukai98@foxmain.com>. All rights reserved.
// // Use of this source code is governed by a MIT style
// // license that can be found in the LICENSE file. The original repo for
// // this file is https://github.com/ProgramKai/miniblog

package user

import (
	"cn.xdmnb/study/miniblog/internal/app/miniblog/biz"
	"cn.xdmnb/study/miniblog/internal/app/miniblog/store"
)

type UserController struct {
	b biz.IBiz
}

func NewUserController(ds store.IStore) *UserController {
	return &UserController{b: biz.NewBiz(ds)}
}
