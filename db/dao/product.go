package dao

import (
	"time"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

type ProductDao struct{}

var ProductDaoImpl = &ProductDao{}

func (imp *ProductDao) Delete(id int64) error {
	cli := db.Get()
	return cli.Delete(&model.Product{Id: uint64(id)}).Error
}

func (imp *ProductDao) Upsert(m *model.Product) error {
	now := time.Now()
	m.UpdateTime = &now
	cli := db.Get()
	return cli.Save(m).Error
}

func (imp *ProductDao) Get(id int64) (*model.Product, error) {
	var err error
	var m = new(model.Product)

	cli := db.Get()
	err = cli.Where("id = ?", id).First(m).Error

	return m, err
}
