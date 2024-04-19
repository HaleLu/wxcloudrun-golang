package dao

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

var (
	MissionDaoImpl    = &Dao[model.Mission, uint64]{}
	ProductDaoImpl    = &Dao[model.Product, uint64]{}
	UserDaoImpl       = &UserDao{}
	OpLogDaoImpl      = &OpLogDao{}
	FinanceLogDaoImpl = &FinanceLogDao{}
)

type Dao[T model.DbModel[TID], TID uint64 | int64 | string] struct{}

func (imp *Dao[T, TID]) Get(id TID) (*T, error) {
	var err error
	var m = new(T)
	err = db.Get().Where("id = ?", id).First(m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return m, err
}

func (imp *Dao[T, TID]) Delete(id TID) error {
	cli := db.Get()
	var m T
	m.SetId(id)
	return cli.Delete(&m).Error
}

func (imp *Dao[T, TID]) Upsert(m *T) error {
	if m == nil {
		return nil
	}
	(*m).SetUpdateTime(time.Now())
	return db.Get().Save(m).Error
}

func (imp *Dao[T, TID]) List() ([]*T, error) {
	var err error
	var list []*T
	err = db.Get().Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return list, err
}
