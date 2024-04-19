package dao

import (
	"errors"
	"gorm.io/gorm"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

type FinanceLogDao struct {
	Dao[model.FinanceLog, uint64]
}

func (*FinanceLogDao) GetByMonth(month string) (*model.FinanceLog, error) {
	var err error
	var m = new(model.FinanceLog)
	err = db.Get().
		Where("month = ?", month).
		First(m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return m, err
}
