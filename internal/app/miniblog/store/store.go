// // Copyright 2023 Innkeeper Belm(郁凯) <yukai98@foxmain.com>. All rights reserved.
// // Use of this source code is governed by a MIT style
// // license that can be found in the LICENSE file. The original repo for
// // this file is https://github.com/ProgramKai/miniblog

package store

import (
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once

	S *datasource
)

type IStore interface {
	Users() IUserStore
}

type datasource struct {
	db *gorm.DB
}

var _ IStore = (*datasource)(nil)

func NewStore(db *gorm.DB) *datasource {
	once.Do(func() {
		S = &datasource{db: db}
	})
	return S
}

func (d *datasource) Users() IUserStore {
	return newUserStore(d.db)
}
