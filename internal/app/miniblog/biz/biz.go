// // Copyright 2023 Innkeeper Belm(郁凯) <yukai98@foxmain.com>. All rights reserved.
// // Use of this source code is governed by a MIT style
// // license that can be found in the LICENSE file. The original repo for
// // this file is https://github.com/ProgramKai/miniblog

package biz

import "cn.xdmnb/study/miniblog/internal/app/miniblog/store"

type IBiz interface {
	UserBiz() IUserBiz
}

var _ IBiz = (*biz)(nil)

type biz struct {
	ds store.IStore
}

var _ IUserBiz = (*UserBizImpl)(nil)

func (b biz) UserBiz() IUserBiz {
	return newUserBiz(b.ds)
}

func NewBiz(ds store.IStore) IBiz {
	return &biz{ds: ds}
}
