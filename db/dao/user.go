package dao

import (
	"errors"
	"gorm.io/gorm"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

type UserDao struct {
	Dao[model.User, uint64]
}

func (*UserDao) GetByOpenId(openId string) (*model.User, error) {
	var err error
	var m = new(model.User)
	err = db.Get().Where("open_id = ?", openId).First(m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return m, err
}
