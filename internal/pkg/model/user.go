// // Copyright 2023 Innkeeper Belm(郁凯) <yukai98@foxmain.com>. All rights reserved.
// // Use of this source code is governed by a MIT style
// // license that can be found in the LICENSE file. The original repo for
// // this file is https://github.com/ProgramKai/miniblog

package model

import (
	"cn.xdmnb/study/miniblog/pkg/auth"
	"time"
)

type UserM struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Username  string    `gorm:"column:username;not null"`
	Password  string    `gorm:"column:password;not null"`
	Nickname  string    `gorm:"column:nickname"`
	Email     string    `gorm:"column:email"`
	Phone     string    `gorm:"column:phone"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}

func (UserM) TableName() string {
	return "user"
}

func (u *UserM) BeforeCreate() (err error) {
	// Encrypt the user password.
	u.Password, err = auth.Encrypt(u.Password)
	if err != nil {
		return err
	}

	return nil
}
