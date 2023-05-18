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
