// // Copyright 2023 Innkeeper Belm(郁凯) <yukai98@foxmain.com>. All rights reserved.
// // Use of this source code is governed by a MIT style
// // license that can be found in the LICENSE file. The original repo for
// // this file is https://github.com/ProgramKai/miniblog

package store

import (
	"cn.xdmnb/study/miniblog/internal/pkg/model"
	"context"
	"gorm.io/gorm"
)

type IUserStore interface {
	Create(ctx context.Context, user *model.UserM) error
}

type UserStoreImpl struct {
	db *gorm.DB
}

var _ IUserStore = (*UserStoreImpl)(nil)

func newUserStore(db *gorm.DB) *UserStoreImpl {
	return &UserStoreImpl{db: db}
}

func (u *UserStoreImpl) Create(ctx context.Context, user *model.UserM) error {
	return u.db.WithContext(ctx).Create(user).Error
}
