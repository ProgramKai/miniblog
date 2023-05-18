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
