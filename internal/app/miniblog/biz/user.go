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
	"cn.xdmnb/study/miniblog/internal/pkg/token"
	"cn.xdmnb/study/miniblog/pkg/auth"
	"context"
	"regexp"
	"sync"
)

var (
	userBiz IUserBiz
	once    sync.Once
)

type IUserBiz interface {
	Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error)
	Create(ctx context.Context, createUserReq *v1.CreateUserReq) error
	ChangePassword(ctx context.Context, username string, changePasswordReq *v1.ChangePasswordReq) error
}

type UserBizImpl struct {
	ds store.IStore
}

var _ IUserBiz = (*UserBizImpl)(nil)

func newUserBiz(ds store.IStore) IUserBiz {
	once.Do(func() {
		userBiz = &UserBizImpl{ds: ds}
	})
	return userBiz
}

// Create 创建用户
func (u *UserBizImpl) Create(ctx context.Context, createUserReq *v1.CreateUserReq) error {
	user := u.createUserReqToUserM(createUserReq)
	if err := u.ds.Users().Create(ctx, user); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry .* for key 'user.username'", err.Error()); match {
			return errno.ErrUserAlreadyExists
		}
		return err
	}
	return nil
}

// ChangePassword 修改密码
func (u *UserBizImpl) ChangePassword(ctx context.Context, username string, changePasswordReq *v1.ChangePasswordReq) error {
	user, err := u.ds.Users().Get(ctx, username)
	if err != nil {
		return err
	}
	if err := auth.Compare(user.Password, changePasswordReq.OldPassword); err != nil {
		return errno.ErrPasswordIncorrect
	}
	newPassword, _ := auth.Encrypt(changePasswordReq.NewPassword)
	changePasswordReq.NewPassword = newPassword
	return u.ds.Users().Update(ctx, user)
}

// Login 登录
func (u *UserBizImpl) Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error) {
	// 获取登录用户的所有信息
	user, err := u.ds.Users().Get(ctx, r.Username)
	if err != nil {
		return nil, errno.ErrUserNotFound
	}

	// 对比传入的明文密码和数据库中已加密过的密码是否匹配
	if err := auth.Compare(user.Password, r.Password); err != nil {
		return nil, errno.ErrPasswordIncorrect
	}

	// 如果匹配成功，说明登录成功，签发 token 并返回
	t, err := token.Sign(r.Username)
	if err != nil {
		return nil, errno.ErrSignToken
	}

	return &v1.LoginResponse{Token: t}, nil
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
