package model

import (
	"gorm.io/gorm"
	"time"
)

var _ DbModel[uint64] = User{}

// User 用户表
//CREATE TABLE `user` (
// `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
// `open_id` varchar(100) NOT NULL,
// `amount` bigint(20) unsigned NOT NULL COMMENT '余额',
// `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
// `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// `delete_time` timestamp NULL DEFAULT NULL,
// PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表'
//go:generate sqlgen -file db/model/user.go -type Mission -db mysql
type User struct {
	Id         uint64         `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	OpenId     string         `gorm:"column:open_id;NOT NULL" json:"open_id"`
	Amount     uint64         `gorm:"column:amount;NOT NULL" json:"amount"` // 余额
	CreateTime *time.Time     `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`
	UpdateTime *time.Time     `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time"`
}

func (p User) GetId() uint64 {
	return p.Id
}

func (p User) SetId(id uint64) {
	p.Id = id
}

func (p User) TableName() string {
	return "user"
}

func (p User) SetUpdateTime(time time.Time) {
	p.UpdateTime = &time
}
