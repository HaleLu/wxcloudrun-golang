package service

import (
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

type SimpleService[T model.DbModel[TID], TID uint64 | int64 | string] struct {
	dao.Dao[T, TID]
}
