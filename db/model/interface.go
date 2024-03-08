package model

import "time"

type DbModel[T uint64 | int64 | string] interface {
	GetId() T
	SetId(T)
	TableName() string
	SetUpdateTime(time time.Time)
}
