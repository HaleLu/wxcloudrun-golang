package model

import (
	"gorm.io/gorm"
	"time"
)

var _ DbModel[uint64] = FinanceLog{}

// FinanceLog 财务记录表
//CREATE TABLE `finance_log` (
//`id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
//`month` varchar (20) NOT NULL DEFAULT '',
//`content` longtext NULL,
//`total` bigint NOT NULL COMMENT '余额总计',
//`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
//`update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//`delete_time` timestamp NULL,
//PRIMARY KEY (`id`),
//unique `uk_month` USING btree (`month`)
//) COMMENT = "财务记录表" ENGINE = innodb DEFAULT CHARACTER SET = "utf8mb4"
type FinanceLog struct {
	Id         uint64         `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Month      string         `gorm:"column:month;default:;NOT NULL"`
	Content    string         `gorm:"column:content"`
	Total      int64          `gorm:"column:total;NOT NULL;comment:'余额总计'"`
	CreateTime *time.Time     `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdateTime *time.Time     `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time"`
}

func (p FinanceLog) GetId() uint64 {
	return p.Id
}

func (p FinanceLog) SetId(id uint64) {
	p.Id = id
}

func (p FinanceLog) TableName() string {
	return "finance_log"
}

func (p FinanceLog) SetUpdateTime(time time.Time) {
	p.UpdateTime = &time
}

type Item struct {
	Name   string
	Amount int64
}

type FinanceLogContent struct {
	Savings     []*Item
	Income      []*Item
	Expenditure []*Item
}
