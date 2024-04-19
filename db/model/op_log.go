package model

import (
	"gorm.io/gorm"
	"time"
)

var _ DbModel[uint64] = OpLog{}

type OpType int8

const (
	OpTypeFinishMission OpType = iota
	OpTypeExchangeProduct
)

// OpLog 操作记录表
//CREATE TABLE `op_log` (
//`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
//`type` tinyint(4) NOT NULL COMMENT '操作类型 0-完成任务 1-兑换商品',
//`open_id` varchar(100) NOT NULL,
//`object_id` bigint(20) unsigned NOT NULL COMMENT '对象ID',
//`price` bigint(20) unsigned NOT NULL COMMENT '金额',
//`content` varchar(100) NOT NULL,
//`finish_time` timestamp NOT NULL,
//`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//`update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//`delete_time` timestamp NULL DEFAULT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='操作记录表'
type OpLog struct {
	Id         uint64         `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Type       OpType         `gorm:"type;NOT NULL" json:"type"`
	OpenId     string         `gorm:"column:open_id;NOT NULL" json:"open_id"`
	ObjectId   uint64         `gorm:"column:object_id;NOT NULL" json:"object_id"`
	Price      uint64         `gorm:"column:amount;NOT NULL" json:"price"`
	Content    string         `gorm:"column:content;NOT NULL" json:"content"`
	FinishTime time.Time      `gorm:"column:finish_time;NOT NULL" json:"finish_time"`
	CreateTime *time.Time     `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`
	UpdateTime *time.Time     `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time" json:"delete_time"`
}

func (p OpLog) GetId() uint64 {
	return p.Id
}

func (p OpLog) SetId(id uint64) {
	p.Id = id
}

func (p OpLog) TableName() string {
	return "op_log"
}

func (p OpLog) SetUpdateTime(time time.Time) {
	p.UpdateTime = &time
}
