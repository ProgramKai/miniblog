// // Copyright 2023 Innkeeper Belm(郁凯) <yukai98@foxmain.com>. All rights reserved.
// // Use of this source code is governed by a MIT style
// // license that can be found in the LICENSE file. The original repo for
// // this file is https://github.com/ProgramKai/miniblog

package biz

import (
	"cn.xdmnb/study/miniblog/internal/app/miniblog/store"
	"cn.xdmnb/study/miniblog/internal/pkg/errno"
	"cn.xdmnb/study/miniblog/internal/pkg/model"
	v1 "cn.xdmnb/study/miniblog/internal/pkg/request_body/v1"
	"context"
	"regexp"
)

type IUserBiz interface {
	Create(ctx context.Context, createUserReq *v1.CreateUserReq) error
}

type UserBizImpl struct {
	ds store.IStore
}

var _ IUserBiz = (*UserBizImpl)(nil)

func newUserBiz(ds store.IStore) *UserBizImpl {
	return &UserBizImpl{ds: ds}
}

// Create 创建用户
func (u *UserBizImpl) Create(ctx context.Context, createUserReq *v1.CreateUserReq) error {
	user := u.createUserReqToUserM(createUserReq)
	if err := u.ds.Users().Create(ctx, user); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry .* for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExists
		}
		return err
	}
	return nil
}

// createUserReqToUserM 将 CreateUserReq 转换为 UserM
func (u *UserBizImpl) createUserReqToUserM(createUserReq *v1.CreateUserReq) *model.UserM {
	return &model.UserM{
		Username: createUserReq.Username,
		Password: createUserReq.Password,
		Nickname: createUserReq.Nickname,
		Email:    createUserReq.Email,
		Phone:    createUserReq.Phone,
	}
}
