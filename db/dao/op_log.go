package dao

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/util"
)

type OpLogDao struct {
	Dao[model.OpLog, uint64]
}

func (*OpLogDao) GetByDate(openId string, opType model.OpType, objectId uint64, date time.Time) ([]*model.OpLog, error) {
	var err error
	st := util.GetStartOfDate(date)
	ed := st.AddDate(0, 0, 1)
	var list []*model.OpLog
	err = db.Get().
		Where("type = ?", opType).
		Where("open_id = ?", openId).
		Where("object_id = ?", objectId).
		Where("finish_time >= ?", st).
		Where("finish_time < ?", ed).
		Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return list, err
}
